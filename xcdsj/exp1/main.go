package main

import (
	"fmt"
	"strconv"
	"time"
)

// 接口：一组方法的集合
// OpenCloser 接口定义两个方法 返回 error
type OpenCloser interface {
	Open() error
	Close() error
}

type Door struct {
	open bool // 门的状态是否开启
	lock bool // 门的状态是否上锁
}

func (d *Door) Open() error {
	fmt.Println("door open ...")
	d.open = true
	return nil
}

func (d *Door) Close() error {
	fmt.Println("door close ...")
	d.open = false
	return nil
}

type AutoDoor struct {
	OpenCloser        // 匿名接口
	delay      int    // 延迟多长时间开启
	msg        string // 自动开启时的警报
}

func (a *AutoDoor) Open() error {
	fmt.Println("改写接口方法")
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening:" + a.msg)
	return nil
}

func f1(closer OpenCloser){
	fmt.Printf("结构体嵌套接口就实现了接口的方法\n")
	//closer.Close()
}

func main() {
	door := &AutoDoor{
	}

	//结构体嵌套接口就实现了接口的方法
	f1(door)

	//结构体可以重新实现接口的方法
	door=&AutoDoor{
		OpenCloser: &Door{
			open: false,
			lock: false,
		},
		delay: 3,
		msg:   "warning",
	}

	door.Open()
	door.Close()
}