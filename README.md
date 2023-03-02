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

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-009.2%20-%20more%20receiver.png?raw=true)


Make new function `New Deck`

---

## make newDeck() Function
updated in `deck.go`

whenever yo have some variable that you dont actualyy have to use , we always replace it with an underscore (_).

and run 
```
go run main.go deck.go
```

---

## **Slice Range Syntax**

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-014%20-%20slice%20syntax.png?raw=true)


![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-015%20-%20slice%20selection.png?raw=true)

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-013%20-%20slice%20syntax.png?raw=true)

make deal() function in `deck.go`

array deck will not change , because slice will make separate array.

---


### Make saveToFile() Function

we use package `ioutil` , package `ioutil` implements some I/O utility functions.

```
func WriteFile(filename string, data []byte, perm fs.FileMode) error
```

byteslice :

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-018%20-%20byte%20slice.png?raw=true)

byte slice is , slices like an array ehere every element inside of it corresponds to an ASCII character code.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-019%20-%20type%20casting.png?raw=true)

```
deck -> []string -> string -> []byte
```

if we use writeTOfile() , than new file will be creates same root in main.go , same name with filename that parsing to filename in function writeToFile()

---

### **Reading from Harddrive**

```
func ReadFile(filename string) ([]byte, error)
```

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-018%20-%20error%20handling.png?raw=true)

`nil` is a value in go , which essentially means no value , or this thing has no value .

`package os` in this package there are some functions that are going to work equally well across all those different operating systems.

to Exit the program we use :

```
func Exit(code int)
```

>Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
For portability, the status code should be in the range [0, 125].

because we want ouput as deck type , but the output of `readFile` is byte slice ([]byte) so 

```
[]byte -> string ->[]string -> deck 
```

- to make `[]byte` to `string` we use string([]byte type )
- to make `string` to `[]string` we use :

    ```
    func Split(s string, sep string) []string
    ```

    >Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

- to make `[]string` to deck we use `deck(s)`

--- 

### **Make shuffle function**

we use `math/rand package` to randomize , and wwe use :
```
func Intn(n int) int
```
>Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <= 0.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-021%20-%20psuedorandom.png?raw=true)



if we use `rand` package is used seed value , and the seed value is always same , so the random number is same to , so we must change the `seed` value in random generator we use `type Rand` in package rand , so we use :

```
func New(src Source) *Rand
```
>New returns a new Rand that uses random values from src to generate other random values. 

and to make new type source we use :

```
func NewSource(seed int64) Source
```

because `seed ` must updat every time , so we can use package `time`


---

### **Testing with GO**

- Go testing is not RSpec, mocha,jasmine,selenium,etc! , we just use very small interface or very small set of functions to help us actually test our code.

- To make a test , create a new file ending in _test.go
    example : `desctest.go`

- To run all tests in a package, run the command :
    ```
    go test
    ```
![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-023%20-%20tests.png?raw=true)

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-024%20-%20tests.png?raw=true)

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-025%20-%20more%20tests.png?raw=true)

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-026%20-%20save%20and%20load.png?raw=true)

to remove file we use : 

```
func Remove(name string) error
```
>Remove removes the named file or (empty) directory. If there is an error, it will be of type *PathError.

in `os package`

---

## Organizing Data With Structs


`Struct` is short for structure , it is a data structure in go . like a collection of different properties that are somehow related together of have some type of common purpose,

in javascript is like object.

![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/Untitled%20Diagram-012%20-%20structs.png?raw=true)

if we declare struct like :

```
var alex person
```

then the struc will give Zero Value .




![Chat Preview](https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-015%20-%20slice%20selection.png?raw=true)





