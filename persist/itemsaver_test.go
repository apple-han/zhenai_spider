package persist

import (
	"encoding/json"
	"testing"

	"learn/crawler/engine"
	"learn/crawler/model"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func TestSaver(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	json.Unmarshal(*(resp.Source), &actual)
	actualProfile, _ := model.FromJsonObj(
		actual.Payload)
	actual.Payload = actualProfile

	if err != nil {
		panic(err)
	}
	// Verify result
	if actual != expected {
		t.Errorf("got %v; expect %v",
			actual, expected)
	}
}
