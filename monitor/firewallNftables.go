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
	myFlex := setMainMenu()
    	return myFlex
}

func setMainMenu(flexs ...*cview.Flex) *cview.Flex {
	var myFlex *cview.Flex
	if len(flexs) != 0 {
		myFlex = flexs[0]
		for i := 0; i < len(flexs); i++ {
			if i !=0 {
				myFlex.RemoveItem(flexs[i])
			}
		}
	} else {
    		myFlex = cview.NewFlex()
    		myFlex.SetBorder(true)
    		myFlex.SetBorderColor(tcell.ColorPurple)
	}
    	subFlex := cview.NewFlex()
    	subFlex.SetDirection(cview.FlexRow)
    
    	listButton := Button("List")
    	listButton.SetSelectedFunc(func() {viewFirewall(myFlex, subFlex)})
	addButton := Button("Add")
	addButton.SetSelectedFunc(func() {addRule(myFlex, subFlex)})
	deleteButton := Button("Delete")
	deleteButton.SetSelectedFunc(func() {deleteRule(myFlex, subFlex)})
	
    	subFlex.AddItem(listButton, 0, 1, false)
    	subFlex.AddItem(addButton, 0, 1, false)
    	subFlex.AddItem(Button("Edit"), 0, 1, false)
    	subFlex.AddItem(deleteButton, 0, 1, false)
    	subFlex.AddItem(TextBox("TBD"), 0, 5, false)
    
    	myFlex.AddItem(subFlex, 0, 1, false)

	return myFlex
}


func viewFirewall(myFlex *cview.Flex, subFlex *cview.Flex) {
	myFlex.RemoveItem(subFlex)
	nsubFlex := cview.NewFlex()
	fwFlex := cview.NewFlex()
	
	backButton := Button("Back")
	backButton.SetSelectedFunc(func() {setMainMenu(myFlex, nsubFlex, fwFlex)})

	nsubFlex.SetDirection(cview.FlexRow)
	nsubFlex.AddItem(backButton, 0, 1, false)
	nsubFlex.AddItem(TextBox(""), 0, 8, false)


	cmd := exec.Command("nft", "list", "ruleset")
	out, err := cmd.Output()
	if err != nil{
		log.Fatal(err)
	}
	ruleList := TextBox("FIREWALL RULES")
	ruleList.SetText(string(out))
	ruleList.SetScrollable(true)
	ruleList.SetWrap(true)
	ruleList.SetWordWrap(true)
	
	fwFlex.SetDirection(cview.FlexRow)
	fwFlex.AddItem(ruleList, 0, 1, false)

	myFlex.SetDirection(cview.FlexColumn)
    	myFlex.AddItem(nsubFlex, 0, 1, false)
	myFlex.AddItem(fwFlex, 0, 8, false)
	
	return
}

func deleteRule(flexs ...*cview.Flex) {
	var myFlex *cview.Flex
	myFlex = flexs[0]
	for i := 0; i < len(flexs); i++ {
		if i !=0 {
			myFlex.RemoveItem(flexs[i])
		}
	}

	subFlex := cview.NewFlex()
	delFlex := cview.NewFlex()
	
	backButton := Button("Back")
	backButton.SetSelectedFunc(func() {setMainMenu(myFlex, subFlex, delFlex)})

	subFlex.SetDirection(cview.FlexRow)
	subFlex.AddItem(backButton, 0, 1, false)
	tbox := TextBox("Example")
	tbox.SetText("inet table1 chain3 6\n\n------------\n\n[Family] [Table] [Chain] [Handle#]")
	tbox.SetWrap(true)
	tbox.SetWordWrap(true)
	
	subFlex.AddItem(tbox, 0, 8, false)


	cmd := exec.Command("nft", "--handle", "list", "ruleset")
	out, err := cmd.Output()
	if err != nil{
		log.Fatal(err)
	}
	ruleList := TextBox("FIREWALL RULES")
	ruleList.SetText(string(out))
	ruleList.SetScrollable(true)
	ruleList.SetWrap(true)
	ruleList.SetWordWrap(true)
	
	inputBox := cview.NewForm()
	inputBox.SetHorizontal(true)
	inputBox.AddInputField("Delete Rule", "family table chain handle", 100, nil, nil)
	inputBox.AddButton("Delete", func() {
		pushDelete(inputBox.GetFormItem(0).(*cview.InputField).GetText()) 
		deleteRule(myFlex, subFlex, delFlex)
	})

	delFlex.SetDirection(cview.FlexRow)
	delFlex.AddItem(ruleList, 0, 8, false)
	delFlex.AddItem(inputBox, 0, 1, false)

	myFlex.SetDirection(cview.FlexColumn)
    	myFlex.AddItem(subFlex, 0, 1, false)
	myFlex.AddItem(delFlex, 0, 8, false)
	
	return
}

