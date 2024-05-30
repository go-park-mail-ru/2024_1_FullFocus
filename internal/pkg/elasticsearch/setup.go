package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"golang.org/x/sync/errgroup"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/database"
)

const (
	CategoryIndex = "category"
	ProductIndex  = "product"
)

const (
	_categoryIndexSettings = `
	{
		"settings": {
			"analysis": {
				"filter": {
					"russian_morphology": {
						"type": "snowball",
						"language": "russian"
					},
					"edge_ngram_filter": {
						"type": "edge_ngram",
						"min_gram": 1,
						"max_gram": 15,
						"preserve_original": true
					}
				},
				"analyzer": {
					"category_name_analyzer": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": ["lowercase", "russian_morphology", "edge_ngram_filter"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"category_id": {
					"type": "long"
				},
				"category_name": {
					"type": "text",
					"analyzer": "category_name_analyzer",
					"fields": {
						"keyword": {
							"type": "keyword",
							"ignore_above": 256
						}
					}
				}
			}
		}
	}
	`
	_productIndexSettings = `
    {
        "settings": {
            "analysis": {
                "filter": {
					"russian_morphology": {
						"type": "snowball",
						"language": "russian"
					}
            	},
				"analyzer": {
					"product_name_analyzer": {
						"type": "custom",
						"tokenizer": "whitespace",
						"filter": ["lowercase", "russian_morphology"]
					}
				}
            }
        },
        "mappings": {
			"properties": {
				"product_name": {
					"type": "completion",
					"analyzer": "product_name_analyzer"
				}
			}
		}
    }
    `
)

const _ngramWordsAmount = 4

type product struct {
	Name string `json:"product_name"`
}

type category struct {
	ID   uint   `json:"category_id" db:"id"`
	Name string `json:"category_name" db:"category_name"`
}

func InitElasticData(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	g := errgroup.Group{}
	g.Go(func() error {
		return initCategoryIndex(ctx, db, es)
	})
	g.Go(func() error {
		return initProductIndex(ctx, db, es)
	})
	return g.Wait()
}

// Category

func initCategoryIndex(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	if err := createCategoryIndex(ctx, es); err != nil {
		return err
	}
	return insertCategories(ctx, db, es)
}

func createCategoryIndex(ctx context.Context, es *elasticsearch.Client) error {
	if err := dropIndex(es, CategoryIndex); err != nil {
		return err
	}
	req := esapi.IndicesCreateRequest{
		Index: CategoryIndex,
		Body:  strings.NewReader(_categoryIndexSettings),
	}
	res, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("creating index error: %s", res.String())
	}
	return nil
}

func insertCategories(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	if err := dropIndex(es, CategoryIndex); err != nil {
		return err
	}
	var categories []category
	q := `SELECT id, category_name
		  FROM category;`
	if err := db.Select(ctx, &categories, q); err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  CategoryIndex,
		Client: es,
	})
	if err != nil {
		return err
	}
	for _, c := range categories {
		data, _ := json.Marshal(c)
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

// Product

func initProductIndex(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	if err := createProductIndex(ctx, es); err != nil {
		return err
	}
	return insertProducts(ctx, db, es)
}

func createProductIndex(ctx context.Context, es *elasticsearch.Client) error {
	if err := dropIndex(es, ProductIndex); err != nil {
		return err
	}
	req := esapi.IndicesCreateRequest{
		Index: ProductIndex,
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
	return nil
}

func insertProducts(ctx context.Context, db database.Database, es *elasticsearch.Client) error {
	var products []string
	q := `SELECT product_name
		  FROM product;`
	if err := db.Select(ctx, &products, q); err != nil {
		return err
	}
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  ProductIndex,
		Client: es,
	})
	if err != nil {
		return err
	}
	for _, p := range products {
		for _, productNameSubstring := range ProductNameNgrams(p, _ngramWordsAmount) {
			data, _ := json.Marshal(product{Name: productNameSubstring})
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
	}
	return bi.Close(ctx)
}

func dropIndex(es *elasticsearch.Client, indexName string) error {
	indexExists, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		return err
	}
	if indexExists.StatusCode == 200 {
		_, err = es.Indices.Delete([]string{indexName})
	}
	return err
}
