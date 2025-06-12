{self, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.unit = pkgs.writeShellApplication {
      meta.description = "Runs all unit tests.";
      name = "unit";
      runtimeInputs = [
        self.packages.${system}.snekcheck.go
      ];
      text = ''
        cd "$SOURCE_ROOT"
        go test --cover ./...
      '';
    };
  };
}
