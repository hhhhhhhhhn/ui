package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

type Font graphics.Struct_SS_sfFont

func LoadFont(file string) Font {
	return graphics.SfFont_createFromFile(file)
}

type TextWidget interface {
	Widget
	SetFont(Font)      Widget
	SetContent(string) Widget
	SetFontSize(uint)  Widget
	SetColor(Color)    Widget
}

type Text struct {
	TextWidget
	text      graphics.Struct_SS_sfText
	position  system.SfVector2f
	redraw    bool
	width     float32
	fontSize  uint
	content   string
	font      Font
}

func (t *Text) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
	width float32, height float32) {
		if width != t.width || t.redraw {
			graphics.SfText_setString(
				t.text,
				wrapWords(t.content, t.font, t.fontSize, width),
			)
			t.width = width
			t.redraw = false
		}

		t.position.SetX(x)
		t.position.SetY(y)
		graphics.SfText_setPosition(t.text, t.position)

		graphics.SfRenderWindow_drawText(w, t.text, graphics.SwigcptrSfRenderStates(0))
}

func (t *Text) Clean() {
	graphics.SfText_destroy(t.text)
	system.DeleteSfVector2f(t.position)
}

func (t *Text) Init() ([]string, []func(Event)) {
	return []string{}, []func(Event){}
}

func (t *Text) SetFont(font Font) *Text {
	graphics.SfText_setFont(t.text, font)
	return t
}

func (t *Text) SetContent(content string) *Text {
	t.content = content
	t.redraw = true
	return t
}

func (t *Text) SetFontSize(fontSize uint) *Text {
	graphics.SfText_setCharacterSize(t.text, fontSize)
	return t
}
func (t *Text) SetColor(color Color) *Text {
	graphics.SfText_setFillColor(t.text, color.ToSFColor())
	return t
}

func NewText(font Font, fontSize uint) *Text {
	text := graphics.SfText_create()
	graphics.SfText_setFont(text, font)
	graphics.SfText_setCharacterSize(text, fontSize)

	position := system.NewSfVector2f()
	return &Text{text: text, position: position, fontSize: fontSize, font: font}
}

func wrapWords(content string, font Font, fontSize uint, width float32) string {
	newContent := []rune(content)
	xPosition := float32(0) // Relative to the top-left corner
	wordIndex := 0 // Relative to the current line
	lastSpaceIndex := 0

	for i := 0; i < len(newContent); i++ {
		switch (newContent[i]) {
		case '\n':
			xPosition = 0
			wordIndex = 0
			break
		case ' ':
			lastSpaceIndex = i
			wordIndex++
			break
		default:
			xPosition += graphics.SfFont_getGlyph(font, uint(newContent[i]), fontSize, 0, 0).GetAdvance()
			if i != 0 {
				xPosition += graphics.SfFont_getKerning(font, uint(newContent[i-1]), uint(newContent[i]), fontSize)
			}

			if (wordIndex != 0 && xPosition > width) {
				newContent[lastSpaceIndex] = '\n'
				wordIndex = 0
				xPosition = 0
				i = lastSpaceIndex
			}
		}
	}

	return string(newContent)
}
