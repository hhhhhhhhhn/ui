package main

import (
	"runtime"

	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

func init() { runtime.LockOSThread() }

type Widget interface {
	Draw(
		w graphics.Struct_SS_sfRenderWindow,
		x float32,
		y float32,
		width float32,
		height float32,
	)
	Clean()
}


type Color struct {
	Red   byte
	Green byte
	Blue  byte
	Alpha byte
}

func (color *Color) ToSFColor() graphics.SfColor {
	return graphics.SfColor_fromRGBA(
		color.Red,
		color.Green,
		color.Blue,
		color.Alpha,
	)
}

type Circle struct {
	circle   graphics.Struct_SS_sfCircleShape
	position system.SfVector2f
}

func (c *Circle) Draw(w graphics.Struct_SS_sfRenderWindow ,x float32, y float32,
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

func NewCircle(color Color) Circle {
	circle := graphics.SfCircleShape_create()
	position := system.NewSfVector2f()
	graphics.SfCircleShape_setFillColor(circle, color.ToSFColor())
	return Circle{circle, position}
}

func main() {
	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)
	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)

	/* Create the main window */
	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)
	w := graphics.SfRenderWindow_create(vm, "SFML window", uint(window.SfResize|window.SfClose), cs)
	defer window.SfWindow_destroy(w)
	view := graphics.SfRenderWindow_getDefaultView(w)
	defer graphics.SfView_destroy(view)

	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	circle := graphics.SfCircleShape_create()
	position := system.NewSfVector2f()
	position.SetX(10)
	position.SetY(10)
	graphics.SfCircleShape_setFillColor(circle, graphics.GetSfBlack())
	graphics.SfCircleShape_setPosition(circle, position)
	graphics.SfCircleShape_setRadius(circle, 10)

	c := NewCircle(Color{255, 0, 255, 255})
	c2 := NewCircle(Color{0, 0, 255, 255})
	c3 := NewCircle(Color{0, 255, 255, 100})
	c4 := NewCircle(Color{0, 100, 100, 255})
	c5 := NewCircle(Color{0, 0, 255, 25})

	/* Start the game loop */
	for window.SfWindow_isOpen(w) > 0 {
		/* Process events */
		for window.SfWindow_pollEvent(w, ev) > 0 {
			/* Close window: exit */
			switch(ev.GetXtype()) {
			case window.SfEventType(window.SfEvtClosed):
				return
			case window.SfEventType(window.SfEvtResized):
				size := ev.GetSize()

				newSize := system.NewSfVector2f()
				newSize.SetX(float32(size.GetWidth()))
				newSize.SetY(float32(size.GetHeight()))

				center := system.NewSfVector2f()
				center.SetX(float32(size.GetWidth() / 2))
				center.SetY(float32(size.GetHeight() / 2))

				graphics.SfView_setCenter(view, center)

				graphics.SfView_setSize(view, newSize)
				graphics.SfRenderWindow_setView(w, view)
				break
			}
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfWhite())
		c.Draw(w, 40, 40, 60, 60)
		c2.Draw(w, 100, 40, 100, 100)
		c3.Draw(w, 400, 350, 60, 60)
		c4.Draw(w, 700, 400, 60, 60)
		c5.Draw(w, 600, 300, 1000, 2500)
		graphics.SfRenderWindow_drawShape(w, circle, graphics.SwigcptrSfRenderStates(0))
		graphics.SfRenderWindow_display(w)
	}

	c.Clean()
}
