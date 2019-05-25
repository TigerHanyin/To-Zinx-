package main

import "zinx/net"

func main(){
	s:=net.NewServer("zinxV0.1")
	s.Serve()
}
