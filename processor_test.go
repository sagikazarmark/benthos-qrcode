package benthosqrcode

import (
	"context"
	"testing"

	"github.com/hexops/autogold/v2"
	goqr "github.com/piglig/go-qr"
	"github.com/redpanda-data/benthos/v4/public/service"
	"github.com/stretchr/testify/require"
)

func TestProcessor_SVG(t *testing.T) {
	processor := &processor{
		ecl:    goqr.Medium,
		config: goqr.NewQrCodeImgConfig(10, 4, goqr.WithSVGXMLHeader(true)),
		format: "svg",
		light:  "#FFFFFF",
		dark:   "#000000",
	}

	result, err := processor.Process(
		context.Background(),
		service.NewMessage([]byte("hello world")),
	)
	require.NoError(t, err)

	resBytes, err := result[0].AsBytes()
	require.NoError(t, err)

	autogold.ExpectFile(t, resBytes)
}

func TestProcessor_PNG(t *testing.T) {
	processor := &processor{
		ecl:    goqr.Medium,
		config: goqr.NewQrCodeImgConfig(10, 4),
		format: "png",
	}

	result, err := processor.Process(
		context.Background(),
		service.NewMessage([]byte("hello world")),
	)
	require.NoError(t, err)

	resBytes, err := result[0].AsBytes()
	require.NoError(t, err)

	autogold.ExpectFile(t, resBytes)
}
