package stage

//Tomb ...
type Tomb struct {
	x, y        int
	contentType int
	canOpen     bool
}

//NewTomb ...
func NewTomb(x, y, contentType int) *Tomb {
	tomb := &Tomb{}
	tomb.x = x
	tomb.y = y
	tomb.contentType = contentType

	return tomb
}

//Update ...
func (t *Tomb) Update(stagemap [][]int) {
	x := t.x - 1
	y := t.y - 1

	t.canOpen = true
	for i := 0; i < 5; i++ {
		t.canOpen = t.canOpen && (stagemap[y][x+i] > 0)
		t.canOpen = t.canOpen && (stagemap[y+3][x+i] > 0)
	}

	for i := 0; i < 4; i++ {
		t.canOpen = t.canOpen && (stagemap[y+i][x] > 0)
		t.canOpen = t.canOpen && (stagemap[y+i][x+4] > 0)
	}
}
