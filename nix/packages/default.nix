{inputs, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    packages.default = pkgs.callPackage ./snekcheck.nix {
      inherit (inputs.gomod2nix.legacyPackages.${system}) buildGoApplication;
    };
  };
}
