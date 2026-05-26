// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

// Implementation of Merge Sort
func mergeSort(pairs []Pair) []Pair {
    return mergeSortHelper(pairs, 0, len(pairs)-1)
}

func mergeSortHelper(pairs []Pair, s_idx, e int) []Pair {
    if e-s_idx+1 <= 1 {
        return pairs
    }
    // The middle index of the array
    m := (s_idx + e) / 2
    // Sort the left half
    mergeSortHelper(pairs, s_idx, m)
    // Sort the right half
    mergeSortHelper(pairs, m+1, e)
    // Merge sorted halfs
    merge(pairs, s_idx, m, e)
    return pairs
}

// Merge in-place
func merge(arr []Pair, s_idx, m, e int) {
    // Copy the sorted left & right halfs to temp arrays
    L := make([]Pair, m-s_idx+1)
    copy(L, arr[s_idx:m+1])
    R := make([]Pair, e-m)
    copy(R, arr[m+1:e+1])
   
    i := 0 // index for L
    j := 0 // index for R
    k := s_idx // index for arr
   
    // Merge the two sorted halfs into the original array
    for i < len(L) && j < len(R) {
        if L[i].Key <= R[j].Key {
            arr[k] = L[i]
            i++
        } else {
            arr[k] = R[j]
            j++
        }
        k++
    }
   
    // One of the halfs will have elements remaining
    for i < len(L) {
        arr[k] = L[i]
        i++
        k++
    }
    for j < len(R) {
        arr[k] = R[j]
        j++
        k++
    }
}
