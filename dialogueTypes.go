package main

// type chObj struct {}
type SceneObj struct {
	id    int           `json:"id"`
	scene []DialogueObj `json:"Scene"`
}
type DialogueObj struct {
	// Only Name & Dialogue fields mandatory, rest will remain as empty strings/ slice until conversion to JSON upon where they will be removed.
	Name       string      `json:"Name"`
	Dialogue   string      `json:"Dialogue"`
	Background string      `json:"Background,omitempty"`
	Question   string      `json:"Question,omitempty"`
	Options    []OptionObj `json:"Options,omitempty"`
	//Sprites SpriteObj
}
type OptionObj struct {
	// All fields mandatory
	Text       string `json:"Text"`
	Next       int    `json:"Next"`
	LuckChange int    `json:"LuckChange"`
	MinLuck    int    `json:"MinLuck"`
}
