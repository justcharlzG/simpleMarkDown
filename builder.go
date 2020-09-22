package simpleMarkdown

import (
	"fmt"
	"regexp"
	"strings"
)

type Builder struct {
	title   string
	content []string
}

type HeaderLevel int

const (
	Level1 HeaderLevel = iota + 1
	Level2
	Level3
	Level4
	Level5
	Level6
)

type Row struct {
	string
}

var regWord = regexp.MustCompile(`^\w+(\s+\w+)*$`)

func (b *Builder) String() string {
	if len(b.title) == 0 {
		return fmt.Sprintf("%s", strings.Join(b.content, "\n"))
	}
	return fmt.Sprintf("%s\n%s\n", b.title, strings.Join(b.content, "\n"))
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) SetTitle(title string) *Builder {
	b.title = title
	return b
}

func (b *Builder) AddHeader(level HeaderLevel, header string) *Builder {
	b.content = append(b.content, fmt.Sprintf("%s %s\n", strings.Repeat("#", int(level)), header))
	return b
}

func (b *Builder) AddReference(row *Row) *Builder {
	b.content = append(b.content, fmt.Sprintf("> %s\n", row.string))
	return b
}

func (b *Builder) AddRow(row *Row) *Builder {
	b.content = append(b.content, fmt.Sprintf("%s\n", row.string))
	return b
}

// AddUnList 无序列表
func (b *Builder) AddUoList(list []*Row) *Builder {
	for _, r := range list {
		b.content = append(b.content, fmt.Sprintf("- %s", r.string))
	}
	b.content = append(b.content, "\n")

	return b
}

// AddList 有序列表
func (b *Builder) AddOList(list []*Row) *Builder {
	for i, r := range list {
		b.content = append(b.content, fmt.Sprintf("%d. %s", i + 1, r.string))
	}
	b.content = append(b.content, "\n")

	return b
}

func NewRow(text string) *Row {
	return &Row{
		text,
	}
}

func (r *Row) Add(text string) *Row {
	r.string += handleWord("%s", text)
	return r
}

func (r *Row) Bold(text string) *Row {
	r.string += handleWord(" **%s** ", text)
	return r
}

func (r *Row) Italic(text string) *Row {
	r.string += handleWord(" *%s* ", text)
	return r
}

func (r *Row) Picture(desc, url string) *Row {
	r.string += fmt.Sprintf("![%s](%s)", desc, url)
	return r
}

func (r *Row) Link(desc, url string) *Row {
	r.string += fmt.Sprintf("[%s](%s)", desc, url)
	return r
}

func handleWord(format, text string) string {
	if regWord.MatchString(text) {
		return fmt.Sprintf(" "+format+" ", text)
	}
	return fmt.Sprintf(format, text)
}
