{inputs, ...}: {
  imports = [
    ./scripts

    ./devenv.nix
    ./overlay.nix
    ./packages.nix
  ];
  perSystem = {pkgs, ...}: {
    formatter = pkgs.alejandra;
  };
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
