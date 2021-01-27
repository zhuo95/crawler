package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("test_parseCityList_file.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	//verify
	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d : %s, but got %s", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d : %s, but got %s", i, city, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items, but had %d", resultSize, len(result.Items))
	}
}