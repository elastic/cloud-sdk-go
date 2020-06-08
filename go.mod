module github.com/elastic/cloud-sdk-go

go 1.13

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/blang/semver v3.5.1+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/elastic/ecctl v1.0.0-rc1
	github.com/fatih/color v1.9.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/runtime v0.19.15
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.8
	github.com/hashicorp/go-multierror v1.0.0
	github.com/mitchellh/mapstructure v1.3.1
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/spf13/cobra v0.0.7
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.0
)

replace sourcegraph.com/sourcegraph/go-diff v0.5.1 => github.com/sourcegraph/go-diff v0.5.1
