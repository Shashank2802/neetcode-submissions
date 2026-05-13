func calPoints(operations []string) int {

	stack:=make([]int,0)

	for _,val:=range operations {

		if val != "+" && val != "D" && val !="C"{
		  
		  score,_:= strconv.Atoi(val)
		  stack = append(stack, score)
		  continue
		}

		if val == "+"{
			 a:= stack[len(stack)-1]
			 b:= stack[len(stack)-2]
			 c:= a+b
			 stack = append(stack,c)
		}

		if val == "D" {
			a:= stack[len(stack)-1]
			stack = append(stack,2*a)
		}

		if val == "C"{
			stack = stack[:len(stack)-1]
		}
	}

	score:=0 

	for _,val:= range stack{
		score+=val
	}
	return score
	
	}
	


