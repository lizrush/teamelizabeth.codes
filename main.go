package main // package delcaration

import "github.com/go-martini/martini" // framework
import "github.com/martini-contrib/render" // rendering package to make things like erbs
import "teamelizabeth.codes/string" // i wrote this package

func main() { // when we run "go run main.go" it will call this method
  m := martini.Classic() // sets up the martini framework to handle requests
  m.Use(render.Renderer(render.Options{ // tells martini to use the render package Renderer with the following options
    Directory: "templates", // for our pages that render into the yeild of the template
    Layout: "layout", // this is the main layout into which other pages render
    Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
  }))

  m.Get("/", func() string {
    return "THIS SITE IS FOR ELIZABETH EYES ONLY!"
  })

  m.Get("/hello", func(r render.Render) { //uses the render package
    r.HTML(200, "hello", "elizabeth") // specifies http code, template, content for the {{.}}
  })

  m.Get("/olleh", func(r render.Render) {
    r.HTML(200, "hello", reverse_string.Reverse("elizabeth")) // calls my reverse funtion before passing to template
    })

  m.Run()
}
