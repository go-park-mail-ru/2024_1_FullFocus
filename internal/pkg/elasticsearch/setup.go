package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
)

const (
	_categoryIndex        = "category"
	_productIndex         = "product"
	_productIndexSettings = `
    {
        "settings": {
            "analysis": {
                "filter": {
					"shingle_filter": {
						"type": "shingle",
						"max_shingle_size": 4,
						"output_unigrams": true
					},
					"russian_morphology": {
						"type": "snowball",
						"language": "russian"
					}
            	},
				"analyzer": {
					"product_name_analyzer": {
						"type": "custom",
						"tokenizer": "whitespace",
						"filter": ["shingle_filter", "lowercase", "russian_morphology"]
					}
				}
            }
        },
        "mappings": {
            "properties": {
                "product_name": {
                    "type": "text",
                    "analyzer": "product_name_analyzer"
                }
            }
        }
    }
    `
)

type product struct {
	Name string `json:"product_name"`
}

type categoryTuple struct {
	id   uint   `db:"id"`
	name string `db:"category_name"`
}

func CreateElasticsearchIndices(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	//if err := CreateCategoryIndex(ctx, db, es); err != nil {
	//	return err
	//}
	return CreateProductIndex(ctx, db, es)
}

func CreateCategoryIndex(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	if err := DropIndex(es, _categoryIndex); err != nil {
		return err
	}
	var categories []categoryTuple
	q := `SELECT id, category_name
		  FROM category;`
	if err := db.Select(ctx, &categories, q); err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  _categoryIndex,
		Client: es,
	})
	if err != nil {
		return err
	}
	var countSuccessful uint64
	start := time.Now().UTC()

	for _, c := range categories {
		data, err := json.Marshal(c)
		if err != nil {
			log.Fatalf("Cannot encode article %s", err)
		}
		err = bi.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}
	fmt.Printf("index created (%d items) in %s\n", countSuccessful, time.Since(start))

	if err = bi.Close(ctx); err != nil {
		panic("bulk index close error: " + err.Error())
	}
	return nil
}

func CreateProductIndex(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	if err := DropIndex(es, _productIndex); err != nil {
		return err
	}
	req := esapi.IndicesCreateRequest{
		Index: _productIndex,
		Body:  strings.NewReader(_productIndexSettings),
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("creating index error: %s", res.String())
	}
	var products []string
	q := `SELECT product_name
		  FROM product;`
	if err = db.Select(ctx, &products, q); err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  _productIndex,
		Client: es,
	})
	if err != nil {
		return err
	}
	for _, p := range products {
		data, _ := json.Marshal(product{Name: p})
		if err = bi.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(data),
			},
		); err != nil {
			return err
		}
	}
	return bi.Close(ctx)
}

func DropIndex(es *elasticsearch.Client, indexName string) error {
	indexExists, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		return err
	}
	if indexExists.StatusCode == 200 {
		_, err = es.Indices.Delete([]string{indexName})
	}
	return err
}
