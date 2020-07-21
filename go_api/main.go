package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Fibonacci(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if digit, err := strconv.ParseInt(ps.ByName("digit"), 10, 32); err == nil {
		// var digit int = ps.ByName("digit")
		f := fib()
		if digit < 0 {
			fmt.Fprint(w, "Error! Please use a positive number.\n")
		} else {
			fmt.Fprintf(w, "0 ")
			for n := int(digit); n > 1; n-- {
				fmt.Fprintf(w, "%d ", f())
			}
		}
	} else {
		fmt.Fprint(w, "Error! Please use a valid number.\n")
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	router := httprouter.New()
	router.GET("/api", Index)
	router.GET("/api/hello/:name", Hello)
	router.GET("/api/fibonacci/:digit", Fibonacci)
	log.Fatal(http.ListenAndServe(":8080", router))
}
