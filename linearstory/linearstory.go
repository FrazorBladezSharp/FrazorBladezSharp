package main

import (
	"bufio"
	"fmt"
	"os"
)

// linked list
type bookPage struct {
	text     string
	nextPage *bookPage
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	page1 := bookPage{"The rain came down on a moody night.", nil}
	page2 := bookPage{"Leaving the car door open I rush toward the Stage Entrance.", nil}
	page3 := bookPage{"Clutching my dress bag, sidesteping the wet slippery grate I get inside.", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page1)

	scanner.Scan()

}

func playStory(page *bookPage) {

	if page == nil {
		return
	}

	fmt.Println(page.text)
	playStory(page.nextPage) // use of recursion (calling itself)
}
