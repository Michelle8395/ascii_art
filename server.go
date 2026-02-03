package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Allow safe rendering of ASCII (preserve spaces + newlines)
var funcMap = template.FuncMap{
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
}

var pageTpl = template.Must(template.New("page").Funcs(funcMap).Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>ASCII Art Web</title>
	<style>
		body { font-family: monospace; }
		pre {
			background: #f4f4f4;
			padding: 10px;
			overflow-x: auto;
		}
	</style>
</head>
<body>
	<h1>ASCII Art Web</h1>

	<form method="POST">
		<label>Text:</label><br>
		<textarea name="text" rows="4" cols="50">{{.Input}}</textarea><br><br>

		<label>Banner:</label><br>
		<select name="banner">
			<option value="standard.txt" {{if eq .Banner "standard.txt"}}selected{{end}}>standard</option>
			<option value="shadow.txt" {{if eq .Banner "shadow.txt"}}selected{{end}}>shadow</option>
			<option value="thinkertoy.txt" {{if eq .Banner "thinkertoy.txt"}}selected{{end}}>thinkertoy</option>
		</select><br><br>

		<input type="submit" value="Render ASCII">
	</form>

	{{if .Result}}
	<h2>Result:</h2>
	<pre>{{.Result | safeHTML}}</pre>
	{{end}}
</body>
</html>
`))

type PageData struct {
	Result string
	Input  string
	Banner string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Banner: "standard.txt", // default selection
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		data.Input = text
		data.Banner = banner

		result, err := RenderAscii(text, banner)
		if err != nil {
			data.Result = fmt.Sprintf("Error: %s", err)
		} else {
			data.Result = result
		}
	}

	if err := pageTpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
