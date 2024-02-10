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
}

func setClientRequests(clientRequests chan *SocialMediaPostRequest) {
	instagramPost := &SocialMediaPostRequest{"Hello Instagram!", postToInstagram}
	tikTokPost := &SocialMediaPostRequest{"Hello TikTok!", postToTikTok}
	clientRequests <- instagramPost
	clientRequests <- tikTokPost

	close(clientRequests)
}

func processRequests(clientRequests chan *SocialMediaPostRequest) {
	for request := range clientRequests {
		fmt.Println(request.text, "sent!")
	}
}

func main() {
	clientRequests := make(chan *SocialMediaPostRequest)

	go setClientRequests(clientRequests)

	processRequests(clientRequests)


}
