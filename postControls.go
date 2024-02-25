package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getUserPosts(username string) {
	var db = getDatabaseItem()
	var rows, _ = db.Query("select * from posts where username = '" + username + "'")

	for rows.Next() {
		var postList [15]string

		var (
			photo        string
			community    string
			id_num       string
			post_title   string
			usernamePost string
			details      string

			post_type  string
			visability string
			catagory   string
			off_name   string
			dob        string
			value      string

			like_count  string
			user_public string
			date        string
		)

		rows.Scan(&photo, &community, &id_num, &post_title, &username, &details, &post_type, &visability, &catagory, &off_name, &dob, &value, &like_count, &user_public, &date)

		fmt.Println(community)

		postList[0] = photo
		postList[1] = community
		postList[2] = id_num
		postList[3] = post_title
		postList[4] = usernamePost
		postList[5] = details

		postList[6] = post_type
		postList[7] = visability
		postList[8] = catagory
		postList[9] = off_name
		postList[10] = dob
		postList[11] = value

		postList[12] = like_count
		postList[13] = user_public
		postList[14] = date

		color.Blue(fmt.Sprint(postList))
	}

	color.Red("Out Of Posts")

	db.Close()
}
