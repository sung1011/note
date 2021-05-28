package main

//
type IParserToText interface {
	Parse([]byte)
}

//
type IParserToPic interface {
	Parse([]byte)
}

type sqlParserText struct{}

func (p *sqlParserText) Parse(d []byte) {
	panic("implement me")
}

type sqlParserPic struct{}

func (p *sqlParserPic) Parse(d []byte) {
	panic("implement me")
}

//
type IParseFactory interface {
	GenIParserToText() IParserToText
	GenIParserToPic() IParserToPic
}

type sqlParserFactory struct{}

func (pf *sqlParserFactory) GenIParserToText() IParserToText {
	return &sqlParserText{}
}

func (pf *sqlParserFactory) GenIParserToPic() IParserToPic {
	return &sqlParserPic{}
}
