package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Handle("/v1/healthcheck", http.HandlerFunc(app.healthcheckHandler)).Methods(http.MethodGet)
	router.HandleFunc("/v1/movies", app.listCameraHandler).Methods(http.MethodGet)
	router.Handle("/v1/cameras", http.HandlerFunc(app.createCameraHandler)).Methods(http.MethodPost)
	router.Handle("/v1/cameras/{id:[0-9]+}", http.HandlerFunc(app.showCameraHandler)).Methods(http.MethodGet)
	router.HandleFunc("/v1/movies/{id:[0-9]+}", app.updateCameraHandler).Methods(http.MethodPut)
	router.HandleFunc("/v1/movies/{id:[0-9]+}", app.deleteCameraHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":4000", router)

	return app.recoverPanic(app.rateLimit(router))
}

/*

func (app *application) routes() http.Handler {
	router := mux.NewRouter()


	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Handler(http.MethodGet, "/v1/healthcheck", http.HandlerFunc(app.healthcheckHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listCameraHandler)
	router.Handler(http.MethodPost, "/v1/cameras", http.HandlerFunc(app.createCameraHandler))
	router.Handler(http.MethodGet, "/v1/cameras/:id", http.HandlerFunc(app.showCameraHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateCameraHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteCameraHandler)

	return app.recoverPanic(app.rateLimit(router))
}
*/
