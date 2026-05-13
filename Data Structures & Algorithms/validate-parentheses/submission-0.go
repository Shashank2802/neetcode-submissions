func isValid(s string) bool {

	if len(s)==0{
		return true
	}else if len(s)==1{
		return false
	}
	
	m:= map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}
	var stack []rune

	for _,val:= range s{

		if val == '(' || val == '{' || val == '['{
			stack = append(stack,val)
		}else{
            if len(stack) == 0 {
                return false
            }
			element:= stack[len(stack)-1]
			
			if m[val]==element{
				stack = stack[:len(stack)-1]
			}else {
				return false
			}

	        }
           
	}
	return len(stack) == 0
}

