package LNode

import "fmt"

// LNode 链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func (l *LNode) Reverse() {
	if l == nil || l.Next == nil {
		return
	}
	var pre *LNode    //定义前驱结点
	var cur *LNode    //定义当前结点
	next := l.Next    //把后继结点存起来，防止丢失
	for next != nil { //l.next为空说明遍历到最后一个了，反之说明没有到最后
		cur = next.Next
		next.Next = pre //next指向pre，实现逆序
		pre = next      //后移前驱结点
		next = cur      //后移后驱结点
	}
	l.Next = pre //让链表的第一个结点指向头结点
}

// CreateNode 创建链表
func (l *LNode) CreateNode(max int) {
	cur := l
	for i := 1; i <= max; i++ {
		cur.Next = &LNode{} //分配内存
		cur.Next.Data = i
		cur = cur.Next //向后移动指针

	}

}

// PrintNode 打印链表
func (l *LNode) PrintNode(info string) {
	fmt.Print(info)                                 //打印“逆序前或逆序后”
	for cur := l.Next; cur != nil; cur = cur.Next { //遍历链表打印
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}
