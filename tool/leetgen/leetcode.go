package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// generateFile generate go code file
func generateFile(rootPath, titleSlug string) error {
	question, err := getProblemItem(titleSlug)
	if err != nil {
		return err
	}

	// 文件夹名采用小写字母开头的驼峰风格
	s := strings.Replace(question.TitleSlug, "-", " ", -1)
	s = cases.Title(language.Und).String(s)
	s = strings.Replace(s, " ", "", -1)
	dirName := string(unicode.ToLower(rune(s[0]))) + s[1:]

	// 文件夹路径
	dirPath := path.Join(rootPath, dirName)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	// 获取go代码片段
	var goSnippet string
	for _, snippet := range question.CodeSnippets {
		if snippet.Lang == "Go" {
			goSnippet = snippet.Code
			break
		}
	}

	// 生成go文件
	fileHeader := fmt.Sprintf(`// Source : https://leetcode.cn/problems/%s
// Date   : %s

`, question.TitleSlug, time.Now().Format("2006-01-02"))
	questionContent := getQuestionContent(question.TranslatedContent)
	commentContent := makeComment(questionContent)
	goContent := fileHeader + commentContent + fmt.Sprintf("\npackage %s\n\n", dirName) + goSnippet
	fileName := dirName + ".go"
	err = os.WriteFile(path.Join(dirPath, fileName), []byte(goContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// getQuestionContent make html string into text
func getQuestionContent(str string) string {
	// 去除html
	reg := regexp.MustCompile("<[^<]*>")
	str = reg.ReplaceAllLiteralString(str, "")
	// 多行空行转为一行
	reg = regexp.MustCompile("\n\n")
	str = reg.ReplaceAllLiteralString(str, "\n")
	// NBSP转为空格
	reg = regexp.MustCompile("\u00a0")
	str = reg.ReplaceAllLiteralString(str, " ")
	// 特殊字符替换
	str = strings.ReplaceAll(str, "&nbsp;", " ")
	str = strings.ReplaceAll(str, "&amp;", "&")
	str = strings.ReplaceAll(str, "&lt;", "<")
	str = strings.ReplaceAll(str, "&gt;", ">")
	str = strings.ReplaceAll(str, "&quot;", "\"")
	str = strings.ReplaceAll(str, "&ldquo;", "\"")
	return str
}

// makeComment make question desc into comment
func makeComment(str string) string {
	str = "/**********************************************************************************\n" + str
	str = str + "**********************************************************************************/\n"
	return str
}
