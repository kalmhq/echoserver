package main

import (
	"context"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(w, "HostName: %s\n", name)

	_, _ = fmt.Fprintf(w, "\nRequest Info:\n")
	_, _ = fmt.Fprintf(w, "    content-length: %d\n", req.ContentLength)
	_, _ = fmt.Fprintf(w, "    client address: %s\n", req.RemoteAddr)
	_, _ = fmt.Fprintf(w, "    method: %s\n", req.Method)
	_, _ = fmt.Fprintf(w, "    path: %s\n", req.URL.Path)
	_, _ = fmt.Fprintf(w, "    query: %s\n", req.URL.RawQuery)
	_, _ = fmt.Fprintf(w, "    request_version: %s\n", req.Proto)
	_, _ = fmt.Fprintf(w, "    uri: %s\n", req.RequestURI)
	_, _ = fmt.Fprintf(w, "    tls: %t\n", req.TLS != nil)

	_, _ = fmt.Fprintf(w, "\nHeaders:\n")

	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "    %v: %v\n", name, h)
		}
	}

	_, _ = fmt.Fprintf(w, "\nBody:\n")

	if req.Body != nil && req.ContentLength > 0 {
		_, _ = fmt.Fprintf(w, "Length: %d\n", req.ContentLength)

	} else {
		_, _ = fmt.Fprintf(w, "No Body\n")
	}

	_, _ = fmt.Fprintf(w, "\n")
}

func StartHttpServer(port int) {
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		Handler: http.HandlerFunc(handler),
	}
	addr := "0.0.0.0:" + strconv.Itoa(port)
	fmt.Printf("listening on %s, http1.0, http1.1\n", addr)
	server.ListenAndServe()
}

func StartHttp2CleartextServer(port int) {
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		Handler: h2c.NewHandler(http.HandlerFunc(handler), h2s),
	}
	fmt.Printf("listening on %s, support http1.0, http1.1, non-TLS HTTP/2 (aka h2c, upgrade, prior knowledge)\n", server.Addr)
	server.ListenAndServe()
}

func StartHttp2TLSServer(port int) {
	server := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(port),
		Handler: http.HandlerFunc(handler),
	}

	err := http2.ConfigureServer(server, &http2.Server{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("listening on %s, TLS enbaled, http1.1, http/2\n", server.Addr)
	if err := server.ListenAndServeTLS("./default.pem", "./default.key"); err != nil {
		panic(err)
	}
}

type GrpcServer struct{}

func (*GrpcServer) Greeting(ctx context.Context, msg *GreetingMessage) (*GreetingReply, error) {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	p, ok := peer.FromContext(ctx)

	if !ok {
		panic(fmt.Errorf("get peer from context not success"))
	}

	return &GreetingReply{
		Hostname:      name,
		ClientAddress: p.Addr.String(),
		AuthInfo:      "",
	}, nil
}

func StartGrpcServer(port int) {
	addr := ":" + strconv.Itoa(port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	RegisterHelloWorldServer(s, &GrpcServer{})

	fmt.Printf("listening on %s, grpc\n", addr)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func StartGrpcWithTLSServer(port int) {
	c, err := credentials.NewServerTLSFromFile("./default.pem", "./default.key")

	if err != nil {
		panic(err)
	}

	addr := ":" + strconv.Itoa(port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(c))
	reflection.Register(s)
	RegisterHelloWorldServer(s, &GrpcServer{})

	fmt.Printf("listening on %s, grpc\n", addr)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
