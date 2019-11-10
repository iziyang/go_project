package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 转发 web 请求
// 渲染 html 模板
// 转发 web 请求有两种模式，proxy api
/*
proxy 模式
请求：http://127.0.0.1:8000/upload
实际转发到：http://127.0.0.1:9000/upload
api 模式
请求 http://127.0.0.1/api
解析出
{
URL：“
method
message
}
*/
func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	// 会将本地的文件映射到内部的 api 127.0.0.1/statics/
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}
func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
