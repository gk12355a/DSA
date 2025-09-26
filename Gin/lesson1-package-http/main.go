package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	http.HandleFunc("/demo", demo)
	log.Println("server is starting....")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start")
	}
}

func demo(res http.ResponseWriter, req *http.Request) {
	log.Println("%v", req)
	if req.Method != http.MethodGet {
		http.Error(res, "Không hỗ trợ", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{
		"message": "Kien dang hoc Gin",
	}
	res.Header().Set("Content=type", "application/json")
	res.Header().Set("X-Course", "Gin")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(response)
}
