{
  perSystem = {pkgs, ...}: {
    scripts.bench = pkgs.writeShellApplication {
      meta.description = "Runs all benchmark tests.";
      name = "bench";
      runtimeInputs = [
        pkgs.go
      ];
      text = ''
        cd "$SOURCE_ROOT"
        go test ./... -bench=.
      '';
    };
  };
}
