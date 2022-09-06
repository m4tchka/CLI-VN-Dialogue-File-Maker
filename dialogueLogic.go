package main

import (
	"fmt"
)
/* func newSceneArr() SceneArr {
	sArr := SceneArr[]DialogueObj{
		DialogueObj{
			Name:     "Test name",
			Dialogue: "...",
		},
	}
	return sArr
} 
*/
func createSceneObj() SceneObj {
	reader := bufio.NewReader(os.Stdin)
	id,_:=getInput("Input an id for this scene: ",reader)

	sObj := SceneObj{
		id:    id,
		scene: []DialogueObj{addDialogueObj()},
	}
	fmt.Println("sObj= ", sObj)
	return sObj
}
func (sArr *SceneArr)addDialogueObj() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter name: ", reader)
	dialogue, _ := getInput("Enter dialogue: ", reader)
	numAdditionalFields,_:=getInput("How many additional fields ?",reader)
	addAdditionalFields(numAdditionalFields)
	dObj := {Name:name, Dialogue:dialogue}
	sArr:= append(sArr,dObj)

}
func addAdditionalFields(n int) {
	reader := bufio.NewReader(os.Stdin)
	switch n {
	case 0:
		return nil
	case 1:
		bg,_:=("Enter background url: ",reader)
		return bg
	case 2:
		q,_:=("Enter question: ",reader)
		numOptions,_:=("Enter how many options there should be: ",reader)
		o:=createOptionsArr(numOptions)
		return q,o 
	case 3:
		bg,_:=("Enter background url: ",reader)
		q,_:=("Enter question: ",reader)
		numOptions,_:=("Enter how many options there should be: ",reader)
		o:=createOptionsArr(numOptions)
		return bg,q,o 
	default:addAdditionalFields(n)
	}
}
func createOptionsArr(n int) []OptionObj {
	o:=[]OptionObj
	for i:=0; i<
	o:=append(o,)
	return o
}