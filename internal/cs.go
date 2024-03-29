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
		MapAliasType:    internalTypeToCSSystemType,
		MapMessageType:  internalTypeToCS,
		ProvideTemplate: csharpTemplate,
		CustomTypes:     []TypeDef{},
		Enums:           make(map[EnumName][]EnumValue),
		Messages:        make(map[MessageName][]TypeDef),
		MessageIDs:      map[MessageName]int{},
	}
}

func internalTypeToCS(t string) string {
	switch t {
	case "bool":
		return "System.Boolean"
	case "i32":
		return "int"
	case "i64":
		return "long"
	case "f32":
		return "float"
	case "f64":
		return "double"
	case "str":
		return "string"
	default:
		if isListType(t) {
			return t[1:len(t)-1] + "[]"
		}
	}

	return t
}

func internalTypeToCSSystemType(t string) string {
	switch t {
	case "bool":
		return "System.Boolean"
	case "i32":
		return "System.Int32"
	case "i64":
		return "System.Int64"
	case "f32":
		return "System.Single"
	case "f64":
		return "System.Double"
	case "str":
		return "System.String"
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
		return "ReadStringSbmf"
	default:
		return ""
	}
}

func csharpTemplate(g *Generator) (*template.Template, error) {
	t := template.New("cs-generator")

	t.Funcs(template.FuncMap{
		"readFunc": func(t string) string {
			t1 := g.findAliasType(t)
			t2 := internalTypeToCS(t1)
			t3 := csTypeToBinaryReadFuncName(t2)

			// support Message types for dictionary keys
			if t3 == "" {
				t3 = "Read" + t2
			}

			return t3
		},
		"isString": func(t string) bool {
			t = findPrimitiveType(g.findAliasType)(t)
			return t == "string"
		},
		"isStringList": func(t string, dim int) bool {
			tPrimitive := findPrimitiveType(g.findAliasType)(t)
			return tPrimitive == "string" && dim >= 1
		},
		"isPrimitive": func(t string) bool {
			t = g.findAliasType(t)
			return csTypeToBinaryReadFuncName(t) != ""
		},
		"isList":            func(i int) bool { return i >= 1 },
		"isMap":             g.isMapType,
		"isEnum":            g.isEnum,
		"isMessage":         g.isMessage,
		"findPrimitiveType": findPrimitiveType(g.findAliasType),
		"loop": func(n int) []int {
			a := make([]int, n)
			for i := range a {
				a[i] = i
			}
			return a
		},
		"loopless": func(n int) []int {
			a := make([]int, n-1)
			for i := range a {
				a[i] = i
			}
			return a
		},
		"typeDef": func(typeDef TypeDef) string {
			var t = getCsType(typeDef.Type, g)

			if typeDef.DictKey != "" {
				var k = getCsType(typeDef.DictKey, g)
				return "Dictionary<" + k + ", " + t + strings.Repeat("[]", typeDef.Dim) + ">"
			} else if typeDef.Dim > 0 {
				return t + strings.Repeat("[]", typeDef.Dim)
			}
			return t
		},
	})

	var err error
	t, err = t.Parse(templates.CS)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getCsType(typeName string, g *Generator) string {
	var t string
	if g.isEnum(typeName) || g.isMessage(typeName) {
		t = typeName
	} else if g.isCustomType(typeName) {
		t = internalTypeToCS(g.getCustomType(typeName))
	} else {
		t = internalTypeToCS(typeName)
	}

	return t
}

func findPrimitiveType(findAliasType findAliasTypeFunc) func(t string) string {
	return func(t string) string {
		t = strings.TrimSuffix(t, "[]")
		var t1 = findAliasType(t)
		var t2 = internalTypeToCS(t1)

		return t2
	}
}
