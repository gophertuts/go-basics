package main

import (
	"fmt"
	"log"
	"net/http"
)

const html = `<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="go-import" content="0.0.0.0/user/pkg git https://github.com/steevehook/some-test">
        <meta http-equiv="refresh" content="0; url=https://github.com/steevehook/some-test">
    </head>
    <body>
        Nothing to see here; <a href="https://github.com/steevehook/some-test">move along</a>.
    </body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
