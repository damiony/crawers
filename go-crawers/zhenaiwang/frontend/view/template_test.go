package view

import (
	"frontend/model"
	"log"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")

	page := model.SearchResult{}

	page.Hits = 123

	item := model.Item{
		Url:  "http://album.zhenai.com/u/1214814888",
		Type: "zhenai",
		Id:   1214814888,
		Upload: model.Profile{
			NickName:      "yyy",
			Sex:           1,
			Age:           30,
			Height:        166,
			Constellation: "天蠍座",
			Education:     "大專",
			Marriage:      "未婚",
		},
	}

	for i := 1; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		log.Fatal(err)
	}
}
