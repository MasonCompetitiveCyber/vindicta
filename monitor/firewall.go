package monitor

import (
	"log"
	"strings"
	"os/exec"

	"github.com/google/nftables"
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func ConfigureFirewall(myCviewApp *cview.Application) *cview.Flex {

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


    dropFlex := cview.NewFlex()
    dropFlex.SetDirection(cview.FlexRow)

    dropFlex.AddItem(DropDownMenu("Tables: "), 0, 1, false)
    dropFlex.AddItem(DropDownMenu("Chains: "), 0, 1, false)
    dropFlex.AddItem(DropDownMenu("Rules: "), 0, 1, false)

    myFlex.AddItem(subFlex, 0, 1, false)
    myFlex.AddItem(dropFlex, 0, 8, false)
    //myFlex.AddItem(TextBox("Nftables"), 0, 8, false)

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


func Button(btnName string) *cview.Button {

    button := cview.NewButton(btnName)
    button.SetRect(0, 0, 1, 1)
    button.SetBorder(true)
    button.SetBorderColor(tcell.ColorBlue)
    button.SetTitleColor(tcell.ColorYellow)
    button.SetBackgroundColor(tcell.ColorBlack)
    button.SetLabelColorFocused(tcell.ColorWhite)

    return button
}

func DropDownMenu(label string) *cview.DropDown {
    dropdown := cview.NewDropDown()
    dropdown.SetLabel(label)
    dropdown.SetBorder(true)
    dropdown.SetBorderColor(tcell.ColorPurple)
    dropdown.SetFieldBackgroundColorFocused(tcell.ColorGrey)

    if label == "Tables: " {
        tbls := GetAllTables()
        dropdown.SetOptionsSimple(nil, tbls...)

    } else if label == "Chains: " {
        chns := GetAllChains()
        dropdown.SetOptionsSimple(nil, chns...)

    } else if label == "Rules: " {
        rls := GetAllRules()
	dropdown.SetOptionsSimple(nil, rls...)

    } else {
        cview.NewDropDownOption("NO OPTION AVAILABLE")
    }
    return dropdown

}

func GetAllTables() []string {
    firewall := GetFirewall()

    var tableNames []string

    for f := range firewall {
	for t := range firewall[f] {
        	tableNames = append(tableNames, t)
	}
    }

    return tableNames

}

func GetAllChains() []string {
    firewall := GetFirewall()

    var chainNames []string

    for f := range firewall {
	for t := range firewall[f] {
	    for c := range firewall[f][t] {
        	chainNames = append(chainNames, firewall[f][t][c])
	    }
	}
    }

    return chainNames
}

func GetAllRules() []string {
    firewall := GetFirewall()

    var allRules []string

    for f := range firewall {
	for t := range firewall[f] {
	    for c := range firewall[f][t] {
		cmd := exec.Command("nft", "list", "chain", f, t, firewall[f][t][c]) // calls nft list command with family, table, and chain
		out, err := cmd.Output()
		if err != nil {
		    log.Fatal(err)
		}
		ruleList := string(out)
		rules := strings.Split(ruleList, "\n") // splits each rule into list
		rules = rules[2:len(rules)-3] // gets rid of headers, and footers (table name / excess '}'s
		for r := 0; r < len(rules); r++ {
		    allRules = append(allRules, rules[r][2:len(rules[r])]) // appends rules to list of all rules, removing the initial 2 tabs
		}
	    }
	}
    }
    return allRules
}

func GetFirewall() map[string]map[string][]string {
    cc, err := nftables.New()
    if err != nil {
        log.Fatal(err)
    }

    tab, err := cc.ListTables()
    if err != nil {
        log.Fatal(err)
    }

    ch, err := cc.ListChains()
    if err != nil {
        log.Fatal(err)
    }
    
    // creates a map of all nftables families (number : family name)
    familyMap := map[int]string{
    1: "inet",
    2: "ip",
    3: "arp",
    5: "netdev",
    7: "bridge",
    10: "ip6",
    }

    allNFT := make(map[string]map[string][]string) // creates mapping for map[family][table][chain]. family & table values are strings, chain is a list
    
    for k := range familyMap { // loops through families
        tables := make(map[string][]string) // creates empty mapping of tables
        for t := 0; t < len(tab); t++ { // loops through all tables
            for c := 0; c < len(ch); c++ { // loops through all chains
                if ch[c].Table.Name == tab[t].Name && int(ch[c].Table.Family) == k { // verifies table name & family (no duplicates)
                    tables[tab[t].Name] = append(tables[tab[t].Name], ch[c].Name) // appends all chains of the same table to a mapping
                }
            }
        }
        if len(tables) != 0 { // verifies the table family isnt empty (removes redundant families)
            allNFT[familyMap[k]] = tables // appends current tables & chains mapping to specified family
        }
    }
    return allNFT
}
