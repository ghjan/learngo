package view

import (
	"io"
	"text/template"

	"github.com/ghjan/learngo/crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

//CreateSearchResultView 从文件生成模板对象
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
