package internal

import (
	"os"
)

var (
	CONTENT_DIR = "content"
)

type Topic struct {
	Name string
	Path string
}

type Article struct {
	Name string
	Path string
}

func GetTopics() ([]string, error) {
	var topics []string

	entries, err := os.ReadDir(CONTENT_DIR)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir(){
			topics = append(topics, entry.Name())
		}
	}

	return topics, nil
}

func GetArticles(topic string) ([]string, error) {
	var articles []string

	entries, err := os.ReadDir(CONTENT_DIR + "/" + topic)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		articles = append(articles, entry.Name())
	}

	return articles, nil
}