package main

import (
	"crawler_go/engine"
	"crawler_go/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParseCityList,
	})
}
//	resp, err := http.Get("http://www.zhenai.com/zhenghun")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		fmt.Println("ERROR, http code =", resp.StatusCode)
//		return
//	}
//	e := determinEncoding(resp.Body)
//	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
//	all, err := ioutil.ReadAll(utf8Reader)
//	if err != nil {
//		panic(err)
//	}
//	printCityList(all)
//}
//
///**
//determine the encoding of the web content
// */
//func determinEncoding(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}
//
//func printCityList(content []byte) {
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`) //[not >]
//	matches := re.FindAllSubmatch(content, -1)
//	for _, m := range matches {
//		fmt.Printf("%s\n",m[1])
//	}
//	fmt.Println(len(matches))
//}