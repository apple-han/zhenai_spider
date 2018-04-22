package parser

import (
	"fmt"
	"io/ioutil"
	"testing"

	"learn/crawler/engine"
	"learn/crawler/model"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserProfile(contents, "微凉", "http://album.zhenai.com/u/106573110")

	if len(result.Items) != 1 {
		t.Error("Items should contain 1 "+
			"element; but was $v", result.Items)
	}

	actual := result.Items[0]

	//profile.Name = "违心者"

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/106573110",
		Type: "zhenai",
		Id:   "106573110",
		Payload: model.Profile{
			Name:       "微凉",
			Age:        32,
			Height:     155,
			Weight:     50,
			Income:     "5001-8000元",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "销售专员",
			Hokou:      "贵州安顺",
			Car:        "未购车",
			House:      "和家人同住",
			Gender:     "女",
			Xinzuo:     "牡羊座",
		},
	}

	fmt.Println(actual)
	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}
}
