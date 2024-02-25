package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getUserPosts(username string) {
	var db = getDatabaseItem()
	var rows, _ = db.Query("select * from posts where username = '" + username + "'")

	for rows.Next() {
		var postList [16]string

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
			country    string
			dob        string
			value      string

			like_count  string
			user_public string
			date        string
		)

		rows.Scan(&photo, &community, &id_num, &post_title, &username, &details, &post_type, &visability, &catagory, &off_name, &country, &dob, &value, &like_count, &user_public, &date)

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
		postList[10] = country
		postList[11] = dob
		postList[12] = value

		postList[13] = like_count
		postList[14] = user_public
		postList[15] = date

		color.Blue(fmt.Sprint(postList))
	}

	color.Red("Out Of Posts")

	db.Close()
}

func removeUserPost(postid string) string {
	var db = getDatabaseItem()

	var confirm = ""

	fmt.Print("Are You Sure You Want To Delete Post " + postid + "(y/n): ")

	fmt.Scan(&confirm)

	var rows, _ = db.Query("select * from posts where id_num = '" + postid + "'")

	for rows.Next() {
		if confirm == "y" {
			db.Query("DELETE FROM joint_cposts WHERE post_id ='" + postid + "'")
			db.Query("DELETE FROM posts WHERE id_num ='" + postid + "'")
			db.Close()
			return "Post: " + postid + " Has Been Deleted"
		} else {
			db.Close()
			return "ACTION: CANCELLED"
		}
	}
	db.Close()
	return "Post Not Found"
}
