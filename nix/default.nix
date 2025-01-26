{inputs, ...}: {
  imports = [
    ./devenv.nix
    ./formatter.nix
    ./package.nix
  ];
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
