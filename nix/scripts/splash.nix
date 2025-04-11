{
  perSystem = {pkgs, ...}: {
    scripts.splash = pkgs.writeShellApplication {
      meta.description = "Prints a splash screen.";
      name = "splash";
      runtimeInputs = [
        pkgs.lolcat
        pkgs.uutils-coreutils-noprefix
      ];
      text = ''
        printf "     ▄▄▄▄▄    ▄   ▄███▄   █  █▀ ▄█▄     ▄  █ ▄███▄   ▄█▄    █  █▀
            █     ▀▄   █  █▀   ▀  █▄█   █▀ ▀▄  █   █ █▀   ▀  █▀ ▀▄  █▄█
          ▄  ▀▀▀▀▄ ██   █ ██▄▄    █▀▄   █   ▀  ██▀▀█ ██▄▄    █   ▀  █▀▄
           ▀▄▄▄▄▀  █ █  █ █▄   ▄▀ █  █  █▄  ▄▀ █   █ █▄   ▄▀ █▄  ▄▀ █  █
                   █  █ █ ▀███▀     █   ▀███▀     █  ▀███▀   ▀███▀    █
                   █   ██          ▀             ▀                   ▀\n" | lolcat
      '';
    };
  };
}
