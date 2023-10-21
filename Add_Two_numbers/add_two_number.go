package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) addValue(value int) *ListNode {
	newNode := &ListNode{Val: value, Next: nil}
	current := n

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return n
}

// func getNumberFromListNode(n *ListNode) int {
// 	current := n
// 	result := new(big.Int)
// 	i := new(big.Int)
// 	for current != nil {
// 		aux := big.NewInt(int64(current.Val))
// 		result.Add(result, aux.Mul(aux, new(big.Int).Exp(i, big.NewInt(10))))
// 		i++
// 		current = current.Next
// 	}
// 	return result
// }

func sliceToListNode(slice []int) *ListNode {
	result := &ListNode{Val: slice[0], Next: nil}
	current := result

	for i := 1; i < len(slice); i++ {
		newNode := &ListNode{Val: slice[i], Next: nil}
		current = result

		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode
	}

	return result
}

func listNodeToArray(l1 *ListNode) []int {
	result := []int{}
	for l1 != nil {
		result = append(result, l1.Val)
		l1 = l1.Next
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	slice1 := listNodeToArray(l1)
	slice2 := listNodeToArray(l2)
	minLenght := 0
	maxLenght := 0
	if len(slice1) <= len(slice2) {
		minLenght = len(slice1)
		maxLenght = len(slice2)
	} else {
		minLenght = len(slice2)
		maxLenght = len(slice1)
	}
	result := []int{}

	i := 0
	rest := 0
	for ; i < minLenght; i++ {
		if (slice1[i] + slice2[i] + rest) <= 9 {
			result = append(result, slice1[i]+slice2[i]+rest)
			rest = 0
		} else {
			result = append(result, slice1[i]+slice2[i]-10+rest)
			rest = 1
		}
	}

	for ; i < maxLenght; i++ {
		if len(slice1) <= len(slice2) {
			if (slice2[i] + rest) < 9 {
				result = append(result, slice2[i]+rest)
				rest = 0
			} else {
				result = append(result, slice2[i]-10+rest)
				rest = 1
			}
		} else {
			if (slice1[i] + rest) <= 9 {
				result = append(result, slice1[i]+rest)
				rest = 0
			} else {
				result = append(result, slice1[i]-10+rest)
				rest = 1
			}
		}
	}
	if rest == 1 {
		result = append(result, rest)
	}

	return sliceToListNode(result)
}

// func numberToListNode(sum int) *ListNode {
// 	val := sum % 10
// 	result := &ListNode{Val: val}
// 	sum /= 10
// 	for sum > 0 {
// 		val = sum % 10
// 		result = addValue(result, val)
// 		sum /= 10
// 	}
// 	return result
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	nL1 := sumListNode(l1)
// 	nL2 := getNumberFromListNode(l2)
// 	fmt.Println(">>>>> Sum =", nL1+nL2)
// 	return numberToListNode(nL1 + nL2)
// }

func printNodeList(n *ListNode) {
	current := n
	for current != nil {
		fmt.Print(">>>> val = ", current.Val, " ")
		current = current.Next
	}
}

func main() {
	sliceL1 := []int{8, 6}
	sliceL2 := []int{6, 4, 8}

	laux1 := &ListNode{Val: sliceL1[0]}
	for _, val := range sliceL1[1:] {
		laux1.addValue(val)
	}

	laux2 := &ListNode{Val: sliceL2[0]}
	for _, val := range sliceL2[1:] {
		laux2.addValue(val)
	}

	// fmt.Println(laux1.getNode())
	// fmt.Println(laux2.getNode())
	// fmt.Println()
	printNodeList(addTwoNumbers(laux1, laux2))
	// addTwoNumbers(laux1, laux2)
}
