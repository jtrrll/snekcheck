{
  perSystem = {pkgs, ...}: {
    scripts.e2e = pkgs.writeShellApplication {
      meta.description = "Runs all end-to-end tests.";
      name = "e2e";
      runtimeInputs = [
        pkgs.shellspec
      ];
      text = ''
        build
        shellspec --no-warning-as-failure "$PROJECT_ROOT"
      '';
    };
  };
}
