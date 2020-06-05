package main

import (
	"fmt"
	"runtime"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/profile"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	sprigTheme "git.sr.ht/~whereswaldon/sprig/widget/theme"
)

type ViewManager interface {
	RequestViewSwitch(ViewID)
	RegisterView(ViewID, View)
	RequestClipboardPaste()
	HandleClipboard(contents string)
	UpdateClipboard(string)
	Layout(gtx layout.Context) layout.Dimensions
	SetProfiling(bool)
}

type viewManager struct {
	views   map[ViewID]View
	current ViewID
	window  *app.Window
	Theme   *sprigTheme.Theme

	// runtime profiling data
	profiling   bool
	profile     profile.Event
	lastMallocs uint64
}

func NewViewManager(window *app.Window, theme *sprigTheme.Theme, profile bool) ViewManager {
	vm := &viewManager{
		views:     make(map[ViewID]View),
		window:    window,
		profiling: profile,
		Theme:     theme,
	}
	return vm
}

func (vm *viewManager) RegisterView(id ViewID, view View) {
	vm.views[id] = view
	view.SetManager(vm)
}

func (vm *viewManager) RequestViewSwitch(id ViewID) {
	vm.current = id
}

func (vm *viewManager) RequestClipboardPaste() {
	vm.window.ReadClipboard()
}

func (vm *viewManager) UpdateClipboard(contents string) {
	vm.window.WriteClipboard(contents)
}

func (vm *viewManager) HandleClipboard(contents string) {
	vm.views[vm.current].HandleClipboard(contents)
}

func (vm *viewManager) Layout(gtx layout.Context) layout.Dimensions {
	defer vm.layoutProfileTimings(gtx)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return vm.layoutProfileTimings(gtx)
		}),
		layout.Rigid(func(gtx C) D {
			gtx.Constraints.Min = gtx.Constraints.Max
			vm.views[vm.current].Update(gtx)
			return vm.views[vm.current].Layout(gtx)
		}),
	)
}

func (vm *viewManager) layoutProfileTimings(gtx layout.Context) layout.Dimensions {
	if !vm.profiling {
		return D{}
	}
	for _, e := range gtx.Events(vm) {
		if e, ok := e.(profile.Event); ok {
			vm.profile = e
		}
	}
	profile.Op{Tag: vm}.Add(gtx.Ops)
	var mstats runtime.MemStats
	runtime.ReadMemStats(&mstats)
	mallocs := mstats.Mallocs - vm.lastMallocs
	vm.lastMallocs = mstats.Mallocs
	text := fmt.Sprintf("m: %d %s", mallocs, vm.profile.Timings)
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return sprigTheme.DrawRect(gtx,
				vm.Theme.Background.Light,
				f32.Point{
					X: float32(gtx.Constraints.Min.X),
					Y: float32(gtx.Constraints.Min.Y),
				},
				0)
		}),
		layout.Stacked(func(gtx C) D {
			return layout.Inset{Top: unit.Dp(4), Left: unit.Dp(4)}.Layout(gtx, func(gtx C) D {
				return material.Body1(vm.Theme.Theme, text).Layout(gtx)
			})
		}),
	)
}

func (vm *viewManager) SetProfiling(isProfiling bool) {
	vm.profiling = isProfiling
}
