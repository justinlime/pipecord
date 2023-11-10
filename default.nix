{ lib, buildGoModule }:

buildGoModule (finalAttrs: {
  pname = "pipecord";
  src = ./.;
}
