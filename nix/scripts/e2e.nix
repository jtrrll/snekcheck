{
  go,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Runs all end-to-end tests.";
  name = "e2e";
  runtimeInputs = [go];
  text = ''
    build
    cd "$SPEC_ROOT"
    go test -count 1 ./...
  '';
}
