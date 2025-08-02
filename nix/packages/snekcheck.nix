{
  buildGoModule,
  lib,
}:
buildGoModule {
  meta = {
    description = "An opinionated filename linter that loves snake case.";
    homepage = "https://github.com/jtrrll/snekcheck";
    license = lib.licenses.mit;
  };
  modRoot = "./go";
  pname = "snekcheck";
  src = ../..;
  vendorHash = "sha256-zsMzeDcba92K80iRxvV1YNW1olWjSoC9jNyzfy7t2jI=";
  version = "0.1.0";
}
