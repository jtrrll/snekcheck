{
  buildGoModule,
}:
buildGoModule {
  pname = "snekcheck";
  version = "0.1.0";

  meta = {
    description = "An opinionated filename linter that loves snake case.";
    mainProgram = "snekcheck";
  };
  subPackages = [ "cmd/snekcheck" ];
}
