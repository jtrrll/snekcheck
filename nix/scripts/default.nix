{
  inputs,
  lib,
  ...
}:
{
  imports = [
    (inputs.flake-parts.lib.mkTransposedPerSystemModule {
      name = "scripts";
      option = lib.mkOption {
        type = lib.types.lazyAttrsOf lib.types.package;
        default = { };
        description = "Shell scripts for development";
      };
      file = "./default.nix";
    })
  ];
  perSystem =
    {
      pkgs,
      self',
      ...
    }:
    {
      scripts = {
        bench = pkgs.callPackage ./bench.nix {
          inherit (self'.packages.default) go;
        };
        demo = pkgs.callPackage ./demo.nix {
          snekcheck = self'.packages.default;
        };
        e2e = pkgs.callPackage ./e2e.nix {
          inherit (self'.packages.default) go;
        };
        lint = pkgs.callPackage ./lint.nix {
          inherit (self'.packages.default) go;
          snekcheck = self'.packages.default;
        };
        run = pkgs.callPackage ./run.nix {
          inherit (self'.packages.default) go;
        };
        splash = pkgs.callPackage ./splash.nix { };
        unit = pkgs.callPackage ./unit.nix {
          inherit (self'.packages.default) go;
        };
      };
    };
}
