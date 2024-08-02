package server

import "flag"

var (
	addr = flag.String("addr", "localhost:60003", "the address to connect to")
)
