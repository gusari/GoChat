package main

import (
	"net/http"
)

func main() {
	//マルチプレクサ作成・リクエストをハンドラにリダイレクトする
	mux := http.NewServeMux()

	//静的ファイルを返却
	///static/ではじまる全てのURLから文字列/static/を取り去りディレクトリpublicを起点として残った文字列を名前にもるファイルを探す

	//指定されたディレクトリからファイルを送信するハンドラ作成
	files := http.FileServer(http.Dir("/public"))
	//ハンドラをマルチプレクサmuxのhundleに渡す・リクエストURLからPrefixを削除
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//適切なハンドラへのリダイレクト
	//同package内は参照してくれるからok
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/loogout", loogout)
	mux.HandleFunc("/authenfication", authenfication)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/creat", creatThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
