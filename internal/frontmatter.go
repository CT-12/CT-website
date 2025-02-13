package internal

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	CreateAt string `yaml:"create_at"`
	UpdateAt string `yaml:"update_at"`
	Draft    bool   `yaml:"draft"`
	Tags	 []string `yaml:"tags"`
}

type MarkdownContent string

func ParseMarkdown(filePath string) (FrontMatter, MarkdownContent, error) {
	// 讀取 Markdown 檔案
	md_content, err := os.ReadFile(filePath)
	if err != nil {
		return FrontMatter{}, "", errors.New("無法讀取 Markdown 檔案: " + err.Error())
	}
	
	// 使用正則表達式提取 Front Matter
	re := regexp.MustCompile(`(?s)^---\n(.*?)\n---\n(.*)`)
	matches := re.FindStringSubmatch(string(md_content))
	if len(matches) < 3 {
		// matches[0] 是匹配到的整個字串（Front matter + Markdown 內容）
		// matches[1] 是 Front Matter
		// matches[2] 是 Markdown 內容
		return FrontMatter{}, "", errors.New("invalid markdown format") 
	}

	// 解析 Front Matter (yaml)
	var fm FrontMatter
	err = yaml.Unmarshal([]byte(matches[1]), &fm)
	if err != nil {
		return FrontMatter{}, "", errors.New("無法解析 Front Matter " + filePath + ": " + err.Error())
	}

	// 剩下的是 Markdown 內容
	markdownContent := strings.TrimSpace(matches[2])

	return fm, MarkdownContent(markdownContent), nil
}