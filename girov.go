package girov

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
// 类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法
type HandlerFunc func(*Context)

// Engine is the uni handler for all requests
// Engine implement the interface of ServeHTTP
// 一个空的结构体Engine，实现了方法 ServeHTTP。
// 这个方法有2个参数，第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应
// 第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；
type Engine struct {
	// 路由映射表router
	// key由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello
	// 这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法
	//router map[string]HandlerFunc
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 增加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	c := newContext(writer, req)
	engine.router.handle(c)
}
