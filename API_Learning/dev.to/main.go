package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Version 1 of the code
// type server struct{}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(
// 			`{"message": "GET Called"}`))

// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(
// 			`{"message": "POST Called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(
// 			`{"message": "PUT Called"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(
// 			`{"message": "DELETE Called"}`))
// 	}
// }

// func main() {
// 	s := &server{}
// 	http.Handle("/", s)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// Version 2 of the code
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(
// 			`{"message": "GET Called"}`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(
// 			`{"message": "POST Called"}`))
// 	case "PUT":
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Write([]byte(
// 			`{"message": "PUT Called"}`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(
// 			`{"message": "DELETE Called"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(
// 			`{"message": "Not Found"}`))
// 	}
// }

// func main() {
// 	http.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe("", nil))
// }

// Version 3 With mux routing and home from verion 2
// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe("", r))
// }

// Version 4
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`{"message": "GET Called"}`))
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(
		`{"message": "POST Called"}`))
}
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(
		`{"message": "PUT Called"}`))
}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(
		`{"message": "DELETE Called"}`))
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(
		`{"message": "Not Found"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(
				`{"message":"need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(
				`{"message":"need a number"}`))
			return
		}
	}
	query := r.URL.Query()
	location := query.Get("location")
	w.Write([]byte(
		fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s"}`,
			userID, commentID, location)))
}
func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1.0").Subrouter()

	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}",params).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":80", api))
}
