package main

//go:generate protoc -I.  --go_out=plugins=grpc:. ./hello.proto

func main() {
	go StartHttpServer(8001)
	go StartHttp2CleartextServer(8002)
	go StartHttp2TLSServer(8003)
	go StartGrpcServer(8004)
	go StartGrpcWithTLSServer(8005)
	go StartTCPServer(8006)
	StartUDPServer(8007)
}
