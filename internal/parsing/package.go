package parsing

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var withoutScheme = regexp.MustCompile(`https?\:\/\/`)

func IsGoGet(req *http.Request) bool {
	return req.URL.Query().Get("go-get") == "1"
}

func GetModulePath(req *http.Request) (string, error) {
	return req.URL.Host + req.URL.Path, nil
}

func GetPackagePath(req *http.Request, base string) (string, error) {
	if !strings.HasPrefix(base, "http") {
		base = "https://" + base
	}
	path := req.URL.Path
	if path == "" {
		return "/", nil
	}
	path = strings.ReplaceAll(path, base, "")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = strings.TrimSuffix(path, "/")
	return path, nil
}

func MakeRepoPath(repo, pkg string) string {
	repo = withoutScheme.ReplaceAllString(repo, "")
	out := strings.Join([]string{repo, pkg}, "/")
	out = strings.ReplaceAll(out, "//", "/")
	return fmt.Sprintf("https://%s", out)
}

func PartialMatch(req *http.Request, base string) bool {
	return strings.Contains(req.URL.String(), base)
}
