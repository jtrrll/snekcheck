{
  perSystem = {pkgs, ...}: {
    scripts.unit = pkgs.writeShellApplication {
      meta.description = "Runs all unit tests.";
      name = "unit";
      runtimeInputs = [
        pkgs.go
      ];
      text = ''
        cd "$SOURCE_ROOT"
        go test --cover ./...
      '';
    };
  };
}
