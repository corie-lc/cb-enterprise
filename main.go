package main

import (
	"fmt"
	"strconv"

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

	// REMOVE POST BY IS
	color.Magenta("4) Remove User Post?")

	// CHANGE USER TEIR LEVEL
	color.Magenta("5) Change User Tier Level?")

	// GET USER MEMBERSHIP INFORMATION
	color.Magenta("6) Get User Tier Table?")

	// CHANGE USER AUTO PRICE COUNT INFORMATION
	color.Magenta("7) Change User Auto Price Count?")

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
	} else if option == "4" {

		var idInput = ""
		color.Green("Enter A Post Id: ")
		fmt.Scan(&idInput)

		color.Yellow("Results are Loading ...:")
		removeUserPost(idInput)

		color.Blue(removeUserPost(idInput))
		controller()
	} else if option == "5" {
		var usernameInput = ""
		var levelInput = -1

		color.Green("Enter A Username: ")
		fmt.Scan(&usernameInput)

		color.Cyan("Enter 0 for free, 1 for essentails or 2 for plus.")

		color.Green("Enter A Level : ")
		fmt.Scan(&levelInput)

		color.Yellow("Results are Loading ...:")
		color.Blue(changeUserTierLevel(usernameInput, levelInput))
		controller()
	} else if option == "6" {
		var usernameInput = ""

		color.Green("Enter A Username: ")
		fmt.Scan(&usernameInput)

		color.Blue(getUserTierTable(usernameInput))

		controller()

	} else if option == "7" {
		var usernameInput = ""
		var count = -1

		color.Green("Enter A Username: ")
		fmt.Scan(&usernameInput)

		color.Green("Enter A New Count(NUMBER) : ")
		fmt.Scan(&count)

		color.Yellow("Results are Loading ...:")
		color.Blue(changeUserAutoPriceCount(usernameInput, strconv.Itoa(count)))
		controller()
	} else {

		// FAILURE TO PROVIDE VALID OPTION
		color.Red("Invalid Option...")
		controller()
	}
}
