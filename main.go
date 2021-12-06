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

type App struct {
	window graphics.Struct_SS_sfRenderWindow
	view   graphics.Struct_SS_sfView
	root   Widget
	width  float32
	height float32
}

func NewApp(title string, width uint, height uint, root Widget) (app App) {
	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)
	vm.SetWidth(width)
	vm.SetHeight(height)
	vm.SetBitsPerPixel(32)

	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)

	app.window = graphics.SfRenderWindow_create(vm, title, uint(window.SfResize|window.SfClose), cs)
	app.view   = graphics.SfRenderWindow_getDefaultView(app.window)
	app.root   = root
	app.width  = float32(width)
	app.height = float32(height)
	return app
}

func (app *App) Run() {
	event := window.NewSfEvent()
	defer window.DeleteSfEvent(event)

	for window.SfWindow_isOpen(app.window) > 0 {
		for window.SfWindow_pollEvent(app.window, event) > 0 {
			switch(event.GetXtype()) {
			case window.SfEventType(window.SfEvtClosed):
				return
			case window.SfEventType(window.SfEvtResized):
				size := event.GetSize()

				newSize := system.NewSfVector2f()
				newSize.SetX(float32(size.GetWidth()))
				newSize.SetY(float32(size.GetHeight()))

				center := system.NewSfVector2f()
				center.SetX(float32(size.GetWidth() / 2))
				center.SetY(float32(size.GetHeight() / 2))

				graphics.SfView_setCenter(app.view, center)

				graphics.SfView_setSize(app.view, newSize)
				graphics.SfRenderWindow_setView(app.window, app.view)
				
				app.height = float32(size.GetHeight())
				app.width  = float32(size.GetWidth())
				break
			}
		}
		graphics.SfRenderWindow_clear(app.window, graphics.GetSfWhite())
		app.root.Draw(app.window, 0, 0, app.width, app.height)
		graphics.SfRenderWindow_display(app.window)
	}
}

func (app *App) Clean() {
	app.root.Clean()
	window.SfWindow_destroy(app.window)
	graphics.SfView_destroy(app.view)
}

func main() {
	c := NewCircle(Color{0,255,255,55})
	c2 := NewRectangle(Color{255,255,0,255})
	root := NewFixedLayout([]FixedLayoutArg{
		{&c, 200, 200, 200, 200},
		{&c2, 400, 300, 200, 200},
	})
	app := NewApp("New app", 800, 600, &root)
	app.Run()
	app.Clean()
}

//func main() {
//	vm := window.NewSfVideoMode()
//	defer window.DeleteSfVideoMode(vm)
//	vm.SetWidth(800)
//	vm.SetHeight(600)
//	vm.SetBitsPerPixel(32)
//
//	/* Create the main window */
//	cs := window.NewSfContextSettings()
//	defer window.DeleteSfContextSettings(cs)
//	w := graphics.SfRenderWindow_create(vm, "SFML window", uint(window.SfResize|window.SfClose), cs)
//	defer window.SfWindow_destroy(w)
//	view := graphics.SfRenderWindow_getDefaultView(w)
//	defer graphics.SfView_destroy(view)
//
//	ev := window.NewSfEvent()
//	defer window.DeleteSfEvent(ev)
//
//	circle := graphics.SfCircleShape_create()
//	position := system.NewSfVector2f()
//	position.SetX(10)
//	position.SetY(10)
//	graphics.SfCircleShape_setFillColor(circle, graphics.GetSfBlack())
//	graphics.SfCircleShape_setPosition(circle, position)
//	graphics.SfCircleShape_setRadius(circle, 10)
//
//	c := NewCircle(Color{255, 0, 255, 255})
//	c2 := NewCircle(Color{0, 0, 255, 255})
//	c3 := NewCircle(Color{0, 255, 255, 100})
//	f := LoadFont("/usr/share/fonts/TTF/DejaVuSans.ttf")
//	t := NewText(f, Color{0,0,0,255}, 20, "Text text text text text text text text text text text sdf¿sdf sdf sdf skjdfvbhs kdjfhsdfkghdkfuhgsd fjksgdfquisefhsf sdiufh sidfhsidf sdfiuhsd")
//	c3.SetOutlineColor(Color{255, 0, 255, 255})
//	c3.SetOutlineThickness(20)
//	c4 := NewCircle(Color{0, 100, 100, 255})
//	c5 := NewRectangle(Color{0, 0, 255, 25})
//
//	/* Start the game loop */
//	for window.SfWindow_isOpen(w) > 0 {
//		/* Process events */
//		for window.SfWindow_pollEvent(w, ev) > 0 {
//			/* Close window: exit */
//			switch(ev.GetXtype()) {
//			case window.SfEventType(window.SfEvtClosed):
//				return
//			case window.SfEventType(window.SfEvtResized):
//				size := ev.GetSize()
//
//				newSize := system.NewSfVector2f()
//				newSize.SetX(float32(size.GetWidth()))
//				newSize.SetY(float32(size.GetHeight()))
//
//				center := system.NewSfVector2f()
//				center.SetX(float32(size.GetWidth() / 2))
//				center.SetY(float32(size.GetHeight() / 2))
//
//				graphics.SfView_setCenter(view, center)
//
//				graphics.SfView_setSize(view, newSize)
//				graphics.SfRenderWindow_setView(w, view)
//				break
//			}
//		}
//		graphics.SfRenderWindow_clear(w, graphics.GetSfWhite())
//		c.Draw(w, 40, 40, 60, 60)
//		c2.Draw(w, 100, 40, 100, 100)
//		c3.Draw(w, 400, 350, 60, 60)
//		c4.Draw(w, 700, 400, 60, 60)
//		c5.Draw(w, 600, 300, 1000, 2500)
//		t.Draw(w, 400, 300, 300, 9)
//		graphics.SfRenderWindow_drawShape(w, circle, graphics.SwigcptrSfRenderStates(0))
//		graphics.SfRenderWindow_display(w)
//	}
//
//	c.Clean()
//}
