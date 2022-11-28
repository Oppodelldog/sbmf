package internal

import (
	"sbmf/internal/templates"
	"text/template"
)

func newGoGenerator(version int, p, o string) *Generator {
	return &Generator{
		Version:         version,
		Output:          o,
		Package:         p,
		MapAliasType:    goType,
		MapMessageType:  goType,
		ProvideTemplate: goTemplate,
		Types:           []TypeDef{},
		ListTypes:       map[string][]int{},
		Enums:           make(map[EnumName][]EnumValue),
		Messages:        make(map[MessageName][]FieldDef),
		MessageIDs:      map[MessageName]int{},
	}
}

func goType(t string) string {
	switch t {
	case "bool":
		return "bool"
	case "i32":
		return "int32"
	case "i64":
		return "int64"
		//case "f32":
		//return "float32"
		//case "f64":
		//return "float64"
	case "str":
		return "string"
	default:
		if isListType(t) {
			return "[]" + t[1:len(t)-1]
		}
	}

	return t
}

func goTemplate(_ findAliasTypeFunc, _ isEnumFunc, _ isMessageFunc) (*template.Template, error) {
	t := template.New("go-generator").
		Funcs(template.FuncMap{
			"loop": func(n int) []int {
				a := make([]int, n)
				for i := range a {
					a[i] = i
				}
				return a
			},
		})

	t, err := t.Parse(templates.Go)
	if err != nil {
		return nil, err
	}

	return t, nil
}
