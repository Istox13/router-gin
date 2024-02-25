package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EndpointsList map[string]gin.HandlerFunc
type Endpoints map[string]EndpointsList

type Router struct {
	engine *gin.Engine
	prefix string
}

func NewRouter(engine *gin.Engine, prefix string) *Router {
	router := &Router{
		engine: engine,
		prefix: prefix,
	}

	router.engine.Handle(http.MethodGet, "/", Index)

	return router
}

func (r *Router) Register(endpoints Endpoints) {
	for method, endpointList := range endpoints {
		r.register(method, endpointList)
	}
}

func (r *Router) AddHealthcheckURL(url string) {
	r.engine.Handle(http.MethodGet, url, Index)
}

func (r *Router) register(method string, endpoints EndpointsList) {
	for postfix, handler := range endpoints {
		r.engine.Handle(method, r.resolveUrl(postfix), handler)
	}
}

func (r *Router) resolveUrl(handlerName string) string {
	return r.prefix + handlerName
}
