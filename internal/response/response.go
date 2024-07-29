package response

import (
	"fmt"
)

var _meta1 = `<meta name="go-import" content="%s git %s">`
var _meta2 = `<meta name="go-source" content="%s %s %s/tree/main/{/dir} %s/tree/main/{/dir}/{file}#L{line}">`
var _meta3 = `<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/%s">`
var _body = `<body><pre>You probably want <a href="%s">this</a>.</pre></body>`

var _full = `
  <!DOCTYPE html>
  <html>
    <head>
      %s
      %s
      %s
    </head>
    %s
  </html>
  `

type Data struct {
	RepoPath string
	Package  string
}

func CreateResponse(data *Data) ([]byte, error) {

	meta1 := fmt.Sprintf(_meta1, data.Package, data.RepoPath)
	meta2 := fmt.Sprintf(_meta2, data.Package, data.RepoPath, data.RepoPath, data.RepoPath)
	meta3 := fmt.Sprintf(_meta3, data.Package)
	body := fmt.Sprintf(_body, data.RepoPath)
	full := fmt.Sprintf(_full, meta1, meta2, meta3, body)
	return []byte(full), nil

}
