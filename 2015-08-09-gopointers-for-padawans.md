# Go Pointers for Padawans

Hi my friend, 

like you I'm also trying to learn more about Go language. I've always wanted to know more about it, but as usual I always had an excuse on the tip of my tongue to postpone this  desire. Now I'm done. I found a good introduction book, An Introduction to Programming in Go, and I can read Effective Go online.

I started reading about Go last week because a lot of friends are using it and they are encouraging me (pissing me off =) to do the same.

I'm a complete newbie in this technology, so if you are an advanced jedi hacker this is not for you, go play RPG with your imaginary ORCs. However, if you are just a padawan like me, try to not sleep while reading it.

If you see something wrong just leave me a comment, and I will fix it as soon as possible. Just to be clear "You suck Hobbit" is not a good comment.

## Pointers: What are they?

When I used to develop software, pointers always scared a lot of programmers, mainly java programmers (calm down, just a joke =). It is really common to hear developers linking this subject  to C language.

Although a lot of languages have pointers, in some languages they are more explicit or clear than others. I agree that managing pointers/memory sometimes can be a pain in the ass, but I think this is not the case of Go.

If the concept of pointers is new to you, no worries, let's learn together. In a few words a pointer is: A variable that stores the address of other variable. Come here my little boy, just like an integer variable stores a number/integer, in the same way a pointer stores an address.

 - **Jon Snow**: Hey, I have a question. What is an Address?
 - **Me**: Are you kidding me Jon Snow?
 - **Jon Snow**: Not really!
 - **Me**: You know nothing, Jon Snow.

Ok, when you declare a variable in your favorite language, in fact you are just saying to the Operating System: "Hey my dear OS, I have here this variable called *age* and I need a place, in the memory, to store it.". In other words, a memory address is an unique identifier that points to a place in the memory to be used to store your data. Look this simple example:

```go
package main

func main() {
    var age int
    age = 10

    var max int
    max = 3

}

             MEMORY        ADDRESSES     VARIABLES  
        +--------------+
        |      10      |  0x20818a220      age
        +--------------+
        |              |     ...
        +--------------+
        |      3       |  0x20818a228      max
        +--------------+
```

  - **Me**: Easy peasy lemon squeezy Jon Snow?!? Stop thinking about Ygritte and pay attention on the pointers. 

## Declaring pointers

This is de moment we will become men and start working with pointers. Stop crying :D! Ok, let's declare some of them.

```go
package main

func main() {
    // Pointer To Int
    var pi *int
    // Pointer To Bool
    var pb *bool
    // Pointer to Float
    var pf *float64
}
```

In GoLang a pointer is represented by an asterisk (*). In the example above we are just declaring a pointer to primitive types, such as int, bool and float. We also can create a type and then declare a pointer to our type, I will show how to do that. Hey Padawans, do you have questions or comments about declaring pointers?

 - **Sheldon Cooper**: That is it? I have to lose 100 points of my IQ to find it difficult.
 - **Penny**: I have to agree with Sheldon, this is not difficult at all. I learnt it in the Community College.
 - **Sheldon Cooper**: haaaa, Penny, Penny, Penny.
 - **Penny**: What? What? What?
 - **Sheldon Cooper**: You probably learnt Java at Community College.
 - **Me**: Calm down my friends, Java is not the way.
 - **Howard Wolowitz**: I've a master in engineering from MIT, Go Pointers is easy.

## Operators: * and &

Like I said before, the asterisk character (*) is used to declare a pointer. What you don't know is, it also can be used to dereference.

 - **Leonard**: “Whaaaat?”

It is quite simple, Dereference is when you want to access the value stored in the address. Maybe you will understand better when we show an example.

The second operator is "&" and it is used to get the variable address. All variables have an address, remember? So, when a pointer get/save a variable address it is actually pointing to the variable.

 - **Lagertha**: Finished? This is so complicated!
 - **Ragnar Lothbrok**: This is not the end, it's just the beginning ...
 - **Floki**: hihihi, Ragnar the gods will help us.
 - **Me**: You are Vickings and never give up.

Ok, less talk and more code.

```go
package main

import "fmt"

func main() {
    x := 10     // Declaring a int variable and assigning value 10
    var p *int  // Declaring a pointer to int

    p = &x      // Assigning to variable p the address of x

    fmt.Println(p)   // Printing the content of p
    fmt.Println(&x)  // Printing the address of x
    fmt.Println(*p)  // Dereference. Accessing the value 10
    fmt.Println(x)   // Printing the content of x

    *p = 5           // Assigning value 5 to variable x
    // p = 5         // Error: p is not a int, it is a pointer to int
    fmt.Println(*p)  // Dereference. Will print 5
    fmt.Println(x)   // Printing x variable, value = 5
}
            Variables                    Memory Adrresses
                      +--------------+
              +---> x |      10      |    0x20818a220
              |       +--------------+
              |       |              |
              |       +--------------+
              +---- p | 0x20818a220  |    0x20818a228
                      +--------------+ 
```

Let's understand the code above. 

As you know, all variables have a memory address. In the example the x address is ‘0x20818a220’ and the p address is ‘0x20818a228’. The variable x is an Integer and stores the value 10. The p variable is a pointer to integer and stores the value ‘0x20818a220’. When we do this instruction “*p = 5” we are using dereference to save the value 5 in place of 10.

