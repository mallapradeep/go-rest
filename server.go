package main

import (
	// Standard library packages

	"net/http"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/mallapradeep/go-rest/controllers"
	"github.com/mallapradeep/go-rest/tree/master/controllers"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)

	/////////////////////////////////////////////////////////////////////////////////////////
	// // Add a handler on /test
	// // Get a user resource
	// r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	// Stub an example user
	// 	u := models.User{
	// 		Name:   "Pradeep Malla",
	// 		Gender: "male",
	// 		Age:    50,
	// 		Id:     p.ByName("id"),
	// 	}

	// 	// Marshal provided Interface to JSON structure
	// 	uj, _ := json.Marshal(u)

	// 	// Write content-type, statuscode, payload
	// 	w.Header().Set("Content-type", "application/json")
	// 	w.WriteHeader(200)
	// 	fmt.Fprintf(w, "%s", uj)
	// })

	// r.POST("/user", func(w http.ResponseWriter, r *http.Request, _httprouter.Params){
	// 	// Stub an user to be populated from the body
	// 	u := models.User{}

	// 	// Populate the user data
	// 	json.NewDecoder(r.Body).Decode(&u)

	// 	// Add an Id
	// 	u.Id = "foo"

	// 	// Write content-type, status-code, payload
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(201)
	// 	fmt.Fprintf(w, "%s", uj)
	// })

	// r.DELETE("/user/:id", func(w http.ResponseWriter, r *http.Request, _httprouter.Params){
	// 	// on;y write status for now
	// 	w.WriteHeader(200)
	// })

}
