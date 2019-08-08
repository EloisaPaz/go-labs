package main

import (
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

func main() {
	handleRequests()
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/calc/", homePage).Methods("GET")
	myRouter.HandleFunc("/calc/sum/{firstNum}/{secondNum}", sum).Methods("GET")
	myRouter.HandleFunc("/calc/sub/{firstNum}/{secondNum}", sub).Methods("GET")
	myRouter.HandleFunc("/calc/mul/{firstNum}/{secondNum}", mul).Methods("GET")
	myRouter.HandleFunc("/calc/div/{firstNum}/{secondNum}", div).Methods("GET")
	myRouter.HandleFunc("/calc/history", history).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

type Calculator struct {
	FirstNumber int
	Operator string
	SecondNumber int
	Result int
}

var History []Calculator

func getNumbers(w http.ResponseWriter, req *http.Request) (int, int){
	params := mux.Vars(req)
	firstNum, errorFirstNum := strconv.Atoi(params["firstNum"])
	secondNum, errorSecondNum := strconv.Atoi(params["secondNum"])
	if((errorFirstNum != nil) || (errorSecondNum != nil)) {
		firstNum = 0
		secondNum = 0
	}
	return firstNum, secondNum
}

func homePage(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "\nWelcome to the calculator!\n")
	fmt.Fprintf(w, "\nThe calculator supports the following operations:")
	fmt.Fprintf(w, "\n- Addition")
	fmt.Fprintf(w, "\nFor additions you must use the following URL: /calc/sum/{firstNum}/{secondNum}")
	fmt.Fprintf(w, "\n\n- Subtraction")
	fmt.Fprintf(w, "\nFor subtractions you must use the following URL: /calc/sub/{firstNum}/{secondNum}")
	fmt.Fprintf(w, "\n\n- Multiplication")
	fmt.Fprintf(w, "\nFor multiplications you must use the following URL: /calc/mul/{firstNum}/{secondNum}")
	fmt.Fprintf(w, "\n\n- Division")
	fmt.Fprintf(w, "\nFor divisions you must use the following URL: /calc/div/{firstNum}/{secondNum}")
	fmt.Fprintf(w, "\n\nYou can also check the history through the following url: /calc/history")
}

func sum(w http.ResponseWriter, req *http.Request) {
	firstNum, secondNum := getNumbers(w, req)
	result := firstNum + secondNum
	calculator := Calculator{firstNum, "+", secondNum, result}
	History = append(History, calculator)
	js, errorJs := json.Marshal(calculator)
	 if errorJs != nil {
	 	http.Error(w, errorJs.Error(), http.StatusInternalServerError)
	 	return
	}
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func sub(w http.ResponseWriter, req *http.Request) {
	firstNum, secondNum := getNumbers(w, req)
	result := firstNum - secondNum
	calculator := Calculator{firstNum, "-", secondNum, result}
	History = append(History, calculator)
	js, errorJs := json.Marshal(calculator)
	 if errorJs != nil {
	 	http.Error(w, errorJs.Error(), http.StatusInternalServerError)
	 	return
	}
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func mul(w http.ResponseWriter, req *http.Request) {
	firstNum, secondNum := getNumbers(w, req)
	result := firstNum * secondNum
	calculator := Calculator{firstNum, "*", secondNum, result}
	History = append(History, calculator)
	js, errorJs := json.Marshal(calculator)
	 if errorJs != nil {
	 	http.Error(w, errorJs.Error(), http.StatusInternalServerError)
	 	return
	}
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func div(w http.ResponseWriter, req *http.Request) {
	firstNum, secondNum := getNumbers(w, req)
	result := firstNum / secondNum
	calculator := Calculator{firstNum, "/", secondNum, result}
	History = append(History, calculator)
	js, errorJs := json.Marshal(calculator)
	 if errorJs != nil {
	 	http.Error(w, errorJs.Error(), http.StatusInternalServerError)
	 	return
	}
	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func history(w http.ResponseWriter, req *http.Request) {
	js, errorJs := json.Marshal(History)
	if errorJs != nil {
		http.Error(w, errorJs.Error(), http.StatusInternalServerError)
		return
   }
   w.Header().Set("Content-Type", "application/json")
   w.Write(js)
}