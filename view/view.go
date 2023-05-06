package view

import (
	"bytes"
	"test/model"
	"text/template"
)

func GenCategoriesTemplate(lablePulls model.LablePulls) string {

	t, err := template.ParseFiles("./view/categroies.tmpl")
	if err != nil {
		panic(err)
	}

	var template bytes.Buffer
	err = t.Execute(&template, lablePulls)
	if err != nil {
		panic(err)
	}
	return template.String()
}
