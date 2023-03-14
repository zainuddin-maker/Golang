package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{

		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	//case 2 add channel
	c := make(chan string)

	for _, link := range links {

		// case 1 add go keyword
		//case 2 add channel
		go checkLink(link, c)
	}
	//case 2 add channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// eternal fetch
	// for {
	// 	go checkLink(<-c , c)
	// }

	// alternate eternal fetch
	for l := range c {
		// add pause but not best practise

		// time.Sleep(5 * time.Second)

		// go checkLink(l, c)

		// add literal function

		go func ()  {
			time.Sleep(5 * time.Second)

		    go checkLink(l, c)
		}()
	}

}

// case 2 add channel
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		//case 2 add channel
		// c <- "Migh be down i think"

		// eternal fetch
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//case 2 add channel

	//// eternal fetch
	c <- link
}
