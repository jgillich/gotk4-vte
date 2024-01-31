package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const vteModule = "github.com/jgillich/gotk4-vte/pkg"

var postprocessors = map[string][]girgen.Postprocessor{}

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: vteModule,
		Packages: []genmain.Package{
			{Name: "vte-2.91-gtk4", Namespaces: []string{"Vte-3"}},
		},
		PkgGenerated: []string{"vte"},
		Filters: []types.FilterMatcher{
			types.AbsoluteFilter("C.vte_terminal_get_text"),
			types.AbsoluteFilter("C.vte_terminal_get_text_include_trailing_spaces"),
			types.AbsoluteFilter("C.vte_terminal_get_text_range"),
		},
	},
)

func main() {
	genmain.Run(Data)
}
