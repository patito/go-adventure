package main

import (
    "fmt"
)

type Element struct {
    next *Element
    value interface{}
}

type List struct {
    size int
    head *Element
}

// Create new list
func New() *List {
    return &List{0, nil}
}

// Returns the size of list
func (l *List) Length() int {
    return l.size
}

// Insert the new element in the front of list
func (l *List) Preppend(value interface{}) {
    node := new(Element)
    node.value = value

    node.next = l.head
    l.head = node
    l.size++
}

// Insert the new element in the back of list
func (l *List) Append(value interface{}) {
    node := new(Element)
    node.value = value

    if l.head == nil {
        l.head = node
    } else {
        var previous *Element = nil;
        for current := l.head; current != nil; {
            previous = current
            current = current.next
        }
        previous.next = node

    }
    l.size++
}

// Returns the first element of list
func (l *List) First() *Element {
    if l.Empty() {
        return nil
    }
    return l.head
}

// Returns the last element of list
func (l *List) Last() *Element {
    if l.Empty() {
        return nil
    }
    return l.head
}

// Verify if the list is empty
func (l *List) Empty() bool {
    return l.size == 0
}

// Print the list
func (l *List) Print() {
    for current := l.head; current != nil; current = current.next {
        fmt.Println("Item: ", current.value)
    }
}

// Search for elements
func (l *List) Search(value interface{}) (*Element) {
    for current := l.head; current != nil; current = current.next {
        if current.value == value {
            return  current
        }
    }
    return nil
}

// Remove an element
func (l *List) Remove(value interface{}) *Element {
    var p *Element = nil
    current := l.head

    if l.head.value == value {
        l.head = current.next

        return current
    }

    for current != nil {
       if current.value == value {
           p.next = current.next
           current.next = nil

           return current
       }
       p = current
       current = current.next
    }
    return nil
}


func main() {

    l := New()

    l.Append("titao1")
    l.Append("grazzilinda")
    l.Append("nani3")
    l.Append("cadide4")


    //fmt.Println(l.Search("cadide4").value)
    l.Remove("titao1")
    l.Remove("cadide4")
    //fmt.Println(l.Search("cadide4").value)
    l.Print()

}
