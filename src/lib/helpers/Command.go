package gohack

// Command
/*
A struct to contain details for the main command line tools and options for the
main binary

Fields:
    BinaryName        string
    Aliases           []string
    PossibleArguments []string
    IsFunctional      bool
	SuppliedArguments &map[string]string
*/
type Command struct {
	BinaryName        string
	Aliases           []string
	PossibleArguments []string
	IsFunctional      bool
	SuppliedArguments &map[string]string
}

// Maintain a Command struct for all tools and access in "gohack.go":
var COMMANDS []Command = []Command{
	//portScanner
	Command{
		BinaryName:        "portScanner",
		Aliases:           []string{"ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"},
		PossibleArguments: []string{"start", "end", "timeout", "url", "protocol"},
		IsFunctional:      true,
		SuppliedArguments: nil,
	},

	//bannerGrabber
	Command{
		BinaryName:        "bannerGrabber",
		Aliases:           []string{"bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"},
		PossibleArguments: []string{"url"},
		IsFunctional:      true,
		SuppliedArguments: nil,
	},
}

// The useage has to be like:
/*
$ gohack <tool_name/alias> args ...
*/


// ExecuteCommand: ...
/*

*/
func (c *Command) ExecuteCommand() {

}

// UpdateArgs: ...
/*

*/
func (c *Command) UpdateArgs(map[string]string) {

}


func (c *Command) checkSuppliedArgs() {

}
