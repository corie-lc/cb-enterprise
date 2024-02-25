package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("-------------------------------------------------------------------------")
	color.Red("This Application Is Verified By ENV VAR - If you don't have them assigned, please reach out to management.")
	controller()

}

func controller() {
	fmt.Println("--------------------------------------------------------------------------")

	var option = ""

	color.Magenta("What Would You Like To Do? >")
	fmt.Println("------------------------------------------------------------------------")

	//fmt.Println("0) Send A Site Wide Notification?")

	// GENERAL USER SEARCH
	color.Magenta("1) Search For User")

	// DELETE USER PERM
	color.Magenta("2) Delete A User?")

	// PRINT ALL USER POSTS WITH INFORMATION
	color.Magenta("3) Get User Posts/PostsInformation?")

	//fmt.Println("4) Remove User Collection?")

	color.Green("Enter An Option: ")

	fmt.Scan(&option)

	// GENERAL USER SEARCH
	if option == "1" {
		color.Yellow("---- Loading User Search --")

		var usernameInput = ""
		color.Green("Enter A Username: ")
		fmt.Scan(&usernameInput)

		color.Yellow("Results are Loading ...:")
		color.Blue(getUserInformation(usernameInput))
		controller()

	} else if option == "2" {
		// DELETE USER PERM
		color.Yellow("---- Loading User Delete --")

		var usernameInput = ""
		color.Green("Enter A Username to DELETE FOREVER: ")
		fmt.Scan(&usernameInput)

		color.Blue(removeUser(usernameInput))
		controller()
	} else if option == "3" {
		// PRINT ALL USER POSTS WITH INFORMATION
		color.Yellow("---- Loading User Posts Search --")

		var usernameInput = ""
		color.Green("Enter A Username: ")
		fmt.Scan(&usernameInput)

		color.Yellow("Results are Loading ...:")
		getUserPosts(usernameInput)
		controller()
	} else {
		// FAILURE TO PROVIDE VALID OPTION
		color.Red("Invalid Option...")
		controller()
	}
}
