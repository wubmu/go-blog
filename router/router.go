package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {

	// 路由功能

	// 1. 返回页面  views， 2.返回数据json, api 3 返回静态资源
	http.HandleFunc("/", views.HTML.Index)

	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	// 静态子资源加载
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
