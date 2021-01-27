package parser

import (
	"crawler_go/engine"
	"regexp"
)

const cityList = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityList) //[not >]
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFun: ParseCity})
	}
	return result
}
