package game

import "fs/lib/phys"

type Screen struct {
	Dim       *phys.Rect
	Fullsceen *bool
}

func NewScreen(w, h float64, fullscreen bool) *Screen {
	return &Screen{
		Dim:       phys.NewRect(0, 0, w, h),
		Fullsceen: &fullscreen,
	}
}
