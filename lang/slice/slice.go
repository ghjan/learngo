package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6]=", s)
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1, arr[2:]", s1)
	s2 := arr[:]
	fmt.Println("s2， arr[:]=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)

	fmt.Println("Reslice")
	s2 = reslice(s2, arr)

	fmt.Println("Extending slice")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	extendSlice(arr, s1, s2)

	var s0 []int //Zero value for slice is nil
	fmt.Println("appendSlice")
	s0 = appendSlice(s0)

	printSlice(s0)

	s1 = []int{2, 4, 6, 8}
	printSlice(s1)
	s2 = make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	//删除中间元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0:]
	s2 = front[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("leecode twoSum()")
	fmt.Println(twoSum(nums, target))
	target = 22
	fmt.Println(twoSum(nums, target))
	target = 23
	fmt.Println(twoSum(nums, target))
	nums = []int{-1, -2, -3, -4, -5}
	target = -8
	fmt.Println(twoSum(nums, target))

}

func updateSlice(s []int) {
	s[0] = 100
}

func reslice(s2 []int, arr [8]int) []int {
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	s2 = arr[:]
	fmt.Println(s2)
	return s2
}

//slice可以向后扩展 只要没有超出capicity，就会访问到底层原先的值
func extendSlice(arr [8]int, s1 []int, s2 []int) {
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	//添加元素时候，如果超过原先分配的capacity，就会重新分配一个更大capacity的底层数组
	// s4 and s5 no longer view arr yet
	//由于按值传递， 必须用一个变量来接受append的返回值
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	fmt.Println("arr=", arr)
}

func appendSlice(s []int) []int {
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	return s
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

//https://leetcode.com/problems/two-sum/description/

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.

//You may assume that each input would have exactly one solution, and you may not use the same element twice.
func twoSum(nums []int, target int) []int {
	firstI, secondI := -1, -1
outer:
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			if num1+num2 == target {
				firstI = i
				secondI = j + i + 1
				break outer
			}
		}
	}
	return []int{firstI, secondI}
}
