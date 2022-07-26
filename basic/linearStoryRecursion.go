package main

import "fmt"


type storyPage struct {
	text string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	page.nextPage.playStory()
}

func main(){
	page1 := storyPage{"stooory", nil}
	page2 := storyPage{"stooory 2", nil}
	page3 := storyPage{"stooory 3", nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	page1.playStory()
}