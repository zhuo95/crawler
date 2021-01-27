package parser

import (
	"crawler_go/engine"
	"regexp"
)
const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe) //[not >]
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFun: engine.NilParser})
	}
	return result
}