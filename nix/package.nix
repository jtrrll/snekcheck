{inputs, ...}: {
  perSystem = {system, ...}: let
    pkg = inputs.gomod2nix.legacyPackages.${system}.buildGoApplication {
      modules = ../go/gomod2nix.toml;
      pname = "snekcheck";
      src = ../go;
      version = "1.0.0";
    };
  in {
    apps.snekcheck.program = "${pkg}/bin/snekcheck";
    packages.snekcheck = pkg;
  };
}
