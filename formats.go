package validator

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"golang.org/x/net/html"
	"gopkg.in/yaml.v2"
)

type Format interface {
	IsValidJson(str string) bool
	IsValidXml(str string) bool
	IsValidYaml(str string) bool
	IsValidHtml(str string) bool
}

func IsValidJson(str string) bool {
	return json.Valid([]byte(str))
}

func IsValidXml(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

func IsValidYaml(str string) bool {
	return yaml.Unmarshal([]byte(str), new(interface{})) == nil
}

func IsValidHtml(str string) bool {
	_, err := html.Parse(strings.NewReader(str))
	return err == nil
}
