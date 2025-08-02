{
  snekcheck,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Generates a demo GIF.";
  name = "demo";
  runtimeInputs = [snekcheck];
  text = ''
    mkdir --parents "$PROJECT_ROOT"/demo
    for i in $(seq 1 3); do touch "$PROJECT_ROOT"/demo/"$i"valid; done;
    for i in $(seq 1 3); do touch "$PROJECT_ROOT"/demo/"$i"InVaLiD; done;

    build
    cat <<EOF | vhs -
    Output demo.gif

    Set FontFamily "Hack Nerd Font Mono"
    Set FontSize 28
    Set Padding 10
    Set Theme "catppuccin-frappe"
    Set TypingSpeed 100ms

    Set Width 800
    Set Height 450

    Require snekcheck

    Sleep 1000ms

    Type "snekcheck ."
    Sleep 500ms
    Enter

    Sleep 1000ms

    Type "snekcheck --fix ."
    Sleep 500ms
    Enter

    Sleep 10000ms
    EOF

    rm --force --recursive "$PROJECT_ROOT"/demo
  '';
}
