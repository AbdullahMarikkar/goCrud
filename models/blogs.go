package models

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Blog struct{
	Id int `json:"id"`
	User_Id int `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Created_At string `json:"createdAt"`
}

type CreateBlog struct {
	User_Id int `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
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
		err = rows.Scan(&singleBlog.Id,&singleBlog.User_Id,&singleBlog.Content,&singleBlog.Created_At,&singleBlog.Title)

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
func CreateBlogs(createBlog CreateBlog)([]Blog,error){
	result,err := DB.Exec(`INSERT INTO blogs (user_id, title, content ) VALUES ( ?, ?, ? )`,createBlog.User_Id,createBlog.Title,createBlog.Content)
	var blog Blog

	if err != nil {
        return nil,err
    }

	id, err := result.LastInsertId()

	if err != nil {
        return nil,err
    }
	blogs := make([]Blog,0)


	blog.Id = int(id)
	blog.Created_At = time.Now().GoString()
	blog.Title = createBlog.Title
	blog.Content = createBlog.Content
	blog.User_Id = createBlog.User_Id
	blogs = append(blogs, blog)
	return blogs,nil
}

func GetBlogById(id int)(*Blog,error){
	row,err := DB.Query("SELECT * FROM blogs where id = ?",id)

	if err != nil {
        return nil,err
    }
	defer row.Close()

	var singleBlog Blog
	for row.Next(){
		err = row.Scan(&singleBlog.Id,&singleBlog.User_Id,&singleBlog.Title,&singleBlog.Content,&singleBlog.Created_At)
		if err != nil{
			return nil,err
		}
		break
	}
	
	if err != nil{
		return nil,err
	}

	return &singleBlog,nil

}