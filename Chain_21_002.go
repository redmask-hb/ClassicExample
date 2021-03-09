/*

输入一个链表，反转链表后，输出新链表的表头。

*/
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var curNode, p *ListNode
	if pHead == nil {
		return nil
	}
	head:=ListNode{0,pHead}
	pHead=&head
	curNode=pHead.Next
	for curNode.Next != nil {
		p=curNode.Next
		curNode.Next=p.Next
		p.Next=pHead.Next
		pHead.Next=p

	}
	return pHead.Next
}
