package leet04

import (
	"math"
	"sort"
)

//	Remember:
//	- In Go the integer does not have func max min
//	- the operator / in integer always round down (math.Floor) the result

// Median of Two Sorted Arrays
// Naive Path:
//
//		merge them into third array and handle two cases (length's third array):
//	  - is odd:
//	    median is at length/2 index
//	  - is even:
//	    average of elements at indexes:
//	    . length/2
//	    . length/2 - 1

// Time complexity: O((M + N) * log(M + N)): Time required to sort the array of size M + N.
// Auxiliary space: O(M + N): Creating a new array of size M + N

func findMedianSortedArraysNaive(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	// use func sort.Ints in standard library to get result in increasing order
	sort.Ints(nums3)

	length := len(nums3)
	// length is odd
	if length%2 == 1 {
		return float64(nums3[(length-1)/2])
	}
	// length is even
	return float64((nums3[length/2] + nums3[(length/2)-1])) / 2
}

// Merging Arrays efficiently:
//
//	keep array sorted like Naive Path
//	dive into the process sorting, calculate median when reach half original size of given array
//	two cases like Naive Path

// Time Complexity: O(M + N): to merge both arrays
// Auxiliary Space: O(1): No extra space

func findMedianSortedArraysMergingEfficiently(nums1 []int, nums2 []int) float64 {
	//	init two pointer to 0 index each arrays
	idxA, idxB := 0, 0

	//	default median value init
	m1, m2 := -1, -1

	// set array's name to a, b (a only the shorter array)
	a, b := nums1, nums2
	if len(a) > len(b) {
		a, b = b, a
	}

	m, n := len(a), len(b)

	//	loop stop at (m+n) /2
	for count := 0; count < (m+n)/2+1; count++ {
		//	trick store (m + n)/2 - 1 value when m + n is even
		m2 = m1
		if idxA != m && idxB != n {
			if a[idxA] > b[idxB] {
				m1 = b[idxB]
				idxB++
			} else {
				m1 = a[idxA]
				idxA++
			}
		} else if idxA < m {
			m1 = a[idxA]
			idxA++
		} else {
			m1 = b[idxB]
			idxB++
		}
	}
	if (m+n)%2 == 1 {
		return float64(m1)
	}
	return float64(m1+m2) / 2
}

//	Binary Search
//		divide array and find the median with these steps:
//		- median is the point that will divide sorted merged array into two equal parts
//		=> actual median point at (M+ N + 1)/2 index (count from 0)
//		- given A[] and B[] array will divide to 2 parts: A[] always the array shorter
//		- set left pointer at first index element of A[], right pointer at last index element of A[]
//		find the mid point of A[] => A[] will divide in 2 parts at mid point: left and right
//		simultaneously same as B[] => the sum of count of elements in left both A[] & B[] is the left part of
//		the merged array. vice versa for the right part
//		- now initialize 4 vars indicating values of each part A[] & B[]:
//		. leftA:	Rightmost element in left part of A[]
//		. leftB:	Rightmost element in left part of B[]
//		. rightA:	Leftmost element in right part of A[]
//		. rightB:	Leftmost element in right part of A[]
//
//		- confirm partition is correct by checking main condition:
//			leftA <= rightB && leftB <= rightA
//		- if not => find another mid point of A[] & B[] in cases:
//		. leftA > rightB: decrease size of A's partition = update right pointer to midA - 1
//		. else (leftB > rightA): increase size of A's partition = update left pointer to midA + 1
//
//		- when main condition satisfied:
//		. if size A[] + B[] is odd	=> median = max(leftA, leftB)
//		. if size A[] + B[] is even	=> median = (max(leftA, leftB) + min(rightA, rightB)) / 2

//	Time Complexity: O(min(log M, log N)): Binary search applied on smaller of 2 arrays
//	Auxiliary Space: O(1)

