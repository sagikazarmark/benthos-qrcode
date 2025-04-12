# [Benthos](https://github.com/redpanda-data/benthos) plugin: QR Code

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/benthos-qrcode/ci.yaml?style=flat-square)](https://github.com/sagikazarmark/benthos-qrcode/actions/workflows/ci.yaml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sagikazarmark/benthos-qrcode?style=flat-square&color=61CFDD)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/sagikazarmark/benthos-qrcode/badge?style=flat-square)](https://deps.dev/go/github.com%252Fsagikazarmark%252Fbenthos-qrcode)

## Usage

Create a custom Benthos binary:

```go
package main

import (
	"context"

	"github.com/redpanda-data/benthos/v4/public/service"

	_ "github.com/sagikazarmark/benthos-qrcode"
)

func main() {
	service.RunCLI(context.Background())
}
```

> [!TIP]
> Read more about Benthos plugins [here](https://github.com/redpanda-data/redpanda-connect-plugin-example).

TODO: add an example configuration

## Development

**For an optimal developer experience, it is recommended to install [Nix](https://nixos.org/download.html) and [direnv](https://direnv.net/docs/installation.html).**

```shell
go test -v ./...
```

## License

The project is licensed under the [MIT License](LICENSE).
