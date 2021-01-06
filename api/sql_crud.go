package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strings"
	_ "testing"
)

type user struct {
	ID         int
	Age        int
	Name       string
	Department string
}


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}


//run fst time
//var schema string = "CREATE TABLE `users` (	  	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,	`age` integer(10) NOT NULL	,  	`name` varchar(255) NOT NULL	,	`department` varchar(255) NOT NULL)"



func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					}
					return []byte("secret"), nil
				})
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

//create token
func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	if user.Username == "Admin" && user.Password == "Admin123" {
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	}else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid Username or password"})
	}
}

// create new user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")

	res, err := conn.Exec("INSERT INTO users (age,name,department) VALUES(?,?,?)", p.Age, p.Name, p.Department)
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created user with id:%d", id)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Created",
		"Status":  http.StatusCreated,
	})

}

//delete user {id}
func Userdelete(w http.ResponseWriter, r *http.Request) {

	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
	}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")

	_, err = conn.Exec("DELETE FROM users where id=?", p.ID)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Deleted",
		"Status":  http.StatusOK,
	})

}

//update user {id}{name}{age}{dept}
func Userupdate(w http.ResponseWriter, r *http.Request) {

	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
	}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")

	_, err = conn.Exec("update users set name=?, age=?, department=? where id=?", p.Name, p.Age, p.Department, p.ID)

	fmt.Printf("Updated user")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Updated",
		"Status":  http.StatusOK,
	})

}

//get all users
func Getuser(w http.ResponseWriter, r *http.Request) {

	// var p user
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")
	rows, err := conn.Query("select * from users")
	if err != nil {
	}

	var post = user{}

	for rows.Next() {
		rows.Scan(&post.ID, &post.Age, &post.Name, &post.Department)
		// fmt.Println(post)
		fmt.Fprintf(w, "Users: %+v \n", post)
		json.NewEncoder(w).Encode(post)
	}

}

