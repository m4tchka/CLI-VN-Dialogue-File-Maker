/*
Create a sceneObj

# Within that, add an id and create a scene slice

Within that scene slice, create 1 new dialogueObj
Prompt the user to input Name
Prompt the user to enter Dialogue

	Prompt user for additional keys -
	valid combinations of optional fields:

	None	0

	B		1

	Q,O		2 - Denotes the end of current scene

	B,Q,O	3 - Denotes the end of current scene (rare)

	0 or 1:
	Prompt user if need more dialogueObjs -
		Y: Back to "create 1 new dialogueObj"
		N: Save (Rare - End of chapter/ story)

	2 or 3:
	(2 - Prompt for Background first)
	Prompt user for number of OptionObjs (how many choices should be available)
	Prompt user for each field of OptionObj1
	Add OptionObj1 to Options slice field of current DialogueObj
	Repeat for remaining number of OptionObjs
*/
package main
