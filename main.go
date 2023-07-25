package main

import (
	"e-commerce-postgresdb/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// DB store the db session info. Need to be initialized once
type DBClient struct {
	db *gorm.DB
}

// UserResponse is the response to be send back for User
type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

// GetUserByFirstName fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUserByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("first_name")

	// Handle response details
	var query = "select * from \"user\" where data->>'first_name'=?"
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// responseMap := map[string]interface{}{"url": ""}
	respJSON, _ := json.Marshal(users)
	w.Write(respJSON)
}

// GetUser fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)

	// Handle response details
	driver.db.First(&user, vars["id"])
	var userData interface{}

	// Uumarshal JSON string to interface
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// responseMap := map[string]interface{}{"url": ""}
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

// PostUser adds URL to DB and gives back shortend string
func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := io.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.IniDB()
	if err != nil {
		panic(err)
	}

	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/v1/user/{id:[a-zA-z0-9]*}", dbclient.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods("POST")
	r.HandleFunc("/v1/user", dbclient.GetUserByFirstName).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9000",

		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
