{self, ...}: {
  perSystem = {
    config,
    pkgs,
    system,
    ...
  }: {
    checks = {
      go-lint =
        pkgs.runCommandLocal "go-lint" {
          nativeBuildInputs = [
            pkgs.golangci-lint
            self.packages.${system}.default.go
          ];
        } ''
          export GOCACHE=$(mktemp -d)
          export GOMODCACHE=$(mktemp -d)
          export GOLANGCI_LINT_CACHE=$(mktemp -d)
          cd ${self}/go && golangci-lint run ./...
          cd ${self}/spec && golangci-lint run ./...
          touch $out
        '';
      nix-lint = pkgs.runCommandLocal "nix-lint" {} ''
        ${config.formatter}/bin/* --check $(find ${self} -type f -name '*.nix')
        touch $out
      '';
      snekcheck =
        pkgs.runCommandLocal "snekcheck" {
          nativeBuildInputs = [self.packages.${system}.default];
        } ''
          snekcheck $(find ${self} -mindepth 1)
          touch $out
        '';
    };
  };
}
