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

	return fmt.Sprint(userList)

}
