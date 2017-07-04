package main

import (
	"net/http"

	"github.com/spf13/afero"
)

func main() {
	// initialize the process list
	processList := &ProcessList{fs: afero.NewOsFs()}

	// initialize the http handler
	httpHandler := &HTTPHandler{processLister: processList}

	// register the http handler for the /ps endpoint
	http.Handle("/ps", httpHandler)

	// start the http server
	http.ListenAndServe(":8080", nil)
}
