package twoSum

// TwoSumWithLoop will use double loop all the elements in input array to find two numbers
// that can add up to target
// The firstNum start from the first element in array: nums[0]
// the secondNum start from the second element in array: nums[1]
// Time Complexity:		O(n^2): Use two loop O(n)
// Space Complexity:	O(1): space does not increase with input size
func TwoSumWithLoop(nums []int, target int) []int {
	for firstNum := 0; firstNum < len(nums); firstNum++ { // Time: O(n)
		for secondNum := firstNum + 1; secondNum < len(nums); secondNum++ { // Time: O(n)
			if target == (nums[firstNum] + nums[secondNum]) {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

// TwoSumWithHashMap will create an empty hashMap (key and value) first, Go format hash like map[type]type
// then use a loop through the input array to look for conditions:
// check required number = target - current number is present in hashmap
// If present, return {required number index, current number index} as result
// Otherwise add current number index and its value to hashmap
// Repeat until find out the result.
// Time Complexity:		O(n): Use one loop O(n)
// Space Complexity:	O(n): init a hashMap to store key and value
func TwoSumWithHashMap(nums []int, target int) []int {
	indexMap := make(map[int]int)           // Space: O(n)
	for currentIndex, value := range nums { // Time: O(n)
		if requiredIndex, ok := indexMap[target-value]; ok {
			return []int{requiredIndex, currentIndex}
		}
		indexMap[value] = currentIndex
	}
	return []int{}
}
