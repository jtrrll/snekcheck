{
  go,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs all end-to-end tests.";
  name = "e2e";
  runtimeInputs = [go];
  text = ''
    cd "$SPEC_ROOT"
    go test -count 1 ./...
  '';
}
