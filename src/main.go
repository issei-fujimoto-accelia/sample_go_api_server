package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
)

type Response struct {
	Status int  `json:"status"`
	Data string `json:"data"`
}
type ResponseMulti struct {
	Status int  `json:"status"`
	Data []User `json:"data"`
}
type User struct {
	Name string  `json:"name"`
	Age int `json:"age"`
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	return w
}

func ok(w http.ResponseWriter, req *http.Request) {
	setHeader(w)
	fmt.Println("ok!!")
	response := Response{
		Status: 200,
		Data: "ok",
	}
	json.NewEncoder(w).Encode(response)
}

func list(w http.ResponseWriter, req *http.Request) {
	setHeader(w)
	fmt.Println("list!!")
	users := []User{
		User{
			Name: "aaa",
			Age: 100,
		},
		User{
			Name: "bbb",
			Age: 200,
		},
		User{
			Name: "ccc",
			Age: 300,
		},
	}
	response := ResponseMulti{
		Status: 200,
		Data: users,
	}
	json.NewEncoder(w).Encode(response)
}

func listWithPath(w http.ResponseWriter, req *http.Request) {
	setHeader(w)
	sub := strings.TrimPrefix(req.URL.Path, "/list")
	_, id := filepath.Split(sub)
	if id == "" {
		response := ResponseMulti{
			Status: 200,
			Data: []User{},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	users := []User{
		User{
			Name: "alone!",
			Age: 100,
		},
	}
	response := ResponseMulti{
		Status: 200,
		Data: users,
	}
	json.NewEncoder(w).Encode(response)
	return
}


func serverErr(w http.ResponseWriter, req *http.Request) {
	setHeader(w)
	response := Response{
		Status: 500,
		Data: "server error",
	}
	json.NewEncoder(w).Encode(response)
}
func notFound(w http.ResponseWriter, req *http.Request) {
	setHeader(w)
	response := Response{
		Status: 404,
		Data: "not found",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("start!")
	http.HandleFunc("/ok", ok)
	http.HandleFunc("/list", list)
	http.HandleFunc("/list/", listWithPath)
	http.HandleFunc("/server-error", serverErr)
	http.HandleFunc("/not-found", notFound)
	if err := http.ListenAndServe(":8002", nil); err != nil {
		panic(err)
	}
}
