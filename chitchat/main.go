package main

import (
	"net/http"
)

func main() {
	// マルチプレクサを生成、リクエストをハンドラにリダイレクト
	mux := http.NewServeMux()

	// 静的ファイルの返送を設定
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートURLをハンドラ関数にリダイレクトする
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
