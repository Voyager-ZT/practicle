package main

import "fmt"

func main(){
	a:=[]int{1}
	b:=[]int{2}
	fmt.Println(len(a))
	fmt.Println(len(b))
	if a[0]<b[0]{
		fmt.Println("a<b")
	}


	arr:=[10]int{0,1,2,3,4,5,6,7,8,9}
	s1:=arr[1:]
	s2:=arr[5:]
	s2=append(s2,10)
	s1[4]=99
	fmt.Println(arr)
	fmt.Println(s1)
	fmt.Println(s2)
}
