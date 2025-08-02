{
  inputs,
  lib,
  self,
  ...
}: {
  imports = [
    (inputs.flake-parts.lib.mkTransposedPerSystemModule {
      name = "scripts";
      option = lib.mkOption {
        type = lib.types.lazyAttrsOf lib.types.package;
        default = {};
        description = "Shell scripts for development";
      };
      file = "./default.nix";
    })
  ];
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts = {
      bench = pkgs.callPackage ./bench.nix {
        inherit (self.packages.${system}.default) go;
      };
      build = pkgs.callPackage ./build.nix {
        inherit (self.packages.${system}.default) go;
        inherit (inputs.gomod2nix.legacyPackages.${system}) gomod2nix;
      };
      demo = pkgs.callPackage ./demo.nix {
        snekcheck = self.packages.${system}.default;
      };
      e2e = pkgs.callPackage ./e2e.nix {
        inherit (self.packages.${system}.default) go;
      };
      lint = pkgs.callPackage ./lint.nix {
        inherit (self.packages.${system}.default) go;
        snekcheck = self.packages.${system}.default;
      };
      run = pkgs.callPackage ./run.nix {
        inherit (self.packages.${system}.default) go;
        inherit (inputs.gomod2nix.legacyPackages.${system}) gomod2nix;
      };
      splash = pkgs.callPackage ./splash.nix {};
      unit = pkgs.callPackage ./unit.nix {
        inherit (self.packages.${system}.default) go;
      };
    };
  };
}
