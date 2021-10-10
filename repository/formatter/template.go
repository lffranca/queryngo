package formatter

import (
	"bytes"
	"context"
	"github.com/lffranca/queryngo/pkg/util"
	"text/template"
)

func NewTemplate() *formatterTemplate {
	return &formatterTemplate{}
}

type formatterTemplate struct{}

func (repo *formatterTemplate) Transform(ctx context.Context, templateData []byte, input interface{}) ([]byte, error) {
	t, err := template.New("").Option("missingkey=zero").Funcs(template.FuncMap{
		"concatArrayString": util.ConcatArrayString,
		"toFloat":           util.InterfaceToFloat,
		"toInt":             util.InterfaceToInt,
		"toArrayString":     util.InterfaceToArrayString,
		"UniqueString":      util.UniqueString,
	}).Parse(string(templateData))
	if err != nil {
		return nil, err
	}

	tpl := new(bytes.Buffer)
	if err = t.Execute(tpl, input); err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}
