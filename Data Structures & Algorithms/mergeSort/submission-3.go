// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
	return mergeHelper( pairs , 0 , len(pairs)-1)
}


func mergeHelper( arr []Pair , start , end int)[]Pair {

	if end - start +1 <=1{
		return arr
	}

	m := (start + end) /2 

	mergeHelper(arr , start , m)

	mergeHelper(arr , m+1,end)

	merge(arr, start , m , end)

	return arr 


}


func merge(arr []Pair , s , m ,e int){

	L := make([]Pair, m - s + 1)
	copy(L,arr[s:m+1])

	R:= make([]Pair , e - m)
	copy(R,arr[m+1:e+1])

	i,j,k:=0,0,s


	for i < len(L) && j < len(R) {

		if L[i].Key <= R[j].Key{
			arr[k]=L[i]
			i++
		}else{
			arr[k]=R[j]
			j++
		}

		k++
	}

	for  i < len(L) {

		arr[k]=L[i]
		i++
		k++
	}

	for j < len(R) {

		arr[k]= R[j]
		j++
		k++
	}

}
