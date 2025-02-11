package internal

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func ConvertMdToHtml(topic string, article string) (template.HTML, error) {
	md_path := filepath.Join(CONTENT_DIR, topic, article)

	// 讀取 Markdown 檔案
	md_content, err := os.ReadFile(md_path)
	if err != nil {
		return "", err
	}

	// 轉換 Markdown 到 HTML
	var buf bytes.Buffer

	converter := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // 啟用 GFM（GitHub Flavored Markdown）, for highlight code block
		), 
	)

	if err := converter.Convert(md_content, &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}