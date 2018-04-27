package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Print("Hello highly valued Talent Search customer!\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your customer name: ")

	customer := readAndTrim(reader)

	checkout := Checkout{
		pricingRules: GetRules(customer),
	}

	checkout.Total()
	fmt.Printf("You are a '%v' customer\n", customer)

	option := "1"

	for ok := true; ok; ok = (option != "exit") {
		time.Sleep(1 * time.Second)
		fmt.Println("Would you like to")
		fmt.Println("1. add a job ad 2. delete a job ad 3. get total?")
		option, _ = reader.ReadString('\n')

		switch option {
		case "1\n":
			AddJobAdCommandLineOptions(reader, &checkout)
		case "2\n":
			DeleteJobAdCommandLineOptions(reader, &checkout)
		case "3\n":
			fmt.Printf("$%v\n", checkout.Total())
		case "exit\n":
			break
		default:
			fmt.Println("Invalid option")
		}
	}
}

// AddJobAdCommandLineOptions - Adding job command line options
func AddJobAdCommandLineOptions(reader *bufio.Reader, checkout *Checkout) {
	option := "1"

	fmt.Println("Select a job ad to add")
	fmt.Println("1. Classic, 2. Standout, 3. Premium ad: ")
	option, _ = reader.ReadString('\n')
	switch option {
	case "1\n":
		checkout.Add(Item{classic})
	case "2\n":
		checkout.Add(Item{standout})
	case "3\n":
		checkout.Add(Item{premium})
	default:
		fmt.Println("Invalid option")
	}
	fmt.Println(checkout.Show())
}

// DeleteJobAdCommandLineOptions - Delete a job ad via the command line
func DeleteJobAdCommandLineOptions(reader *bufio.Reader, checkout *Checkout) {
	option := "1"

	fmt.Println("Select an ad to delete")
	fmt.Println("1. Classic, 2. Standout, 3. Premium ad: ")
	option, _ = reader.ReadString('\n')
	switch option {
	case "1\n":
		checkout.Delete(Item{classic})
	case "2\n":
		checkout.Delete(Item{standout})
	case "3\n":
		checkout.Delete(Item{premium})
	default:
		fmt.Println("Invalid option")
	}
	fmt.Println(checkout.Show())
}

// readAndTrim - trims the string
func readAndTrim(reader *bufio.Reader) string {
	readString, _ := reader.ReadString('\n')
	return strings.TrimSuffix(readString, "\n")
}
