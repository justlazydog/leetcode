// Source : https://leetcode.cn/problems/design-linked-list
// Date   : 2023-04-26

/**********************************************************************************
你可以选择使用单链表或者双链表，设计并实现自己的链表。
单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
实现 MyLinkedList 类：

	MyLinkedList() 初始化 MyLinkedList 对象。
	int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
	void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
	void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
	void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
	void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。


示例：

输入
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
输出
[null, null, null, null, 2, null, 3]
解释
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
myLinkedList.get(1);              // 返回 2
myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
myLinkedList.get(1);              // 返回 3


提示：

	0 <= index, val <= 1000
	请不要使用内置的 LinkedList 库。
	调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。
**********************************************************************************/

package designLinkedList

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type MyLinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{Size: 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Size {
		return -1
	}

	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}

	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}

	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.Size {
		return
	}

	if index == 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.Size {
		l.AddAtTail(val)
		return
	}

	node := &Node{Val: val}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}

	node.Prev = cur
	node.Next = cur.Next
	cur.Next.Prev = node
	cur.Next = node

	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size {
		return
	}

	if l.Size == 1 {
		l.Head = nil
		l.Tail = nil
	} else if index == 0 {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	} else if index == l.Size-1 {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	} else {
		cur := l.Head
		for i := 0; i < index; i++ {
			cur = cur.Next
		}

		cur.Prev.Next = cur.Next
		cur.Next.Prev = cur.Prev
	}

	l.Size--
}

//// DoubleNode is a single element in a double linked list
//type DoubleNode struct {
//	Data int
//	Next *DoubleNode
//	Prev *DoubleNode
//}
//
//type MyLinkedList struct {
//	Len   int
//	Dummy *DoubleNode
//}
//
//func Constructor() MyLinkedList {
//	return MyLinkedList{Dummy: &DoubleNode{}}
//}
//
//func (l *MyLinkedList) Get(index int) int {
//	if index >= l.Len {
//		return -1
//	}
//
//	cur := l.Dummy.Next
//	for cur != nil {
//		if index == 0 {
//			return cur.Data
//		}
//		index--
//		cur = cur.Next
//	}
//	return -1
//}
//
//func (l *MyLinkedList) AddAtHead(val int) {
//	node := &DoubleNode{Data: val}
//	head := l.Dummy.Next
//	l.Dummy.Next = node
//
//	if head == nil {
//		node.Prev = node
//		node.Next = node
//		l.Len++
//		return
//	} else {
//		node.Next = head
//		node.Prev = head.Prev
//		head.Prev.Next = node
//		head.Prev = node
//		l.Len++
//		return
//	}
//}
//
//func (l *MyLinkedList) AddAtTail(val int) {
//	node := &DoubleNode{Data: val}
//	head := l.Dummy.Next
//
//	if head == nil {
//		l.Dummy.Next = node
//		node.Prev = node
//		node.Next = node
//		l.Len++
//		return
//	} else {
//		node.Next = head
//		node.Prev = head.Prev
//		head.Prev.Next = node
//		head.Prev = node
//		l.Len++
//		return
//	}
//}
//
//func (l *MyLinkedList) AddAtIndex(index int, val int) {
//	if index > l.Len {
//		return
//	}
//
//	if index == 0 {
//		l.AddAtHead(val)
//		return
//	}
//
//	if index == l.Len {
//		l.AddAtTail(val)
//		return
//	}
//
//	node := &DoubleNode{Data: val}
//	head := l.Dummy.Next
//	for head != nil {
//		if index == 0 {
//			node.Next = head
//			node.Prev = head.Prev
//			head.Prev.Next = node
//			head.Prev = node
//			l.Len++
//			return
//		}
//		head = head.Next
//		index--
//	}
//}
//
//func (l *MyLinkedList) DeleteAtIndex(index int) {
//	if index >= l.Len {
//		return
//	}
//
//	head := l.Dummy.Next
//	for head != nil {
//		if index == 0 {
//			if head == l.Dummy.Next {
//				l.Dummy.Next = head.Next
//			}
//			head.Prev.Next = head.Next
//			head.Next.Prev = head.Prev
//			l.Len--
//			return
//		}
//		head = head.Next
//		index--
//	}
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
