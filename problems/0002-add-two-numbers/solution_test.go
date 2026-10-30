package addtwonumbers

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "example 1",
			l1:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:   &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name: "example 2",
			l1:   &ListNode{Val: 0},
			l2:   &ListNode{Val: 0},
			want: &ListNode{Val: 0},
		},
		{
			name: "example 3",
			l1:   &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			l2:   &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)

			if !equalLists(got, tt.want) {
				t.Errorf(
					"addTwoNumbers() = %s, want %s",
					formatList(got),
					formatList(tt.want),
				)
			}
		})
	}
}

func equalLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func formatList(l *ListNode) string {
	if l == nil {
		return "[]"
	}
	result := ""
	for l != nil {
		result += fmt.Sprintf("%d", l.Val)
		if l.Next != nil {
			result += " -> "
		}
		l = l.Next
	}
	return "[" + result + "]"
}
