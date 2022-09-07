package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputX(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	//Returns everything typed into the terminal, as a string, right before enter is pressed. Must be in SINGLE quotes, and assign it to the variable name (or error if there is an error).
	return strings.TrimSpace(input), err
	//Update the name variable with whitespace removed
}
func newSceneObj(SArr SceneArr) {
	reader := bufio.NewReader(os.Stdin)
	id, _ := getInputX("Input an id for this scene: ", reader)
	idNum64, err := strconv.ParseInt(id, 10, 0)
	idNum := int(idNum64)
	/* 	var abc int = 50 */
	if err != nil {
		fmt.Println("Must be a number")
	}

	var nSO SceneObj = SceneObj{id: idNum, scene: SArr}
	fmt.Println(nSO)
	/*
	   fmt.Printf("idNum64, %v type: %T \n", idNum64, idNum64)
	   fmt.Printf("abc type: %T \n", abc)
	   fmt.Printf("idNum, %v type: %T \n", idNum, idNum)
	*/
	/*
		id, err := strconv.ParseInt(id, 10, 0)
		   	//fmt.Println(id)
		   	sObj := SceneObj{s
		   		id:    id,
		   		scene: SceneArr,
		   	}
	*/
}
func newSceneArr() SceneArr {
	sArr := SceneArr{}
	newDialogueObj()
	sArr.Scene = append(sArr.Scene, newDialogueObj())

	return sArr
}
func newDialogueObjSlice() []DialogueObj {
	var DObjSlice []DialogueObj
	firstDObj := newDialogueObj()
	DObjSlice = append(DObjSlice, firstDObj)
	return DObjSlice
}
func appendToSlice(prev []DialogueObj) []DialogueObj {
	dObjToBeAppended := newDialogueObj()
	new := append(prev, dObjToBeAppended)
	return new
}
func newDialogueObj() DialogueObj {
	reader := bufio.NewReader(os.Stdin)
	n, _ := getInputX("Enter name: ", reader)
	d, _ := getInputX("Enter dialogue: ", reader)
	dObj := DialogueObj{Name: n, Dialogue: d}
	fmt.Println(dObj)
	return dObj
}
func main() {
	var test []DialogueObj
	test2 := appendToSlice(test)
	fmt.Println("test2: ", test2)
	/* newDialogueObj() */
	/* newSceneObj() */
}
