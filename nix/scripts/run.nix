{
  go,
  gomod2nix,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs the project.";
  name = "run";
  runtimeInputs = [
    go
    gomod2nix
  ];
  text = ''
    (cd "$SOURCE_ROOT" && go mod tidy && gomod2nix)
    nix run "$PROJECT_ROOT" -- "$@"
  '';
}
