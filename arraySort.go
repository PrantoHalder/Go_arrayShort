package main

import "fmt"

func rearrange () {

	fmt.Println("rearranging")
	var arr = [10] int {43,12,6,3,8,9,1,4,32,8}
	n :=len(arr)
	var i int 
	var j int
	for i = 0; i < n ; i++ {
		for j= i+1 ; j < n ; j++ {
		if arr[i] > arr[j] {
		   arr[j],arr[i] = arr[i] ,arr[j]
		}
    }
}
	
	for i = 0 ; i< n ;i++ {
		fmt.Println(arr[i])
	}
}