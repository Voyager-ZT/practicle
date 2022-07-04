
package main

import (
	. "awesomeProject/list/method"
)



func main(){
	pHead1:=CreateList()
	PrintList(pHead1)
	pHead2:=CreateList()
	PrintList(pHead2)


	//node:=ListNode{Val: 0}
	//node.Next=pHead
	//pHead=&node
	//PrintList(pHead)


	//pHead=ReverseBetween(pHead,2,4)
	//pHead=ReverseKGroup(pHead,2)

	pHead3:=Merge(pHead1,pHead2)
	PrintList(pHead3)


}
