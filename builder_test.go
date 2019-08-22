package simpleMarkdown

import (
	"fmt"
	"os"
	"testing"
)

func TestBuilder_AddHeader(t *testing.T) {
	b := NewBuilder()
	b.AddHeader(Level1, "SimpleMarkDown").AddReference(NewRow("你好,").Bold("我们").Add("公司").Italic("欢迎各路").Bold("大神")).AddRow(NewRow("").Bold("SimpleMarkDown").
		Add("is a editor for").Italic("MarkDown").Bold("try it!")).
		AddRow(NewRow("").Bold("SimpleMarkDown").Add("是一个").Italic("简单的").Bold("markdown").Add("编辑器").Link("gitlab", "https://github.com/guowenshuai/simpleMarkDown")).
		AddUoList([]*Row{
			NewRow("钉钉"),
			NewRow("微信"),
			NewRow("gitlab"),
			NewRow("email"),
		}).
		AddOList([]*Row{
			NewRow("钉钉"),
			NewRow("微信"),
			NewRow("gitlab"),
			NewRow("email"),
		}).
		AddRow(NewRow("").Link("github", "https://github.com/guowenshuai/simpleMarkDown")).
		AddRow(NewRow("").Picture("cute dog", "https://www.bing.com/th?id=OIP.E_AdAbXzVnCSVPNpJRRaXQHaLF&w=118&h=177&c=7&o=5&dpr=1.25&pid=1.7"))


	b.AddHeader(Level2, "北京天气").
		AddRow(NewRow("9度,").Bold("西北风1级,").Italic("空气良89，相对温度73%")).
		AddRow(NewRow("").Picture("天气","https://cms-bucket.nosdn.127.net/catchpic/a/ac/ac5ed48a4e9a9b153216af8a54c42c6e.jpg?imageView&thumbnail=550x0"))


	f, err := os.Create("README.md")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(b.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, " bytes written successfully")

}
