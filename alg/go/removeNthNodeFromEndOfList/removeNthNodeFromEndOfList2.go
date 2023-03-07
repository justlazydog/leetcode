package removeNthNodeFromEndOfList

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	return removeNode(head, &n)
}

/*
在实现递归删除链表倒数第n个节点的算法时，使用指针参数n的目的是为了在递归过程中可以对n的值进行修改，从而达到找到链表倒数第n个节点的效果。
具体来说，我们需要在递归函数的每一层中对n的值进行减1操作，以表示当前节点是从链表末尾开始数的第几个节点。
由于在递归函数中，我们需要将链表中的某个节点删除，因此我们不能直接使用一个普通的整数类型来存储n的值，否则在递归函数返回时无法更新n的值。
*/
func removeNode(node *ListNode, n *int) *ListNode {
	if node == nil {
		return nil
	}
	node.Next = removeNode(node.Next, n)
	*n--
	if *n == 0 {
		return node.Next
	}
	return node
}
