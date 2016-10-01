package expressgo

import (
	"log"
	"net/http"
)

// Handler - Alias of the handler func
type Handler func(w http.ResponseWriter, r *http.Request)

// MethodHandler - Struct to manage handler functions
type MethodHandler struct {
	Get       Handler
	Put       Handler
	Post      Handler
	Delete    Handler
	Option    Handler
	Head      Handler
	methodMap map[string]*Handler
}

// NewMethodHandler - Creates a pointer to a new method handler
func NewMethodHandler() *MethodHandler {
	mh := &MethodHandler{}

	mh.configure()

	return mh
}

func (mh *MethodHandler) configure() {
	mh.methodMap = map[string]*Handler{
		"GET":    &mh.Get,
		"PUT":    &mh.Put,
		"POST":   &mh.Post,
		"DELETE": &mh.Delete,
		"OPTION": &mh.Option,
		"HEAD":   &mh.Head,
	}
}

// SetMethod - Sets the handler for a method
func (mh *MethodHandler) SetMethod(method string, h Handler) {
	handler := mh.methodMap[method]
	if handler != nil {
		log.Fatalf("Method %v has already been set", method)
	}

	*handler = h
}

// GetMethod - Gets a method handler by method name
func (mh *MethodHandler) GetMethod(method string) *Handler {
	return mh.methodMap[method]
}