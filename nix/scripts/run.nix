{
  go,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs the project.";
  name = "run";
  runtimeInputs = [go];
  text = ''
    (cd "$SOURCE_ROOT" && go mod tidy)
    nix run "$PROJECT_ROOT" -- "$@"
  '';
}
