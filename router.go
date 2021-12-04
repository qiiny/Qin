//Package qin
/**
* Created with GoLand
* User: QIN
* Date: 2021/12/3
* Time: 22:08
* Description
 */
package qin

import (
	"log"
	"net/http"
)

//路由的结构体
type router struct {
	handlers map[string]HandlerFunc
}

//创建路由结构体
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

//添加路由
func (r *router) addRoute(method string, patten string, handler HandlerFunc) {
	log.Printf("route %4s - %s", method, patten)
	key := method + "-" + patten
	r.handlers[key] = handler
}

//处理路由
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404没有找到:%s\n", c.Path)
	}
}
