{self, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.bench = pkgs.writeShellApplication {
      meta.description = "Runs all benchmark tests.";
      name = "bench";
      runtimeInputs = [
        self.packages.${system}.snekcheck.go
      ];
      text = ''
        cd "$SOURCE_ROOT"
        go test ./... -bench=.
      '';
    };
  };
}
