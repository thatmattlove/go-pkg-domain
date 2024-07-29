package main

import (
	"net/http"

	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare"
	"github.com/thatmattlove/go-pkg-domain/internal/parsing"
	"github.com/thatmattlove/go-pkg-domain/internal/response"
)

var base string
var repo string

func init() {
	b := cloudflare.Getenv("BASE")
	if b == "" {
		panic("missing BASE environment variable")
	}
	r := cloudflare.Getenv("REPO")
	if r == "" {
		panic("missing REPO environment variable")
	}
	base = b
	repo = r
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if !parsing.PartialMatch(req, base) {
			http.Redirect(w, req, repo, http.StatusPermanentRedirect)
			return
		}
		pkg, err := parsing.GetPackagePath(req, base)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		path := parsing.MakeRepoPath(repo, pkg)
		mod, err := parsing.GetModulePath(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if parsing.IsGoGet(req) {
			data := &response.Data{
				RepoPath: path,
				Package:  mod,
			}
			b, err := response.CreateResponse(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
		http.Redirect(w, req, path, http.StatusPermanentRedirect)
	})
	workers.Serve(mux)
}
