package main

import (
	"zinx/net"
	"zinx/ziface"
	"fmt"
)

type PingRouter struct {
	net.BaseRouter
}
func (r *PingRouter)PreHandle(request ziface.IRequest){
	fmt.Println("Call Router PreHandler...")
	//给客户端回写一个 数据
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("\n before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")

	}
}

func (r *PingRouter)Handler(re ziface.IRequest){
	fmt.Println("call Router Hnader ....")
	_,err:=re.GetConnection().GetTCPConnection().Write([]byte("Ping...Ping...PIng....\n"))
	if err!=nil{
		fmt.Println("Router Hander error :",err)
	}

}

func (r *PingRouter)PostHandle(re ziface.IRequest){
	fmt.Println("call Router PostHnade ....")
	_,err:=re.GetConnection().GetTCPConnection().Write([]byte("Router Post PIng....\n"))
	if err!=nil{
		fmt.Println("Router PostHande error :",err)
	}
}
func main(){
	s:=net.NewServer("zinxV0.3")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
