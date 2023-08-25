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

func (app *application) handleCustomer(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "GET":
        // get customer handler 'return app.handleGetAllCustomers(w, r)'
        // get customer by id
    case r.Method == "POST":
        // post customer handler 
    case r.Method == "PATCH": 
        // patch request
    case r.Method == "DELETE":
        // delete request
    }
}


