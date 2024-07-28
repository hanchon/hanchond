package hermes

import "github.com/hanchon/hanchond/playground/filesmanager"

type Hermes struct{}

func NewHermes() *Hermes {
	_ = filesmanager.CreateHermesFolder()
	h := &Hermes{}
	h.initHermesConfig()
	return h
}
