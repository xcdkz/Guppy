package tiles

import (
	"errors"

	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term"
)

type tile struct {
	Parent                  *tile
	Children                []*tile
	posX, posY              int
	sizeX, sizeY            int
	cursorPosX, cursorPosY  int
	canBeFocused, isFocused bool
	vSplit                  bool
}

func NewRoot() (tile, error) {
	sX, sY, err := term.GetSize()
	return tile{
		Parent:       nil,
		Children:     make([]*tile, 0),
		posX:         0,
		posY:         0,
		sizeX:        sX,
		sizeY:        sY,
		canBeFocused: false,
		isFocused:    false,
		vSplit:       true,
	}, err
}

func (t *tile) NewChild(vSplit, canBeFocused bool) {
	var posX, posY, sizeX, sizeY int
	if len(t.Children) == 0 {
		posX = 0
		posY = 0
		sizeX = t.sizeX
		sizeY = t.sizeY
	} else {
		if t.vSplit {
			childWidth := t.sizeX / (len(t.Children) + 1)
			for i, c := range t.Children {
				c.posX = i * childWidth
				c.sizeX = childWidth
			}
			posX = len(t.Children) * childWidth
			posY = t.Children[0].posY
			sizeX = childWidth
			sizeY = t.Children[0].sizeY
			sizeX += t.sizeX - (childWidth * (len(t.Children) + 1))
		} else {
			childHeight := t.sizeY / (len(t.Children) + 1)
			for i, c := range t.Children {
				c.posY = i * childHeight
				c.sizeY = childHeight
			}
			posY = len(t.Children) * childHeight
			posX = t.Children[0].posX
			sizeY = childHeight
			sizeX = t.Children[0].sizeX
			sizeY += t.sizeY - (childHeight * (len(t.Children) + 1))
		}
	}
	t.Children = append(t.Children, &tile{
		Parent:       t,
		Children:     make([]*tile, 0),
		posX:         posX,
		posY:         posY,
		sizeX:        sizeX,
		sizeY:        sizeY,
		canBeFocused: canBeFocused,
		isFocused:    false,
		vSplit:       vSplit,
	})
}

func (t *tile) GetPos() (int, int) {
	return t.posX, t.posY
}

func (t *tile) GetSize() (int, int) {
	return t.sizeX, t.sizeY
}

func (t *tile) getRootPosition() (int, int) {
	if t.Parent != nil {
		parentX, parentY := t.Parent.getRootPosition()
		return parentX + t.posX, parentY + t.posY
	}
	return t.posX, t.posY
}

func (t *tile) getRootSize() (int, int) {
	if t.Parent != nil {
		return t.Parent.getRootSize()
	}
	return t.sizeX, t.sizeY
}

func (t *tile) GetCursorPos() (int, int) {
	return t.cursorPosX, t.cursorPosY
}

func (t *tile) DrawBorder() {
	x, y := t.getRootPosition()
	sx, sy := t.GetSize()
	cursor.MoveTo(x, y)
	for i := 0; i < sx; i++ {
		print("-")
	}
	cursor.MoveTo(x, y+sy-1)
	for i := 0; i < sx; i++ {
		print("-")
	}
	cursor.MoveTo(x, y+1)
	for i := 0; i < sy; i++ {
		print("|")
		cursor.MoveTo(x, y+1+i)
	}
	cursor.MoveTo(x+sx, y+1)
	for i := 0; i < sy; i++ {
		print("|")
		cursor.MoveTo(x+sx, y+1+i)
	}
}

func (t *tile) FindFocused() (*tile, error) {
	for t.Parent != nil {
		t = t.Parent
	}
	return t.findFocusedFromRoot()
}

func (t *tile) findFocusedFromRoot() (*tile, error) {
	if t.isFocused {
		return t, nil
	}
	for _, c := range t.Children {
		tmp, err := c.findFocusedFromRoot()
		if err == nil {
			return tmp, nil
		}
	}
	return t, errors.New("no focused tiles")
}