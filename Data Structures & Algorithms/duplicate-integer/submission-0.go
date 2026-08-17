func hasDuplicate(nums []int) bool {

	m:= make(map[int]struct{})

	for _,num := range nums {
		m[num]=struct{}{}
	}

	return len(m)<len(nums)
    
}