I don't know if I was clear enough. I hope so /o\. Let's do one more example, now using a struct.

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    name string
    age int
}

func main() {
    var p *Person              // Declaring a pointer to struct Person
    fmt.Println(p)             // Printing nil

    p = &Person{"Titao", 18}   // Creating a new object Person and assigning to p
    fmt.Println(p)             // Printing p

    fmt.Println("reflect = ", reflect.TypeOf(p))
}
```

We are using the package reflect just to verify the type of variable p.

 - **Athelstan**: Why when the print p in the first time its nil?
 - **Me**: I thought you were dead priest. Good question.

If you declare variables and not initialize them, golang will do it for you. Your variable will be assigned with zero value for its type. For example if your variable is an integer it will be assigned with 0, for booleans false, floats 0.0, strings “” and pointers nil.

 - **Me**: Clear now Athelstan? 
 - **Floki**: I can help to open his mind, hihihi

## Function parameters

I think in the golang website they explain really well about passing parameters by value.

"As in all languages in the C family, everything in Go is passed by value. That is, a function always gets a copy of the thing being passed, as if there were an assignment statement assigning the value to the parameter. For instance, passing an int value to a function makes a copy of the int, and passing a pointer value makes a copy of the pointer, but not the data it points to." (golang.org)

To understand better, let's create a function that receives two int parameters and exchange their values. The name of this function will be swap() and it is a really common example to explain passing parameters by value.

```go
package main

import "fmt"

// Passing parameters by value
func swap(x, y int) {
    var z int
    z = x
    x = y
    y = z
}
func main() {
    x := 2
    y := 3

    // It will not exchange the values
    swap(x, y)
    fmt.Println("X = ", x)
    fmt.Println("Y = ", y)
}
```

Now let’s do the same function using pointers to int.

```go
package main

import "fmt"

// Using pointers
func swap(x, y *int) {
    var z int
    z = *x
    *x = *y
    *y = z
}

func main() {
    x := 2
    y := 3

    // It will exchange the values
    swap(&x, &y)
    fmt.Println("X = ", x)
    fmt.Println("Y = ", y)
}
```

I understand it can be a little bit confusing, but I found a comment from Dylan Beattie in the stack overflow that explains really well this .

“Say I want to share a web page with you.
If I tell you the URL, I'm passing by reference. You can use that URL to see the same web pageI can see. If that page is changed, we both see the changes. If you delete the URL, all you're doing is destroying your reference to that page - you're not deleting the actual page itself.
If I print out the page and give you the printout, I'm passing by value. Your page is a disconnected copy of the original. You won't see any subsequent changes, and any changes that you make (e.g. scribbling on your printout) will not show up on the original page. If you destroy the printout, you have actually destroyed your copy of the object - but the original web page remains intact.”
Ok, how is it going until now? Everybody happy? More questions?

 - **Charlie Harper: You are a nerd hobbit and I’m done with this bullshit.
 - **Me**: Where are you going?
 - **Charlie Harper**: Someplace where the bottles are full and the women are empty. Who comes? Penny? My treat.
 - **Howard Wolowitz**: I’m in.
 - **Penny**: Yeah, why not? Long time I don’t drink.
 - **Me**: Guys, calm down. Let me finish and all of us will go to the pub.

## Function New

This is the last part about pointers.

 - **Penny**: Thanks God, I need wine.
 - **Me**: Is not  that bad. Go is amazing language.
 - **Leonard**: Yeah, I love Go.
 - **Sheldon Cooper**: Let’s change Hailo night to Go night.
 - **Penny**: Or maybe we can just enjoy our life and go out talk to people.
 - **Me**: Guys, more 10 minutes. Let’s talk about function new().

**New()** is a buit-in function that receives a type as argument, allocate memory enough to that type and returns a pointer to the type. In the example bellow we have a struct Person with two differente attributes: name and age. We will use the function new to allocate memory to the type Person.

The amazing part is we don’t have to realease the memory because Go has a gargabe collector:

 - **Rajesh Koothrappali**: No way!
 - **Ragnar Lothbrok**: no more double-free and/or memory leaks?
 - **Sheldon Cooper**: Peeny can do memory leaks in all languages.
 - **Penny**: The memory is mine, I can do everything I want.
 - **Me**: This is the spirit Penny :D.

Let’s see an example using function new.

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    name string
    age int
}

func main() {
    p := new(Person)

    p.name = "Titao"
    fmt.Println(p.name)

    fmt.Println("reflect = ", reflect.TypeOf(p))
}
```

 - **Penny**: When you explained about operators * and & we used the operator & to allocate a struct Person. What is the difference bettween &Type{} and new()?
 - **Me**: Owwww, Penny is on fire, good question.
 - **Penny**: Thanks.
 - **Howard Wolowitz**: Penny, I can explain to you saturday night, drinking some wine. I can show my pointer to you.
 - **Penny**: disgusting Howard.

There is no difference. Both will return the same, pointer to the type. Buuutttt you can use &Type{} only for struct, map, array or a slice. The new function you can use for all types.

That is it guys. You are free to drink now.

 - **Charlie Harper**: uhuuuuuu! Party in my house.

## References

Like I said, I’m learning this new amazing language and I tried to share a little bit what I read in books and Internet. I hope you enjoyed reading it and let me know my mistakes.

 - Effective Go
 - An Introduction to Programming in Go
 - Google
 - StackOverFlow
 - Nerd friends :D
