package utils

import (
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ReplaceHTML check replace template
// symbol maybe {{{num}}} || {{{ num }}}
// or {{{ num}}} || {{{num }}}
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
