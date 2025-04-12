// Package benthosqrcode provides a processor for Benthos that generates QR codes from message content.
package benthosqrcode

import (
	"bytes"
	"context"
	"fmt"
	"maps"
	"slices"

	goqr "github.com/piglig/go-qr"
	"github.com/redpanda-data/benthos/v4/public/service"
)

var eclEnum = map[string]goqr.Ecc{
	"low":      goqr.Low,
	"medium":   goqr.Medium,
	"quartile": goqr.Quartile,
	"high":     goqr.High,
}

func init() {
	configSpec := service.NewConfigSpec().
		Summary("Creates a QR code from the current content of the message.").
		Fields(
			service.NewStringEnumField("format", "png", "svg").
				Description("The image format of the QR code"),
			service.NewStringEnumField("ecl", slices.Collect(maps.Keys(eclEnum))...).
				Default("medium").
				Description("The error correction level in a QR code symbol"),
			service.NewIntField("scale").
				Default(10).
				Description("Scale of image"),
			service.NewIntField("border").
				Default(4).
				Description("Border of image"),
			service.NewObjectField(
				"svg",

				service.NewStringField("light").
					Default("#FFFFFF").
					Description("Color to use for light sections of the QR code"),
				service.NewStringField("dark").
					Default("#000000").
					Description("Color to use for dark sections of the QR code"),
				service.NewBoolField("xmlHeader").
					Default(false),
				service.NewBoolField("optimal").
					Default(false).
					Description("Output optimized SVG file (regions with connected black pixels are merged into a single path)"),
			),
		)

	err := service.RegisterProcessor("qrcode", configSpec, newProcessor)
	if err != nil {
		panic(err)
	}
}

func newProcessor(conf *service.ParsedConfig, _ *service.Resources) (service.Processor, error) {
	ecl, err := conf.FieldString("ecl")
	if err != nil {
		return nil, err
	}

	format, err := conf.FieldString("format")
	if err != nil {
		return nil, err
	}

	scale, err := conf.FieldInt("scale")
	if err != nil {
		return nil, err
	}

	border, err := conf.FieldInt("border")
	if err != nil {
		return nil, err
	}

	light, err := conf.FieldString("svg", "light")
	if err != nil {
		return nil, err
	}

	dark, err := conf.FieldString("svg", "dark")
	if err != nil {
		return nil, err
	}

	var options []func(config *goqr.QrCodeImgConfig)

	xmlHeader, err := conf.FieldBool("svg", "xmlHeader")
	if err != nil {
		return nil, err
	}

	if xmlHeader {
		options = append(options, goqr.WithSVGXMLHeader(true))
	}

	optimal, err := conf.FieldBool("svg", "optimal")
	if err != nil {
		return nil, err
	}

	if optimal {
		options = append(options, goqr.WithOptimalSVG())
	}

	return &processor{
		ecl:    eclEnum[ecl],
		config: goqr.NewQrCodeImgConfig(scale, border, options...),
		format: format,

		light: light,
		dark:  dark,
	}, nil
}

type processor struct {
	ecl    goqr.Ecc
	config *goqr.QrCodeImgConfig
	format string

	light string
	dark  string
}

func (r *processor) Process(_ context.Context, m *service.Message) (service.MessageBatch, error) {
	bytesContent, err := m.AsBytes()
	if err != nil {
		return nil, err
	}

	qr, err := goqr.EncodeText(string(bytesContent), r.ecl)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	switch r.format {
	case "svg":
		err = qr.WriteAsSVG(r.config, &buf, r.light, r.dark)
		if err != nil {
			return nil, err
		}
	case "png":
		err = qr.WriteAsPNG(r.config, &buf)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", r.format)
	}

	m.SetBytes(buf.Bytes())
	return []*service.Message{m}, nil
}

func (r *processor) Close(_ context.Context) error {
	return nil
}
