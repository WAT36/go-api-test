package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 共通レスポンス構造体
type response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// GET /hello?name=XXX の name を取得（未指定なら "world"）
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	json.NewEncoder(w).Encode(response{Message: "Hello, " + name + "!"})
}

func main() {
	// ルーティング設定
	http.HandleFunc("/hello", helloHandler)

	log.Println("▶︎ Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
