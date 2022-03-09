package util

import (
	"fmt"
	"github.com/ghjan/learngo/pkg/utils"
	"go.uber.org/zap/buffer"
	"strings"
	"text/template"
)

func RenderStringWithTemplateString(templateString string,
	context interface{}) (*template.Template, string, error) {
	tmpl, err := template.New("test").Parse(templateString)
	if err != nil {
		panic(err)
	}
	buffer := buffer.Buffer{}
	err = tmpl.Execute(&buffer, context)
	result := ""
	if err == nil {
		result = buffer.String()
	}
	return tmpl, result, err
}

func RenderStringWithTemplate(tmpl *template.Template, context interface{}) (result string, err error) {
	buffer := buffer.Buffer{}
	err = tmpl.Execute(&buffer, context)
	if err == nil {
		result = buffer.String()
	}
	return
}

func RenderString4Object(object interface{}, needRenderFields string, context interface{}) (err error) {
	if needRenderFields != "" {
		fieldNames := strings.Split(needRenderFields, ",")
		for _, fieldName := range fieldNames {
			var (
				result string
			)
			fieldName = strings.TrimSpace(fieldName)
			value := ReflectGetFieldValue(object, fieldName)
			_, result, err = RenderStringWithTemplateString(value.(string), context)
			ReflectSetFieldValue(object, fieldName, result)
			msg := fmt.Sprintf("processBasicSheetItem. RenderString4Object error,%s:%v", fieldName, result)
			utils.Checkerr2(err, msg)
		}
	}
	return
}
