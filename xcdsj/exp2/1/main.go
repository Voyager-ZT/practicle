//方法提升
package main

import "fmt"

//******************************** start
type ip interface{
	PReceive()
}

type TP struct{
}

func (*TP)PReceive(){
	fmt.Println("可调用【指针接收者】方法")
}

type iv interface{
	VReceive()
}

type TV struct{

}
func (TV)VReceive(){
	fmt.Println("可调用【值接收者】方法")
}
//********************************* end

//嵌入类型为T
type Sa struct{
	iv
}
//嵌入类型为*T
type Sb struct{
	ip
}



func main(){
	//1.值接收者
	//嵌入T 调用S
	sa:=Sa{TV{}}
	sa.VReceive()
	//嵌入T 调用*S
	psa:=&Sa{TV{}}
	psa.VReceive()
	//嵌入*T 调用S
	sa1:=Sa{&TV{}}
	sa1.VReceive()
	//嵌入*T 调用*S
	psa1:=&Sa{&TV{}}
	psa1.VReceive()

	//2.指针接收者
	//嵌入T 调用S
	sb:=Sb{TP{}}
	sb.PReceive()
	//嵌入T 调用*Sgo
	psb:=&Sb{TP{}}
	psb.PReceive()
	//嵌入*T 调用S
	sb1:=Sb{&TP{}}
	sb1.PReceive()
	//嵌入*T 调用*S
	psb1:=&Sb{&TP{}}
	psb1.PReceive()
}
