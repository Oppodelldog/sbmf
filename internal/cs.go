package internal

import (
	"sbmf/internal/templates"
	"strings"
	"text/template"
)

func newCSGenerator(version int, ns, o string) *Generator {
	return &Generator{
		Version:         version,
		Output:          o,
		Namespace:       ns,
		MapAliasType:    csAliasType,
		MapMessageType:  csType,
		ProvideTemplate: csharpTemplate,
		Types:           []TypeDef{},
		ListTypes:       map[string]string{},
		Enums:           make(map[EnumName][]EnumValue),
		Messages:        make(map[MessageName][]FieldDef),
	}
}

func csType(t string) string {
	switch t {
	case "bool":
		return "System.Boolean"
	case "i32":
		return "int"
	case "<i32>":
		return "int[]"
	case "i64":
		return "long"
	case "<i64>":
		return "long[]"
	case "f32":
		return "float"
	case "<f32>":
		return "float[]"
	case "f64":
		return "double"
	case "<f64>":
		return "double[]"
	case "str":
		return "string"
	case "<str>":
		return "string[]"
	default:
		if strings.HasPrefix(t, "<") && strings.HasSuffix(t, ">") {
			return t[1:len(t)-1] + "[]"
		}
	}

	return t
}

func csAliasType(t string) string {
	switch t {
	case "bool":
		return "System.Boolean"
	case "i32":
		return "System.Int32"
	case "<i32>":
		return "System.Int32[]"
	case "i64":
		return "System.Int64"
	case "<i64>":
		return "System.Int64[]"
	case "f32":
		return "System.Single"
	case "<f32>":
		return "System.Single[]"
	case "f64":
		return "System.Double"
	case "<f64>":
		return "System.Double[]"
	case "str":
		return "System.String"
	case "<str>":
		return "System.String[]"
	default:
		if strings.HasPrefix(t, "<") && strings.HasSuffix(t, ">") {
			return t[1:len(t)-1] + "[]"
		}
	}

	return t
}

func csTypeToBinaryReadFuncName(t string) string {
	switch t {
	case "System.Int32":
		return "ReadInt32"
	case "System.Int64":
		return "ReadInt64"
	case "System.Single":
		return "ReadSingle"
	case "System.Double":
		return "ReadDouble"
	case "System.Boolean":
		return "ReadBoolean"
	case "int":
		return "ReadInt32"
	case "long":
		return "ReadInt64"
	case "float":
		return "ReadSingle"
	case "double":
		return "ReadDouble"
	case "string":
		return "ReadString"
	default:
		return ""
	}
}

func csharpTemplate(findAliasType findAliasTypeFunc, isEnum isEnumFunc, isMessage isMessageFunc) (*template.Template, error) {
	t := template.New("cs-generator")

	t.Funcs(template.FuncMap{
		"readFunc": func(t string) string {
			t1 := findAliasType(t)
			t2 := csType(t1)
			t3 := csTypeToBinaryReadFuncName(t2)

			return t3
		},
		"isString": func(t string) bool {
			t = findPrimitiveType(findAliasType)(t)
			return t == "string"
		},
		"isStringList": func(t string) bool {
			tPrimitive := findPrimitiveType(findAliasType)(t)
			return tPrimitive == "string" && isList(t)
		},
		"isPrimitive": func(t string) bool {
			t = findAliasType(t)
			return csTypeToBinaryReadFuncName(t) != ""
		},
		"isList":            isList,
		"isEnum":            isEnum,
		"isMessage":         isMessage,
		"findPrimitiveType": findPrimitiveType(findAliasType),
	})

	var err error
	t, err = t.Parse(templates.CS)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func isList(t string) bool {
	return strings.HasSuffix(t, "[]")
}

func findPrimitiveType(findAliasType findAliasTypeFunc) func(t string) string {
	return func(t string) string {
		if strings.HasSuffix(t, "[]") {
			t = t[0 : len(t)-2]
		}
		var t1 = findAliasType(t)
		var t2 = csType(t1)
		//fmt.Printf("%s->%s->%s\n", t, t1, t2)
		return t2
	}
}
