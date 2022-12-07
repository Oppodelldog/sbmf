package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
					gen.addType(TypeDef{Name: k2.(string), Type: v2.(string)})
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
					gen.addEnum(EnumName(k2.(string)), values)
				}
			case "messages":
				for _, v2 := range v1.([]interface{}) {

					var fields []FieldDef
					for k3, v3 := range v2.(map[interface{}]interface{}) {
						for _, v4 := range v3.([]interface{}) {
							for k5, v5 := range v4.(map[interface{}]interface{}) {
								var t = v5.(string)
								var dim = strings.Count(t, "<")
								t = t[dim : len(t)-dim]
								fields = append(fields, FieldDef{Name: k5.(string), Type: t, Dim: dim})
							}
						}
						gen.addMessage(MessageName(k3.(string)), fields)
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
