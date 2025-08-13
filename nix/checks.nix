{ self, ... }:
{
  perSystem =
    {
      pkgs,
      self',
      ...
    }:
    {
      checks = {
        go-lint =
          pkgs.runCommandLocal "go-lint"
            {
              nativeBuildInputs = [
                pkgs.golangci-lint
                self'.packages.default.go
              ];
            }
            ''
              export GOCACHE=$(mktemp -d)
              export GOMODCACHE=$(mktemp -d)
              export GOLANGCI_LINT_CACHE=$(mktemp -d)
              cd ${self}/go && golangci-lint run ./...
              cd ${self}/spec && golangci-lint run ./...
              touch $out
            '';
        snekcheck =
          pkgs.runCommandLocal "snekcheck"
            {
              nativeBuildInputs = [ self'.packages.default ];
            }
            ''
              snekcheck $(find ${self} -mindepth 1)
              touch $out
            '';
      };
    };
}
