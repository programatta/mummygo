package gamelevel

//GameLevel constiene la información del nivel
type GameLevel struct {
	ID      int `json:"id"`
	Mummies int `json:"mummies"`
	Spells  int `json:"spells"`
	Potins  int `json:"potions"`
}
