package main

import (
	"fmt"
)

func main() {
	f := LoadFont("/usr/share/fonts/TTF/DejaVuSans.ttf")
	root := 
		NewGridLayout([]GridLayoutArg{
			{NewTextInput().SetFont(f).SetFontSize(20).SetOutlineColor(Color{0,0,0,255}).SetOutlineThickness(5).SetColor(Color{0,0,0,255}).SetContent("some text").OnChange(func(text string) {fmt.Println(text)}), 1, 1, 3, 3},
			{NewButton().SetFont(f).SetFontSize(50).SetOutlineColor(Color{0,150,0,255}).SetOutlineThickness(2).SetBackgroundColor(Color{0,0,0,255}).SetHoverBackgroundColor(Color{50,50,50,255}).SetColor(Color{255,255,255,255}).SetContent("Button").OnClick(func(){fmt.Println("Pog")}), 0, 0, 1, 1},
			{NewText().SetFont(f).SetFontSize(20).SetContent("Hello there asd asd hasd asd asd asd asd asd asdkshdfsd fssdf sdf sdfkjshdf sdfjk shdjkfhsdkfh sdf sdfh skjdfh skdfhskjdfhskfh skdfjh skdfh skdfhs kdfhsd kfhskdjfhskj dfhskdf hskdf hello there hello there hello there hello there").SetColor(Color{0,0,0,255}), 3, 3, 1, 1},
			{NewCircle().SetOutlineThickness(30).SetBackgroundColor(Color{0,0,0,255}).SetOutlineColor(Color{255,0,0,255}), 4, 4, 1, 1},
		}, 5, 5)
	app := NewApp("New app", 800, 600, root)
	app.Run()
	app.Clean()

	fmt.Println("Done")
}
