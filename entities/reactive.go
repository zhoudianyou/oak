package entities

import (
	"bitbucket.org/oakmoundstudio/plasticpiston/plastic/collision"
	"bitbucket.org/oakmoundstudio/plasticpiston/plastic/event"
	"strconv"
)

type Interactive struct {
	Reactive
	moving
}

func (i *Interactive) Init() event.CID {
	cID := event.NextID(i)
	i.CID = cID
	return cID
}

func (i *Interactive) String() string {
	st := "Interactive: \n{"
	st += i.Reactive.String()
	st += "}\n " + i.moving.String()
	return st
}

type Reactive struct {
	Doodad
	W, H   float64
	RSpace *collision.ReactiveSpace
}

func (r *Reactive) SetDim(w, h float64) {
	r.SetLogicDim(w, h)
	r.RSpace.SetDim(w, h)
}

func (r *Reactive) GetLogicDim() (float64, float64) {
	return r.W, r.H
}

func (r *Reactive) SetLogicDim(w, h float64) {
	r.W = w
	r.H = h
}

func (r *Reactive) SetSpace(sp *collision.ReactiveSpace) {
	collision.Remove(r.RSpace.Space())
	r.RSpace = sp
	collision.Add(r.RSpace.Space())
}

func (r *Reactive) GetSpace() *collision.ReactiveSpace {
	return r.RSpace
}

// Overwrites

func (r *Reactive) Init() event.CID {
	cID := event.NextID(r)
	r.CID = cID
	return cID
}

func (r *Reactive) SetPos(x float64, y float64) {
	r.SetLogicPos(x, y)
	r.R.SetPos(x, y)

	if r.RSpace != nil {
		collision.UpdateSpace(r.X, r.Y, r.W, r.H, r.RSpace.Space())
	}
}

func (r *Reactive) Destroy() {
	r.R.UnDraw()
	collision.Remove(r.RSpace.Space())
	r.CID.UnbindAll()
	event.DestroyEntity(int(r.CID))
}

func (r *Reactive) String() string {
	st := "Reactive:\n{"
	st += r.Doodad.String()
	st += " }, \n"
	w := strconv.FormatFloat(r.W, 'f', 2, 32)
	h := strconv.FormatFloat(r.H, 'f', 2, 32)
	st += "W: " + w + ", H: " + h
	st += ",\nS:{ "
	st += r.RSpace.Space().String()
	st += "}"
	return st
}