package work1013

import (
	"testing"
)

// 定义一个链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 转换成链表
func newList(args ...int) *ListNode {
	var head ListNode
	var p *ListNode = &head
	for _, v := range args {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return &head
}

// 链表转换切片
func listToSlice(head *ListNode) []int {
	ans := []int{}
	for p := head; p != nil; p = p.Next {
		ans = append(ans, p.Val)
	}
	return ans
}

func TestMiddleNode(t *testing.T) {
	data := []struct {
		val    []int
		answer []int
	}{
		{val: []int{1, 2, 3, 4, 5}, answer: []int{3, 4, 5}},
		{val: []int{1, 2, 3, 4, 5, 6}, answer: []int{4, 5, 6}},
	}
	for _, v := range data {
		q := newList(v.val...)
		ans := listToSlice(middleNode(q.Next))
		//如果序列长度不相等
		if len(ans) != len(v.answer) {
			t.Fatalf("expect:[%v] != result[%v]", ans, v.answer)
		} else {
			//长度相等后判断值是否相等
			for i := 0; i < len(ans); i++ {
				if ans[i] != v.answer[i] {
					t.Fatalf("expect:[%v] != result[%v]", sortedSquares(v.val), v.answer)
				}
			}
		}
	}
}

// middleNode返回链表中间节点
func middleNode(head *ListNode) *ListNode {
	l := 1
	//指向头节点的指针
	p := head
	//获得链表长度
	for ; p.Next != nil; p = p.Next {
		l++
	}
	p = head
	//使指针p指向链表中间
	for i := 0; i < l/2; i++ {
		p = p.Next
	}
	return p
}
