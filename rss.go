package main

import (
	"bytes"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
	"golang/util"
	"text/template"
)

func parseRss(fp *gofeed.Parser, RssUrl string) []string {
	var result []string
	if RssUrl != "" {
		feed, err := fp.ParseURL(RssUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Source: ", feed.Title)
		for _, item := range feed.Items {
			item.Description = util.FormatRSSContent(item.Description)
			item.Content = util.FormatRSSContent(item.Content)
			upTmpl := template.Must(template.New("user-prompt").Parse(viper.GetString("config.prompt.user-prompt")))
			userPrompt := new(bytes.Buffer)
			upErr := upTmpl.Execute(userPrompt, item)
			if upErr != nil {
				panic(upErr)
			}
			result = append(result, userPrompt.String())
		}
	}
	return result
}
