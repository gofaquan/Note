package LNode_test

import (
	"fmt"
	"github.com/gofaquan/Note/DataStructure/LNode"
	"testing"
)

func TestLNode_Reverse(t *testing.T) {
	head := &LNode.LNode{}
	fmt.Println("就地逆序")
	head.CreateNode(7)
	head.PrintNode("逆序前")

	head.Reverse()
	head.PrintNode("逆序后")

	for head.Next != nil {
		fmt.Println(head.Next.Data)
		head = head.Next
	}
}

//结果：
//就地逆序
//逆序前1 2 3 4 5 6 7
//逆序后7 6 5 4 3 2 1
