{inputs, ...}: {
  imports = [
    ./packages
    ./scripts

    ./checks.nix
    ./devenv.nix
    ./overlays.nix
  ];
  perSystem = {pkgs, ...}: {
    formatter = pkgs.alejandra;
  };
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
