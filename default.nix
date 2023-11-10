{ lib, buildGoModule }:

buildGoModule {
  pname = "pipecord";
  src = ./.;
}
