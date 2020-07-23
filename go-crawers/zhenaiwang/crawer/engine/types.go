package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url    string
	Type   string
	Id     int
	Upload interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
