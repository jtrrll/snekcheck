{
  go,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs all unit tests.";
  name = "unit";
  runtimeInputs = [go];
  text = ''
    cd "$SOURCE_ROOT"
    go test --cover ./...
  '';
}
