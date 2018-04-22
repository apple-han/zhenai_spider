package persist

import (
	"log"
	//"net/http"

	"learn/crawler/engine"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(
	index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}

		}
	}()
	return out, nil
}

func save(
	client *elastic.Client,
	index string,
	item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	// %+v 是打印结构体的结构
	//fmt.Printf("%+v", resp)
	return nil
}
