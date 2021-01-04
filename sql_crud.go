package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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


func main() {

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", CreateTokenEndpoint).Methods("POST")
	api.HandleFunc("/create", ValidateMiddleware(userCreate)).Methods("POST")
	api.HandleFunc("/delete", ValidateMiddleware(userdelete)).Methods("DELETE")
	api.HandleFunc("/update", ValidateMiddleware(userupdate)).Methods("PUT")
	api.HandleFunc("/test", ValidateMiddleware(getuser)).Methods("GET")
	fmt.Printf("Starting server at port 8091\n")
	http.ListenAndServe("localhost:8091", r)

}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
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


//create token
func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	if user.Username == "Admin" && user.Password == "Admin123" {
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	}else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid Username or password"})
	}
}

// create new user
func userCreate(w http.ResponseWriter, r *http.Request) {
	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")
	if err != nil {
		panic(err.Error())
	}
	res, err := conn.Exec("INSERT INTO users (age,name,department) VALUES(?,?,?)", p.Age, p.Name, p.Department)
	if err != nil {
		panic(err)
	}
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
func userdelete(w http.ResponseWriter, r *http.Request) {

	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")
	if err != nil {
		panic(err.Error())
	}
	_, err = conn.Exec("DELETE FROM users where id=?", p.ID)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Deleted",
		"Status":  http.StatusOK,
	})

}

//update user {id}{name}{age}{dept}
func userupdate(w http.ResponseWriter, r *http.Request) {

	var p user
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")
	if err != nil {
		panic(err.Error())
	}
	_, err = conn.Exec("update users set name=?, age=?, department=? where id=?", p.Name, p.Age, p.Department, p.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Updated user")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Updated",
		"Status":  http.StatusOK,
	})

}

//get all users
func getuser(w http.ResponseWriter, r *http.Request) {

	// var p user
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/authpass")
	if err != nil {
		panic(err.Error())
	}
	rows, err := conn.Query("select * from users")
	if err != nil {
		panic(err)
	}

	var post = user{}

	for rows.Next() {
		rows.Scan(&post.ID, &post.Age, &post.Name, &post.Department)
		// fmt.Println(post)
		fmt.Fprintf(w, "Users: %+v \n", post)
		json.NewEncoder(w).Encode(post)
	}

}

