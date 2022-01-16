package main

type IParser interface {
	Parse([]byte)
}

type jsonParser struct{}

func (p *jsonParser) Parse(d []byte) {
	panic("implement me")
}

type yamlParser struct{}

func (p *yamlParser) Parse(d []byte) {
	panic("implement me")
}

func NewIParser(t string) IParser {
	switch t {
	case "json":
		return &jsonParser{}
	case "yaml":
		return &yamlParser{}
	}
	return nil
}
