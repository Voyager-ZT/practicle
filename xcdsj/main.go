package main

import (
	"fmt"
	"sort"
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
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening:" + a.msg)
	return nil
}

func main() {
	door := &AutoDoor{
		OpenCloser: &Door{
			open: false,
			lock: false,
		},
		delay: 3,
		msg:   "warning",
	}

	door.Open()
	if v, ok := door.OpenCloser.(*Door); ok { // 类型断言
		fmt.Println("door.Open()=", v)
	}

	door.OpenCloser.Open()
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println("door.OpenCloser.Open()=", v)
	}

	door.OpenCloser.Close()
	door.Close()
	if v, ok := door.OpenCloser.(*Door); ok { //类型断言
		fmt.Println("door.Close()=", v)
	}

	reverse := sort.Reverse()
	reverse.Len()
}
