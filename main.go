package main

import (
  "github.com/codegangsta/martini"
  "net/http"
  "github.com/russross/blackfriday"
  gh "github.com/shurcooL/go/github_flavored_markdown"
)

func main () {
  m := martini.Classic()
  m.Post("/parse-github-markdown", func(r *http.Request) []byte {
    formBody := r.FormValue("parse")
    return gh.Markdown([]byte(formBody))
  })

  m.Post("/parse", func(r *http.Request) []byte {
    formBody := r.FormValue("parse")
    return blackfriday.MarkdownBasic([]byte(formBody))
  })

  m.Run()
}
