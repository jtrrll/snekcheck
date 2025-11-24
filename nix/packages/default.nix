{
  perSystem =
    { lib, pkgs, ... }:
    {
      packages =
        let
          buildGoModule =
            args:
            pkgs.buildGoModule (
              lib.recursiveUpdate {
                meta = {
                  homepage = "https://github.com/jtrrll/snekcheck";
                  license = lib.licenses.mit;
                };
                src = lib.cleanSource (pkgs.nix-gitignore.gitignoreRecursiveSource [ ] ../../go);
                vendorHash = "sha256-uVjfU3XRqqLj9RwP/eyi+zsi1qYQvWmyCX93SONR3yw=";
              } args
            );
        in
        {
          default = pkgs.callPackage ./snekcheck.nix {
            inherit buildGoModule;
          };
        };
    };
}
