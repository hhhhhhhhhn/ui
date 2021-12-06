package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
)

type Button struct {
	Widget
	TextWidget
	ShapeWidget
	rectangle    *Rectangle
	text         *Text
	normalColor  Color
	hoverColor   Color
	isHovering   bool
	x            float32
	y            float32
	width        float32
	height       float32
}

func (b *Button) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
	width float32, height float32) {
		b.x = x
		b.y = y
		b.width = width
		b.height = height
		b.rectangle.Draw(w, x, y, width, height)
		b.text.Draw(w, x, y, width, height)
}

func (b *Button) Clean() {
	b.rectangle.Clean()
	b.text.Clean()
}

func (b *Button) Init() ([]string, []func(Event)) {
	onMouseMove := func(event Event) {
		mouseX := float32(event["x"].(int))
		mouseY := float32(event["y"].(int))

		if (isWithin(mouseX, mouseY, b.x, b.y, b.width, b.height)) {
			if !b.isHovering {
				b.rectangle.SetBackgroundColor(b.hoverColor)
				b.isHovering = true
			}
		} else {
			if b.isHovering {
				b.rectangle.SetBackgroundColor(b.normalColor)
				b.isHovering = false
			}
		}
	}
	return []string{"mouseMove"}, []func(Event){onMouseMove}
}

func isWithin(pointX, pointY, hitboxX, hitboxY, hitboxWidth, hitboxHeight float32) bool {
	if pointX < hitboxX ||
		pointY < hitboxY ||
		pointX > (hitboxX + hitboxWidth) ||
		pointY > (hitboxY + hitboxHeight) {
			return false
	}
	return true
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

func NewButton(font Font, fontSize uint) *Button {
	text := NewText(font, fontSize)
	rectangle := NewRectangle()
	return &Button{text: text, rectangle: rectangle}
}
