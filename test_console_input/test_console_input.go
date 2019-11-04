package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	// Test_001()
	//Test_002()
	Test_003(nil)
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func Test_001(t *testing.T) {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

func Test_002(t *testing.T) {
	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func Test_003(t *testing.T) {
	fmt.Println("--- input begin")
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		msg := sender.Text()
		if msg == "stop" {
			return
		} else {
			fmt.Println("--- input:", msg)
		}
	}
	fmt.Println("--- exit")
}
