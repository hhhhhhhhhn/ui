package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"

	"unicode"
)

type Button struct {
	Widget
	TextWidget
	ShapeWidget
	rectangle        *Rectangle
	text             *Text
	normalColor      Color
	hoverColor       Color
	isHovering       bool
	onClickListeners []func()
	bounds           BoundingBox
}

func (b *Button) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
	width float32, height float32) {
		b.bounds.Update(x, y, width, height)
		b.rectangle.Draw(w, x, y, width, height)
		b.text.Draw(w, x, y, width, height)
}

func (b *Button) Clean() {
	b.rectangle.Clean()
	b.text.Clean()
}

func (b *Button) Init() ([]string, []func(Event)) {
	onMouseMove := CreateHoverEventListener(
		&b.bounds,
		&b.isHovering,
		func() {b.rectangle.SetBackgroundColor(b.hoverColor)},
		func() {b.rectangle.SetBackgroundColor(b.normalColor)},
	)
	
	leftClickUp := func(Event) {
		if b.isHovering {
			b.Click()
		}
	}

	return []string{"mouseMove", "leftClickUp"}, []func(Event){onMouseMove, leftClickUp}
}

func (b *Button) SetFont(font Font) *Button {
	b.text.SetFont(font)
	return b
}

func (b *Button) SetContent(content string) *Button {
	b.text.SetContent(content)
	return b
}

func (b *Button) SetFontSize(fontSize uint) *Button {
	b.text.SetFontSize(fontSize)
	return b
}

func (b *Button) SetColor(color Color) *Button {
	b.text.SetColor(color)
	return b
}

func (b *Button) SetBackgroundColor(color Color) *Button {
	b.normalColor = color
	if (!b.isHovering) {
		b.rectangle.SetBackgroundColor(color)
	}
	return b
}

func (b *Button) SetHoverBackgroundColor(color Color) *Button {
	b.hoverColor = color
	if (b.isHovering) {
		b.rectangle.SetBackgroundColor(color)
	}
	return b
}

func (b *Button) SetOutlineColor(color Color) *Button {
	b.rectangle.SetOutlineColor(color)
	return b
}

func (b *Button) SetOutlineThickness(thickness float32) *Button {
	b.rectangle.SetOutlineThickness(thickness)
	return b
}

func (b *Button) Click() *Button {
	for _, listener := range b.onClickListeners {
		listener()
	}
	return b
}

func (b *Button) OnClick(handler func()) *Button {
	b.onClickListeners = append(b.onClickListeners, handler)
	return b
}

func NewButton(font Font, fontSize uint) *Button {
	text := NewText(font, fontSize)
	rectangle := NewRectangle()
	return &Button{text: text, rectangle: rectangle}
}

type TextInput struct {
	Widget
	TextWidget
	ShapeWidget
	rectangle         *Rectangle
	text              *Text
	bounds            BoundingBox
	isHovering        bool
	active            bool
	content           string
	cursor            int
	onChangeListeners []func(string)
}

func (t *TextInput) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
	width float32, height float32) {
		t.bounds.Update(x, y, width, height)
		t.rectangle.Draw(w, x, y, width, height)
		t.text.Draw(w, x, y, width, height)
}

func (t *TextInput) Clean() {
	t.text.Clean()
	t.rectangle.Clean()
}

func (t *TextInput) Init() ([]string, []func(Event)) {
	onMouseMove := CreateHoverEventListener(
		&t.bounds,
		&t.isHovering,
		func(){},
		func(){},
	)

	onClick := func(event Event) {
		t.active = t.isHovering
		t.updateText()
	}

	onKeyUp := func(event Event) {
		keyCode := event["keyCode"].(int)
		if t.active {
			switch keyCode {
			case window.SfKeyLeft:
				if t.cursor > 0 {
					t.cursor--
				}
			case window.SfKeyRight:
				if t.cursor < len(t.content) {
					t.cursor++
				}
				break
			case window.SfKeyBack:
				if t.cursor > 0 {
					t.content = t.content[:t.cursor-1] + t.content[t.cursor:]
					t.cursor--
					t.Change()
				}
				break
			}
			t.updateText()
		}
	}

	onText := func(event Event) {
		char := event["char"].(rune)
		if unicode.IsPrint(char) && t.active {
			t.content = t.content[:t.cursor] + string(char) + t.content[t.cursor:]
			t.cursor++
			t.Change()
			t.updateText()
		}
	}

	return []string{"mouseMove", "leftClickUp", "keyUp", "text"}, []func(Event){onMouseMove, onClick, onKeyUp, onText}
}

func (t *TextInput) updateText() {
	if t.cursor >= 0 && t.cursor <= len(t.content) && t.active {
		t.text.SetContent(t.content[:t.cursor] + "|" + t.content[t.cursor:])
	} else {
		t.text.SetContent(t.content)
	}
}

func (t *TextInput) SetFont(font Font) *TextInput {
	t.text.SetFont(font)
	return t
}

func (t *TextInput) SetContent(content string) *TextInput {
	t.content = content
	t.cursor = len(content)
	t.Change()
	t.updateText()
	return t
}

func (t *TextInput) Change() *TextInput {
	for _, listener := range t.onChangeListeners {
		listener(t.content)
	}
	return t
}

func (t *TextInput) OnChange(handler func(string)) *TextInput {
	t.onChangeListeners = append(t.onChangeListeners, handler)
	return t
}

func (t *TextInput) SetFontSize(fontSize uint) *TextInput {
	t.text.SetFontSize(fontSize)
	return t
}

func (t *TextInput) SetColor(color Color) *TextInput {
	t.text.SetColor(color)
	return t
}

func (t *TextInput) SetBackgroundColor(color Color) *TextInput {
	t.rectangle.SetBackgroundColor(color)
	return t
}

func (t *TextInput) SetOutlineColor(color Color) *TextInput {
	t.rectangle.SetOutlineColor(color)
	return t
}

func (t *TextInput) SetOutlineThickness(thickness float32) *TextInput {
	t.rectangle.SetOutlineThickness(thickness)
	return t
}

func NewTextInput(font Font, fontSize uint) *TextInput {
	text := NewText(font, fontSize)
	rectangle := NewRectangle()
	return &TextInput{text: text, rectangle: rectangle}
}
