package gomarkdown

import (
	"fmt"
	"regexp"
)

type MarkdownInterface interface {
	DefaultTransform(text string) (string, error)
	transform(text string) (string, error)
}

const (
	VERSION = "1.0.0"
)

type Suffix string
type URL string
type Size int
type Character rune
type Gamut map[string]int
type FilterSetting func(m *Markdown)
type CallBack func(string) string

type Markdown struct {
	TabWidth                 Size  //The width of indentation of the output markup
	PredefUrls               []URL //Predefined URLs and titles for reference links and images.
	PredefTitles             []URL
	escapeChars              []Character
	nestedBracketsDepth      Size
	nestedUrlParenthsisDepth Size
	urls                     []URL
	titles                   []URL
	htmlHashes               []Character
	documentGamut            Gamut
	blockGamut               Gamut
	spanGamut                Gamut
}

func NewDefaultMarkdown() *Markdown {
	return &Markdown{
		TabWidth:                 4,
		nestedBracketsDepth:      6,
		nestedUrlParenthsisDepth: 4,
		escapeChars:              []Character{'\\', '*', '_', '{', '}', '[', ']', '(', ')', '>', '#', '+', '-', '.', '!', '`'},
		documentGamut:            Gamut{"stripLinkDefinitions": 20, "runBasicBlockGamut": 30},
		blockGamut:               Gamut{"doHeaders": 10, "doHorizontalRules": 20, "doLists": 40, "doCodeBlocks": 50, "doBlockQuotes": 60},
		spanGamut:                Gamut{"parseSpan": -30, "doImages": 10, "doAnchors": 20, "doAutoLinks": 30, "encodeAmpsAndAngles": 40, "doItalicsAndBold": 50, "doHardBreaks": 60},
	}
}

func NewMarkdown() *Markdown {
	markdown := NewDefaultMarkdown()
	// markdown.
	return markdown
}

func (m Markdown) DoHeaders(text string) string {
	pat := `^(.+?)[ ]*\n(=+|-+)[ ]*\n+`
	reg := regexp.MustCompile(pat)
	text = reg.ReplaceAllStringFunc(text, m.doHeadersCallbackSetext)
	return text
}

func (m Markdown) doHeadersCallbackSetext(text string) string {
	return text
}
