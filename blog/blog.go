package blog

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"io/ioutil"
	"html/template"
	"path/filepath"
	"strings"
)

type Post struct {
	Date	time.Time
	Id	string
	Title	string
	Body	template.HTML
}

type Posts []Post

var (
	All 				Posts
	indexTemplate, detailTemplate 	*template.Template
	postIndex 			= make(map[string]Post)
)

func init() {
	var indexErr, detailErr error
	indexTemplate, indexErr = template.ParseFiles("tmpl/blog/index.html")
	if indexErr != nil {
		log.Fatal(indexErr)
	}

	detailTemplate, detailErr = template.ParseFiles("tmpl/blog/show.html")
	if detailErr != nil {
		log.Fatal(detailErr)
	}

	postFiles, err := filepath.Glob("tmpl/blog/posts/*.html")
	if err == nil {
		for _, path := range postFiles {
			base := filepath.Base(path)
			if len(base) > 8 {
				date := base[:8]
				id := strings.Split(base, ".")[0]
				body, ioErr := ioutil.ReadFile(path)
				if ioErr == nil {
					dateFmtd, dateErr := time.Parse("02-01-06", date)
					if dateErr == nil {
						bodySegment := strings.Split(string(body), "\n")
						post := Post{
							Date:  dateFmtd,
							Id:    id,
							Title: bodySegment[0],
							Body:  template.HTML(strings.Join(bodySegment[1:], "\n")),
						}

						All = append(All, post)
						postIndex[id] = post
					} else {
						log.Fatal("Error formatting date")
					}
				} else {
					log.Fatal("Error reading " + path)
				}
			} else {
				log.Fatal("Getting base path failed for some reason")
			}
		}
	} else {
		panic(err)
	}
}

func (p Post) DateFormatted() string {
	return p.Date.Format("02/01/06")
}

func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	indexTemplate.Execute(w, struct {
		AllPosts	Posts
	}{
		AllPosts:	All,
	})
}

func PostHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	entry, cool := postIndex[id]

	if cool {
		detailTemplate.Execute(w, struct {
			AllPosts	[]Post
			Post		Post
		}{
			AllPosts: 	All,
			Post: 		entry,
		})
	} else {
		w.WriteHeader(404)
		return
	}
}