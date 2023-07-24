package mod04

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Run(lab int) {
	switch lab {
	case 1:
		runLab01()
	case 2:
		runLab02()
	case 3:
		runLab03()
	case 4:
		runLab04()
	default:
		println("no lab found")
	}
}

func runLab01() {
	messages := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		duration := time.Duration(rand.Intn(5)*int(time.Second) + 1) // duration is randomly in range [1,5]
		println("you're gonna got a quote every " + duration.String())
		for {
			time.Sleep(duration)
			messages <- GenerateRandomMessage()
		}
	}()

	println("press ^c to exit")
	for message := range messages {
		println(message)
	}
}

func runLab02() {
	red := make(chan bool)
	green := make(chan bool)
	yellow := make(chan bool)
	current := make(chan string)

	go WaitColor(Color{
		Current: Red,
		Ttl:     15,
		Wait:    red,
		Notify:  green,
	}, current)
	go WaitColor(Color{
		Current: Green,
		Ttl:     30,
		Wait:    green,
		Notify:  yellow,
	}, current)
	go WaitColor(Color{
		Current: Yellow,
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

func runLab03() {
	var reader = bufio.NewReader(os.Stdin)
	ports := make(chan int, 65535)
	results := make(chan Port)

	print("give me a host: ")
	host, _ := reader.ReadString('\n')
	host = strings.Replace(host, "\n", "", -1)
	println("scanning open ports for " + host + "...")

	checkers := 69
	timeout := 6
	for i := 0; i < checkers; i++ {
		go func() {
			for port := range ports {
				results <- Port{
					Number: port,
					Open:   CheckPort(host, port, timeout),
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

func runLab04() {
	print("which port your server will listen on? ")
	var port int
	fmt.Scan(&port)
	fmt.Printf("ok, server listening on %d...\n", port)
	println("press ^c to exit")
	StartServer(port)
}
