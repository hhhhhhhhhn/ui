package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/system"
)

//"gopkg.in/teh-cmc/go-sfml.v24/graphics"

type FixedLayoutArg struct {
	widget Widget
	x      float32
	y      float32
	width  float32
	height float32
}

type FixedLayout struct {
	widgets []FixedLayoutArg
}

func (f *FixedLayout) Draw(t Texture, x float32,
	y float32, width float32, height float32) {
		for _, widgetStruct := range f.widgets {
			widgetStruct.widget.Draw(
				t,
				widgetStruct.x + x,
				widgetStruct.y + y,
				widgetStruct.width,
				widgetStruct.height,
			)
		}
}

func (f *FixedLayout) Clean() {
	for _, widgetStruct := range f.widgets {
		widgetStruct.widget.Clean()
	}
}

func (f *FixedLayout) Init() (eventTypes []string, listeners []func(Event)) {
	for _, widgetStruct := range f.widgets {
		widgetEventTypes, widgetListeners := widgetStruct.widget.Init()
		eventTypes = append(eventTypes, widgetEventTypes...)
		listeners = append(listeners, widgetListeners...)
	}
	return eventTypes, listeners
}

func NewFixedLayout(widgets []FixedLayoutArg) *FixedLayout {
	return &FixedLayout{widgets: widgets}
}

type GridLayoutArg struct {
	widget Widget
	x      int
	y      int
	width  int
	height int
}

type GridLayout struct {
	rows        int
	columns     int
	rowHeight   float32
	columnWidth float32
	lastWidth   float32
	lastHeight  float32
	widgets     []GridLayoutArg
}

func (g *GridLayout) Draw(t Texture, x float32,
	y float32, width float32, height float32) {
		if width != g.lastWidth || height != g.lastHeight {
			g.calculateRowsAndColumns(width, height)
		}
		for _, widgetStruct := range g.widgets {
			widgetStruct.widget.Draw(
				t,
				float32(widgetStruct.x) * g.columnWidth + x,
				float32(widgetStruct.y) * g.rowHeight + x,
				float32(widgetStruct.width) * g.columnWidth,
				float32(widgetStruct.height) * g.rowHeight,
			)
		}
		g.lastWidth = width
		g.lastHeight = height
}

func (g *GridLayout) calculateRowsAndColumns(width, height float32) {
	g.columnWidth = width / float32(g.columns)
	g.rowHeight = height / float32(g.rows)
}

func (g *GridLayout) Clean() {
	for _, widgetStruct := range g.widgets {
		widgetStruct.widget.Clean()
	}
}

func (g *GridLayout) Init() (eventTypes []string, listeners []func(Event)) {
	for _, widgetStruct := range g.widgets {
		widgetEventTypes, widgetListeners := widgetStruct.widget.Init()
		eventTypes = append(eventTypes, widgetEventTypes...)
		listeners = append(listeners, widgetListeners...)
	}
	return eventTypes, listeners
}

func NewGridLayout(widgets []GridLayoutArg, rows, columns int) *GridLayout {
	return &GridLayout{widgets: widgets, rows: rows, columns: columns}
}

type FixedScrollLayoutArg struct {
	widget Widget
	x      float32
	y      float32
	width  float32
	height float32
}

type FixedScrollLayout struct {
	widgets           []FixedScrollLayoutArg
	texture           Texture
	view              graphics.Struct_SS_sfView
	sprite            graphics.Struct_SS_sfSprite
	position          system.SfVector2f
	bottom            float32
	scroll            float32
	width             float32
	height            float32
	onScrollListeners []func(float32)
}

func (f *FixedScrollLayout) resize(x, y, width, height float32) {
	size := system.NewSfVector2f()
	size.SetX(width); size.SetY(height)

	center := system.NewSfVector2f()
	center.SetX(width / 2 + x); center.SetY(height / 2 + y)

	graphics.SfView_setCenter(f.view, center)
	graphics.SfView_setSize(f.view, size)

	f.texture = graphics.SfRenderTexture_create(uint(width), uint(height), 0)

	graphics.SfRenderTexture_setView(f.texture, f.view)
}

func (f *FixedScrollLayout) Draw(t Texture, x float32,
	y float32, width float32, height float32) {
		if width != f.width || height != f.height {
			f.resize(x, y, width, height)
			f.width = width
			f.height = height
		}

		graphics.SfRenderTexture_clear(f.texture, graphics.GetSfWhite())
		for _, widgetStruct := range f.widgets {
			// Is in the top
			if widgetStruct.y - f.scroll + widgetStruct.height < 0 {
				continue
			}
			// Is in the bottom
			if widgetStruct.y - f.scroll > height {
				continue
			}
			widgetStruct.widget.Draw(
				f.texture,
				x + widgetStruct.x,
				y + widgetStruct.y - f.scroll,
				widgetStruct.width,
				widgetStruct.height,
			)
		}

		graphics.SfRenderTexture_display(f.texture)
		graphics.SfSprite_setTexture(f.sprite, graphics.SfRenderTexture_getTexture(f.texture), 1)

		f.position.SetX(x)
		f.position.SetY(y)
		graphics.SfSprite_setPosition(f.sprite, f.position)

		graphics.SfRenderTexture_drawSprite(t, f.sprite, graphics.SwigcptrSfRenderStates(0))
}

func (f *FixedScrollLayout) Clean() {
	for _, widgetStruct := range f.widgets {
		widgetStruct.widget.Clean()
	}
	graphics.SfRenderTexture_destroy(f.texture)
}

func (f *FixedScrollLayout) Init() (eventTypes []string, listeners []func(Event)) {
	for _, widgetStruct := range f.widgets {
		widgetEventTypes, widgetListeners := widgetStruct.widget.Init()
		eventTypes = append(eventTypes, widgetEventTypes...)
		listeners = append(listeners, widgetListeners...)
	}
	return eventTypes, listeners
}

func (f *FixedScrollLayout) Scroll(scroll float32) *FixedScrollLayout {
	if scroll > f.bottom {
		scroll = f.bottom
	}
	f.scroll = scroll
	for _, handler := range f.onScrollListeners {
		handler(scroll)
	}
	return f
}

func (f *FixedScrollLayout) OnScroll(handler func(float32)) *FixedScrollLayout {
	f.onScrollListeners = append(f.onScrollListeners, handler)
	return f
}

func (f *FixedScrollLayout) SetBottom(bottom float32) *FixedScrollLayout {
	f.bottom = bottom
	if f.scroll > f.bottom {
		f.scroll = f.bottom
	}
	return f
}

func NewFixedScrollLayout(widgets []FixedScrollLayoutArg) *FixedScrollLayout {
	layout := &FixedScrollLayout{}
	layout.position = system.NewSfVector2f()
	layout.sprite = graphics.SfSprite_create()
	layout.widgets = widgets
	layout.view = graphics.SfView_create()
	return layout
}
