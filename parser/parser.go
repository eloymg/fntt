package parser

import (
	"errors"
	"regexp"
	"strings"
)

type Parser struct {
	Regex    *regexp.Regexp
	Regexcom *regexp.Regexp
	Line     string
	Elements []string
}

func NewParser() *Parser {
	s := &Parser{}
	return s
}

func (p *Parser) Read(s string) (r string, e error) {

	p.Line = s
	re := regexp.MustCompile(`\w+`)
	recom := regexp.MustCompile(`\/\/.*`)
	p.Regex = re
	s = recom.ReplaceAllString(s,"")
	p.Elements = re.FindAllString(s, -1)

	if len(p.Elements) == 0{
		return "delete",nil
	}

	if len(p.Elements) == 1 {
		if p.Elements[0] == "add" {
		} else {
			return "", errors.New("error arithmetic!")
		}
		code_block := dict["add"]
		return code_block, nil
	} else if len(p.Elements) == 3 {
		if p.Elements[0] == "push" && p.Elements[1] == "constant"{
			code_block := dict["push_constant"]
			code_block = strings.Replace(code_block,"{num}",p.Elements[2],-1)
			return code_block, nil
		} else {
			return "", errors.New("error memory!")
		}

		return "", nil
	} else {
		return "", errors.New("error!")
	}

}
