package router

import (
	"log"
	"net/http"

	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth"
	pkgResource "github.com/DestroyerAlpha/COMSW4156-AuthDemo/pkg/resource"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user"
	"github.com/gorilla/mux"
)

func setResourceHeader(h http.Handler, resource string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("resource", resource)
		log.Printf("Setting resource header: %s", w.Header().Get("resource"))
		h.ServeHTTP(w, r)
	})
}

func applyMiddleware(h http.Handler, middlewares ...mux.MiddlewareFunc) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func SetupRoutes(middleWares ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	setupRoute := func(path string, handler http.HandlerFunc, resource string, method string) {
		log.Printf("Setting up route: %s %s", method, path) // Debug log

		wrappedHandler := setResourceHeader(handler, resource)
		log.Printf("Handler wrapped with setResourceHeader for %s", path) // Debug log

		finalHandler := applyMiddleware(wrappedHandler, middleWares...)
		log.Printf("Middlewares applied for %s", path) // Debug log

		router.Handle(path, finalHandler).Methods(method)
		log.Printf("Route added to router: %s %s", method, path) // Debug log
	}

	// Set up routes
	setupRoute("/api/auth/authenticate", auth.Authenticate, pkgResource.Authenticate, http.MethodPost)
	setupRoute("/api/user/create", user.CreateUser, pkgResource.CreateUser, http.MethodPost)
	setupRoute("/api/user/friend/add", user.AddFriend, pkgResource.AddFriend, http.MethodPost)
	setupRoute("/api/user/friend/get", user.GetFriends, pkgResource.GetFriends, http.MethodGet)
	setupRoute("/api/user/get", user.GetUser, pkgResource.GetUser, http.MethodGet)

	return router
}
