package controller

import (
	"encoding/json"
	"fmt"
	"frontend/model"
	"frontend/view"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

type SearchResultHandler struct {
	View   view.SearchResultView
	Client *elastic.Client
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := s.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.View.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.Client.Search().
		RestTotalHitsAsInt(true).
		Index("zhenaiwang").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Query = q
	result.Items = ParseItemFromJson(resp.Each(reflect.TypeOf(model.Item{})))
	return result, nil
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	return SearchResultHandler{
		View:   view.CreateSearchResultView(template),
		Client: client,
	}
}

func ParseItemFromJson(j []interface{}) []model.Item {
	var items []model.Item
	data, err := json.Marshal(j)
	if err != nil {
		log.Println("marshal error:", err)
		return items
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return items
	}

	return items
}
