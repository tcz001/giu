// Package main presents usage of giu.RangeBuilder.
package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func loop() {
	g.SingleWindow().Layout(
		g.Label("Below buttons are generated by RangeBuilder in one line"),
		g.Row(
			g.RangeBuilder("Buttons in one line", []string{"Button1", "Button2", "Button3"}, func(i int, v string) g.Widget {
				return g.Layout{
					g.Button(v).OnClick(func() {
						fmt.Println(v, "is clicked", i)
					}),
					g.Tooltip("Tooltip for " + v),
				}
			})...,
		),
		g.Dummy(0, 8),
		g.Label("Below buttons are generated by RangeBuilder"),
		g.RangeBuilder("Buttons", []string{"Button1", "Button2", "Button3"}, func(_ int, str string) g.Widget {
			return g.Button(str).OnClick(func() {
				fmt.Println(str, "is clicked")
			})
		}),
	)
}

func main() {
	wnd := g.NewMasterWindow("Range builder demo", 800, 600, 0)
	wnd.Run(loop)
}
