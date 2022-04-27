/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var fast, slow *ListNode

	fast = head.Next
	slow = head

	// for i:=0; i<100; i++ {
	for {
		if fast == nil || fast.Next == nil {
			return false
		}

		fmt.Println(fast, slow, fast == slow)

		if fast == slow {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}