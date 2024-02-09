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
	text []byte
	function func([]byte) bool
	notifyChannel chan bool
}
