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
	case 22:
		runLab022()
	case 23:
		runLab023()
	case 24:
		runLab024()
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
		for {
			time.Sleep(duration)
			messages <- lab02.GenerateRandomMessage()
		}
	}()

	println("press ^c to exit")
	for message := range messages {
		println(message)
	}
}

func runLab022() {
	red := make(chan bool)
	green := make(chan bool)
	yellow := make(chan bool)
	current := make(chan string)

	go lab02.WaitColor(lab02.Color{
		Current: lab02.Red,
		Ttl:     15,
		Wait:    red,
		Notify:  green,
	}, current)
	go lab02.WaitColor(lab02.Color{
		Current: lab02.Green,
		Ttl:     30,
		Wait:    green,
		Notify:  yellow,
	}, current)
	go lab02.WaitColor(lab02.Color{
		Current: lab02.Yellow,
		Ttl:     3,
		Wait:    yellow,
		Notify:  red,
	}, current)

	println("press ^c to exit")
	yellow <- true
	for color := range current {
		println(time.Now().Format(time.Stamp) + ": " + color)
	}
}

func runLab023() {
	ports := make(chan int, 65535)
	results := make(chan lab02.Port)

	print("give me a host: ")
	host, _ := reader.ReadString('\n')
	host = strings.Replace(host, "\n", "", -1)
	println("scanning open ports for " + host + "...")

	checkers := 69
	timeout := 6
	for i := 0; i < checkers; i++ {
		go func() {
			for port := range ports {
				results <- lab02.Port{
					Number: port,
					Open:   lab02.CheckPort(host, port, timeout),
				}
			}
		}()
	}

	for port := 1; port <= cap(ports); port++ {
		ports <- port
	}

	println("press ^c to exit")
	for result := range results {
		if result.Open {
			println(result.Number)
		}
	}
}

func runLab024() {
	print("which port your server will listen on? ")
	var port int
	fmt.Scan(&port)
	fmt.Printf("ok, server listening on %d...\n", port)
	println("press ^c to exit")
	lab02.StartServer(port)
}
