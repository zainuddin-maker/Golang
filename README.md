# **GOLANG**
https://github.com/StephenGrider/GoCasts

#### **Environtment Setup**
---
 ![Chat Preview]( https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-001%20-%20install.png?raw=true)

 to install GO , go to https://go.dev/dl/

 to check if go is already installed , open terminal and write `go` , then a long help message will be appear on the screen.

 and install GO extenction in VS code , just search Go and install.

 after install , close directory , and open again , and make file . so in the bottom right there is `plaintext` change that to `GO` , after that in bottom right we will see yellow text something that says `analysis tools missing` , and click that , and install.

 ---
 in `main.go` we have several questions :

 1. How do we run the code in our project?
 2. What does `package main` mean?
 3. What does `import "fmt"` mean ?
 4. What that `func` thing?
 5. How is the `main.go` file organized ?

 Answer:

 1. `go run main.go`


 ![Chat Preview](  https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-003%20-%20go%20cli.png?raw=true)

 different `go build` and `go run` is , go build just compile the program not execute.

 if we run `go build` than we will get `main.exe` so this is an actual executable file that was built out of our source code of main go.

 we can run it by `./main` or `./main.exe`

 2. A package is a collection of common source code files.

 ![Chat Preview](   https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-008%20-%20what%20package.png?raw=true)

 package can have many related file instide of it ,each file ending with a file extension of GO (.go). the only requirement for every file inside of a package is that the very first line of each file must declare the package that it belongs to.

  ![Chat Preview](    https://github.com/zainuddin-maker/Golang/blob/master/imgdiagram/diagrams-014%20-%20packages.png?raw=true)

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





