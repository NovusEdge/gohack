package gohack

// Maintain a Command struct for all tools and access in "gohack.go":
var COMMANDS []CommandTemplate = []CommandTemplate{
	//portScanner
	CommandTemplate{
		Aliases:           []string{"ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"},
		BinaryName:        "portScanner",
		IsFunctional:      true,
		PossibleArguments: []string{"start", "end", "timeout", "url", "protocol"},
	},

	//bannerGrabber
	CommandTemplate{
		Aliases:           []string{"bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"},
		BinaryName:        "bannerGrabber",
		IsFunctional:      true,
		PossibleArguments: []string{"url"},
	},
}
