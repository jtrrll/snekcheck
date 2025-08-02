{
  inputs,
  self,
  ...
}: {
  flake.overlays.default = _: prev: {
    writePythonApplication = {
      description,
      libraries ? [],
      name,
      python ? prev.pkgs.python3,
      ruffIgnoredRules ? [
        "D100"
        "D203"
        "D212"
        "FBT"
        "I001"
        "S101"
        "S603"
        "S607"
        "T201"
      ],
      runtimeInputs ? [],
      script,
    }:
      assert builtins.isString description;
      assert builtins.isList libraries && builtins.all (pkg: builtins.isAttrs pkg && pkg ? outPath) libraries;
      assert builtins.isString name;
      assert builtins.isAttrs python && python ? withPackages;
      assert builtins.isList ruffIgnoredRules && builtins.all (rule: builtins.isString rule) ruffIgnoredRules;
      assert builtins.isList runtimeInputs && builtins.all (pkg: builtins.isAttrs pkg && pkg ? outPath) runtimeInputs;
      assert builtins.isString script;
        prev.pkgs.runCommandLocal name {
          buildInputs = [
            prev.pkgs.makeWrapper
            prev.pkgs.ruff
            prev.pkgs.uutils-coreutils-noprefix
          ];
          meta.description = description;
        } ''
          mkdir -p "$out/bin"

          cat > "$out/bin/${name}.py" <<EOF
          #!${python.withPackages (_: libraries)}/bin/python
          ${script}
          EOF

          chmod +x "$out/bin/${name}.py"

          ruff check "$out/bin/${name}.py" \
            --select "ALL" \
            ${builtins.concatStringsSep " " (builtins.map (rule: "--ignore ${rule}") ruffIgnoredRules)}

          makeWrapper "$out/bin/${name}.py" "$out/bin/${name}" \
            --set "PATH" "${prev.lib.makeBinPath runtimeInputs}"
        '';
  };

  perSystem = {system, ...}: {
    _module.args.pkgs = import inputs.nixpkgs {
      inherit system;
      overlays = [self.overlays.default];
    };
  };
}
