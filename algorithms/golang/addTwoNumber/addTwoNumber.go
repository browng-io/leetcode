package addtwonumber

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// addTwoNumbers func:
//
//  1. Init a dummy ListNode as Head
//
//  2. Use currentNode as an pointer of dummy ListNode for formula : sum 2 linked list from input
//
//  3. init a total var for summing l1.Value l2.Value carry
//     init a carry var for case: the sum of 2 node value greater than 10
//     then carry should have value 1 and add to next sum of 2 node
//     the current sum of 2 node will have value = mod 10
//     Note: carry value calculated by formula: floor division of total with 10
//
//     init a while loop calculating sum l1 & l2 with conditions:
//     - l1 or l2 not None
//     - or carry not zero value
//     currentNode will pointer next sum of 2 node each loop
//     Result is Head.Next (the first Node of Head has nothing)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	currentNode := head
	var total, l1Value, l2Value, carry int

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		} else {
			l1Value = 0
		}
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		} else {
			l2Value = 0
		}
		total = l1Value + l2Value + carry
		currentNode.Next = &ListNode{Val: total % 10}
		carry = total / 10
		currentNode = currentNode.Next
	}
	return head.Next
}
