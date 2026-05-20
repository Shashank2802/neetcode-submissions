type Node struct {
    Val  int
    Next *Node
}

type MyLinkedList struct {
    head *Node
    tail *Node
    size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    temp := this.head
    for i := 0; i < index; i++ {
        temp = temp.Next
    }
    return temp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    newNode := &Node{Val: val}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
    } else {
        newNode.Next = this.head
        this.head = newNode
    }
    this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
    newNode := &Node{Val: val}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
    } else {
        this.tail.Next = newNode
        this.tail = newNode
    }
    this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
    } else if index == this.size {
        this.AddAtTail(val)
    } else {
        newNode := &Node{Val: val}
        temp := this.head
        for i := 0; i < index-1; i++ {
            temp = temp.Next
        }
        newNode.Next = temp.Next
        temp.Next = newNode
        this.size++
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    if index == 0 {
        this.head = this.head.Next
        if this.size == 1 {
            this.tail = nil
        }
    } else {
        temp := this.head
        for i := 0; i < index-1; i++ {
            temp = temp.Next
        }
        temp.Next = temp.Next.Next
        if index == this.size-1 {
            this.tail = temp
        }
    }
    this.size--
}