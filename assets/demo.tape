# Regenerate with: vhs assets/demo.tape (from the repository root)
# Requires HOTPEPPER_API_KEY in the environment and jq on PATH.

Output assets/demo.gif

Set Shell "bash"
Set FontSize 22
Set Width 1440
Set Height 760
Set Padding 24
Set Theme "Catppuccin Mocha"
Set TypingSpeed 45ms

Hide
Type `export D=/tmp/hpp-demo && rm -rf "$D" && mkdir -p "$D" && go build -o "$D/hpp" . && cd "$D" && export PATH="$D:$PATH" PS1='\[\e[1;35m\]$ \[\e[0m\]' && clear`
Enter
Sleep 4s
Show

Type "hpp area middle --keyword 渋谷 --format table"
Sleep 400ms
Enter
Sleep 2s

Type "hpp genre --format table"
Sleep 400ms
Enter
Sleep 4s

Type "clear"
Enter
Type "hpp search --middle-area Y030 --genre G013 --lunch --count 3 | jq '.results.shop[] | {name, budget: .budget.average, url: .urls.pc}'"
Sleep 400ms
Enter
Sleep 4s

Type "clear"
Enter
Type `hpp search --middle-area Y030 --genre G001 --private-room --free-drink --count 5 | jq -r '.results.shop[] | .name'`
Sleep 400ms
Enter
Sleep 4s
