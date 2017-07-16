package index

import (
	"html/template"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"blog/blog"
)

var homepage *template.Template

func init() {
	var err error
	homepage, err = template.ParseFiles("tmpl/index.html")
	if err != nil {
		log.Fatal(err)
	}
}

func HomepageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	homepage.Execute(w, struct {
		AllPosts	blog.Posts
	}{
		AllPosts:	blog.All,
	})
}
