{
  go,
  golangci-lint,
  snekcheck,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Lints the project.";
  name = "lint";
  runtimeInputs = [
    go
    golangci-lint
    snekcheck
  ];
  text = ''
    snekcheck --fix "$PROJECT_ROOT"
    nix fmt "$PROJECT_ROOT" -- --quiet
    cd "$SOURCE_ROOT" && go mod tidy && go fmt ./... && go vet ./... && golangci-lint run ./...
    cd "$SPEC_ROOT" && go mod tidy && go fmt ./... && go vet ./... && golangci-lint run ./...
  '';
}
