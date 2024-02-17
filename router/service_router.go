package router

import (
	"github.com/gin-gonic/gin"
)

type EndpointsList map[string]gin.HandlerFunc
type Endpoints map[string]EndpointsList

type Router struct {
	engine *gin.Engine
	prefix string
}

func NewRouter(engine *gin.Engine, prefix string) *Router {
	return &Router{
		engine: engine,
		prefix: prefix,
	}
}

func (r *Router) Register(endpoints Endpoints) {
	r.engine.Handle(http.MethodGet, "/", Index)

	for method, endpointList := range endpoints {
		r.register(method, endpointList)
	}
}

func (r *Router) register(method string, endpoints EndpointsList) {
	for postfix, handler := range endpoints {
		r.engine.Handle(method, r.resolveUrl(postfix), handler)
	}
}

func (r *Router) resolveUrl(handlerName string) string {
	return r.prefix + handlerName
}
