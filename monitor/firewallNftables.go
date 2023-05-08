package monitor

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func ConfigureFirewall(myCviewApp *cview.Application) *cview.Flex {

    myFlex := cview.NewFlex() 
    myFlex.SetBorder(true)
    myFlex.SetBorderColor(tcell.ColorPurple)

    subFlex := cview.NewFlex()
    subFlex.SetDirection(cview.FlexRow)
    subFlex.AddItem(FirewallButton("List"), 0, 1, false)
	subFlex.AddItem(FirewallButton("Add"), 0, 1, false)
	subFlex.AddItem(FirewallButton("Edit"), 0, 1, false)
	subFlex.AddItem(FirewallButton("Delete"), 0, 1, false)
    subFlex.AddItem(TextBox("TBD"), 0, 7, false)


    myFlex.AddItem(subFlex, 0, 1, false)
	myFlex.AddItem(TextBox("Nftables"), 0, 8, false)

    return myFlex
}

func TextBox(title string) *cview.TextView {
	b := cview.NewTextView()
	b.SetBorder(true)
    b.SetBorderColor(tcell.ColorBlue)
    b.SetTitleColor(tcell.ColorYellow)
	b.SetTitle(title)
	return b
}


func FirewallButton(btnName string) *cview.Button {

    button := cview.NewButton(btnName)
    button.SetRect(0, 0, 1, 1)
    button.SetBorder(true)
    button.SetBorderColor(tcell.ColorBlue)
    button.SetTitleColor(tcell.ColorYellow)
    button.SetBackgroundColor(tcell.ColorBlack)
    button.SetLabelColorFocused(tcell.ColorWhite)

    return button
}
