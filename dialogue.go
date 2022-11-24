package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func getInput(prompt string) string {
	// Function that prompts the user, with the string passed in, and returns a string
	r := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(input)
}
func newSceneObj(dOS []DialogueObj) (SceneObj, string) {
	id := getInput("Input an id for this scene: ")
	idNum, err := strconv.Atoi(id)
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
	n := getInput("Enter name: ")
	d := getInput("Enter dialogue: ")
	dO := addExtraFields(n, d)
	fmt.Println(dO, "Added to scene !")
	return dO
}
func addExtraFields(n string, d string) DialogueObj {
	numExtraFields := getInput("Enter number of extra fields \n0 - name & dialogue only,\n1 - name, dialogue & background,\n2 - name, dialogue, question & options,\n3 - name, dialogue, question, options & background.\n: ")
	switch numExtraFields {
	case "0":
		dO := DialogueObj{Name: n, Dialogue: d}
		return dO
	case "1":
		bg := getInput("Enter background url: ")
		dO := DialogueObj{Name: n, Dialogue: d, Background: bg}
		return dO
	case "2":
		q := getInput("Enter question: ")
		numOptions := getInput("Enter how many options there should be (max 8): ")
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
	case "3":
		bg := getInput("Enter background url: ")
		q := getInput("Enter question: ")
		numOptions := getInput("Enter how many options there should be (max 8): ")
		nO64, err := strconv.ParseInt(numOptions, 10, 0)
		if err != nil {
			fmt.Println("Must be a number")
		}
		nO := int(nO64)
		if nO == 0 || nO > 8 {
			return addExtraFields(n, d)
		}
		options := newOptionObjSlice(nO)
		dO := DialogueObj{Name: n, Dialogue: d, Background: bg, Question: q, Options: options}
		return dO
	default:
		fmt.Println("Invalid option !")
		return addExtraFields(n, d)
	}
}
func main() {
	newDOS := newDialogueObjSlice()
	finalDOS := optionsPrompt(newDOS)
	nSO, id := newSceneObj(finalDOS) // NOTE: Remove id later
	// fmt.Printf("nSO = %v, of type: %T\n", nSO, nSO)
	fmt.Printf("id = %v, of type: %T\n", id, id)
	scene, _ := json.MarshalIndent(nSO, "", " ")
	action := getInput("Save to DB ?\n[Y] Yes / [N] No : ")
	if strings.ToLower(action) == "y" {
		PostToAPI(scene)
	}
	action = getInput("Write to file ?\n[Y] Yes / [N] No : ")
	if strings.ToLower(action) == "y" {
		WriteToFile(scene, id)
	}
	fmt.Println("Exiting ...")
}
func optionsPrompt(dOS []DialogueObj) []DialogueObj {
	opt := getInput("Choose option (a - add new dialogue entry, s - save scene): ")
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
	var oOS []OptionObj = make([]OptionObj, numOptions) // NOTE: Could refactor to make array instead of slice
	/* for i:=1;i<=numOptions;i++ {
		prompt:=fmt.Sprintf("Option %v:")
		text,_=getInputX()
	} */
	for index := range oOS { // For each option in the newly-created slice
		/*  NOTE: Can use 2nd answer from https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go
		to loop through the fields of the current optionObj struct
		Loop through and declare a new variable of same name to be the field's name,
		and prompt for that field's value. Convert afterwards.
		*/

		fmt.Printf("For option %v: \n", index+1)
		t := getInput("Enter option text: ")
		next := getInput("Enter id of the next scene: ")
		n, _ := strconv.Atoi(next)
		luckChange := getInput("Enter option luck change: ")
		lc, _ := strconv.Atoi(luckChange)
		maxLuck := getInput("Enter option min luck requirement: ")
		ml, _ := strconv.Atoi(maxLuck)

		oOS[index].Text = t
		oOS[index].Next = n
		oOS[index].LuckChange = lc
		oOS[index].MinLuck = ml

	}
	return oOS
}
func WriteToFile(s []byte, id string) {
	_ = os.WriteFile("scenes/scene_"+id+".json", s, 0644)
	fmt.Println("Scene successfully written to folder: /scenes")
}
func PostToAPI(s []byte) {
	fmt.Println("---------------------POSTED TO API --------------------")
	uri := "http://localhost:8081/scenes"
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(s))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// uri := "http://localhost:8081/articles"
	// res, err := http.Get(uri)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)

	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// client := &http.Client{}
	// res, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
}
