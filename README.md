# CLI-VN-Dialogue-File-Maker

## Overview
This CLI tool is used to write dialogue for "planned-vn-style-game" !
bills.go & billsLogic.go were created using the Go tutorial series by TheNetNinja, found [here.](https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM)
Afterwards, inspired by this, I created my own version in order to make json files of dialogue.

## Detail
Instead of making methods which directly mutate the struct variable that they're called on, as in the bills CLI tutorial, I instead took the approach of making normal functions, which took in a (copy of a) struct variable, prompted users for input to put into the fields of that struct, prompted users if they wanted to add more entries (in a slice), and returned the updated copy of that struct variable.

## User Flow diagram
TBD

## Example of usage 
To launch, type into the terminal:
```bash
clear ; go run dialogue.go dialogueTypes.go
```

## 
