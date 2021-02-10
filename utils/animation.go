package utils

//Animation ...
type Animation struct {
	frames       []Rect
	count        int
	currentFrame int
}

func NewAnimation() *Animation {
	anim := &Animation{}

	anim.frames = make([]Rect, 0)
	anim.count = 0

	return anim
}

//AddFrame allow adds frames to do movement.
func (a *Animation) AddFrame(rect Rect) {
	a.frames = append(a.frames, rect)
}

//UpdateFrame update logic frame.
func (a *Animation) UpdateFrame() {
	a.count++
	a.currentFrame = (a.count / 16) % 4
}

//GetFrame returns the current frame.
func (a *Animation) GetFrame() Rect {
	return a.frames[a.currentFrame]
}

//GetFrameIndex returns the frame by index
func (a *Animation) GetFrameIndex(pos int) Rect {
	return a.frames[pos]
}
