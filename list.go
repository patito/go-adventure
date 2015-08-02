package main

import "fmt"

type Node struct {
    item int
    next *Node
}

type LinkedList struct {
    head *Node
    size int
}

func (l *LinkedList) Preppend(item int) {
    node := new(Node)
    node.item = item

    if l.head == nil {
        l.head = node
    } else {
        node.next = l.head
        l.head = node
    }
    l.size++
}

func (l *LinkedList) Size() int {
    return l.size
}

func (l *LinkedList) Empty() bool {
    return l.size == 0
}

func (l *LinkedList) Print() {
    
    for current := l.head; current != nil; current = current.next {
        fmt.Println(current.item)
    }
}

func (l *LinkedList) Search(item int) *Node {
    
    for current := l.head; current != nil; current = current.next {
        if current.item == item {
            return current
        }
    }
    return nil
}

func main() {

    l := new(LinkedList)
    l.Preppend(2)
    l.Preppend(3)
    l.Preppend(4)
    l.Preppend(5)
    fmt.Println("Size list = ", l.Size())
    fmt.Println("Empty = ", l.Empty())
    fmt.Println("Node = ", l.Search(1))
    l.Print()
}