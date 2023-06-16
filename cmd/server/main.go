package main

import (
	"fmt"

	"wormhole/server"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	// flag.Usage = server.Usage
	// server := flag.Bool("server", true, "Run server")
	// protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	// framed := flag.Bool("framed", false, "Use framed transport")
	// buffered := flag.Bool("buffered", false, "Use buffered transport")
	// addr := flag.String("addr", "localhost:20001", "Address to listen to")
	// secure := flag.Bool("secure", false, "Use tls secure transport")

	// flag.Parse()
	addr := "10.67.0.15:1320"
	secure := false

	var protocolFactory thrift.TProtocolFactory
	// switch *protocol {
	// case "compact":
	protocolFactory = thrift.NewTCompactProtocolFactoryConf(nil)

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if err := server.RunServer(transportFactory, protocolFactory, addr, secure); err != nil {
		fmt.Println("error running server:", err)
	}
}
