package models

import (
	"fmt"
	"time"

	"github.com/AbdullahMarikkar/goCrud/utils"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Created_At string `json:"createdAt"`
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUsers() ([]User, error) {
	rows, err := DB.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.Id, &singleUser.Name, &singleUser.Email, &singleUser.Password, &singleUser.Role, &singleUser.Created_At)

		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user CreateUserDto) ([]User, error) {
	result, err := DB.Exec(`INSERT INTO users (name, email, password, role) VALUES ( ?, ?, ?, ? )`, user.Name, user.Email,user.Password,user.Role)

	var newUser User

	if err != nil {
		fmt.Println("Error in CreateUser",err)
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Error in CreateUser",err)
		return nil, err
	}
	users := make([]User, 0)

	newUser.Id = int(id)
	newUser.Name = user.Name
	newUser.Password = utils.HashPassword(user.Password)
	newUser.Email = user.Email
	newUser.Role = user.Role
	newUser.Created_At = time.Now().GoString()
	users = append(users,newUser)
	return users,nil
}

func AuthorizeUser(email string,password string)(*User,error){
	row, err := DB.Query("SELECT * FROM users where email = ?",email)

	if err != nil {
		fmt.Println("Error in AuthorizeUser",err)
		return nil, err
	}

	defer row.Close()

	var singleUser User
	for row.Next(){
		err = row.Scan(&singleUser.Id,&singleUser.Email,&singleUser.Name,&singleUser.Password,&singleUser.Role,&singleUser.Created_At)
		if err != nil{
			return nil,err
		}
		break
	}

	if(!utils.VerifyPassword(password,singleUser.Password)){
		fmt.Println("Error in AuthorizeUser, Password is Incorrect",err)
		return nil,err
	}

	if err != nil{
		return nil,err
	}

	return &singleUser,nil
}