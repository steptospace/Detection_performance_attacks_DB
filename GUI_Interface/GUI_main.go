package GUI_Interface

import (
	"Coursework_DB/DB_Comunicate"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func CreateWindow() error {
	var inTE, userId, passTe, outTE *walk.TextEdit

	MainWindow{
		Title:   "Correct request",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				MaxSize: Size{10, 20},
				Children: []Widget{
					TextLabel{Text: "User"},
					TextLabel{Text: "Password"},
					TextLabel{Text: ""},
				},
			},
			HSplitter{
				MaxSize: Size{10, 20},
				Children: []Widget{
					TextEdit{AssignTo: &userId, Name: "User"},
					TextEdit{AssignTo: &passTe, Name: "Password"},
					PushButton{
						Text: "Check request",
						OnClicked: func() {
							db := DB_Comunicate.Connect(userId.Text(), passTe.Text())
							requestDB, err := DB_Comunicate.StartCommunicate(db, inTE.Text())
							if err != nil {
								log.Default()
							}
							DB_Comunicate.Close(db)
							outTE.SetText(requestDB) // Print in window
						},
					},
				},
			},
			HSplitter{
				MinSize: Size{540, 480},
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
		},
	}.Run()
	return nil
}
