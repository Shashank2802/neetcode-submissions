/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
	if list1==nil{
        return list2
    }
    
    if list2==nil {
        return list1
    }
   
    var head *ListNode
    
    if list1.Val > list2.Val{
        head=list2
        list2=list2.Next
    }else{
         head=list1
        list1=list1.Next
    }
    
    current:=head
    
    for list1!=nil && list2!=nil {
        
        if list1.Val>list2.Val {
            current.Next=list2
            current=current.Next
            list2=list2.Next
        }else{
            current.Next=list1
            current=current.Next
            list1=list1.Next
        }
    }
   
    if list1!=nil {
        current.Next=list1
    }
    
    if list2!=nil{
        current.Next=list2
    }
    
    return head
}
