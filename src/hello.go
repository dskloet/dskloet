package sample

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type TemplateParameters struct {
	Name string
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	tpl, err := template.ParseFiles("templates/hello.html")

	if err != nil {
		fmt.Fprintf(writer, "Error parsing template: %v", err)
		return
	}

	var params TemplateParameters
	params.Name = request.FormValue("name")

	buf := bytes.NewBuffer(make([]byte, 0, 0))
	err = tpl.Execute(buf, &params)
	if err != nil {
		fmt.Fprintf(writer, "Error executing template: %v", err)
		return
	}

	writer.Write(buf.Bytes())
}
