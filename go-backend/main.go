package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	wordpress "github.com/sogko/go-wordpress"
	"github.com/spf13/viper"
	"golang/config"
	"golang/openai"
	"log"
	"os"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

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

	article, err := openai.GenerateArticle(allPrompts[0], viper.GetString("config.prompt.system-prompt"))
	if err != nil {
		panic(err)
	}

	title, err := openai.GenerateTitle(article, viper.GetString("config.prompt.title-prompt"), viper.GetString("config.prompt.system-prompt"))
	if err != nil {
		panic(err)
	}
	fmt.Println(title)
	fmt.Println(article)

	//file, err := os.ReadFile("text.txt")
	//if err != nil {
	//	panic(err)
	//}
	//article := string(file)

	post := Post{
		Title:   title,
		Content: article,
	}
	postToWordpress(post)
}

func postToWordpress(post Post) *wordpress.Post {
	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: "http://127.0.0.6:8080/wp-json/wp/v2",
		Username:   os.Getenv("WP_USERNAME"),
		Password:   os.Getenv("WP_PASSWORD"),
	})

	newPost := &wordpress.Post{
		Title:   wordpress.Title{Raw: post.Title},
		Content: wordpress.Content{Raw: post.Content},
	}
	newPost, res, _, err := client.Posts().Create(newPost)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Printf("%+v\n", post)
	return newPost
}
