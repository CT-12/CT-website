package internal

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func ConvertMdToHtml(md_content MarkdownContent) (template.HTML, error) {
	// 轉換 Markdown 到 HTML
	var buf bytes.Buffer

	converter := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // 啟用 GFM（GitHub Flavored Markdown）, for highlight code block
		), 
	)

	if err := converter.Convert([]byte(md_content), &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}