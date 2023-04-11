package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
	"golang/config"
	"golang/openai"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ParseConfig()

	var allPrompts []string

	RssUrls := config.GetRssUrls()
	fp := gofeed.NewParser()
	for _, RssUrl := range RssUrls {
		prompts := parseRss(fp, RssUrl)
		allPrompts = append(allPrompts, prompts...)
	}

	result, err := openai.GenerateArticle(allPrompts[0], viper.GetString("config.prompt.system-prompt"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:")
	fmt.Println(result)
}
