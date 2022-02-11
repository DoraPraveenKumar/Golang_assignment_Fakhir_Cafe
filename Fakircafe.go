package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var amt_day int
var cust_total int
var amt_bill int

func calculate_bill(input string) {
	for i := 0; ; i++ {
		fmt.Println("---------------\ncode")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input := scanner.Text()
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)
		if input == "*" {
			amt_bill = 0
			fmt.Println("\nYour total bill is : ", cust_total)
			fmt.Println("Wishing you had a great time with us\n\n============================\n\n")
			time.Sleep(3 * time.Second)
			cust_total = 0
			menu()
			continue
		} else if input == "END" {
			fmt.Println("=====================================\nTotal income at end of the day is:", amt_day)
			time.Sleep(2 * time.Second)
			fmt.Println("Wishing you had a great day\n====================================\n")
			break
		} else if len([]rune(input)) == 3 {
			item_costing(input)
		} else {
			fmt.Println("* WRONG INPUT **\n===========================\nEnter Again\n")
		}
	}
}

func menu() {
	fmt.Println("\n\t    WElCOME TO FAKIR CAFE\n")
	fmt.Println("\t\t    MENU\n\n\tCODE\tITEM\t\tPRICE\n")
	fmt.Println("\tC: coffee\t\t40rs\n\tD: dosa \t\t80rs\n\tT: tomato soup \t\t20rs\n\tI: idli  \t\t20rs\n\tV: vada \t\t25rs\n\tB: bhature &chhole \t30rs\n\tP: paneer pakoda \t30rs\n\tM: manchurian \t\t90rs\n\tH: hakka noodle \t70rs\n\tF: French fries \t30rs\n\tJ: jalebi \t\t30rs\n\tL: lemonade \t\t15rs\n\tS: spring roll \t\t20rs\n")
	fmt.Println("Enter your quantity and code eg:-[ 4 C ]")
	fmt.Println("Press [ * ] to get bill of a customer")
	fmt.Println("Enter End to close the day and get total income of a day")
}

func item_costing(input string) {
	items := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	quant, _ := strconv.Atoi(string(input[0]))
	id := string(input[2])
	amt_bill = quant * items[id]
	cust_total += amt_bill
	fmt.Println("Cost of the dish : ", amt_bill)
	amt_day += amt_bill
}
func main() {
	fmt.Println("HELLO  this is DORA")
	menu()
	calculate_bill("Starting")
}
