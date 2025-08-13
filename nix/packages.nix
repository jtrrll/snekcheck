{
  perSystem =
    { pkgs, ... }:
    {
      packages.default =
        pkgs.callPackage
          (
            {
              buildGoModule,
              lib,
            }:
            buildGoModule {
              meta = {
                description = "An opinionated filename linter that loves snake case.";
                homepage = "https://github.com/jtrrll/snekcheck";
                license = lib.licenses.mit;
                mainProgram = "snekcheck";
              };
              modRoot = "./go";
              pname = "snekcheck";
              src = ./..;
              vendorHash = "sha256-uVjfU3XRqqLj9RwP/eyi+zsi1qYQvWmyCX93SONR3yw=";
              version = "0.1.0";
            }
          )
          {
            buildGoModule = pkgs.buildGo124Module;
          };
    };
}
