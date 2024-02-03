package api

import (
	"html/template"
	"net/http"
)

// func respondWithHTML(w http.ResponseWriter, code int, tmplString string, data any) error {
// 	tmpl, err := template.New("template").Parse(tmplString)
// 	if err != nil {
// 		return err
// 	}

// 	w.WriteHeader(code)
//  w.Header().Set("content-type", "text/html")

// 	if err := tmpl.Execute(w, data); err != nil {
// 		return err
// 	}
// 	return nil
// }

func respondWithHTMLFromFile(w http.ResponseWriter, code int, tmplPath string, data any) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	w.Header().Set("content-type", "text/html")

	if err := tmpl.Execute(w, data); err != nil {
		return err
	}
	return nil
}
