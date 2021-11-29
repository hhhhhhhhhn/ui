package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

type ShapeWidget interface {
	Widget
	SetOutlineColor(Color)
	SetOutlineThickness(float32)
}

type Circle struct {
	ShapeWidget
	circle   graphics.Struct_SS_sfCircleShape
	position system.SfVector2f
}

func (c *Circle) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
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
		graphics.SfRenderWindow_drawShape(w, c.circle, graphics.SwigcptrSfRenderStates(0))
}

func (c *Circle) Clean() {
	graphics.SfCircleShape_destroy(c.circle)
	system.DeleteSfVector2f(c.position)
}

func (c *Circle) SetOutlineColor(color Color) {
	graphics.SfCircleShape_setOutlineColor(c.circle, color.ToSFColor())
}

func (c *Circle) SetOutlineThickness(thickness float32) {
	graphics.SfCircleShape_setOutlineThickness(c.circle, thickness)
}

func NewCircle(color Color) Circle {
	circle := graphics.SfCircleShape_create()
	position := system.NewSfVector2f()
	graphics.SfCircleShape_setFillColor(circle, color.ToSFColor())
	return Circle{circle: circle, position: position}
}

type Rectangle struct {
	ShapeWidget
	rectangle graphics.Struct_SS_sfRectangleShape
	position  system.SfVector2f
	size      system.SfVector2f
}

func (r *Rectangle) Draw(w graphics.Struct_SS_sfRenderWindow, x float32, y float32,
	width float32, height float32) {
		r.position.SetX(x)
		r.position.SetY(y)
		graphics.SfRectangleShape_setPosition(r.rectangle, r.position)

		r.size.SetX(width)
		r.size.SetY(height)
		graphics.SfRectangleShape_setSize(r.rectangle, r.size)

		graphics.SfRenderWindow_drawShape(w, r.rectangle, graphics.SwigcptrSfRenderStates(0))
}

func (r *Rectangle) Clean() {
	graphics.SfRectangleShape_destroy(r.rectangle)
	system.DeleteSfVector2f(r.position)
	system.DeleteSfVector2f(r.size)
}

func (r *Rectangle) SetOutlineColor(color Color) {
	graphics.SfRectangleShape_setOutlineColor(r.rectangle, color.ToSFColor())
}

func (r *Rectangle) SetOutlineThickness(thickness float32) {
	graphics.SfRectangleShape_setOutlineThickness(r.rectangle, thickness)
}

func NewRectangle(color Color) Rectangle {
	rectangle := graphics.SfRectangleShape_create()
	position := system.NewSfVector2f()
	size := system.NewSfVector2f()
	graphics.SfRectangleShape_setFillColor(rectangle, color.ToSFColor())
	return Rectangle{rectangle: rectangle, position: position, size: size}
}

