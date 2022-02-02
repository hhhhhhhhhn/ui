package main

import (
	//"gopkg.in/teh-cmc/go-sfml.v24/graphics"
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

func (b *Button) Draw(t Texture, x float32, y float32,
	width float32, height float32) {
		b.bounds.Update(x, y, width, height)
		b.rectangle.Draw(t, x, y, width, height)
		b.text.Draw(t, x + b.rectangle.outlineThickness, y + b.rectangle.outlineThickness, width, height)
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

func NewButton() *Button {
	text := NewText()
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

func (t *TextInput) Draw(tx Texture, x float32, y float32,
	width float32, height float32) {
		t.bounds.Update(x, y, width, height)
		t.rectangle.Draw(tx, x, y, width, height)
		t.text.Draw(tx, x + t.rectangle.outlineThickness, y + t.rectangle.outlineThickness, width, height)
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

func NewTextInput() *TextInput {
	text := NewText()
	rectangle := NewRectangle()
	return &TextInput{text: text, rectangle: rectangle}
}

type Slider struct {
	background        *Rectangle
	slider            *Rectangle
	isHovering        bool
	isDragging        bool
	onChangeListeners []func(float32)
	sliderBounds      BoundingBox
	bounds            BoundingBox
	value             float32
	width             float32
}

func (s *Slider) Draw(t Texture, x float32, y float32,
	width float32, height float32) {
		s.background.Draw(t, x, y + ((height - s.width) / 2), width, s.width)
		s.slider.Draw(t, x + ((width - 2 * s.width) * s.value), y + ((height - 4 * s.width) / 2), 2 * s.width, 4 * s.width)
		s.sliderBounds.Update(x + ((width - 2 * s.width) * s.value), y + ((height - 4 * s.width) / 2), 2 * s.width, 4 * s.width)
		s.bounds.Update(x, y, width, height)
}

func (s *Slider) Clean() {
	s.background.Clean()
	s.slider.Clean()
}

func (s *Slider) Init() ([]string, []func(Event)) {
	detectHover := CreateHoverEventListener(
		&s.sliderBounds,
		&s.isHovering,
		func() {},
		func() {},
	)

	onMouseMove := func(event Event) {
		detectHover(event)
		x := float32(event["x"].(int))

		if s.isDragging {
			s.Change((x - s.bounds.X) / s.bounds.Width)
		}
	}

	leftClickDown := func(Event) {
		if s.isHovering {
			s.isDragging = true
		}
	}

	leftClickUp := func(Event) {
		s.isDragging = false
	}

	return []string{"mouseMove", "leftClickUp", "leftClickDown"},
		[]func(Event){onMouseMove, leftClickUp, leftClickDown}
}

func (s *Slider) SetBackgroundColor(color Color) *Slider {
	s.background.SetBackgroundColor(color)
	return s
}

func (s *Slider) SetSliderColor(color Color) *Slider {
	s.slider.SetBackgroundColor(color)
	return s
}

func (s *Slider) SetWidth (width float32) *Slider {
	s.width = width
	return s
}

func (s *Slider) Change(value float32) *Slider {
	s.value = value
	if s.value < 0 {
		s.value = 0
	} else if s.value > 1 {
		s.value = 1
	}
	for _, listener := range s.onChangeListeners {
		listener(s.value)
	}
	return s
}

func (s *Slider) OnChange(handler func(float32)) *Slider {
	s.onChangeListeners = append(s.onChangeListeners, handler)
	return s
}

func NewSlider() *Slider {
	background := NewRectangle()
	slider := NewRectangle()
	return &Slider{background: background, slider: slider}
}
