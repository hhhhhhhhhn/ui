package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

type ShapeWidget interface {
	Widget
	SetBackgroundColor(Color)    Widget
	SetOutlineColor(Color)       Widget
	SetOutlineThickness(float32) Widget
}

type Circle struct {
	ShapeWidget
	circle   graphics.Struct_SS_sfCircleShape
	position system.SfVector2f
}

func (c *Circle) Draw(t Texture, x float32, y float32,
	width float32, height float32) {
		if width < height {
			graphics.SfCircleShape_setRadius(c.circle, width / 2)
			y += (height - width) / 2
		} else {
			graphics.SfCircleShape_setRadius(c.circle, height / 2)
			x += (width - height) / 2
		}
		c.position.SetX(x)
		c.position.SetY(y)

		graphics.SfCircleShape_setPosition(c.circle, c.position)
		graphics.SfRenderTexture_drawShape(t, c.circle, graphics.SwigcptrSfRenderStates(0))
}

func (c *Circle) Clean() {
	graphics.SfCircleShape_destroy(c.circle)
	system.DeleteSfVector2f(c.position)
}

func (c *Circle) Init() ([]string, []func(Event)) {
	return []string{}, []func(Event){}
}

func (c *Circle) SetOutlineColor(color Color) *Circle {
	graphics.SfCircleShape_setOutlineColor(c.circle, color.ToSFColor())
	return c
}

func (c *Circle) SetOutlineThickness(thickness float32) *Circle {
	graphics.SfCircleShape_setOutlineThickness(c.circle, thickness)
	return c
}

func (c *Circle) SetBackgroundColor(color Color) *Circle {
	graphics.SfCircleShape_setFillColor(c.circle, color.ToSFColor())
	return c
}

func NewCircle() *Circle {
	circle := graphics.SfCircleShape_create()
	position := system.NewSfVector2f()
	return &Circle{circle: circle, position: position}
}

type Rectangle struct {
	ShapeWidget
	rectangle graphics.Struct_SS_sfRectangleShape
	position  system.SfVector2f
	size      system.SfVector2f
}

func (r *Rectangle) Draw(t Texture, x float32, y float32,
	width float32, height float32) {
		r.position.SetX(x)
		r.position.SetY(y)
		graphics.SfRectangleShape_setPosition(r.rectangle, r.position)

		r.size.SetX(width)
		r.size.SetY(height)
		graphics.SfRectangleShape_setSize(r.rectangle, r.size)

		graphics.SfRenderTexture_drawShape(t, r.rectangle, graphics.SwigcptrSfRenderStates(0))
}

func (r *Rectangle) Clean() {
	graphics.SfRectangleShape_destroy(r.rectangle)
	system.DeleteSfVector2f(r.position)
	system.DeleteSfVector2f(r.size)
}

func (r *Rectangle) Init() ([]string, []func(Event)) {
	return []string{}, []func(Event){}
}

func (r *Rectangle) SetOutlineColor(color Color) *Rectangle {
	graphics.SfRectangleShape_setOutlineColor(r.rectangle, color.ToSFColor())
	return r
}

func (r *Rectangle) SetOutlineThickness(thickness float32) *Rectangle {
	graphics.SfRectangleShape_setOutlineThickness(r.rectangle, thickness)
	return r
}

func (r *Rectangle) SetBackgroundColor(color Color) *Rectangle {
	graphics.SfRectangleShape_setFillColor(r.rectangle, color.ToSFColor())
	return r
}

func NewRectangle() *Rectangle {
	rectangle := graphics.SfRectangleShape_create()
	position := system.NewSfVector2f()
	size := system.NewSfVector2f()
	return &Rectangle{rectangle: rectangle, position: position, size: size}
}

