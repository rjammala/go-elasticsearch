package utils

import "strings"

var (
	// Options: boolean, enum, list, number, string, time
	typesToGo = map[string]string{
		"boolean": "*bool",
		"enum":    "string", // TODO: Custom "enum" type
		"list":    "[]string",
		"number":  "*int",
		"string":  "string",
		"time":    "time.Duration",
	}
)

func NameToGo(s string) string {
	exceptions := map[string]string{"index": "Index", "id": "DocumentID", "type": "DocumentType"}
	if value, ok := exceptions[s]; ok {
		return value
	}

	ep := strings.Split(s, "_")
	ns := make([]string, len(ep))
	for _, v := range ep {
		if v == "id" {
			v = "ID"
		}
		ns = append(ns, strings.Title(v))
	}
	return strings.Join(ns, "")
}

func TypeToGo(s string, comment ...bool) string {
	// If the string contains a pipe character, it's a polymorphic parameter,
	// ie. it takes heterogeous values, such as "boolean" and "number"
	if strings.Contains(s, "|") {
		return "interface{}"
	}

	if v, ok := typesToGo[s]; ok {
		return v
	}

	return "interface{}"
}
