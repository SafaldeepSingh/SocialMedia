package DAO

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"project/Model"
	"strconv"
)

func GetPosts() []Model.Post{
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/social_media")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	rc, err := db.Query("select count(*) from posts")
	rc.Next()
	rowCount:=0
	rc.Scan(&rowCount)
	results, err := db.Query("select post_id,user_id,caption from posts")
	if err != nil {
		//fmt.Println(err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var posts =make([]Model.Post,rowCount)
	i:=0
	for results.Next() {
		var post Model.Post

		err = results.Scan(&post.Id, &post.UserId,&post.Caption)
		posts[i] = post
		i++
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return posts
	//for i:=0;i<len(posts);i++{
	//	fmt.Println(posts[i].Id,posts[i].UserId,posts[i].Caption)
	//}
}
func GetPostsByUserId(userId int) []Model.Post{
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/social_media")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	rc, err := db.Query("select count(*) from posts where user_id="+strconv.Itoa(userId))
	rc.Next()
	rowCount:=0
	rc.Scan(&rowCount)
	results, err := db.Query("select post_id,user_id,caption from posts where user_id="+strconv.Itoa(userId))
	if err != nil {
		//fmt.Println(err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var posts =make([]Model.Post,rowCount)
	i:=0
	for results.Next() {
		var post Model.Post

		err = results.Scan(&post.Id, &post.UserId,&post.Caption)
		posts[i] = post
		i++
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return posts
	//for i:=0;i<len(posts);i++{
	//	fmt.Println(posts[i].Id,posts[i].UserId,posts[i].Caption)
	//}
}