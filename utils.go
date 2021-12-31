package main

type BoundingBox struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func (b *BoundingBox) Update(x, y, width, height float32) {
	b.X = x
	b.Y = y
	b.Width = width
	b.Height = height
}

func CreateHoverEventListener(bounds *BoundingBox, isHovering *bool, onHover, onLeave func()) func(Event) {
	return func(event Event) {
		mouseX := float32(event["x"].(int))
		mouseY := float32(event["y"].(int))

		if (isWithin(mouseX, mouseY, bounds.X, bounds.Y, bounds.Width, bounds.Height)) {
			if !*isHovering {
				onHover()
				*isHovering = true
			}
		} else {
			if *isHovering {
				onLeave()
				*isHovering = false
			}
		}
	}
}

func isWithin(pointX, pointY, hitboxX, hitboxY, hitboxWidth, hitboxHeight float32) bool {
	if pointX < hitboxX ||
		pointY < hitboxY ||
		pointX > (hitboxX + hitboxWidth) ||
		pointY > (hitboxY + hitboxHeight) {
			return false
	}
	return true
}
