package storyHandler

import (
	"net/http"
	"text/template"

	"github.com/gopher-training/choose_your_own_adventure/structs"
)

var defaultTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
        <h1>{{.ArcData.Title}}</h1>
        {{range .ArcData.Story}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .ArcData.Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>
            {{end}}
        </ul>
	</body>
	<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
</html>`

type handler struct {
	s map[string]structs.ArcStory
}

//StoryHandler provides us a Handler interface to put our story on some socket
func StoryHandler(s map[string]structs.ArcStory) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templateStory := template.Must(template.New("").Parse(defaultTemplate))
	nextArc := r.URL.Path
	if nextArc == "" || nextArc == "/" {
		nextArc = "/intro"
	}
	err := templateStory.Execute(w, h.s[nextArc[1:]])
	if err != nil {
		panic(err)
	}
}