// Solutions are many type but 2 type most seen:
//
// they use -infinity and +infinity when mid point A[] or B[] is out of index
// 1. define left point, right pointer as normal (recommend)
func findMedianSortedArraysBinaryNormal(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	total := m + n

	// +1 because we want the size include mid point. ie math.Cecil
	half := (total + 1) / 2
	l, r := 0, m-1
	for {
		midA := (l + r) >> 1

		// -2 cause length is last index + 1 (count from 0) => midA + 1 + midB + 1 = half
		midB := half - midA - 2

		var l1, l2, r1, r2 int

		if midA >= 0 {
			l1 = nums1[midA]
		} else {
			l1 = -math.MaxInt
		}
		if midB >= 0 {
			l2 = nums2[midB]
		} else {
			l2 = -math.MaxInt
		}
		if midA+1 < m {
			r1 = nums1[midA+1]
		} else {
			r1 = math.MaxInt
		}
		if midB+1 < n {
			r2 = nums2[midB+1]
		} else {
			r2 = math.MaxInt
		}

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 0 {
				return float64(max(l1, l2)+min(r1, r2)) / 2
			}
			return float64(max(l1, l2))
		}
		if l1 > r2 {
			r = midA - 1
		} else {
			l = midA + 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 2. define left pointer, right pointer as start(low), end(high)
func findMedianSortedArraysBinaryLowHigh(a []int, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	low, high, medianPos, total := 0, len(a), (m+n+1)/2, m+n

	for low <= high {
		cut1 := (low + high) / 2
		cut2 := medianPos - cut1
		l1 := ternary(cut1 == 0, math.MinInt, a, cut1-1)
		l2 := ternary(cut2 == 0, math.MinInt, b, cut2-1)
		r1 := ternary(cut1 == m, math.MaxInt, a, cut1)
		r2 := ternary(cut2 == n, math.MaxInt, b, cut2)
		if l1 <= r2 && l2 <= r1 {
			if total%2 == 1 {
				return maxX(l1, l2)
			} else {
				return (maxX(l1, l2) + minN(r1, r2)) / 2
			}
		} else if l1 > r2 {
			high = cut1 - 1
		} else {
			low = cut1 + 1
		}
	}
	// when the input invalid not match any cases
	return 0
}

func ternary(exp bool, a int, nums []int, idx int) int {
	if exp {
		return a
	}
	return nums[idx]
}

func minN(a, b int) float64 {
	if a < b {
		return float64(a)
	}
	return float64(b)
}

func maxX(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

// 3. my research (not recommend):
// - combine 1. & 2. condition
// - using trick -infinity +infinity to get value
// - still pass all test cases in leetcode
func findMedianSortedArraysResearch(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	if len(nums1) > len(nums2) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	total := m + n
	half := total / 2

	left, right := 0, m

	for left <= right {
		midPointA := (left + right) / 2
		midPointB := half - midPointA

		leftA := getValue(midPointA == 0, math.MinInt, a, midPointA-1)
		rightA := getValue(midPointA == m, math.MaxInt, a, midPointA)
		leftB := getValue(midPointB == 0, math.MinInt, b, midPointB-1)
		rightB := getValue(midPointB == n, math.MaxInt, b, midPointB)

		if leftA <= rightB && leftB <= rightA {
			if total%2 == 1 {
				return minQ(rightA, rightB)
			}

			return (maxQ(leftA, leftB) + minQ(rightA, rightB)) / 2
		} else if leftA > rightB {
			right = midPointA - 1
		} else {
			left = midPointA + 1
		}
	}
	return 0
}

func getValue(res bool, infVal int, nums []int, idx int) int {
	if res {
		return infVal
	}
	return nums[idx]
}

func maxQ(x, y int) float64 {
	if x > y {
		return float64(x)
	}
	return float64(y)
}

func minQ(x, y int) float64 {
	if x < y {
		return float64(x)
	}
	return float64(y)
}
