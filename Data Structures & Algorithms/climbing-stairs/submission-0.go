func climbStairs(n int) int {
    
	secondLast,last := 1,1 

	for i:= 0 ; i<n-1 ; i++ {
		secondLast ,last = secondLast + last ,secondLast
	}

	return secondLast
	}
