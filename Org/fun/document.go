package fun

import (
	"baliance.com/gooxml/document"
	"gitee.com/johng/gf/g"
	"strings"
)

func ReadWord(path string) (g.SliceStr, error) {
	doc, err := document.Open(path)
	list := g.SliceStr{}
	if err == nil {
		//doc.Paragraphs()得到包含文档所有的段落的切片
		for _, para := range doc.Paragraphs() {
			str := g.SliceStr{}
			//run为每个段落相同格式的文字组成的片段
			for _, run := range para.Runs() {
				str = append(str, run.Text())
			}
			list = append(list, strings.Join(str, ""))
		}
	}
	return list, err
}
