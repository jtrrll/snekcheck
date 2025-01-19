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
      modules = [
        inputs.env-help.devenvModule
      ];
      shells.default = let
        buildInputs = config.packages.snekcheck.nativeBuildInputs;
        goPkg = lib.findFirst (pkg: builtins.match "go" pkg.pname != null) pkgs.go buildInputs;
      in {
        enterShell = ''
          printf "   ▄▄▄▄▄    ▄   ▄███▄   █  █▀ ▄█▄     ▄  █ ▄███▄   ▄█▄    █  █▀
            █     ▀▄   █  █▀   ▀  █▄█   █▀ ▀▄  █   █ █▀   ▀  █▀ ▀▄  █▄█
          ▄  ▀▀▀▀▄ ██   █ ██▄▄    █▀▄   █   ▀  ██▀▀█ ██▄▄    █   ▀  █▀▄
           ▀▄▄▄▄▀  █ █  █ █▄   ▄▀ █  █  █▄  ▄▀ █   █ █▄   ▄▀ █▄  ▄▀ █  █
                   █  █ █ ▀███▀     █   ▀███▀     █  ▀███▀   ▀███▀    █
                   █   ██          ▀             ▀                   ▀\n" | ${pkgs.lolcat}/bin/lolcat
          printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
        '';

        env-help.enable = true;

        languages = {
          go = {
            enable = true;
            package = goPkg;
          };
          nix.enable = true;
        };

        pre-commit = {
          default_stages = ["pre-push"];
          hooks = {
            actionlint.enable = true;
            alejandra.enable = true;
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
            gofmt.enable = true;
            golangci-lint.enable = true;
            govet.enable = true;
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
            snekcheck = {
              enable = true;
              entry = "${self.packages.${system}.snekcheck}/bin/snekcheck";
              name = "snekcheck";
            };
            statix.enable = true;
          };
        };

        scripts = {
          bench = {
            description = "Runs all benchmark tests.";
            exec = ''
              ${goPkg}/bin/go test "$DEVENV_ROOT"/... -bench=.
            '';
          };
          build = {
            description = "Builds the project binary.";
            exec = ''
              ${inputs.gomod2nix.legacyPackages.${system}.gomod2nix}/bin/gomod2nix && \
              nix build "$DEVENV_ROOT"#snekcheck
            '';
          };
          demo = {
            description = "Generates a demo GIF.";
            exec = ''
              PATH="$DEVENV_ROOT/result/bin:${pkgs.bashInteractive}/bin:$PATH"

              mkdir --parents "$DEVENV_ROOT"/demo
              for i in $(seq 1 3); do touch "$DEVENV_ROOT"/demo/"$i"valid; done;
              for i in $(seq 1 3); do touch "$DEVENV_ROOT"/demo/"$i"InVaLiD; done;

              build && \
              ${pkgs.vhs}/bin/vhs "$DEVENV_ROOT"/demo.tape

              rm --force --recursive "$DEVENV_ROOT"/demo
            '';
          };
          e2e = {
            description = "Runs all end-to-end tests.";
            exec = ''
              build && \
              ${pkgs.shellspec}/bin/shellspec --no-warning-as-failure "$DEVENV_ROOT"
            '';
          };
          lint = {
            description = "Lints the project.";
            exec = ''
              nix fmt "$DEVENV_ROOT" -- --quiet && \
              ${goPkg}/bin/go mod tidy && \
              ${goPkg}/bin/go fmt "$DEVENV_ROOT"/... && \
              ${goPkg}/bin/go vet "$DEVENV_ROOT"/... && \
              ${pkgs.golangci-lint}/bin/golangci-lint run "$DEVENV_ROOT"/...
            '';
          };
          run = {
            description = "Runs the project.";
            exec = ''
              ${inputs.gomod2nix.legacyPackages.${system}.gomod2nix}/bin/gomod2nix && \
              nix run "$DEVENV_ROOT"#snekcheck -- "$@"
            '';
          };
          unit = {
            description = "Runs all unit tests.";
            exec = ''
              ${goPkg}/bin/go test --cover "$DEVENV_ROOT"/...
            '';
          };
        };
      };
    };
  };
}
