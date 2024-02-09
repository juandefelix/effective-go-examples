package main

import "fmt"

// functions that post to different social media
func postToInstagram(text string) bool {
	fmt.Println("Posted to Instagram successfully")
	return true
}

func postToTikTok(text string) bool {
	fmt.Println("Posted to TikTok successfully")
	return true
}

type SocialMediaPostRequest struct {
	text string
	function func(string) bool
	// notifyChannel chan int
}

func sendRequests(clientRequests chan *SocialMediaPostRequest) {
	instagramPost := &SocialMediaPostRequest{"Hello Instagram!", postToInstagram}
	tikTokPost := &SocialMediaPostRequest{"Hello TikTok!", postToTikTok}
	clientRequests <- instagramPost
	clientRequests <- tikTokPost
}

func main() {
	clientRequests := make(chan *SocialMediaPostRequest)

	go sendRequests(clientRequests)

	request1, request2 := <-clientRequests, <-clientRequests // receive from c

	fmt.Println(request1, request2)

}
