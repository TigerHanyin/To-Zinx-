package main

import "fmt"

type CPU interface {
	Cal()
}
type Card interface {
	Dis()
}
type Memory interface {
	Sto()
}
type Computer struct {
	cpu CPU
	card Card
	mem Memory
}

func NewComputer(cpu CPU,card Card,mem Memory)*Computer {
	c:=&Computer{
		cpu:cpu,
		card:card,
		mem:mem,
	}
	return c
}

func (this *Computer) Work() {

	this.cpu.Cal()
	this.card.Dis()
	this.mem.Sto()
}

type IntelCPU struct {}
func(this *IntelCPU)Cal(){
	fmt.Println("intel cpu ......")
}

type IntelCard struct {}
func(this *IntelCard)Dis(){
	fmt.Println("intel card ......")
}

type IntelMem struct {}
func(this *IntelMem)Sto(){
	fmt.Println("intel mem ......")
}

type KinCard struct {}
func(this *KinCard)Dis(){
	fmt.Println("kin card ......")
}

type NoMem struct {}
func(this *NoMem)Sto(){
	fmt.Println("nNo mem ......")
}

func main(){
	com1:=NewComputer(&IntelCPU{},&IntelCard{},&IntelMem{})
	com1.Work()
	fmt.Println("........................................")
	com2:=NewComputer(&IntelCPU{},&KinCard{},&NoMem{})
	com2.Work()

}