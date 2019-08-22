package simpleMarkdown

import (
	"fmt"
	"os"
	"testing"
)

func TestBuilder_AddHeader(t *testing.T) {
	b := NewBuilder()
	b.AddHeader(Level1, "SimpleMarkDown").AddReference(NewRow("你好,").Bold("我们").Add("公司").Italic("欢迎各路").Bold("大神"))
	b.AddRow(NewRow("").Bold("SimpleMarkDown").Add("is a editor for").Italic("MarkDown").Bold("try it!"))
	b.AddRow(NewRow("").Bold("SimpleMarkDown").Add("是一个").Italic("简单的").Bold("markdown").Add("编辑器").Link("gitlab", "https://github.com/guowenshuai/simpleMarkDown"))

	b.AddUoList([]*Row{
		NewRow("钉钉"),
		NewRow("微信"),
		NewRow("gitlab"),
		NewRow("email"),
	})

	b.AddOList([]*Row{
		NewRow("钉钉"),
		NewRow("微信"),
		NewRow("gitlab"),
		NewRow("email"),
	})

	b.AddRow(NewRow("").Link("github", "https://github.com/guowenshuai/simpleMarkDown"))
	b.AddRow(NewRow("").Picture("cute dog", "https://www.bing.com/th?id=OIP.E_AdAbXzVnCSVPNpJRRaXQHaLF&w=118&h=177&c=7&o=5&dpr=1.25&pid=1.7"))


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
