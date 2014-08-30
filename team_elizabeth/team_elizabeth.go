package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "templates",
    Layout: "layout",
    Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
  }))

  m.Get("/", func() string {
    return "THIS SITE IS FOR ELIZABETH EYES ONLY!"
  })

   m.Get("/hello", func(r render.Render) {
    r.HTML(200, "hello", "elizabeth")
  })

  m.Run()
}
