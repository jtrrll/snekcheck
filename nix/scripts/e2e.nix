{self, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.e2e = pkgs.writeShellApplication {
      meta.description = "Runs all end-to-end tests.";
      name = "e2e";
      runtimeInputs = [
        self.packages.${system}.snekcheck.go
      ];
      text = ''
        build
        cd "$SPEC_ROOT"
        go test -count 1 ./...
      '';
    };
  };
}
