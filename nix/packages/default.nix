{
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    packages.default = pkgs.callPackage ./snekcheck.nix {
      buildGoModule = pkgs.buildGo124Module;
    };
  };
}
