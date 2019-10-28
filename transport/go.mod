module github.com/AlexGustafsson/upmon/rpc

go 1.13

require (
  google.golang.org/grpc v1.24.0
  github.com/AlexGustafsson/upmon/core v0.0.0
)

replace github.com/AlexGustafsson/upmon/core => ../core
