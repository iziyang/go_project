package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	// 在 gorotine 中处理每一个请求，一个占 4k，一个 4U8G 可以起几千个协程
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r) // 注册函数，阻塞函数

}

// listen->RegisterHandlers->handlers->validation{1.request, 2.user}->business logic->response.
// validation{1.request, 2.user} 内容设计
// 1. data model
// 2. error handling

//main->middleware->defs(message,err)->handlers->dbops->response
// middleware 需要做一些校验工作
