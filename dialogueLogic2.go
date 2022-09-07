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
func newSceneObj(sA SceneArr) SceneObj {
	reader := bufio.NewReader(os.Stdin)
	id, _ := getInputX("Input an id for this scene: ", reader)
	idNum64, err := strconv.ParseInt(id, 10, 0)
	idNum := int(idNum64)
	if err != nil {
		fmt.Println("Must be a number")
	}
	var nSO SceneObj = SceneObj{id: idNum, scene: sA}
	fmt.Println(nSO)
	return nSO
}
func newSceneArr(dOS []DialogueObj) SceneArr {
	sArr := SceneArr{}
	newDialogueObj()
	sArr.Scene = dOS
	return sArr
}
func newDialogueObjSlice() []DialogueObj {
	var dOS []DialogueObj
	dO1 := newDialogueObj()
	dOS = append(dOS, dO1)
	return dOS
}
func appendToSlice(prev []DialogueObj) []DialogueObj {
	dO := newDialogueObj()
	newDOS := append(prev, dO)
	return newDOS
}
func newDialogueObj() DialogueObj {
	reader := bufio.NewReader(os.Stdin)
	n, _ := getInputX("Enter name: ", reader)
	d, _ := getInputX("Enter dialogue: ", reader)
	dO := DialogueObj{Name: n, Dialogue: d}
	fmt.Println(dO)
	return dO
}
func main() {
	newDOS := newDialogueObjSlice()
	fmt.Printf("newDOS = %v, of type = %T\n", newDOS, newDOS)
	finalDOS := optionsPrompt(newDOS)
	fmt.Printf("finalDOS = %v, of type: %T\n", finalDOS, finalDOS)
	nSA := newSceneArr(finalDOS)
	fmt.Printf("nSA = %v, of type: %T\n", nSA, nSA)
	nSO := newSceneObj(nSA)
	fmt.Printf("nSO = %v, of type: %T\n", nSO, nSO)
	/* newDialogueObj() */
	/* newSceneObj() */
}
func optionsPrompt(dOS []DialogueObj) []DialogueObj {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInputX("Choose option (a - add new dialogue entry, s - save scene): ", reader)
	switch opt {
	case "a":
		new := appendToSlice(dOS)
		optionsPrompt(new)
	case "s":
		fmt.Println("Implemented soon (tm) !")
		fmt.Println("Returning slice of dObjs:", dOS)
		return dOS
	default:
		fmt.Println("Invalid option !")
		optionsPrompt(dOS)
	}
	fmt.Println("How did you get here !?")
	return dOS
}
