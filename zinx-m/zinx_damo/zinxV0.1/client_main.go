package main

import (
	"fmt"
	"net"
	"os"
)

func main(){
	fmt.Println("client start................")
	conn,err:=net.Dial("tcp","127.0.0.1:8999")
	defer conn.Close()
	if err!=nil{
		fmt.Println("客户端连接错误.......",err)
		return
	}

	for{
		buff:=make([]byte,512)
		n,err:=os.Stdin.Read(buff)
		if err!=nil{
			fmt.Println("Stdin.Read error :",err)
			return
		}
		_,err=conn.Write(buff[:n])
		if err!=nil{
			fmt.Println("Write error :",err)
			continue
		}

		buf :=make([]byte,512)
		cnt,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("Read error:",err)
			continue
		}
		fmt.Printf(" servar call back : %s, cnt = %d\n", buf[:cnt-1], cnt)

	}

}
