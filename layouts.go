package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
)

type FixedLayoutArg struct {
	widget Widget
	x      float32
	y      float32
	width  float32
	height float32
}

type FixedLayout struct {
	Widget
	widgets []FixedLayoutArg
}

func (f *FixedLayout) Draw(w graphics.Struct_SS_sfRenderWindow, x float32,
	y float32, width float32, height float32) {
		for _, widgetStruct := range f.widgets {
			widgetStruct.widget.Draw(
				w,
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
	Widget
	rows        int
	columns     int
	rowHeight   float32
	columnWidth float32
	lastWidth   float32
	lastHeight  float32
	widgets     []GridLayoutArg
}

func (g *GridLayout) Draw(w graphics.Struct_SS_sfRenderWindow, x float32,
	y float32, width float32, height float32) {
		if width != g.lastWidth || height != g.lastHeight {
			g.calculateRowsAndColumns(width, height)
		}
		for _, widgetStruct := range g.widgets {
			widgetStruct.widget.Draw(
				w,
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
