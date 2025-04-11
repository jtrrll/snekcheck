{inputs, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: {
    scripts.build = pkgs.writePythonApplication {
      description = "Builds the snekcheck package.";
      name = "build";
      libraries = [pkgs.python3.pkgs.typer];
      runtimeInputs = [
        inputs.gomod2nix.legacyPackages.${system}.gomod2nix
        pkgs.go
        pkgs.nix
      ];
      script = ''
        from pathlib import Path
        import datetime
        import os
        import subprocess
        import typer

        app = typer.Typer()

        PROJECT_ROOT: Path = Path(os.environ.get("PROJECT_ROOT", "."))
        SOURCE_ROOT: Path = Path(os.environ.get("SOURCE_ROOT"))
        RESULT_DIR: Path = PROJECT_ROOT / "result"

        def format_time(epoch: float) -> str:
          """Format a timestamp as a string."""
          return datetime.datetime.fromtimestamp(
            epoch, tz=datetime.UTC,
          ).strftime("%Y-%m-%dT%H:%M:%S")

        def should_rebuild() -> bool:
          """Determine if the project should be rebuilt."""
          output_binary = RESULT_DIR / "bin" / "snekcheck"

          if not output_binary.exists():
            return True

          go_files: list[Path] = list(SOURCE_ROOT.rglob("*.go"))
          if not go_files:
            return False

          latest_src_time: float = max(path.stat().st_mtime for path in go_files)
          latest_build_time: float = RESULT_DIR.lstat().st_mtime

          print(f"{"":20} | {"Human Time":20} | Epoch Time")
          print(f"{"-"*20} | {"-"*20} | {"-"*11}")
          print(
            f"{"Source changed":20} | "
            f"{format_time(latest_src_time):20} | "
            f"{int(latest_src_time)}",
          )
          print(
            f"{"Last build":20} | "
            f"{format_time(latest_build_time):20} | "
            f"{int(latest_build_time)}",
          )

          return latest_src_time > latest_build_time

        @app.command()
        def main(
          force: bool = typer.Option(False, "--force", "-f",
                                     help="Force rebuild regardless of timestamps"),
        ) -> None:
          """Rebuild the snekcheck binary if the source has changed."""
          if force or should_rebuild():
            if RESULT_DIR.exists():
              RESULT_DIR.unlink()

            print("Rebuilding...")
            subprocess.run(["go", "mod", "tidy"], cwd=SOURCE_ROOT, check=True)
            subprocess.run(["gomod2nix"], cwd=SOURCE_ROOT, check=True)
            subprocess.run(["nix", "build", f"{PROJECT_ROOT}#snekcheck"], check=True)
          else:
            print("Build is up to date.")

        if __name__ == "__main__":
          app()
      '';
    };
  };
}
