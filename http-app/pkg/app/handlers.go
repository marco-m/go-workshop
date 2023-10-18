package app

import "net/http"

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// TODO add middleware for logging the request! app.log.Info()
	
	w.Write([]byte("OK"))
}
