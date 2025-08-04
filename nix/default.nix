{inputs, ...}: {
  imports = [
    ./packages
    ./scripts

    ./checks.nix
    ./devenv.nix
  ];
  perSystem = {pkgs, ...}: {
    formatter = pkgs.alejandra;
  };
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
