module github.com/AlexGustafsson/upmon

go 1.16

require (
  github.com/AlexGustafsson/upmon/api v0.0.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/gofiber/fiber/v2 v2.15.0
	github.com/google/btree v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.2.4
	github.com/knadh/koanf v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/AlexGustafsson/upmon/api => ./api
