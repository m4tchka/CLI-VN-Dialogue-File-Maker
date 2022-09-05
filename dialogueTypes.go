package main

type SceneObj struct {
	Scene []DialogueObj `json:"Scene"`
}
type DialogueObj struct { // Only Name & Dialogue fields mandatory
	Name       string    `json:"Name"`
	Dialogue   string    `json:"Dialogue"`
	Background string    `json:"Background,omitempty"`
	Question   string    `json:"Question,omitempty"`
	Options    OptionObj `json:"Options,omitempty"`
	//Sprites SpriteObj
}
type OptionObj struct { // All fields mandatory
	Text       string `json:"Text"`
	Next       int    `json:"Next"`
	LuckChange int    `json:"LuckChange"`
	MinLuck    int    `json:"MinLuck"`
}
