package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

var RC redis.Conn
var ID = 0

func getRedisConnection() redis.Conn {
	rc, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v\n", err)
	}
	return rc
}

func handleRequests() {
	RC = getRedisConnection()
	defer RC.Close()

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/fruits/", homePage).Methods("GET")
	myRouter.HandleFunc("/fruits/apple", apple).Methods("GET")
	myRouter.HandleFunc("/fruits/strawberry", strawberry).Methods("GET")
	myRouter.HandleFunc("/fruits/history", history).Methods("GET")

	log.Println("Fruits : Listening on localhost:8080 for requests...")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "\nWelcome to fruits!\n")
	fmt.Fprintf(w, "\nYou can choose between:")
	fmt.Fprintf(w, "\n- Apple")
	fmt.Fprintf(w, "\nFor apple you must use the following URL: /fruits/apple")
	fmt.Fprintf(w, "\n\n- Strawberry")
	fmt.Fprintf(w, "\nFor strawberry you must use the following URL: /fruits/strawberry")
	fmt.Fprintf(w, "\n\nYou can also check the history through the following url: /fruits/history")
}

func makeId(w http.ResponseWriter, req *http.Request) int {
	ID++
	return ID
}

func apple(w http.ResponseWriter, req *http.Request) {
	id := makeId(w, req)
	fruit := "You choose apple"
	fmt.Printf("%+v\n", fruit)
	_, err := RC.Do("SET", id, fruit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(id, ":", fruit)
	stringReturn := strconv.Itoa(id) + ":" + fruit
	js, errorJs := json.Marshal(stringReturn)
	if errorJs != nil {
		http.Error(w, errorJs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func strawberry(w http.ResponseWriter, req *http.Request) {
	id := makeId(w, req)
	fruit := "You choose strawberry"
	fmt.Printf("%+v\n", fruit)
	_, err := RC.Do("SET", id, fruit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(id, ":", fruit)
	stringReturn := strconv.Itoa(id) + ":" + fruit
	js, errorJs := json.Marshal(stringReturn)
	if errorJs != nil {
		http.Error(w, errorJs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func history(w http.ResponseWriter, req *http.Request) {
	value, err := RC.Do("GET", "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(value)

	stringReturn := value
	js, errorJs := json.Marshal(stringReturn)
	if errorJs != nil {
		http.Error(w, errorJs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
