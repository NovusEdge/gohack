package gohack

// Maintain a Command struct for all tools and access in "gohack.go":
var COMMANDS []CommandTemplate = []CommandTemplate{
	//portScanner
	CommandTemplate{
		Aliases:      []string{"ps", "pscanner", "PORTSCANNER", "portscanner", "PortScanner"},
		BinaryName:   "portScanner",
		Discription: "A port scanner -_-",
		IsFunctional: true,
	},

	//bannerGrabber
	CommandTemplate{
		Aliases:      []string{"bg", "bgrabber", "BANNERGRABBER", "bannergrabber", "BannerGrabber"},
		BinaryName:   "bannerGrabber",
		Discription: "A banner grabber",
		IsFunctional: true,
	},

	CommandTemplate{
		Aliases:      []string{"doc", "doctor", "DOCTOR", "Doctor"},
		BinaryName: "doctor",
		Discription: "The doctor tool helps fix broken/missing binaries.",
		IsFunctional: true,
	},
}
