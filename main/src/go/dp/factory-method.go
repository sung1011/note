package main

// factory

type IParserFactory interface {
	GenParser() iParser
}

type jsonParserFactory struct{}

func (j *jsonParserFactory) GenParser() iParser {
	return &jParser{}
}

type yamlParserFactory struct{}

func (j *yamlParserFactory) GenParser() iParser {
	return &yParser{}
}

// instance

type iParser interface {
	Parse([]byte)
}

type jParser struct{}

func (j *jParser) Parse(b []byte) {

}

type yParser struct{}

func (j *yParser) Parse(b []byte) {

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
