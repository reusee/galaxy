package gen

//go:generate protoc -I . etc.proto --go_out=plugins=grpc:etc
