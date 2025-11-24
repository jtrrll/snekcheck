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
      lib,
      pkgs,
      self',
      ...
    }:
    {
      scripts =
        let
          inherit (self'.packages.default) go;
          snekcheck = self'.packages.default;

          writeNuApplication =
            {
              name,
              meta ? { },
              runtimeInputs ? [ ],
              text,
            }:
            (
              (pkgs.writers.writeNuBin name {
                makeWrapperArgs = [ "--prefix PATH : ${lib.makeBinPath runtimeInputs}" ];
              } text).overrideAttrs
              (oldAttrs: {
                meta = (oldAttrs.meta or { }) // meta;
              })
            );
        in
        {
          bench = writeNuApplication {
            name = "bench";
            meta.description = "Runs all benchmark tests.";
            runtimeInputs = [ go ];
            text = lib.readFile ./bench.nu;
          };

          demo = writeNuApplication {
            name = "demo";
            meta.description = "Generates a demo GIF.";
            runtimeInputs = [
              pkgs.bashInteractive
              pkgs.vhs
              snekcheck
            ];
            text = lib.readFile ./demo.nu;
          };

          e2e = writeNuApplication {
            name = "e2e";
            meta.description = "Runs all end-to-end tests.";
            runtimeInputs = [ go ];
            text = lib.readFile ./e2e.nu;
          };

          lint = writeNuApplication {
            name = "lint";
            meta.description = "Lints the project.";
            runtimeInputs = [
              go
              pkgs.findutils
              pkgs.golangci-lint
              snekcheck
            ];
            text = lib.readFile ./lint.nu;
          };

          run = writeNuApplication {
            name = "run";
            meta.description = "Runs the project.";
            runtimeInputs = [ go ];
            text = lib.readFile ./run.nu;
          };

          splash = writeNuApplication {
            name = "splash";
            meta.description = "Prints a splash screen.";
            runtimeInputs = [ pkgs.lolcat ];
            text = lib.readFile ./splash.nu;
          };

          unit = writeNuApplication {
            name = "unit";
            meta.description = "Runs all unit tests.";
            runtimeInputs = [ go ];
            text = lib.readFile ./unit.nu;
          };
        };
    };
}
