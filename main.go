package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	//返回json数据
	w.Header().Set("Content-Type", "application/json")

	var indexData IndexData
	indexData.Title = "Go博客"
	indexData.Desc = "现在刚入门"

	jsonStr, _ := json.Marshal(indexData)

	w.Write(jsonStr)
}

func IndexHtml(w http.ResponseWriter, request *http.Request) {
	var indexData IndexData
	indexData.Title = "Go博客"
	indexData.Desc = "页面模板使用"
	t := template.New("index.html")

	// 1.拿到当前路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	//程序入口，一个项目只能由一个入口
	//web程序,http协议, ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	http.HandleFunc("/index.html", IndexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}

}
