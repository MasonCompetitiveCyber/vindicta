package monitor

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func ConfigureServices(myCviewApp *cview.Application) *cview.Flex {

    myFlex := cview.NewFlex() 
    myFlex.SetBorder(true)
    myFlex.SetBorderColor(tcell.ColorPurple)

    subFlex := cview.NewFlex()
    subFlex.SetDirection(cview.FlexRow)
    subFlex.AddItem(Button("List"), 0, 1, false)
	subFlex.AddItem(Button("Add"), 0, 1, false)
	subFlex.AddItem(Button("Edit"), 0, 1, false)
	subFlex.AddItem(Button("Delete"), 0, 1, false)
    subFlex.AddItem(TextBox("TBD"), 0, 5, false)


    myFlex.AddItem(subFlex, 0, 1, false)
	myFlex.AddItem(TextBox("Services"), 0, 8, false)

    return myFlex
}
