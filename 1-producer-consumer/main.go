//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"errors"
	"fmt"
	"time"
)

func producer(stream Stream, tweetChan chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if errors.Is(err, ErrEOF) {
			close(tweetChan)
			return
		}

		tweetChan <- tweet
	}
}

func consumer(tweetChan <-chan *Tweet) {
	for {
		if tweet, ok := <-tweetChan; ok {
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets about golang")
			} else {
				fmt.Println(tweet.Username, "\tdoes not tweet about golang")
			}

		} else {
			return
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweetChan := make(chan *Tweet, 10)

	// Producer
	go producer(stream, tweetChan)

	// Consumer
	go consumer(tweetChan)

	fmt.Printf("Process took %s\n", time.Since(start))
}
