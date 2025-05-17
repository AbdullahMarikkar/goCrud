package models

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Blog struct{
	Id int `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Created_At string `json:"createdAt"`
}

var DB *sql.DB

func ConnectDatabase()error{
	db,err := sql.Open("sqlite","./goCrud.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func GetBlogs()([]Blog,error){
	rows,err := DB.Query("SELECT * FROM blogs")

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	blogs := make([]Blog,0)

	for rows.Next(){
		singleBlog := Blog{}
		err = rows.Scan(&singleBlog.Id,&singleBlog.Author,&singleBlog.Content,&singleBlog.Created_At)

		if err != nil{
			return nil,err
		}

		blogs = append(blogs,singleBlog)
	}

	err = rows.Err()

	if err != nil{
		return nil,err
	}
	return blogs,nil
}