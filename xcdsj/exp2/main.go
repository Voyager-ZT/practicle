//方法提升
package main

import "fmt"

//******************************** start
type T struct{
	//i
}

//type i interface{
//	PReceive()
//	VReceive()
//}

func (*T)PReceive(){
	fmt.Println("可调用【指针接收者】方法")
}

func (T)VReceive(){
	fmt.Println("可调用【值接收者】方法")
}
//********************************* end

//嵌入类型为T
type Sa struct{
	T
}
//嵌入类型为*T
type Sb struct{
	*T
}

//Q:嵌入类型为T 指针接收者的方法 使用S为什么也能调用？请思考
func main(){
	sa:=Sa{T{}}
	sb:=Sb{&T{}}
	psa:=&Sa{T{}}
	psb:=&Sb{&T{}}
	//1.嵌入类型为T
	sa.VReceive()
	sa.PReceive()   //Q:嵌入类型为T 指针接收者的方法 使用S为什么也能调用？请思考  //因为sa可寻址,详细见下
	psa.VReceive()
	psa.PReceive()

	//2.嵌入类型为*T
	sb.VReceive()
	sb.PReceive()
	psb.VReceive()
	psb.PReceive()
}
//ANSWER:
//这是因为在一个为值类型的变量调用接收器的指针类型的方法时，golang会进行对该变量的取地址操作，从而产生出一个指针，之后再用这个指针调用方法。
//前提是这个变量要能取地址。

//再看下面函数，此时嵌入类型为T 指针接收者的方法 使用S
//因为结构体字面量的字段值不可取址
func f(){
	//1.嵌入类型为T
	Sa{T{}}.VReceive()
	Sa{T{}}.PReceive()
	(&Sa{T{}}).VReceive()
	(&Sa{T{}}).PReceive()

	//2.嵌入类型为*T
	Sb{&T{}}.VReceive()
	Sb{&T{}}.PReceive()
	(&Sb{&T{}}).VReceive()
	(&Sb{&T{}}).PReceive()
}