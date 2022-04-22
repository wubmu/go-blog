package main

import (
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
	var indexData IndexData
	indexData.Title = "Go博客"
	indexData.Desc = "页面模板使用"
	t := template.New("index.html")

	// 1.拿到当前路径
	path, _ := os.Getwd()

	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其所有设计到的模板进行解析

	home := path + "/template/home.html"
	header := path + "/template/layout/home.html"
	footer := path + "/template/layout/home.html"
	personal := path + "/template/layout/home.html"
	post := path + "/template/layout/home.html"
	pagination := path + "/template/layout/home.html"

	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	// 页面上设计到的所有数据，必须定义

	t.Execute(w, indexData)
}

func main() {
	//程序入口，一个项目只能由一个入口
	//web程序,http协议, ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	//http.HandleFunc("/index.html", IndexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}

}
