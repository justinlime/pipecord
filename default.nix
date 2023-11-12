{ lib, buildGoModule }:

buildGoModule {
  name = "pipecord";
  src = ./.;
  vendorHash = "sha256-bEGHHUzqI8Bk6ync3t9gY0zqfZ4mV5usdcqkrI5aGjQ=";
  buildPhase = ''
    go build
    mkdir -p $out/bin
    install -m755 m $out/bin/pipecord
  '';
}
