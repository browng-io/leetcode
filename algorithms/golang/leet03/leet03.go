package leet03

import "strings"

// Longest Substring Without Repeating Characters

// Brute force
// Consider all substrings one by one and check for each substring
// whether it contains all unique character or not

// Will be n*(n+1)/2 substrings
//  1. make 2 pointers (all start from left to right):
//     - index: position each character in string
//     - sub: position each character from index to the end of string
//  2. make a map storing all ascii code: 128 unique chars (no extended ascii)
//     with default value: 0
//     - when character is visited: map[char] += 1
//  3. check conditions:
//     - if character is visited more than 1, skip that substring
//     - else consider its length and maximize it
//     Time Complexity: O(n^3): processing n^2 substrings with maximum length n
//     Auxiliary Space: O(1)
func lengthOfLongestSubstringBruteForce(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 0

	for index := 0; index < len(s); index++ { // O(n)
		for sub := index; sub < len(s); sub++ { // O(n)
			if isDistinct(s, index, sub) {
				length := sub - index + 1
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}
	return maxLength
}

func isDistinct(s string, start, end int) bool {
	// Default values in visited are false
	visited := make([]int, 128)
	for i := start; i < end+1; i++ { // O(n): worst case
		char := s[i]
		visited[char] += 1
		if visited[char] > 1 {
			return false
		}
	}
	return true
}

// Sliding Window Technique
// When see repetition, remove the previous occurrence and slide the window
// Time Complexity: O(n^2)
// Auxiliary Space: O(1)
func lengthOfLongestSubstringSlideWindow(s string) int {
	maxLength := 0

	for i := 0; i < len(s); i++ { // O(n)
		visited := make([]int, 128)

		for j := i; j < len(s); j++ { // O(n)
			if visited[s[j]] == 1 {
				break
			} else {
				length := j - i + 1
				if length > maxLength {
					maxLength = length
				}
				visited[s[j]] = 1
			}
		}
		visited[s[i]] = 0
	}
	return maxLength
}

// Linear Time: Sliding Window
// Whenever see repetition, remove the window till the repeated string
// Time Complexity: O(n)
// Auxiliary Space: O(1)
func lengthOfLongestSubstringLinearTime(s string) int {
	test := ""
	maxLength := -1

	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	for _, val := range s {
		current := "" + string(val)
		if strings.Contains(test, current) {
			test = test[strings.Index(test, current)+1:]
		}
		test = test + string(val)
		if len(test) > maxLength {
			maxLength = len(test)
		}
	}
	return maxLength
}

// Sliding Window: Optimize (recommend)
// uses extra space to store last indexes already visited characters (hashMap)
// make two pointer
//   - rightIndex: iterates through all char in string
//   - leftIndex: whenever visited char found in hashMap:
//     . update leftIndex to current visited char's index stored in hashMap
//     if leftIndex < current visited char's index
//     calculate maxLength = rightIndex - leftIndex + 1
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func lengthOfLongestSubstringOptimizeSlidingWindow(s string) int {
	if len(s) == 0 {
		return 0
	}
	hashMap := make(map[byte]int)
	maxLength, left := 0, 0

	for right, char := range s {
		if index, found := hashMap[byte(char)]; found && index >= left {
			left = index + 1
		}
		hashMap[byte(char)] = right
		length := right - left + 1
		if length > maxLength {
			maxLength = length
		}

	}
	return maxLength
}

// Linear Time: KMP Technique (all test cases from leetcode still result time & space bad)
// Maintain an Unordered Set to keep track maximum non repeating char substring (instead standard LPS)
// When find a repeating char --> clear the Set and reset len to zero
// Time Complexity: O(n)
// Space Complexity: O(m): m - length of resultant substring
func lengthOfLongestSubstringKMP(s string) int {
	if len(s) == 0 {
		return 0
	}

	st := map[byte]bool{s[0]: true}
	i := 1

	length := 1
	maxLength := 0

	for i < len(s) {
		// check if consecutive chars are distinct and non repeating
		_, found := st[s[i]]
		if s[i] != s[i-1] && !found {
			st[s[i]] = true
			length += 1
			i += 1

			// back up the maxLength
			if length > maxLength {
				maxLength = length
			}
		} else {
			// move forward for repeating chars
			if length == 1 {
				i += 1
			} else {
				// reset the substring and set the pivot for next substring
				st = map[byte]bool{}
				i = i - length + 1
				length = 0
			}
		}
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
