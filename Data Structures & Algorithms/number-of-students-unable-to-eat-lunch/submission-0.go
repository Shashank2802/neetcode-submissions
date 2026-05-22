func countStudents(students []int, sandwiches []int) int {
 
  res:= len(students)
  count := make([]int,2)


   for _, v := range students {

	  count[v]++
   }

   for _ , v := range sandwiches{

	  if count[v] > 0{
		 
		 count[v]--
		 res --
	  }else{
		break
	  }
	  
   }

   return res



	
    
}