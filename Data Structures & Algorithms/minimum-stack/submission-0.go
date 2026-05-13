type MinStack struct {
    s []int
    min []int
    }


func Constructor() MinStack {
    return MinStack{ []int{},[]int{}}
    }


func (this *MinStack) Push(val int)  {

    if len(this.s)==0{
        this.s = append(this.s,val)
        this.min = append(this.min,val)
    }else{
        this.s = append(this.s,val)
        this.min = append(this.min, min(val,this.min[len(this.min)-1]))
    }
    
}


func (this *MinStack) Pop()  {

    this.s = this.s[:len(this.s)-1]
    this.min = this.min[:len(this.min)-1]
    
}


func (this *MinStack) Top() int {
    
    return this.s[len(this.s)-1]
}


func (this *MinStack) GetMin() int {
    
    return this.min[len(this.min)-1]
}

func min(a,b int ) int {

    if a < b {
        return a
    }

    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */