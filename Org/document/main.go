package main

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
)

func main() {
	doc, err := document.Open("E:/桌面/新建文件夹/中国人民银行规章制.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	//doc.Paragraphs()得到包含文档所有的段落的切片
	for i, para := range doc.Paragraphs() {
		//run为每个段落相同格式的文字组成的片段
		fmt.Println("-----------第", i, "段-------------")
		for j, run := range para.Runs() {
			fmt.Print("\t-----------第", j, "格式片段-------------")
			fmt.Print(run.Text())

		}
		fmt.Println()
	}
}