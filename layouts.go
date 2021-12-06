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
				widgetStruct.x,
				widgetStruct.y,
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

func NewFixedLayout(widgets []FixedLayoutArg) FixedLayout {
	return FixedLayout{widgets: widgets}
}
