package GUI_Interface

import (
	"Coursework_DB/DB_Comunicate"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func CreateWindow() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "Correct request",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Check request",
				OnClicked: func() {
					outTE.SetText(DB_Comunicate.StartCommunicate(inTE.Text()))
				},
			},
		},
	}.Run()
}
