package main

import (
	"strings"

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
	text      graphics.Struct_SS_sfText
	position  system.SfVector2f
	redraw    bool
	width     float32
	height    float32
	fontSize  uint
	content   string
	font      Font
}

func (t *Text) Draw(w Texture, x float32, y float32,
	width float32, height float32) {
		if width != t.width || height != t.height || t.redraw {
			graphics.SfText_setString(
				t.text,
				wrapWords(t.content, t.font, t.fontSize, width, height),
			)
			t.width = width
			t.height = height
			t.redraw = false
		}

		t.position.SetX(x)
		t.position.SetY(y)
		graphics.SfText_setPosition(t.text, t.position)

		graphics.SfRenderTexture_drawText(w, t.text, graphics.SwigcptrSfRenderStates(0))
}

func (t *Text) Clean() {
	graphics.SfText_destroy(t.text)
	system.DeleteSfVector2f(t.position)
}

func (t *Text) Init() ([]string, []func(Event)) {
	return []string{}, []func(Event){}
}

func (t *Text) SetFont(font Font) *Text {
	t.font = font
	graphics.SfText_setFont(t.text, font)
	return t
}

func (t *Text) SetContent(content string) *Text {
	t.content = content
	t.redraw = true
	return t
}

func (t *Text) SetFontSize(fontSize uint) *Text {
	t.fontSize = fontSize
	graphics.SfText_setCharacterSize(t.text, fontSize)
	return t
}

func (t *Text) SetColor(color Color) *Text {
	graphics.SfText_setFillColor(t.text, color.ToSFColor())
	return t
}

func NewText() *Text {
	text := graphics.SfText_create()

	position := system.NewSfVector2f()
	return &Text{text: text, position: position}
}

func wrapWords(content string, font Font, fontSize uint, width, height float32) string {
	newContent := []rune(content)
	lineSpacing := graphics.SfFont_getLineSpacing(font, fontSize)
	xPosition := float32(0) // Relative to the top-left corner
	yPosition := lineSpacing
	wordIndex := 0 // Relative to the current line
	lastSpaceIndex := 0

	for i := 0; i < len(newContent); i++ {
		if float32(yPosition) > height {
			return string(newContent[:i])
		}
		switch (newContent[i]) {
		case '\n':
			xPosition = 0
			yPosition += lineSpacing
			wordIndex = 0
			break
		case ' ':
			lastSpaceIndex = i
			wordIndex++
			xPosition += graphics.SfFont_getGlyph(font, uint(newContent[i]), fontSize, 0, 0).GetAdvance()
			if i != 0 {
				xPosition += graphics.SfFont_getKerning(font, uint(newContent[i-1]), uint(newContent[i]), fontSize)
			}
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
				yPosition += lineSpacing
				i = lastSpaceIndex
			}
		}
	}

	return string(newContent)
}

type AdvancedText struct {
	texts     []graphics.Struct_SS_sfText
	position  system.SfVector2f
	width     float32
	height    float32
	alignment float32
	fontSize  uint
	redraw    bool
	content   string
	color     Color
	font      Font
}

func NewAdvancedText() *AdvancedText {
	position := system.NewSfVector2f()
	return &AdvancedText{position: position, redraw: true}
}

func (a *AdvancedText) Draw(w Texture, x float32, y float32,
	width float32, height float32) {
		if width != a.width || height != a.height || a.redraw {
			a.update(width, height)
		}

		a.position.SetX(x)
		a.position.SetY(y)

		for _, text := range a.texts {
			graphics.SfText_setPosition(text, a.position)
			graphics.SfRenderTexture_drawText(w, text, graphics.SwigcptrSfRenderStates(0))
		}

}

func (a *AdvancedText) update(width, height float32) {
	wrappedText := wrapWords(a.content, a.font, a.fontSize, width, height)
	lines := strings.Split(wrappedText, "\n")

	for len(a.texts) < len(lines) {
		text := graphics.SfText_create()
		graphics.SfText_setColor(text, a.color.ToSFColor())
		graphics.SfText_setFont(text, a.font)
		graphics.SfText_setCharacterSize(text, a.fontSize)
		a.texts = append(a.texts, text)
	}

	lineSpacing := graphics.SfFont_getLineSpacing(a.font, a.fontSize)

	for i, text := range a.texts {
		var line string
		if i < len(lines) {
			line = lines[i]
		} else {
			line = ""
		}
		graphics.SfText_setString(text, line)
		bounds := graphics.SfText_getLocalBounds(text)
		textWidth := bounds.GetWidth()

		origin := system.NewSfVector2f()
		origin.SetY(-lineSpacing * float32(i))
		origin.SetX((textWidth - width) * a.alignment)

		graphics.SfText_setOrigin(text, origin)
	}
}

const (
	Left   float32 = 0
	Center float32 = 0.5
	Right  float32 = 1
)

func (a *AdvancedText) SetAlignment(alignment float32) *AdvancedText {
	a.alignment = alignment
	a.redraw = true
	return a
}

func (a *AdvancedText) SetContent(content string) *AdvancedText {
	a.content = content
	a.redraw = true
	return a
}

func (a *AdvancedText) SetFont(font Font) *AdvancedText {
	a.font = font
	for _, text := range a.texts {
		graphics.SfText_setFont(text, font)
	}
	return a
}

func (a *AdvancedText) SetFontSize(fontSize uint) *AdvancedText {
	a.fontSize = fontSize
	for _, text := range a.texts {
		graphics.SfText_setCharacterSize(text, fontSize)
	}
	return a
}

func (a *AdvancedText) SetColor(color Color) *AdvancedText {
	a.color = color
	for _, text := range a.texts {
		graphics.SfText_setFillColor(text, color.ToSFColor())
	}
	return a
}


func (a *AdvancedText) Clean() {
	for _, text := range a.texts {
		graphics.SfText_destroy(text)
	}
	system.DeleteSfVector2f(a.position)
}

func (a *AdvancedText) Init() ([]string, []func(Event)) {
	return []string{}, []func(Event){}
}
