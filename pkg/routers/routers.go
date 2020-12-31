package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router ::
type Router struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

// Generate ::
func Generate() *mux.Router {
	newRouter := mux.NewRouter()

	return Configure(newRouter)
}

// Configure ::
func Configure(muxRouter *mux.Router) *mux.Router {

	for _, routerOptions := range routersPlugin {
		muxRouter.HandleFunc(routerOptions.URI, routerOptions.Func).Methods(routerOptions.Method)
	}

	return muxRouter
}
