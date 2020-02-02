package framework

import (
	gkey "github.com/gop9/olt/gio/io/key"
	"github.com/gop9/olt/gio/io/pointer"
	"image"

	"github.com/gop9/olt/framework/rect"
)

type mouseButton struct {
	Down       bool
	Clicked    bool
	ClickedPos image.Point
}

type MouseInput struct {
	valid       bool
	clip        rect.Rect
	Buttons     [6]mouseButton
	Pos         image.Point
	Prev        image.Point
	Delta       image.Point
	ScrollDelta int
}

type KeyboardInput struct {
	Keys []gkey.Event
	Text string
}

type Input struct {
	Keyboard KeyboardInput
	Mouse    MouseInput
}

func (win *Window) Input() *Input {
	if !win.toplevel() {
		return &Input{}
	}
	win.ctx.Input.Mouse.clip = win.cmds.Clip
	return &win.ctx.Input
}

func (win *Window) scrollwheelInput() *Input {
	if win.ctx.scrollwheelFocus == win.idx {
		return &win.ctx.Input
	}
	return &Input{}
}

func (win *Window) KeyboardOnHover(bounds rect.Rect) KeyboardInput {
	if !win.toplevel() || !win.ctx.Input.Mouse.HoveringRect(bounds) {
		return KeyboardInput{}
	}
	return win.ctx.Input.Keyboard
}

func (i *MouseInput) HasClickInRect(id pointer.Buttons, b rect.Rect) bool {
	btn := &i.Buttons[id]
	return unify(b, i.clip).Contains(btn.ClickedPos)
}

func (i *MouseInput) IsClickInRect(id pointer.Buttons, b rect.Rect) bool {
	return i.IsClickDownInRect(id, b, false)
}

func (i *MouseInput) IsClickDownInRect(id pointer.Buttons, b rect.Rect, down bool) bool {
	btn := &i.Buttons[id]
	return i.HasClickInRect(id, b) && btn.Down == down && btn.Clicked
}

func (i *MouseInput) AnyClickInRect(b rect.Rect) bool {
	return i.IsClickInRect(pointer.ButtonLeft, b) || i.IsClickInRect(pointer.ButtonMiddle, b) || i.IsClickInRect(pointer.ButtonRight, b)
}

func (i *MouseInput) HoveringRect(rect rect.Rect) bool {
	return i.valid && unify(rect, i.clip).Contains(i.Pos)
}

func (i *MouseInput) PrevHoveringRect(rect rect.Rect) bool {
	return i.valid && unify(rect, i.clip).Contains(i.Prev)
}

func (i *MouseInput) Clicked(id pointer.Buttons, rect rect.Rect) bool {
	if !i.HoveringRect(rect) {
		return false
	}
	return i.IsClickInRect(id, rect)
}

func (i *MouseInput) Down(id pointer.Buttons) bool {
	return i.Buttons[id].Down
}

func (i *MouseInput) Pressed(id pointer.Buttons) bool {
	return i.Buttons[id].Down && i.Buttons[id].Clicked
}

func (i *MouseInput) Released(id pointer.Buttons) bool {
	return !(i.Buttons[id].Down) && i.Buttons[id].Clicked
}

func (i *KeyboardInput) Pressed(key gkey.Code) bool {
	for _, k := range i.Keys {
		if k.Code == key {
			return true
		}
	}
	return false
}

func (win *Window) inputMaybe(widgetValid bool) *Input {
	if widgetValid && win.toplevel() && win.flags&windowEnabled != 0 {
		win.ctx.Input.Mouse.clip = win.cmds.Clip
		return &win.ctx.Input
	}
	return &Input{}
}

func (win *Window) toplevel() bool {
	if win.moving {
		return false
	}
	if win.ctx.dockedWindowFocus != 0 && win.idx == win.ctx.dockedWindowFocus {
		return true
	}
	return win.idx == win.ctx.floatWindowFocus
}
