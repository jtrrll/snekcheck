{
  inputs,
  self,
  ...
}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.run = pkgs.writeShellApplication {
      meta.description = "Runs the project.";
      name = "run";
      runtimeInputs = [
        inputs.gomod2nix.legacyPackages.${system}.gomod2nix
        self.packages.${system}.snekcheck.go
      ];
      text = ''
        (cd "$SOURCE_ROOT" && go mod tidy && gomod2nix)
        nix run "$PROJECT_ROOT"#snekcheck -- "$@"
      '';
    };
  };
}
