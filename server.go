package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Input   string
	Banner  string
	Result  string
	Error   string
	Banners []string
}

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>ASCII Art Web</title>
	<style>
		body {
			background: #0f172a;
			color: #e5e7eb;
			font-family: system-ui, sans-serif;
		}
		.container {
			max-width: 900px;
			margin: 40px auto;
			padding: 30px;
			background: #020617;
			border-radius: 12px;
		}
		h1 {
			text-align: center;
			color: #38bdf8;
		}
		label {
			font-weight: bold;
		}
		textarea, select {
			width: 100%;
			margin-top: 8px;
			margin-bottom: 20px;
			padding: 10px;
			border-radius: 6px;
			border: none;
			background: #020617;
			color: #e5e7eb;
			box-shadow: inset 0 0 0 1px #1e293b;
			font-family: monospace;
		}
		input[type=submit] {
			width: 100%;
			padding: 12px;
			border-radius: 6px;
			border: none;
			background: #38bdf8;
			color: #020617;
			font-weight: bold;
			cursor: pointer;
		}
		pre {
			margin-top: 30px;
			padding: 20px;
			background: #020617;
			border-radius: 12px;
			box-shadow: inset 0 0 0 1px #1e293b;
			white-space: pre;
			overflow-x: auto;
		}
		.error {
			color: #f87171;
			font-weight: bold;
		}
	</style>
</head>
<body>
<div class="container">
	<h1>ASCII Art Web</h1>

	<form method="POST">
		<label>Text</label>
		<textarea name="text" rows="4">{{.Input}}</textarea>

		<label>Banner</label>
		<select name="banner">
			{{range .Banners}}
				<option value="{{.}}" {{if eq . $.Banner}}selected{{end}}>{{.}}</option>
			{{end}}
		</select>

		<input type="submit" value="Render">
	</form>

	{{if .Error}}
		<p class="error">{{.Error}}</p>
	{{end}}

	{{if .Result}}
		<pre>{{.Result}}</pre>
	{{end}}
</div>
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Banners: []string{"standard.txt", "shadow.txt", "thinkertoy.txt"},
		Banner:  "standard.txt",
	}

	if r.Method == http.MethodPost {
		data.Input = r.FormValue("text")
		data.Banner = r.FormValue("banner")

		if data.Input == "" {
			data.Error = "Please enter text"
		} else {
			result, err := RenderAscii(data.Input, data.Banner)
			if err != nil {
				data.Error = err.Error()
			} else {
				data.Result = result
			}
		}
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
