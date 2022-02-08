package main
import (
	"encoding/json"
	"net/http"
)
type Response struct {
	Status int  `json:"status"`
	Message string `json:"message"`
}
func ok(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: 200,
		Message: "ok",
	}
	json.NewEncoder(w).Encode(response)
}
func serverErr(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: 500,
		Message: "server error",
	}
	json.NewEncoder(w).Encode(response)
}
func notFound(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status: 404,
		Message: "not found",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/ok", ok)
	http.HandleFunc("/server-error", serverErr)
	http.HandleFunc("/not-found", notFound)
	http.ListenAndServe(":8002", nil)
}
