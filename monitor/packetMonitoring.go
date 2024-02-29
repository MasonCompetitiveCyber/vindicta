package monitor

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
	"github.com/google/gopacket/pcap"
)

// To Be Used Later
func JustTextBoxForNow(title string, textColor tcell.Color) *cview.TextView {
	b := cview.NewTextView()
	b.SetBorder(true)
	b.SetTitle(title)
	b.SetTextColor(textColor)
	b.SetBorderColor(tcell.ColorOrange)
	b.SetTitleAlign(cview.AlignLeft)
	return b
}

func interfaceInfo() *cview.DropDown {

	// Some styling
	dropDown := cview.NewDropDown()
	dropDown.SetDropDownBackgroundColor(tcell.ColorBlue)
	dropDown.SetDropDownTextColor(tcell.ColorBlack)
	dropDown.SetBorder(true)
	dropDown.SetBorderColor(tcell.ColorRed)
	dropDown.SetPadding(1, 0, 0, 0)
	dropDown.SetTitle("[black:aqua]Interface")
	dropDown.SetTitleAlign(cview.AlignLeft)
	dropDown.SetFieldWidth(0)
	dropDown.SetFieldBackgroundColor(tcell.ColorDarkOliveGreen)
	dropDown.SetFieldTextColor(tcell.ColorBlack)
	dropDown.SetDropDownSelectedBackgroundColor(tcell.ColorPurple)
	dropDown.SetAlwaysDrawDropDownSymbol(false)

	// Find all available network interfaces
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	// Create options for the DropDown using the network interface names
	var options []*cview.DropDownOption

	// Range through all available interfaces
	for _, iface := range ifaces {
		options = append(options, cview.NewDropDownOption(" "+iface.Name))
	}

	// Set the options in the dropdown
	dropDown.SetOptions(nil, options...)

	// Return the names of those interfaces
	return dropDown

}

func filterInfo() *cview.InputField {
	// Just some styling
	vw := cview.NewInputField()
	vw.SetTitle("[black:aqua]Filter")
	vw.SetBorder(true)
	vw.SetBorderColor(tcell.ColorRed)
	vw.SetTitleAlign(cview.AlignLeft)
	vw.SetFieldWidth(0)
	vw.SetFieldBackgroundColor(tcell.ColorBlueViolet)
	vw.SetFieldBackgroundColorFocused(tcell.ColorOrange)
	vw.SetFieldTextColorFocused(tcell.ColorBlack)
	vw.SetFieldTextColor(tcell.ColorWhite)
	vw.SetPlaceholderTextColor(tcell.ColorWhite)
	vw.SetPlaceholderTextColorFocused(tcell.ColorBlack)
	vw.SetPlaceholder("tcp port 80")
	vw.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			return
		}
	})

	return vw
}

func PacketMonitor(cviewApp *cview.Application) *cview.Flex {
	// Main Flex UI
	view := cview.NewFlex()
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorBlue)
	view.SetDirection(cview.FlexRow)

	// Create SubFlex for first row
	firstRow := cview.NewFlex()
	firstRow.SetBorderColor(tcell.ColorAquaMarine)
	firstRow.SetDirection(cview.FlexColumn)

	// Call Em
	iff := interfaceInfo()
	rF := filterInfo()

	// Add interfaces info and packet filter rules
	firstRow.AddItem(iff, 0, 1, false)
	firstRow.AddItem(rF, 0, 4, false)

	// Add them to main Flex
	view.AddItem(firstRow, 0, 1, false)
	view.AddItem(JustTextBoxForNow("[aqua:black] Packet Dump Goes Here With Network Layers [1-4]", tcell.ColorGreenYellow), 0, 10, false)

	return view
}
