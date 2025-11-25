#!/usr/bin/env nu

mkdir ($env.PROJECT_ROOT | path join demo)
for i in 1..3 {
    touch ($env.PROJECT_ROOT | path join demo $"($i)valid")
}
for i in 1..3 {
    touch ($env.PROJECT_ROOT | path join demo $"($i)InVaLiD")
}

let pwd = $env.PWD
cd ($env.PROJECT_ROOT | path join demo)

let vhs_script = "Output demo.gif

Set FontFamily \"Hack Nerd Font Mono\"
Set FontSize 28
Set Padding 10
Set Theme \"catppuccin-frappe\"
Set TypingSpeed 100ms

Set Width 800
Set Height 450

Require snekcheck
Require find

Sleep 1000ms

Type \"find . -exec snekcheck {} +\"
Sleep 500ms
Enter

Sleep 1000ms

Type \"find . -exec snekcheck --fix {} +\"
Sleep 500ms
Enter

Sleep 10000ms"

$vhs_script | vhs -

mv demo.gif ($env.PROJECT_ROOT | path join demo.gif)

cd $pwd
rm -rf ($env.PROJECT_ROOT | path join demo)
