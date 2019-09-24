module github.com/petuhovskiy/grpc-hydra-bench/hydracli

go 1.13

replace github.com/ory/hydra/sdk/go/hydra => ./

require (
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.6
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.3
	github.com/ory/hydra/sdk/go/hydra v0.0.0-00010101000000-000000000000
)
