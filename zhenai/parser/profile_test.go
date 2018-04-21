package parser

import (
	"io/ioutil"
	"testing"

	"learn/crawler/model"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserProfile(contents, "违心者")

	if len(result.Items) != 1 {
		t.Error("Items should contain 1 "+
			"element; but was $v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	//profile.Name = "违心者"

	expected := model.Profile{
		Name:       "违心者",
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
		Xinzou:     "牡羊座",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v",
			expected, profile)
	}
}
