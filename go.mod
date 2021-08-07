module github.com/AlexGustafsson/upmon

go 1.16

require (
	github.com/AlexGustafsson/upmon/api v0.0.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-ping/ping v0.0.0-20210506233800-ff8be3320020
	github.com/gofiber/fiber/v2 v2.15.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/memberlist v0.2.4
	github.com/knadh/koanf v1.2.0
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/tools v0.1.5
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/AlexGustafsson/upmon/api => ./api
