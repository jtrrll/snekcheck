{
  inputs,
  self,
  ...
}: {
  imports = [
    inputs.devenv.flakeModule
  ];
  perSystem = {
    config,
    lib,
    pkgs,
    system,
    ...
  }: {
    devenv = {
      shells.default = {
        containers = lib.mkForce {}; # Workaround to remove containers from flake checks.
        enterShell = "${pkgs.writeShellApplication {
          name = "splashScreen";
          runtimeInputs = [
            self.scripts.${system}.splash
            pkgs.lolcat
            pkgs.uutils-coreutils-noprefix
          ];
          text = ''
            splash
            printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
          '';
        }}/bin/splashScreen";

        env = let
          PROJECT_ROOT = config.devenv.shells.default.env.DEVENV_ROOT;
        in {
          inherit PROJECT_ROOT;
          SOURCE_ROOT = "${PROJECT_ROOT}/go";
          SPEC_ROOT = "${PROJECT_ROOT}/spec";
        };

        languages = {
          go = {
            enable = true;
            package = self.packages.${system}.default.go;
          };
          nix.enable = true;
        };

        git-hooks = {
          default_stages = ["pre-push"];
          hooks = {
            actionlint.enable = true;
            check-added-large-files = {
              enable = true;
              stages = ["pre-commit"];
            };
            check-yaml.enable = true;
            deadnix.enable = true;
            detect-private-keys = {
              enable = true;
              stages = ["pre-commit"];
            };
            end-of-file-fixer.enable = true;
            flake-checker.enable = true;
            lint = {
              enable = true;
              entry = "lint";
              name = "lint";
              pass_filenames = false;
            };
            markdownlint.enable = true;
            mixed-line-endings.enable = true;
            nil.enable = true;
            no-commit-to-branch = {
              enable = true;
              stages = ["pre-commit"];
            };
            ripsecrets = {
              enable = true;
              stages = ["pre-commit"];
            };
            statix.enable = true;
          };
        };

        scripts =
          lib.mapAttrs (_: pkg: {
            inherit (pkg.meta) description;
            exec = "${pkg}/bin/${pkg.name} $@";
          })
          self.scripts.${system};
      };
    };
  };
}
