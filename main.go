package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"tuanm.dev/golab/lab01"
	"tuanm.dev/golab/lab02"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	args := os.Args[1:]
	var i int = 1
	if len(args) > 0 {
		j, err := strconv.Atoi(args[0])
		if err != nil {
			println(err.Error())
		} else {
			i = j
		}
	}

	switch i {
	case 1:
		runLab011()
	case 2:
		runLab012()
	case 3:
		runLab013()
	case 21:
		runLab021()
	default:
		println("no lab found")
	}
}

func runLab011() {
	print("tell me something: ")
	text, _ := reader.ReadString('\n')
	println("you can count on me like:")
	for word, count := range lab01.CountWords(text) {
		fmt.Printf("- %s: %d\n", word, count)
	}
}

func runLab012() {
	print("tell me a number: ")
	var n uint
	fmt.Scan(&n)
	primes, err := lab01.GetAllPrimesUpTo(n, "./lab/first-1000000-primes.json")
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("these are all primes up to %d: ", n)
		for _, prime := range primes {
			fmt.Printf("%d, ", prime)
		}
		if len(primes) > 0 {
			println("that's all")
		} else {
			println("none")
		}
	}
}

func runLab013() {
	print("please enter a url: ")
	url, _ := reader.ReadString('\n')
	url = strings.Replace(url, "\n", "", -1) // remove last newline
	hasProtocol, _ := regexp.MatchString("^(http:\\/\\/)|(https:\\/\\/)", url)
	if !hasProtocol {
		url = "https://" + url
	}
	links, err := lab01.CrawlAnchorsFrom(url)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("all links i found at %s:\n", url)
		for _, link := range links {
			fmt.Printf("- %s\n", link)
		}
	}
}

func runLab021() {
	messages := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		duration := time.Duration(rand.Intn(5)*int(time.Second) + 1) // duration is randomly in range [1,5]
		println("you're gonna got a quote every " + duration.String())
		println("press ^c to exit")
		for {
			time.Sleep(duration)
			messages <- lab02.GenerateRandomMessage()
		}
	}()

	for message := range messages {
		println(message)
	}
}
