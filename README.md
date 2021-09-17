# Prop Logic Symbol Swapper
Swaps ASCII characters for the proper utf8 logic symbols
### For Example:
((A or C) => (B and !E))
#### Becomes:
((A ∨ C) ⇒ (B ∧ ¬E))

## Accepted Syntax:
**All variables must be in lowercase**

¬ = "~", "!", "not"

⇒ = "=>"

∧ = "&", "^", "and"

### To Install:
clone the repository and run "go build"

### Running:
run ./proplogicsymbolswapper filename
and you'll receive swapped-filename.txt in the directory that the program is ran in
