package main

import "fmt"


type storyPage struct {
	text string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func (page *storyPage) addAfter(text string){
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func main() {
	page1 := storyPage{"stooory", nil}
	page1.addToEnd("stoory 2")
	page1.addAfter("stoory 4")
	page1.addToEnd("stoory 3")


	page1.playStory()
}