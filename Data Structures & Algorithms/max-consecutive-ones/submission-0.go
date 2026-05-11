func findMaxConsecutiveOnes(nums []int) int {

	    maxCount,currCount:=0,0

     for _,v:= range nums {

        if v == 1 {
            currCount++
        }else {
            if maxCount < currCount {
                maxCount = currCount
            }
            currCount = 0 
        }
     }

     if maxCount < currCount {
        maxCount = currCount
     }

     return maxCount
}
