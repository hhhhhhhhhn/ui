package main

import (
	"runtime"

	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

func init() { runtime.LockOSThread() }

type Texture graphics.Struct_SS_sfRenderTexture

type Widget interface {
	Draw(
		t Texture,
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
	texture        graphics.Struct_SS_sfTexture
	textureSprite  graphics.Struct_SS_sfSprite
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

	app.textureSprite = graphics.SfSprite_create()
	app.texture	      = graphics.SfRenderTexture_create(width, height, 0)
	app.window        = graphics.SfRenderWindow_create(vm, title, uint(window.SfResize|window.SfClose), cs)
	app.view          = graphics.SfRenderWindow_getDefaultView(app.window)
	app.root          = root

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
				graphics.SfTexture_destroy(app.texture)
				app.texture = graphics.SfRenderTexture_create(size.GetWidth(), size.GetHeight(), 0)
				break
			case window.SfEventType(window.SfEvtMouseMoved):
				mouseMove := event.GetMouseMove()
				x := mouseMove.GetX()
				y := mouseMove.GetY()
				for _, eventListener := range app.eventListeners["mouseMove"] {
					eventListener(Event{"x": x, "y": y})
				}
				break
			case window.SfEventType(window.SfEvtKeyPressed):
				keyPressed := event.GetKey()
				code := keyPressed.GetCode()
				shift := keyPressed.GetShift()
				control := keyPressed.GetControl()
				alt := keyPressed.GetAlt()
				for _, eventListener := range app.eventListeners["keyUp"] {
					eventListener(Event{
						"keyCode": int(code),
						"shift":   shift > 0,
						"control": control > 0,
						"alt":     alt > 0,
					})
				}
				break
			case window.SfEventType(window.SfEvtTextEntered):
				text := event.GetText()
				unicode := rune(text.GetUnicode())
				for _, eventListener := range app.eventListeners["text"] {
					eventListener(Event{
						"char": unicode,
					})
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
		graphics.SfRenderTexture_clear(app.texture, graphics.GetSfWhite())
		//graphics.SfRenderWindow_clear(app.window, graphics.GetSfWhite())
		app.root.Draw(app.texture, 0, 0, app.width, app.height)
		graphics.SfRenderTexture_display(app.texture)
		graphics.SfSprite_setTexture(app.textureSprite, graphics.SfRenderTexture_getTexture(app.texture), 1)
		graphics.SfRenderWindow_drawSprite(app.window, app.textureSprite, graphics.SwigcptrSfRenderStates(0))
		graphics.SfRenderWindow_display(app.window)
	}
}

func (app *App) AddEventListener(eventType string, listener func(Event)) {
	app.eventListeners[eventType] = append(app.eventListeners[eventType], listener)
}

func (app *App) Clean() {
	app.root.Clean()
	//graphics.SfView_destroy(app.view)
	graphics.SfTexture_destroy(app.texture)
	window.SfWindow_destroy(app.window)
}
