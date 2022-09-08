package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getInputX(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func newSceneObj(dOS []DialogueObj) (SceneObj, string) {
	reader := bufio.NewReader(os.Stdin)
	id, _ := getInputX("Input an id for this scene: ", reader)
	idNum64, err := strconv.ParseInt(id, 10, 0)
	idNum := int(idNum64)
	if err != nil {
		fmt.Println("Must be a number")
	}
	var nSO SceneObj = SceneObj{Id: idNum, Scene: dOS}
	fmt.Println(nSO)
	return nSO, id
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
	// fmt.Printf("newDOS = %v, of type = %T\n", newDOS, newDOS)
	finalDOS := optionsPrompt(newDOS)
	// fmt.Printf("finalDOS = %v, of type: %T\n", finalDOS, finalDOS)
	nSO, id := newSceneObj(finalDOS) // NOTE: Remove id later
	// fmt.Printf("nSO = %v, of type: %T\n", nSO, nSO)
	fmt.Printf("id = %v, of type: %T\n", id, id)
	scene, _ := json.MarshalIndent(nSO, "", " ")
	_ = ioutil.WriteFile("dialogues/scene_"+id+".json", scene, 0644)
}
func optionsPrompt(dOS []DialogueObj) []DialogueObj {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInputX("Choose option (a - add new dialogue entry, s - save scene): ", reader)
	switch opt {
	case "a":
		new := appendToSlice(dOS)
		return optionsPrompt(new)
	case "s":
		fmt.Println("Implemented soon (tm) !")
		fmt.Println("Returning slice of dObjs:", dOS)
		return dOS
	default:
		fmt.Println("Invalid option !")
		return optionsPrompt(dOS)
	}
}
