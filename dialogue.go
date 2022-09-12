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

/*
TODO: Update README.md
	TODO: Update userFlow Diagram
TODO: Prompt for scene Id first
TODO: Consider using the "make() function isntead of blank values
	REVIEW:- "https://www.includehelp.com/golang/make-function-with-examples.aspx"
TODO: Put blank name + dialogue for questions (requires large refactoring)
TODO: Add chapterObj struct
	TODO: After entering each scene, return the "dialogue forks", all of the "Next" fields from the choices
*/

func getInputX(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func newSceneObj(dOS []DialogueObj) (SceneObj, string) {
	reader := bufio.NewReader(os.Stdin)
	id, _ := getInputX("Input an id for this scene: ", reader)
	idNum64, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		fmt.Println("Must be a number")
	}
	idNum := int(idNum64)
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
	dO := addExtraFields(n, d)
	fmt.Println(dO, "Added to scene !")
	return dO
}
func addExtraFields(n string, d string) DialogueObj {
	reader := bufio.NewReader(os.Stdin)
	numExtraFields, _ := getInputX("Enter number of extra fields \n0 - name & dialogue only,\n1 - name, dialogue & background,\n2 - name, dialogue, question & options,\n3 - name, dialogue, question, options & background.\n: ", reader)
	switch numExtraFields {
	case "0":
		dO := DialogueObj{Name: n, Dialogue: d}
		return dO
	case "1":
		bg, _ := getInputX("Enter background url: ", reader)
		dO := DialogueObj{Name: n, Dialogue: d, Background: bg}
		return dO
	case "2":
		q, _ := getInputX("Enter question: ", reader)
		numOptions, _ := getInputX("Enter how many options there should be (max 8): ", reader)
		nO64, err := strconv.ParseInt(numOptions, 10, 0)
		if err != nil {
			fmt.Println("Must be a number")
		}
		nO := int(nO64)
		if nO == 0 || nO > 8 {
			return addExtraFields(n, d)
			/* Reprompt for number of fields if invalid amount
			- needs investigating: instead of reprompting back to the start, instead prompt for number options again.
			But, then question, q, would be added with no options... */
		}
		options := newOptionObjSlice(nO)
		dO := DialogueObj{Name: n, Dialogue: d, Question: q, Options: options}
		return dO
		/* case "3":
		   		bg,_:=("Enter background url: ",reader)
		   		q,_:=("Enter question: ",reader)
		   		numOptions,_:=("Enter how many options there should be: ",reader)
		   		options:=newOptionObjSlice(numOptions)
				dO := DialogueObj{Name: n, Dialogue: d, Background: bg Question: q, Options:options}
		   		return dO  */
	default:
		fmt.Println("Invalid option !")
		return addExtraFields(n, d)
	}
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
	_ = ioutil.WriteFile("scenes/scene_"+id+".json", scene, 0644)
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
func newOptionObjSlice(numOptions int) []OptionObj {
	reader := bufio.NewReader(os.Stdin)
	var oOS []OptionObj = make([]OptionObj, numOptions) // NOTE: Could refactor to make array instead of slice
	/* for i:=1;i<=numOptions;i++ {
		prompt:=fmt.Sprintf("Option %v:")
		text,_=getInputX()
	} */
	for index, _ := range oOS { // For each option in the newly-created slice
		/*  NOTE: Can use 2nd answer from https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go
		to loop through the fields of the current optionObj struct
		Loop through and declare a new variable of same name to be the field's name,
		and prompt for that field's value. Convert afterwards.
		*/

		fmt.Printf("For option %v: \n", index+1)
		t, _ := getInputX("Enter option text: ", reader)
		next, _ := getInputX("Enter id of the next scene: ", reader)
		n, _ := strconv.Atoi(next)
		luckChange, _ := getInputX("Enter option luck change: ", reader)
		lc, _ := strconv.Atoi(luckChange)
		maxLuck, _ := getInputX("Enter option min luck requirement: ", reader)
		ml, _ := strconv.Atoi(maxLuck)

		oOS[index].Text = t
		oOS[index].Next = n
		oOS[index].LuckChange = lc
		oOS[index].MinLuck = ml

	}
	return oOS
}
