func replaceElements(arr []int) []int { 

	ans:= make([]int, len(arr))
    rightMax:=-1 

	for i:= len(arr)-1;i>=0;i-- {
        
		ans[i]=rightMax
		if rightMax < arr[i]{
			rightMax = arr[i]
		}
	}


	return ans

}
