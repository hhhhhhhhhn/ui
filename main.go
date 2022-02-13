package main

import (
	"fmt"
)

func main() {
	f := LoadFont("/usr/share/fonts/TTF/DejaVuSans.ttf")

	scrollBottom := float32(1000)

	scroll := NewFixedScrollLayout([]FixedScrollLayoutArg{
				{NewRectangle().SetBackgroundColor(Color{1, 255, 1, 255}), 0, 0, 10, 10},
				{NewText().SetFont(f).SetColor(Color{0, 0, 0, 255}).SetFontSize(10).SetContent("Use the scrollbar!"), 0, 0, 200, 12},
				{NewText().SetFont(f).SetColor(Color{0, 0, 0, 255}).SetFontSize(10).SetContent("Below"), 0, 1000, 200, 12},
				{NewButton().SetFont(f).SetFontSize(50).SetOutlineColor(Color{1,150,1,255}).SetOutlineThickness(2).SetBackgroundColor(Color{1,1,1,255}).SetHoverBackgroundColor(Color{50,50,50,255}).SetColor(Color{255,255,255,255}).SetContent("Button"), 50, 100, 100, 100},
			}).SetBottom(scrollBottom)

	root := 
		NewGridLayout([]GridLayoutArg{
			{scroll, 4, 0, 1, 3},
			{NewTextInput().SetFont(f).SetFontSize(20).SetOutlineColor(Color{1,1,1,255}).SetOutlineThickness(5).SetColor(Color{1,1,1,255}).SetContent("some text").OnChange(func(text string) {fmt.Println(text)}), 1, 1, 3, 3},
			{NewButton().SetFont(f).SetFontSize(50).SetOutlineColor(Color{1,150,1,255}).SetOutlineThickness(2).SetBackgroundColor(Color{1,1,1,255}).SetHoverBackgroundColor(Color{50,50,50,255}).SetColor(Color{255,255,255,255}).SetContent("Button"), 0, 0, 1, 1},
			{NewText().SetFont(f).SetFontSize(20).SetContent("Hello there asd asd hasd asd asd asd asd asd asdkshdfsd fssdf sdf sdfkjshdf sdfjk shdjkfhsdkfh sdf sdfh skjdfh skdfhskjdfhskfh skdfjh skdfh skdfhs kdfhsd kfhskdjfhskj dfhskdf hskdf hello there hello there hello there hello there").SetColor(Color{0,0,0,255}), 3, 3, 1, 1},
			{NewSlider().SetBackgroundColor(Color{1,1,1,255}).SetWidth(5).SetSliderColor(Color{100,100,1,255}).OnChange(func(value float32) {scroll.Scroll(value * scrollBottom)}), 2, 2, 2, 2},
			{NewCircle().SetOutlineThickness(30).SetBackgroundColor(Color{1,1,1,255}).SetOutlineColor(Color{255,1,1,255}), 4, 4, 1, 1},
		}, 5, 5)

	app := NewApp("New app", 800, 600, root)
	app.Run()
	app.Clean()

	fmt.Println("Done")
}
