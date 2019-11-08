package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type appHander func(w http.ResponseWriter, r *http.Request) error
func wrapError(handler appHander) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := Handler(w, r)
		if err != nil {
			log.Printf("error occureend:%s",err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
			return
		}
	}
}
func Handler(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path[len("/list/"):]
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}
func main() {
	http.HandleFunc("/list/", wrapError(Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
