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
        <meta name="go-import" content="0.0.0.0/user/pkg %[1]s %[2]s"/>
        <meta http-equiv="refresh" content="0; url=%[2]s"/>
    </head>
    <body>
        Nothing to see here; <a href="%[2]s">move along</a>.
    </body>
</html>`

var projects = map[string]struct {
	vcs       string
	remoteURL string
}{
	"github": {
		vcs:       "git",
		remoteURL: "https://github.com/steevehook/some-test",
	},
	"gitlab": {
		vcs:       "git",
		remoteURL: "https://gitlab.com/steevehook/some-test",
	},
	"gitlab-docker": {
		vcs:       "git",
		remoteURL: "http://localhost:8000/root/some-test.git",
	},
	"bitbucket-git": {
		vcs:       "git",
		remoteURL: "https://bitbucket.org/steevehook/some-test-git",
	},
	// BitBucket says they will stop Mercurial support by Feb 2020. Be aware
	"bitbucket-hg": {
		vcs:       "hg",
		remoteURL: "https://bitbucket.org/steevehook/some-test-hg",
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := fmt.Sprintf(html, projects["github"].vcs, projects["gitlhub"].remoteURL)
		fmt.Fprint(w, tpl)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
