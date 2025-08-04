{
  bashInteractive,
  snekcheck,
  vhs,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Generates a demo GIF.";
  name = "demo";
  runtimeInputs = [
    bashInteractive
    snekcheck
    vhs
  ];
  text = ''
    mkdir --parents "$PROJECT_ROOT"/demo
    for i in $(seq 1 3); do touch "$PROJECT_ROOT"/demo/"$i"valid; done;
    for i in $(seq 1 3); do touch "$PROJECT_ROOT"/demo/"$i"InVaLiD; done;

    pwd="$PWD"
    cd "$PROJECT_ROOT"/demo

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
    Require find

    Sleep 1000ms

    Type "find . -exec snekcheck {} +"
    Sleep 500ms
    Enter

    Sleep 1000ms

    Type "find . -exec snekcheck --fix {} +"
    Sleep 500ms
    Enter

    Sleep 10000ms
    EOF

    mv demo.gif "$PROJECT_ROOT"/demo.gif
    cd "$pwd"
    rm --force --recursive "$PROJECT_ROOT"/demo
  '';
}
