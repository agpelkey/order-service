package main

import "net/http"


func (app *application) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    
    data := map[string]string{
        "status": "available",
        "environment": app.config.env,
        "version": version,
    }

    err := writeJSON(w, envelope{"status": data}, http.StatusOK)
    if err != nil {
        http.Error(w, "the server encountered a problem", http.StatusInternalServerError)
        return
    }

}
