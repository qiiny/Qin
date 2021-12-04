//Package qin
/**
* Created with GoLand
* User: QIN
* Date: 2021/12/3
* Time: 1:54
* Description
 */
package qin

import (
	"net/http"
)

// HandlerFunc w,req参数的类型别名
type HandlerFunc func(*Context)

// Engine 结构体 一个键为字符串，值为HandlerFunc 的map
type Engine struct {
	router *router
}

// New 实例化Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET get方法
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST post方法
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// Run 启动监听
func (e *Engine) Run(address string) (err error) {
	return http.ListenAndServe(address, e)
}

//实现serveHTTP方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
