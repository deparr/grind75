/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

    curr, runner := head, head.Next

    for runner != nil {
        curr = curr.Next
        runner = runner.Next
        if runner != nil {
            runner = runner.Next
        }
    }

    return curr
}

