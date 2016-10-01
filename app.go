package expressgo

import (
	"fmt"
	"log"
	"net/http"
)

// Get - Constant GET
const Get = "GET"

// Put - Constant PUT
const Put = "PUT"

// Post - Constant POST
const Post = "POST"

// Delete - Constant DELETE
const Delete = "DELETE"

// App - Struct to manage routing handlers
type App struct {
	handlerMap map[string]*MethodHandler
}

func (a *App) initialize() {
	a.handlerMap = make(map[string]*MethodHandler)
}

func (a *App) registerHandler(method string, route string, handler Handler) {
	methodHandler := a.handlerMap[route]
	if methodHandler == nil {
		a.handlerMap[route] = NewMethodHandler()
		methodHandler = a.handlerMap[route]
	}

	methodHandler.SetMethod(method, handler)
}

func (a *App) configureHandlers() {
	for k, v := range a.handlerMap {
		http.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			if v == nil {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			methodHandler := v.GetMethod(r.Method)
			if methodHandler == nil {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			handler := *methodHandler
			handler(w, r)
		})
	}
}

// ExpressGo - Creates a new app instance
func ExpressGo() *App {
	app := &App{}

	app.initialize()

	return app
}

// Get - Registers a GET handler for a route
func (a *App) Get(route string, handler Handler) {
	a.registerHandler(Get, route, handler)
}

// Put - Registers a PUT handler for a route
func (a *App) Put(route string, handler Handler) {
	a.registerHandler(Put, route, handler)
}

// Post - Registers a POST handler for a route
func (a *App) Post(route string, handler Handler) {
	a.registerHandler(Post, route, handler)
}

// Delete - Registers a DELETE handler for a route
func (a *App) Delete(route string, handler Handler) {
	a.registerHandler(Delete, route, handler)
}

// Start - Configures handlers for routes and starts the server
func (a *App) Start(host string, port int) {
	a.configureHandlers()

	addr := fmt.Sprintf("%v:%v", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
