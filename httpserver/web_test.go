package main

import "testing"

func TestQuickSort(t *testing.T) {
	arr1 := []int {3,4,5,6,1}
	arr2 := []int {1,3,4,5,6}
	QuickSort(arr1,0,len(arr1)-1)

	flag := true
	for i:=0; i<len(arr1); i++ {
		if arr1[i] != arr2[i] {
			flag = false
		}
	}

	if flag == false {
		t.Error("QuickSort arr is false!")
	}
}
