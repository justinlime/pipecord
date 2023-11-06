{ lib, buildGoModule }:

buildGoModule {
  name = "pipecord";
  src = ./.;
  buildPhase = ''
    
  '';
}
