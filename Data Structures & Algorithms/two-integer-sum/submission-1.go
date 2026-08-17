func isAnagram(s string, t string) bool {
if len(s) != len(t) {
        return false
    }
	
	var arr[26] int

	for _, char := range s {
		arr[char-'a']++
	}

	for _, char := range t {

		if arr[char -'a'] <=0 {
			return false
		}else{
           arr[char-'a']--
		}
	}

	return true

}
