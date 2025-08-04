{inputs, ...}: {
  imports = [
    ./scripts

    ./checks.nix
    ./devenv.nix
    ./packages.nix
  ];
  perSystem = {pkgs, ...}: {
    formatter = pkgs.alejandra;
  };
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
