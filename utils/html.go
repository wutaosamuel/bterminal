package utils

import (
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// AppendHTML append content 
// the symbol is <!-- {{{ }}} -->
func AppendHTML(template, pattern string) (string, error) {
	symbol := "<!-- {{{ }}} -->"
	if !strings.Contains(template, symbol) {
		return "", Err("Template do not contain " + symbol)
	}
	pattern = pattern+"\n"+symbol+"\n"
	return strings.Replace(template, symbol, pattern, 1), nil
}

// AppendObj append obj into html and return by string
func AppendObj(body interface{}, name, pattern string) string {
	t, err := ReadHTML(name)
	CheckPanic(err)

	p, err := ReplacePattern(pattern, body)
	CheckPanic(err)

	html, err := AppendHTML(t, p)
	CheckPanic(err)
	return html
}

// UpdatePage update html after append a obj
func UpdatePage(body interface{}, name, pattern string) {
	html := AppendObj(body, name, pattern)
	err := SaveHTML(name, html)
	CheckPanic(err)
}

// ReplaceHTML check replace template
func ReplaceHTML(template string, num int, pattern string) (string, error) {
	symbol := "{{{ " + strconv.Itoa(num) + " }}}"
	if !strings.Contains(template, symbol) {
		return "", Err("Template do not contain " + symbol)
	}
	return strings.Replace(template, symbol, pattern, 1), nil
}

// ReplacePattern replace {{{ num }}} in pattern by struct
func ReplacePattern(pattern string, body interface{}) (string, error) {
	var p string
	t := reflect.TypeOf(body)
	if t.Kind() != reflect.Struct {
		return "", Err("ReplacePattern not valid type")
	}
	v := reflect.ValueOf(body)
	for i := 0; i < v.NumField(); i++ {
		if i == 0 {
			p, _ = ReplaceHTML(pattern, i+1, v.Field(i).String())
		}
		if i != 0 {
			p, _ = ReplaceHTML(p, i+1, v.Field(i).String())
		}
	}
	return p, nil
}

// SaveHTML save html file from string
func SaveHTML(name, html string) error {
	if err := ioutil.WriteFile(name, []byte(html), 0777); err != nil {
		return err
	}
	return nil
}

// ReadHTML read html file
func ReadHTML(name string) (string,error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
