package main

import (
	"fmt"
)

func getUserInformation(username string) string {
	var db = getDatabaseItem()
	var rows, _ = db.Query("select * from users where username = '" + username + "'")

	var (
		user     string
		email    string
		password string
		bio      string
		public   string
		avatar   string
	)

	var userList [6]string

	for rows.Next() {

		rows.Scan(&user, &email, &password, &bio, &public, &avatar)

		userList[0] = user
		userList[1] = email
		userList[2] = password
		userList[3] = bio
		userList[4] = public
		userList[5] = avatar

		return fmt.Sprint(userList)
	}

	userList[0] = "Not Found!"
	userList[1] = "Not Found!"
	userList[2] = "Not Found!"
	userList[3] = "Not Found!"
	userList[4] = "Not Found!"
	userList[5] = "Not Found!"

	db.Close()

	return fmt.Sprint(userList)

}

func removeUser(username string) string {
	var db = getDatabaseItem()

	var confirm = ""

	fmt.Print("Are You Sure You Want To Delete User " + username + "(y/n): ")

	fmt.Scan(&confirm)

	var rows, _ = db.Query("select * from users where username = '" + username + "'")

	for rows.Next() {
		if confirm == "y" {
			db.Query("DELETE FROM users WHERE username ='" + username + "'")
			db.Query("DELETE FROM posts WHERE username ='" + username + "'")
			db.Close()
			return "User: " + username + " Has Been Deleted"
		} else {
			db.Close()
			return "ACTION: CANCELLED"
		}
	}
	db.Close()
	return "User Not Found"

}
