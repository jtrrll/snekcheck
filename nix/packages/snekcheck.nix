{
  buildGoApplication,
  go,
  lib,
}:
buildGoApplication {
  inherit go;
  meta = {
    description = "An opinionated filename linter that loves snake case.";
    homepage = "https://github.com/jtrrll/snekcheck";
    license = lib.licenses.mit;
  };
  modules = ../../go/gomod2nix.toml;
  pname = "snekcheck";
  src = ../../go;
  version = "0.1.0";
}
