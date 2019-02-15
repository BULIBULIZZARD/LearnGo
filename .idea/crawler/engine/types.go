package engine


type Request struct {
	Url stirng
	ParserFunf func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Items []interface{}
}