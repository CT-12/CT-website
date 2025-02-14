package internal

import (
	"os"
	"strings"
)

var (
	CONTENT_DIR = "content"
)

type Topic struct {
	Name string
	Path string
	CreateAt string
	UpdateAt string
}

var Topics []Topic

type Article struct {
	FileName string
	Name string
	Path string
	CreateAt string
	UpdateAt string
	FrontMatterObj FrontMatter
	Markdown MarkdownContent
}

var Topic2Articles = make(map[string][]Article)

var ContentInitError error

// 取得所有主題的名字
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

// 取得某個主題下的所有文章的名字
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

func init(){
	// 初始化一些內部資料
	// 取得所有主題的名字
	topicNames, err := GetTopics()
	if err != nil {
		AddInitError(PrintErrorWithLine("Error getting names of topics -> ", err))
		return
	}
	
	for _, topicName := range topicNames {
		// 初始化 Topics 變數
		topicObj := Topic{
			Name: topicName,
			Path: "/article/" + topicName,
		}
		Topics = append(Topics, topicObj)

		// 初始化 Articles 變數
		var Articles []Article
		// 取得某個主題下的所有文章的名字
		articleNames, err := GetArticles(topicName)
		if err != nil {
			AddInitError(PrintErrorWithLine("Error getting names of articles -> ", err))
			return 
		}
		
		for _, articleName := range articleNames {
			// 萃取出 article 名稱
			fileName := articleName // E.g. fileName = 1_HelloWorld.md
			name := strings.Split(fileName, "_")[1] // E.g. name = HelloWorld.md 
			name = strings.TrimSuffix(name, ".md")  // E.g. name = HelloWorld

			// 取得文章的 front matter, markdown 內容
			filePath := CONTENT_DIR + "/" + topicName + "/" + fileName
			frontMatter, markdownContent, err := ParseMarkdown(filePath)
			if err != nil {
				AddInitError(PrintErrorWithLine("Error parsing markdown -> ", err))
				return 
			}

			if frontMatter.Draft == true {
				continue
			}

			// 建立 Article 物件
			articleObj := Article{
				FileName: fileName,
				Name: name,
				Path: "/article/" + topicName + "/" + fileName,
				CreateAt: frontMatter.CreateAt,
				UpdateAt: frontMatter.UpdateAt,
				FrontMatterObj: frontMatter,
				Markdown: markdownContent,
			}
			
			Articles = append(Articles, articleObj)
		}

		// 將 Articles 資料存入 Topic2Articles 變數
		Topic2Articles[topicName] = Articles
	}
}