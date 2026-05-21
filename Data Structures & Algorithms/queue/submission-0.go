type Deque struct {

	arr []int

}


func NewDeque() *Deque {
   return &Deque{}
}

func (d *Deque) IsEmpty() bool {
 
  if len(d.arr)!=0{
	return false
  }

  return true

}

func (d *Deque) Append(value int) {

	d.arr = append(d.arr,value)

}

func (d *Deque) AppendLeft(value int) {

	d.arr = append([]int{value},d.arr...)

}

func (d *Deque) Pop() int {
     
	 if d.IsEmpty(){
		return -1
	 }
	last := d.arr[len(d.arr)-1]
	d.arr= d.arr[:len(d.arr)-1]

	return last

}

func (d *Deque) PopLeft() int {

	 if d.IsEmpty(){
		return -1
	 }
	first:=d.arr[0]

	d.arr = d.arr[1:]

	return first

}
