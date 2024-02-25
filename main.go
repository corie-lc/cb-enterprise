package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("This Application Is Verified By ENV VAR - If you don't have them assigned, please reach out to management.")
	controller()

}

func controller() {
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------")

	var option = ""

	fmt.Println("What Would You Like To Do? >")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------")

	//fmt.Println("0) Send A Site Wide Notification?")
	fmt.Println("1) Search For User")
	//fmt.Println("2) Delete A User?")
	//fmt.Println("3) Remove User Post?")
	//fmt.Println("4) Remove User Collection?")

	fmt.Print("Enter An Option: ")

	fmt.Scan(&option)

	if option == "1" {
		fmt.Println("---- Loading User Search --")

		var usernameInput = ""
		fmt.Print("Enter A Username: ")
		fmt.Scan(&usernameInput)

		fmt.Println("Results are Loading ...:")
		color.Blue(getUserInformation(usernameInput))
		controller()
	} else {
		color.Red("Invalid Option...")
		controller()
	}
}
