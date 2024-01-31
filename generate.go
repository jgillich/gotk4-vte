package main

//go:generate go run . -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
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
		// PkgExceptions: []string{
		// 	"go.mod",
		// 	"go.sum",
		// 	// "LICENSE",
		// },
	},
)

func main() {
	genmain.Run(Data)
}
