module github.com/AlexGustafsson/upmon

go 1.13

require (
	github.com/AlexGustafsson/upmon/cli v0.0.0
	github.com/AlexGustafsson/upmon/core v0.0.0
	github.com/AlexGustafsson/upmon/rpc v0.0.0
	github.com/AlexGustafsson/upmon/transport v0.0.0
	github.com/google/uuid v1.1.1
	google.golang.org/grpc v1.24.0
)

replace github.com/AlexGustafsson/upmon/core => ./core

replace github.com/AlexGustafsson/upmon/rpc => ./rpc

replace github.com/AlexGustafsson/upmon/transport => ./transport

replace github.com/AlexGustafsson/upmon/cli => ./cli
