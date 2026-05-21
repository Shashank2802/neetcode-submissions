type ListNode struct {

	val string
	next *ListNode
	back *ListNode
}


type BrowserHistory struct {
    
	cur *ListNode
}


func Constructor(homepage string) BrowserHistory {

	return BrowserHistory{
		cur : &ListNode{ val:homepage},
	}
    
}


func (this *BrowserHistory) Visit(url string)  {

	this.cur.next = &ListNode{val:url,back :this.cur}
	this.cur = this.cur.next
	
    
}


func (this *BrowserHistory) Back(steps int) string {

	for this.cur.back != nil && steps>0 {
		this.cur = this.cur.back
		steps--
	}

	return this.cur.val
    
}


func (this *BrowserHistory) Forward(steps int) string {
    for this.cur.next != nil && steps>0 {
		this.cur = this.cur.next
		steps--
	}

	return this.cur.val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */