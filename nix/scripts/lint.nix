{self, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.lint = pkgs.writeShellApplication {
      meta.description = "Lints the project.";
      name = "lint";
      runtimeInputs = [
        pkgs.golangci-lint
        self.packages.${system}.snekcheck
        self.packages.${system}.snekcheck.go
      ];
      text = ''
        snekcheck --fix "$PROJECT_ROOT"
        nix fmt "$PROJECT_ROOT" -- --quiet
        cd "$SOURCE_ROOT" && go mod tidy && go fmt ./... && go vet ./... && golangci-lint run ./...
        cd "$SPEC_ROOT" && go mod tidy && go fmt ./... && go vet ./... && golangci-lint run ./...
      '';
    };
  };
}
