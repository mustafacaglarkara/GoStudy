package main

import (
	"fmt"
	"time"
)

func CalculateOne(channel chan string) {
	fmt.Println("Calculation phase one...")
	time.Sleep(time.Second * 2)
	channel <- "phase one is done"
}

func CalculateTwo(channel chan string) {
	fmt.Println("Calculation phase two...")
	time.Sleep(time.Second * 5)
	channel <- "phase two is done"
}

func EvaluateTestData(channel chan string) {
	fmt.Println("Creting test data...")
	time.Sleep(time.Second * 3)
	channel <- "Evaluation is done"
}

func main() {
	channelOne := make(chan string, 1)
	channelTwo := make(chan string, 1)
	channelEval := make(chan string, 1)

	go CalculateOne(channelOne)
	go CalculateTwo(channelTwo)
	go EvaluateTestData(channelEval)

	for i := 0; i < 3; i++ {
		select {
		case messageOne := <-channelOne:
			fmt.Println(messageOne)
		case messageTwo := <-channelTwo:
			fmt.Println(messageTwo)
		case messageEval := <-channelEval:
			fmt.Println(messageEval)
		}
	}
}
