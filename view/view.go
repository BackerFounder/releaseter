package view

import (
	"bytes"

	"text/template"

	"github.com/JZGoopi/releaseter/model"
)

func GenCategoriesTemplate(lablePulls model.LablePulls) string {

	t, err := template.ParseFiles("./view/categroies.tmpl")
	if err != nil {
		panic(err)
	}

	var template bytes.Buffer
	err = t.Execute(&template, map[string]any{
		"lablePulls": lablePulls,
	})
	if err != nil {
		panic(err)
	}
	return template.String()
}
