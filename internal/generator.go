package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type typeMapper func(t string) string
type findAliasTypeFunc func(t string) string
type isEnumFunc func(t string) bool
type isMessageFunc func(t string) bool
type templateProvider func(findAliasTypeFunc, isEnumFunc, isMessageFunc) (*template.Template, error)

type Generator struct {
	Version int
	// Output file path
	Output string
	// CSharp namespace
	Namespace string
	// Go package
	Package         string
	Types           []TypeDef
	InternalTypes   []TypeDef
	Enums           map[EnumName][]EnumValue
	Messages        map[MessageName][]FieldDef
	MapAliasType    typeMapper
	MapMessageType  typeMapper
	ProvideTemplate templateProvider
	ListTypes       map[string]string
	MessageIDs      map[MessageName]int
}

func (g *Generator) AddType(type_ TypeDef) {
	g.Types = append(g.Types, type_)
}

func (g *Generator) AddEnum(name EnumName, values []EnumValue) {
	g.Enums[name] = values
}

func (g *Generator) AddMessage(name MessageName, fields []FieldDef) {
	g.Messages[name] = fields
	for _, f := range fields {
		if strings.HasPrefix(f.Type, "<") && strings.HasSuffix(f.Type, ">") {
			g.AddListType(g.MapAliasType(f.Type[1 : len(f.Type)-1]))
		}
	}
}

func (g *Generator) Generate() string {
	t, err := g.ProvideTemplate(g.findAliasType, g.isEnum, g.isMessage)
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

func (g *Generator) WriteFile() error {
	fmt.Println("Writing file: " + g.Output)
	return os.WriteFile(g.Output, []byte(g.Generate()), 0655)
}

func (g *Generator) patchMessageTypes() {
	for name, values := range g.Messages {
		for i, v := range values {
			g.Messages[name][i].Type = g.MapMessageType(v.Type)
		}
	}
}

func (g *Generator) patchAliasTypes() {
	for i, values := range g.Types {
		g.Types[i].Type = g.MapAliasType(values.Type)
		g.Types[i].OriginalType = values.Type
	}
}

func (g *Generator) findAliasType(t string) string {
	for _, v := range g.Types {
		if v.Name == t {
			return v.OriginalType
		}
	}

	return t
}

func (g *Generator) AddInternalType(def TypeDef) {
	g.InternalTypes = append(g.InternalTypes, def)
}

func (g *Generator) AddListType(aliasType string) {
	g.ListTypes[aliasType] = aliasType
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

type EnumName string
type EnumValue struct {
	Name  string
	Value int
}
type TypeName string
type TypeDef struct {
	Name         string
	Type         string
	OriginalType string
}
type MessageName string
type FieldDef struct {
	Name string
	Type string
}

func Generate(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error3: %v", err)
	}

	var gens []*Generator
	attr, hasAttrbiutes := m["attributes"].(map[interface{}]interface{})
	var version int
	if hasAttrbiutes {
		var ok bool
		version, ok = attr["version"].(int)
		if !ok {
			log.Fatal("Version missing")
		}
	} else {
		log.Fatal("attributes missing")

	}

	if v, ok := attr["csharp"]; ok {
		ns, ok := v.(map[interface{}]interface{})["namespace"].(string)
		if !ok {
			ns = ""
		}
		o, ok := v.(map[interface{}]interface{})["output"].(string)
		if !ok {
			o = "gen.cs"
		}
		err = os.MkdirAll(filepath.Dir(o), 0755)
		if err != nil {
			log.Fatal(err)
		}
		gens = append(gens, newCSGenerator(version, ns, o))
	}
	if v, ok := attr["go"]; ok {
		p, ok := v.(map[interface{}]interface{})["package"].(string)
		if !ok {
			p = ""
		}
		o, ok := v.(map[interface{}]interface{})["output"].(string)
		if !ok {
			o = "gen.go"
		}
		err = os.MkdirAll(filepath.Dir(o), 0755)
		if err != nil {
			log.Fatal(err)
		}
		gens = append(gens, newGoGenerator(version, p, o))
	}

	for _, gen := range gens {
		m = make(map[interface{}]interface{})

		err = yaml.Unmarshal(data, &m)
		if err != nil {
			log.Fatalf("error3: %v", err)
		}

		for k1, v1 := range m {
			switch k1 {
			case "types":
				var internalTypes = map[string]string{}
				for k2, v2 := range v1.(map[interface{}]interface{}) {
					gen.AddType(TypeDef{Name: k2.(string), Type: v2.(string)})
					internalTypes[v2.(string)] = v2.(string)
				}
				for internalType := range internalTypes {
					gen.AddInternalType(TypeDef{Name: internalType, Type: goType(internalType)})
				}

			case "enums":
				for k2, v2 := range v1.(map[interface{}]interface{}) {
					var values []EnumValue
					for _, v3 := range v2.([]interface{}) {
						for k4, v4 := range v3.(map[interface{}]interface{}) {
							values = append(values, EnumValue{Name: k4.(string), Value: v4.(int)})
						}
					}
					gen.AddEnum(EnumName(k2.(string)), values)
				}
			case "messages":
				for _, v2 := range v1.([]interface{}) {

					var fields []FieldDef
					for k3, v3 := range v2.(map[interface{}]interface{}) {
						for _, v4 := range v3.([]interface{}) {
							for k5, v5 := range v4.(map[interface{}]interface{}) {
								fields = append(fields, FieldDef{Name: k5.(string), Type: v5.(string)})
							}
						}
						gen.AddMessage(MessageName(k3.(string)), fields)
					}
				}
			}
		}
		gen.CreateMessageIDs()

		err = gen.WriteFile()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	}
}

func isListType(t string) bool {
	return strings.HasPrefix(t, "<") && strings.HasSuffix(t, ">")
}

func IncreaseVersion(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error3: %v", err)
	}

	a, hasAttributes := m["attributes"].(map[interface{}]interface{})
	if !hasAttributes {
		a = make(map[interface{}]interface{})
		a["version"] = 1
		m["attributes"] = a
	} else {
		v, hasVersion := a["version"].(int)
		if hasVersion {
			a["version"] = v + 1
		} else {
			a["version"] = 1
		}
		m["attributes"].(map[interface{}]interface{})["version"] = a["version"]
	}

	data, err = yaml.Marshal(m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
