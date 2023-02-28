# **GOLANG**

https://github.com/StephenGrider/GoCasts

#### **Environtment Setup**

---

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-001%20-%20install.png?raw=true)

to install GO , go to https://go.dev/dl/

to check if go is already installed , open terminal and write `go` , then a long help message will be appear on the screen.

and install GO extenction in VS code , just search Go and install.

after install , close directory , and open again , and make file . so in the bottom right there is `plaintext` change that to `GO` , after that in bottom right we will see yellow text something that says `analysis tools missing` , and click that , and install.

---

in `main.go` we have several questions :

1.  How do we run the code in our project?
2.  What does `package main` mean?
3.  What does `import "fmt"` mean ?
4.  What that `func` thing?
5.  How is the `main.go` file organized ?

Answer:

1.  `go run main.go`

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-003%20-%20go%20cli.png?raw=true)

different `go build` and `go run` is , go build just compile the program not execute.

if we run `go build` than we will get `main.exe` so this is an actual executable file that was built out of our source code of main go.

we can run it by `./main` or `./main.exe`

2.  A package is a collection of common source code files.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-008%20-%20what%20package.png?raw=true)

package can have many related file instide of it ,each file ending with a file extension of GO (.go). the only requirement for every file inside of a package is that the very first line of each file must declare the package that it belongs to.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-014%20-%20packages.png?raw=true)

inside of go , there are two different types of packages.

there is an executable type and a reusable type.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-012%20-%20types%20of%20packages.png?raw=true)

so how do we know if we are making an executable package or reusable?

the word `main` is ued to make executable type package.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-013%20-%20compiling.png?raw=true)

if we used any other name for our package besides Main , it would not spit out an actual excecutable file.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-013.5%20-%20reusable.png?raw=true)

so if we were trying to make some library of reusable code, or if we wanted share with our friend so they can use our code , than we would start using a more specialized package name.

so for summary :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-008%20-%20packages.png?raw=true)

3. `Import FMT `specially means give my package main access to all of the code and all the functonality that is contained inside of this other package called `fmt`.

FMT is the name of standard library package that is included with the go programming language by default.

FMT itself is a kind of a shortened form of the word format.

The FMT library is used to print out a lot af different information specially to the terminal , just to give you a better sense of debugging and stuff like that.

so for example main package could import FMT, but it could also immport in laike calculator or uploader pacakge.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-010%20-%20many%20imports.png?raw=true)

we can check library in https://pkg.go.dev/

4. Func is short for function inside of go .

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-015%20-%20func.png?raw=true)

5.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-007%20-%20code.png?raw=true)

---

we will make some project.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-000%20-%20cards.png?raw=true)

we make the project in folder `CARDS`

deal is get the card from deck.

> note: if we just write `fmt.Println` than vs code automate write `import "fmt"`
---
how to assignment :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-003%20-%20assignment.png?raw=true)

or change that to `colon equal (:=)`

Go is a statically typed languange .

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-005%20-%20types.png?raw=true)

A dynamically typed languange is one in which you and i , the developers , essentially do not care what values we are assigning to any given varible.

---
basic go types :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-004%20-%20types.png?raw=true)

we just use colon equal (:=) when we are defining a new variable , if we are reassigning a existing variable we just use equal (=)

varibale can initialize outside function , but cannot assign the value outside function.

---
for the function :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-007%20-%20func.png?raw=true)

files in the same package do not have be imported into each other.example imagine we have 2 file :

main.go

```Golang
    package main

    func main() {
        printState()
    }

```

and
state.go

```Golang
package main

import "fmt"

func printState() {
    fmt.Println("California")
}

```

so if we run `go run main.go state.go` then all code will successfully compile because files in the same package do not have to be imported into each other.

---

GO has two basic data structures for handling lists of records :

1. Array -> fixed length list of things , cant add new value.
2. Slice -> An array that can grow or shrink.

whenever we have an array or a slice , every single record inside of it must be of an identical type.

range :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-006%20-%20range.png?raw=true)

---
Object Orientation 

Go is not an object oriented programming language.

if we use dinamic language like python java etc , OO will be right below :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-002%20-%20oo%20approach.png?raw=true)

if we use golong , OO is like below:

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-003.1%20-%20go%20approach.png?raw=true)

we will makee deck type inside our application ,essentially the same thing as a slice of string , because we are making this extra type , it gives us the ability to create a bunc of custom functions that only work with that deck type.

we will make deck type in deck.go

and run 
`go run deck.go main.go`

---

## Receiver

right before the function name of print , is referred to as a receiver.
so any varibale of type "deck" now gets access to the print method.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-009.3%20-%20method%20receiver.png?raw=true)

Make new function `New Deck`





