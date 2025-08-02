{
  lolcat,
  uutils-coreutils-noprefix,
  writeShellApplication,
}:
writeShellApplication {
  meta.description = "Prints a splash screen.";
  name = "splash";
  runtimeInputs = [
    lolcat
    uutils-coreutils-noprefix
  ];
  text = ''
    printf "     ▄▄▄▄▄    ▄   ▄███▄   █  █▀ ▄█▄     ▄  █ ▄███▄   ▄█▄    █  █▀
        █     ▀▄   █  █▀   ▀  █▄█   █▀ ▀▄  █   █ █▀   ▀  █▀ ▀▄  █▄█
      ▄  ▀▀▀▀▄ ██   █ ██▄▄    █▀▄   █   ▀  ██▀▀█ ██▄▄    █   ▀  █▀▄
        ▀▄▄▄▄▀  █ █  █ █▄   ▄▀ █  █  █▄  ▄▀ █   █ █▄   ▄▀ █▄  ▄▀ █  █
                █  █ █ ▀███▀     █   ▀███▀     █  ▀███▀   ▀███▀    █
                █   ██          ▀             ▀                   ▀\n" | lolcat
  '';
}
