package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/helper"
	"github.com/go-park-mail-ru/2024_1_FullFocus/pkg/logger"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	elasticsetup "github.com/go-park-mail-ru/2024_1_FullFocus/internal/pkg/elasticsearch"
	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/repository/dao"
)

const (
	_categoriesLimit = 2
	_productsLimit   = 5
)

type SuggestRepo struct {
	client *elasticsearch.Client
}

func NewSuggestRepo(c *elasticsearch.Client) *SuggestRepo {
	return &SuggestRepo{
		client: c,
	}
}

func (r *SuggestRepo) GetCategorySuggests(ctx context.Context, query string) ([]models.CategorySuggest, error) {
	q := `{
		"query": {
			"match": {
				"category_name": {
					"query": "%s",
					"fuzziness": 2,
					"operator": "or"
				}
			}
		},
		"size": %d
	}`
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(elasticsetup.CategoryIndex),
		r.client.Search.WithBody(strings.NewReader(fmt.Sprintf(q, query, _categoriesLimit))))
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		logger.Error(ctx, res.String())
		return nil, fmt.Errorf("categories query error: " + res.String())
	}
	// Elasticsearch response structure
	var searchResult struct {
		Hits struct {
			Hits []struct {
				ID     string              `json:"_id"`
				Source dao.CategorySuggest `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err = json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	var categories []models.CategorySuggest
	for _, hit := range searchResult.Hits.Hits {
		categories = append(categories, models.CategorySuggest{
			ID:   hit.Source.ID,
			Name: hit.Source.Name,
		})
	}
	return categories, nil
}

func (r *SuggestRepo) GetProductSuggests(ctx context.Context, query string) ([]string, error) {
	q := `{
    "_source": "suggest",
		"suggest": {
			"product_suggest": {
				"prefix": "%s",
				"completion": {
					"field": "product_name",
					"skip_duplicates": true,
					"fuzzy": {
						"fuzziness": 2
					},
					"size": %d
				}
			}
		}
	}`
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(elasticsetup.ProductIndex),
		r.client.Search.WithBody(strings.NewReader(fmt.Sprintf(q, query, _productsLimit))))
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		logger.Error(ctx, res.String())
		return nil, fmt.Errorf("products query error: " + res.String())
	}
	// Elasticsearch response structure
	var searchResult struct {
		Suggest struct {
			ProductSuggest []struct {
				Text    string `json:"text"`
				Options []struct {
					Text string `json:"text"`
				} `json:"options"`
			} `json:"product_suggest"`
		} `json:"suggest"`
	}
	if err = json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	var products []string
	for _, suggestResult := range searchResult.Suggest.ProductSuggest {
		for _, option := range suggestResult.Options {
			products = append(products, option.Text)
		}
	}
	return helper.SortSuggests(products), nil
}
