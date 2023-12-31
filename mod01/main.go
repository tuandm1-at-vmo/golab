package mod01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Run(lab int) {
	switch lab {
	case 1:
		runLab01()
	case 2:
		runLab02()
	case 3:
		runLab03()
	default:
		println("no lab found")
	}
}

func runLab01() {
	var reader = bufio.NewReader(os.Stdin)
	print("tell me something: ")
	text, _ := reader.ReadString('\n')
	println("you can count on me like:")
	for word, count := range CountWords(text) {
		fmt.Printf("- %s: %d\n", word, count)
	}
}

func runLab02() {
	print("tell me a number: ")
	var n uint
	fmt.Scan(&n)
	primes, err := GetAllPrimesUpTo(n, "./mod01/first-1000000-primes.json")
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

func runLab03() {
	var reader = bufio.NewReader(os.Stdin)
	print("please enter a url: ")
	url, _ := reader.ReadString('\n')
	url = strings.Replace(url, "\n", "", -1) // remove last newline
	hasProtocol, _ := regexp.MatchString("^(http:\\/\\/)|(https:\\/\\/)", url)
	if !hasProtocol {
		url = "https://" + url
	}
	links, err := CrawlAnchorsFrom(url)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("all links i found at %s:\n", url)
		for _, link := range links {
			fmt.Printf("- %s\n", link)
		}
	}
}
