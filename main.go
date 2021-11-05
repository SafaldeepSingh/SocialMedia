package main

import (
	"fmt"
	"net/http"
	"project/DAO"
)

func main(){
	//http.HandleFunc("/",home)
	//http.ListenAndServe(":8080",nil)
	//DAO.GetPosts()
	posts:=DAO.GetPostsByUserId(2)
 	for i:=0;i<len(posts);i++{
		fmt.Println(posts[i].Id,posts[i].UserId,posts[i].Caption)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"hello")
}


