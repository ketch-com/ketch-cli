module go.ketch.com/cli/ketch-cli

go 1.16

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/jinzhu/copier v0.3.2
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/xeipuuv/gojsonschema v1.2.0
	go.ketch.com/lib/app v0.3.0
	go.ketch.com/lib/oid v1.2.0
	go.ketch.com/lib/orlop v1.40.0
	go.ketch.com/lib/webhook-client v0.3.1
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

replace go.ketch.com/lib/app => ../app/
