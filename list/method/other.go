package method

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

//有头结点
func ReverseList1( pHead *ListNode ) *ListNode {
	// write code here
	var s *ListNode
	var p *ListNode
	p=pHead.Next
	pHead.Next=nil
	for p!=nil{
		s=p.Next
		p.Next=pHead.Next
		pHead.Next=p
		p=s
	}
	return pHead
}


//无头结点
func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	var s *ListNode
	var p *ListNode
	p=pHead
	pHead=nil   //一定要，不然最后一个节点指向自身
	for p!=nil{
		s=p.Next
		p.Next=pHead
		pHead=p
		p=s
	}
	return pHead
}

func CreateList()(pHead *ListNode){
	var i int
	var n int
	var pTail *ListNode
	fmt.Scanln(&i)
	for i!=999{
		p:=ListNode{Val:i }
		if n==0{
			pHead=&p
			pTail=&p
		}else{
			pTail.Next=&p
			pTail=&p
		}
		n++
		fmt.Scanln(&i)
	}
	return pHead
}

func PrintList(pHead *ListNode){
	p:=pHead
	for p!=nil{
		fmt.Printf("%v\t",p.Val)
		p=p.Next
	}
	fmt.Println()
}