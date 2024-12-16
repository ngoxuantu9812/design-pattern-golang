package main

import "fmt"

func main() {
	mailjet, err := getProvider("mailjet")
	if err != nil {
		fmt.Println(err)
	}
	ses, err := getProvider("ses")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(mailjet.getName(), mailjet.getKey())
	fmt.Print(ses.getName(), ses.getKey())
}
