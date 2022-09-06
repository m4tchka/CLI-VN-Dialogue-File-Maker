package main

import (
	"fmt"
)

func newSceneObj(id int) SceneObj {
	sObj := SceneObj{
		id:    id,
		scene: SceneArr{},
	}
	fmt.Println("sObj= ", sObj)
	return sObj
}
func newSceneArr() SceneArr {
	sArr := SceneArr[]DialogueObj{
		DialogueObj{
			Name:     "Test name",
			Dialogue: "...",
		},
	}
	return sArr
}

/* func newDialogueObj() {
	fmt.Println
} */
