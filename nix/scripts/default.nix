{
  inputs,
  lib,
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

    ./bench.nix
    ./build.nix
    ./demo.nix
    ./e2e.nix
    ./lint.nix
    ./run.nix
    ./splash.nix
    ./unit.nix
  ];
}
