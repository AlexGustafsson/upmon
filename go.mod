module github.com/AlexGustafsson/upmon

go 1.16

require (
	github.com/AlexGustafsson/upmon/api v0.0.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-ping/ping v0.0.0-20210506233800-ff8be3320020
	github.com/gofiber/fiber/v2 v2.15.0
	github.com/hashicorp/memberlist v0.2.4
	github.com/knadh/koanf v1.2.0
	github.com/mitchellh/mapstructure v1.4.1
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

replace github.com/AlexGustafsson/upmon/api => ./api
