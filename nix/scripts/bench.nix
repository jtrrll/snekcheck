{
  go,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs all benchmark tests.";
  name = "bench";
  runtimeInputs = [go];
  text = ''
    cd "$SOURCE_ROOT"
    go test ./... -bench=.
  '';
}
