package main

import (
	"runtime"
	"fmt"

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
	Init() ([]string, []func(Event))
}


type Color struct {
	Red   byte
	Green byte
	Blue  byte
	Alpha byte
}

type Event map[string]interface{}

func (color *Color) ToSFColor() graphics.SfColor {
	return graphics.SfColor_fromRGBA(
		color.Red,
		color.Green,
		color.Blue,
		color.Alpha,
	)
}

type App struct {
	window         graphics.Struct_SS_sfRenderWindow
	view           graphics.Struct_SS_sfView
	root           Widget
	eventListeners map[string][]func(Event)
	width          float32
	height         float32
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

	eventTypes, listeners := root.Init()
	app.eventListeners = make(map[string][]func(Event))

	for i := range listeners {
		app.eventListeners[eventTypes[i]] = append(
			app.eventListeners[eventTypes[i]],
			listeners[i],
		)
	}

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
			case window.SfEventType(window.SfEvtMouseMoved):
				mouseMove := event.GetMouseMove()
				x := mouseMove.GetX()
				y := mouseMove.GetY()
				for _, eventListener := range app.eventListeners["mouseMove"] {
					eventListener(Event{"x": x, "y": y})
				}
				break
			case window.SfEventType(window.SfEvtMouseButtonReleased):
				button := event.GetMouseButton()
				if button.GetButton() == window.SfMouseButton(window.SfMouseLeft) {
					for _, eventListener := range app.eventListeners["leftClickUp"] {
						eventListener(Event{})
					}
				}
				break
			}
		}
		graphics.SfRenderWindow_clear(app.window, graphics.GetSfWhite())
		app.root.Draw(app.window, 0, 0, app.width, app.height)
		graphics.SfRenderWindow_display(app.window)
	}
}

func (app *App) AddEventListener(eventType string, listener func(Event)) {
	app.eventListeners[eventType] = append(app.eventListeners[eventType], listener)
}

func (app *App) Clean() {
	app.root.Clean()
	window.SfWindow_destroy(app.window)
	graphics.SfView_destroy(app.view)
}

func main() {
	f := LoadFont("/usr/share/fonts/TTF/DejaVuSans.ttf")
	root := NewFixedLayout([]FixedLayoutArg{
		{NewCircle().SetBackgroundColor(Color{255,255,0,255}), 200, 200, 200, 200},
		{NewRectangle().SetBackgroundColor(Color{0,0,255,255}), 400, 300, 200, 200},
		{NewText(f, 40).SetContent("Color").SetColor(Color{255,255,0,100}), 400, 300, 200, 200},
		{NewButton(f, 40).SetContent("Button").SetColor(Color{255,255,0,255}).SetBackgroundColor(Color{0,0,0,255}).SetHoverBackgroundColor(Color{100,100,100,255}).OnClick(func(){fmt.Println("sada")}), 700, 300, 200, 200},
	})
	app := NewApp("New app", 800, 600, root)
	app.Run()
	app.Clean()

	fmt.Println("Done")
}
