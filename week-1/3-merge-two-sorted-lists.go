/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
       return list2 
    }

    if list2 == nil {
       return list1 
    }

    var (
        running *ListNode
        head *ListNode
    )
    if list1.Val < list2.Val {
        running = list1
        head = list1
        list1 = list1.Next
    } else {
        running = list2
        head = list2
        list2 = list2.Next
    }
    
    for list1 != nil && list2 != nil {
        if list2.Val < list1.Val {
            running.Next = list2
            list2 = list2.Next

        } else {
            running.Next = list1
            list1 = list1.Next
        }
        running = running.Next
    }

    if list1 == nil {
        running.Next = list2
    } else {
        running.Next = list1
    }

    return head
}

