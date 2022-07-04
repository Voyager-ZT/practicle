package method

//不带头结点
func ReverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
	// write code here
	if head==nil||m==n{
		return head
	}

	node:=ListNode{}
	node.Next=head
	head=&node

	i:=0
	p:=head
	for p!=nil&&i<m-1{
		p=p.Next
		i++
	}
	pre:=p
	p=pre.Next
	pre.Next=nil
	i++
	tmp:=p
	var s *ListNode
	for p!=nil&&i<=n{
		s=p.Next
		p.Next=pre.Next
		pre.Next=p
		p=s
		i++
	}
	//if tmp!=nil{
		tmp.Next=p
	//}
	return head.Next
}

func ReverseKGroup( head *ListNode ,  k int ) *ListNode {
	// write code here
	start:=head
	tail:=head
	i:=0
	for i<k{
		if tail==nil{
			return head
		}
		tail=tail.Next
		i++
	}
	cur:=head
	head=nil
	for cur!=tail{
		curNext:=cur.Next
		cur.Next=head
		head=cur
		cur=curNext
	}

	start.Next=ReverseKGroup(tail,k)
	return head
}

func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	// write code here
	var newHead *ListNode
	newHead=&ListNode{}
	cur:=newHead

	p:=pHead1
	q:=pHead2
	for p!=nil&&q!=nil{
		if p.Val<=q.Val{
			cur.Next=p
			p=p.Next
		}else{
			cur.Next=q
			q=q.Next
		}
		cur=cur.Next
	}
	if p==nil{
		cur.Next=q
	}
	if q==nil{
		cur.Next=p
	}
	return newHead.Next
}

func MergeKLists( lists []*ListNode ) *ListNode {
	// write code here

}

func divideMerge(lists []*ListNode, left int, right int) *ListNode{
	if right<left{
		return nil
	}else if left==right{
		return lists[left]
	}
	mid:=(left+right)/2
	return Merge(divideMerge(lists,left,mid),divideMerge(lists,mid+1,right))
}