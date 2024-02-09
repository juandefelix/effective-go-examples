package main

import "fmt"

// functions that post to different social media
func postToInstagram([]byte) bool {
	fmt.Println("Posted to Instagram successfully")
	return true
}

func postToTikTok([]byte) bool {
	fmt.Println("Posted to TikTok successfully")
	return true
}

type SocialMediaPostRequest struct {
	text string
	function func([]byte) bool
	notifyChannel chan bool
}

func prepareRequests() chan SocialMediaPostRequest {
	clientRequests := make(chan SocialMediaPostRequest)

	instagramRequestSuccessChannel := make(chan bool)
	tikTokRequestSuccessChannel := make(chan bool)

	instagramPost := SocialMediaPostRequest{"Hello Instagram!", postToInstagram, instagramRequestSuccessChannel}
	tikTokPost := SocialMediaPostRequest{"Hello TikTok!", postToTikTok, tikTokRequestSuccessChannel}

	clientRequests <- instagramPost
	clientRequests <- tikTokPost

	return clientRequests
}

func sendRequests(clientRequests chan SocialMediaPostRequest) {
	// TODO
}

func main() {
	clientRequests := prepareRequests()

	sendRequests(clientRequests)
}
