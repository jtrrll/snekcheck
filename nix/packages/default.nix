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
                vendorHash = "sha256-eeipkAobSq4Nh8zClL5HBRN5wXc2oxkjeqYVh04Zf3c=";
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
