package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immoc",
		"quality": "notbad",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)

	fmt.Println("Tranversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}
	if courseName, ok := m["couse"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	//map的遍历
	//使用range遍历key，或者遍历key,value对
	//不保证遍历的次序，需要手动对key排序（slice类型）

	//map的key
	//map使用哈希表，必须可以比较相等
	//除了slice map function的内建类型都可以作为key
	//struct类型不包含上述字段，也可以作为key

	fmt.Println(getLongestSubstringWithoutRepeating("abcabcbb"))
	fmt.Println(getLongestSubstringWithoutRepeating("bbbbb"))
	fmt.Println(getLongestSubstringWithoutRepeating("pwwkew"))
	fmt.Println(getLongestSubstringWithoutRepeating(""))
	fmt.Println(getLongestSubstringWithoutRepeating("b"))
	fmt.Println(getLongestSubstringWithoutRepeating("abcedfg"))
	fmt.Println(getLongestSubstringWithoutRepeating("这里是慕课网"))
}

//寻找最长不含有重复字符的子串
//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func getLongestSubstringWithoutRepeating(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}
