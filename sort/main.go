package main

import "fmt"

func merge(list1 []int,list2 []int)[]int{
	len1:=len(list1)
	len2:=len(list2)
	i:=0
	j:=0
	list:=make([]int,0)
	for i<len1&&j<len2{
		if list1[i]<list2[j]{
			list=append(list,list1[i])
			i++
		}else{
			list=append(list,list2[j])
			j++
		}
	}
	if i==len1{
		list=append(list,list2...)
	}
	if j==len2{
		list=append(list,list1...)
	}
	return list
}


func sort(list []int,left,right int,n int)[]int{
	fmt.Println(n,left,right)
	if right<left{
		return nil
	}else if left==right{
		return []int{list[left]}
	}
	mid:=(right+left)/2
	return merge(sort(list,left,mid,n+1),sort(list,mid+1,right,n+1))
}

func test(l []int){
	l[1]=9
	fmt.Printf("l addr:%p\n",&l)
	l=append(l,10)
	fmt.Printf("la addr:%p\n",&l)
	fmt.Println("l:",l)
}

func main(){
	x:=(2+3)/2
	fmt.Println(x)

	a:=[]int{1,2,4,3,9,8}
	list:=sort(a,0,5,0)
	fmt.Println(list)
}
