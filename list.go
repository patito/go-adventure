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

func (l *LinkedList) Append(item int) {
    node := new(Node)
    node.item = item

    if l.head == nil {
        l.head = node
    } else {
        var previous *Node = nil;
        for current := l.head; current != nil; {
            previous = current
            current = current.next
        }
        previous.next = node

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
        fmt.Println("Item: ", current.item)
    }
}

func (l *LinkedList) Search(item int) (*Node, int) {
    pos := 0
    for current := l.head; current != nil; current = current.next {
        if current.item == item {
            return  current, pos
        }
        pos += 1
    }
    return nil, -1
}


func main() {

    l := new(LinkedList)
    l.Preppend(2)
    l.Preppend(5)
    l.Append(1)
    l.Append(4)
    l.Print()
}