func addRule(flexs ...*cview.Flex) {
	var myFlex *cview.Flex
	myFlex = flexs[0]
	for i := 0; i < len(flexs); i++ {
		if i !=0 {
			myFlex.RemoveItem(flexs[i])
		}
	}
	
	subFlex := cview.NewFlex()
	addFlex := cview.NewFlex()
	
	backButton := Button("Back")
	backButton.SetSelectedFunc(func() {setMainMenu(myFlex, subFlex, addFlex)})

	subFlex.SetDirection(cview.FlexRow)
	subFlex.AddItem(backButton, 0, 1, false)
	tbox := TextBox("Example")
	tbox.SetText("inet table1 chain3 ip saddr 8.8.8.8 accept\n\n------------\n\n[Family] [Table] [Chain] [New rule]")
	tbox.SetWrap(true)
	tbox.SetWordWrap(true)
	subFlex.AddItem(tbox, 0, 8, false)


	cmd := exec.Command("nft", "list", "ruleset")
	out, err := cmd.Output()
	if err != nil{
		log.Fatal(err)
	}
	ruleList := TextBox("FIREWALL RULES")
	ruleList.SetText(string(out))
	ruleList.SetScrollable(true)
	ruleList.SetWrap(true)
	ruleList.SetWordWrap(true)
	
	inputBox := cview.NewForm()
	inputBox.SetHorizontal(true)
	inputBox.AddInputField("Add Rule", "family table chain rule", 100, nil, nil)
	inputBox.AddButton("Add", func() {
		pushAdd(inputBox.GetFormItem(0).(*cview.InputField).GetText()) 
		addRule(myFlex, subFlex, addFlex)
	})

	addFlex.SetDirection(cview.FlexRow)
	addFlex.AddItem(ruleList, 0, 8, false)
	addFlex.AddItem(inputBox, 0, 1, false)

	myFlex.SetDirection(cview.FlexColumn)
    	myFlex.AddItem(subFlex, 0, 1, false)
	myFlex.AddItem(addFlex, 0, 8, false)
	
	return
}

func pushAdd(ruleSet string) {
	ruletemp := strings.Split(ruleSet, " ")
	var newrule string
	var rule []string
	for i := 0; i < len(ruletemp); i++ {
		if i < 3 {
			rule = append(rule, ruletemp[i])
		} else {
			newrule += ruletemp[i]
			if i != len(ruletemp)-1 {
				newrule +=  " "
			}
		}
	}
	rule = append(rule, newrule)
	cmd := exec.Command("nft", "add", "rule", rule[0], rule[1], rule[2], rule[3])
	if err := cmd.Run(); err != nil {
		return
	}
}

func pushDelete(ruleSet string) {
	rule := strings.Split(ruleSet, " ")
	if len(rule) != 4 {
		return
	}
	cmd := exec.Command("nft", "delete", "rule", rule[0], rule[1], rule[2], "handle", rule[3])
	if err := cmd.Run(); err != nil {
		return
	}
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
