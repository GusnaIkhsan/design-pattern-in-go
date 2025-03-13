package main

import (
	"fmt"
	"sync"
)

type TextStyle struct {
	name string
}

func (t *TextStyle) Format(text string) string {
	return fmt.Sprintf("[%s %s]", t.name, text)
}

type TextStyleFactory struct {
	styles map[string]*TextStyle
	mu     sync.Mutex
}

func (f *TextStyleFactory) GetTextStyle(name string) *TextStyle {
	f.mu.Lock()
	defer f.mu.Unlock()

	if style, ok := f.styles[name]; ok {
		return style
	}

	style := &TextStyle{name: name}
	f.styles[name] = style
	return style
}

type TextRange struct {
	Start, End int
	Text       string
	Style      *TextStyle
}

func NewTextRange(start, end int, text string, style *TextStyle) *TextRange {
	return &TextRange{
		Start: start,
		End:   end,
		Text:  text,
		Style: style,
	}
}

func (t *TextRange) Format() string {
	return t.Style.Format(t.Text)
}

func main() {
	factory := &TextStyleFactory{
		styles: make(map[string]*TextStyle),
	}

	textRange1 := NewTextRange(0, 5, "Hello", factory.GetTextStyle("Bold"))
	textRange2 := NewTextRange(6, 11, "World", factory.GetTextStyle("Bold"))

	fmt.Println(textRange1.Format())
	fmt.Println(textRange2.Format())
}
