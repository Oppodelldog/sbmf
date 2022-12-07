package internal

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

type typeMapper func(t string) string
type findAliasTypeFunc func(t string) string
type templateProvider func(generator *Generator) (*template.Template, error)

type (
	EnumName  string
	EnumValue struct {
		Name  string
		Value int
	}
	TypeDef struct {
		Name         string
		Type         string
		Dim          int
		DictKey      string
		OriginalType string
	}
	MessageName string
	Generator   struct {
		Version int
		// Output file path
		Output string
		// CSharp namespace
		Namespace string
		// Go package
		Package         string
		CustomTypes     []TypeDef
		InternalTypes   []TypeDef
		Enums           map[EnumName][]EnumValue
		Messages        map[MessageName][]TypeDef
		MapAliasType    typeMapper
		MapMessageType  typeMapper
		ProvideTemplate templateProvider
		MessageIDs      map[MessageName]int
	}
)

func (g *Generator) WriteFile() error {
	fmt.Println("Writing file: " + g.Output)
	return os.WriteFile(g.Output, []byte(g.generate()), 0655)
}

func (g *Generator) addType(type_ TypeDef) {
	g.CustomTypes = append(g.CustomTypes, type_)
}

func (g *Generator) addEnum(name EnumName, values []EnumValue) {
	g.Enums[name] = values
}

func (g *Generator) addMessage(name MessageName, fields []TypeDef) {
	g.Messages[name] = fields
}

func (g *Generator) generate() string {
	t, err := g.ProvideTemplate(g)
	if err != nil {
		log.Fatal(err)
	}

	g.patchAliasTypes()
	g.patchMessageTypes()

	var sb strings.Builder
	err = t.Execute(&sb, g)
	if err != nil {
		log.Fatalf("error2: %v", err)
	}

	return sb.String()
}

func (g *Generator) patchMessageTypes() {
	for name, values := range g.Messages {
		for i, v := range values {
			g.Messages[name][i].Type = g.MapMessageType(v.Type)
		}
	}
}

func (g *Generator) patchAliasTypes() {
	for i, values := range g.CustomTypes {
		g.CustomTypes[i].Type = g.MapAliasType(values.Type)
		g.CustomTypes[i].OriginalType = values.Type
	}
}

func (g *Generator) findAliasType(t string) string {
	for _, v := range g.CustomTypes {
		if v.Name == t {
			return v.OriginalType
		}
	}

	return t
}

func (g *Generator) AddInternalType(def TypeDef) {
	g.InternalTypes = append(g.InternalTypes, def)
}

func (g *Generator) isEnum(t string) bool {
	if _, exists := g.Enums[EnumName(t)]; exists {
		return true
	}

	return false
}

func (g *Generator) isMessage(t string) bool {
	if _, exists := g.Messages[MessageName(t)]; exists {
		return true
	}

	return false
}

func (g *Generator) CreateMessageIDs() {
	var messageNames []string
	for name := range g.Messages {
		messageNames = append(messageNames, string(name))
	}
	sort.Strings(messageNames)
	for _, name := range messageNames {
		g.MessageIDs[MessageName(name)] = len(g.MessageIDs) + 1
	}
}

func (g *Generator) hasType(s string) bool {
	for _, t := range g.CustomTypes {
		if t.Name == s {
			return true
		}
	}

	return false
}

func (g *Generator) listTypes() []TypeDef {
	var types []TypeDef
	for _, t := range g.InternalTypes {
		if t.Dim > 0 {
			types = append(types, t)
		}
	}
	for _, t := range g.Messages {
		for _, f := range t {
			if f.Dim > 0 {
				types = append(types, f)
			}
		}
	}
	return types
}

func (g *Generator) mapTypes() []TypeDef {
	var types []TypeDef
	for _, t := range g.InternalTypes {
		if t.DictKey != "" {
			types = append(types, t)
		}
	}
	for _, t := range g.Messages {
		for _, f := range t {
			if f.DictKey != "" {
				types = append(types, f)
			}
		}
	}
	return types
}

func contains(dims []int, dimensions int) bool {
	for _, d := range dims {
		if d == dimensions {
			return true
		}
	}

	return false
}
