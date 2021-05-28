package main

type IParserFactory interface {
	GenParser() IParser
}

type jsonParserFactory struct{}

func (j *jsonParserFactory) GenParser() IParser {
	return &jsonParser{}
}

type yamlParserFactory struct{}

func (j *yamlParserFactory) GenParser() IParser {
	return &yamlParser{}
}

func NewParserFactory(t string) IParserFactory {
	switch t {
	case "json":
		return &jsonParserFactory{}
	case "yaml":
		return &yamlParserFactory{}
	}
	return nil
}
