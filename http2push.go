package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	PAGE_URL   = "/"
	PAGE       = `<!DOCTYPE html>
<html>
<script type="text/javascript" src="script.js" defer></script>
<title>HTTP/2 Push example</title>
<h1>HTTP/2 Push example</h1>
<p id="main"></p>
</html>`
	SCRIPT_URL = "/script.js"
	SCRIPT = `(function() {
	document.querySelector("#main").innerHTML = 'HTTP/2 Push example content';	
})();`
)

func script(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	io.WriteString(w, SCRIPT)
}

func page(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		// HTTP/2 Push is supported
		if err := pusher.Push(SCRIPT_URL, nil); err != nil {
			fmt.Printf("Failder to push: %v\n", err)
		}
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, PAGE)
}

func main() {
	http.HandleFunc(PAGE_URL, page)
	http.HandleFunc(SCRIPT_URL, script)
	server := &http.Server {
		Addr:         ":8443",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServeTLS("server.crt", "server.key")
}
