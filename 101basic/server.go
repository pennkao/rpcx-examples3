package main

import (
	"flag"

	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	s := server.Server{}
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)
}
