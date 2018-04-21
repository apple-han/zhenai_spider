package persist

import (
	"encoding/json"
	"testing"

	"learn/crawler/model"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func TestSaver(t *testing.T) {
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
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*(resp.Source), &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expect %v",
			actual, expected)
	}
}
