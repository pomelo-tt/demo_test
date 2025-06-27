package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Demo1()
	// Demo2()
	res := Demo3()
	expiration := 10*time.Second
	fmt.Println(expiration)
	expireMillis := expiration.Milliseconds()
	fmt.Println(expireMillis)
	fmt.Println(res)
	res1 := [][]int{}
    path := []int{}
	res1 = append(res1, append([]int(nil), path...))
	fmt.Println(res1)
}


// ========================================================================Demo6================================================================================
// ========================================================================Demo5================================================================================
// ========================================================================Demo4================================================================================

// ========================================================================Demo3================================================================================
func Demo3() []string {
	fmt.Println("===================Demo3 start=====================")
	s := "0000"
	if len(s) < 4 {
		return []string{}
	}
	res := []string{}
	path := []string{}
	var dfs func(int, int)
	dfs = func(start, seq int) {
		if seq == 4 && start == len(s) {
			res = append(res, strings.Join(path, "."))
			return
		}
		if seq == 4 || start == len(s) {
			return
		}
		for end := start + 1; end <= len(s) && end <= start+3; end++ {
			part := s[start:end]
			// if !isValid(part) || (len(part) > 0 && part[0] == '0') {
			if !isValid(part) || (len(part) > 1 && part[0] == '0') {
				continue
			}
			path = append(path, part)
			dfs(end, seq+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	fmt.Println("===================Demo3 end=====================")
	return res
}

func isValid(s string) bool {
	num, _ := strconv.Atoi(s)
	return num >= 0 && num <= 255
}

// ========================================================================Demo2================================================================================
func Demo2() {
	fmt.Println("===================Demo2 start=====================")
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	res := sortList(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println("===================Demo2 end=====================")
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	midNode := findMid(head)
	right := sortList(midNode.Next)
	midNode.Next = nil
	left := sortList(head)
	return merge(left, right)
}

func findMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}

// ========================================================================Demo1================================================================================
func Demo1() {
	fmt.Println("===================Demo1 start=====================")
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	newHead := reverseKGroup(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
	fmt.Println("===================Demo1 end=====================")
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		nHead := tail.Next
		rHead, rTail := reverse(head, tail)
		rTail.Next = nHead

		pre.Next = rHead
		pre = rTail
		head = nHead
	}
	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	end := tail.Next
	for cur != end {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}

func RenewalWorker(ctx context.Context, key, value string, exp time.Duration) {
	// 设置最大续期时长
	maxRenewal := 5 * time.Minute
	expTicker := time.NewTicker(exp / 3)
	defer expTicker.Stop()
	start := time.Now()

	for {
		select {
		case <-ctx.Done(): // 业务完成
			return
		case <-expTicker.C:
			if time.Since(start) > maxRenewal {
				return // 自动停止续期
			}
			// 执行续期逻辑
			// RenewLock(ctx, key, value, exp)
		}
	}
}
