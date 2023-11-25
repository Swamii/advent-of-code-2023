{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    nixpkgsUnstable.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flakeUtils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, nixpkgsUnstable, flakeUtils }:
    flakeUtils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        pkgsUnstable = nixpkgsUnstable.legacyPackages.${system};
      in {
        packages = flakeUtils.lib.flattenTree {
          go = pkgs.go;
          gotools = pkgs.gotools;
          gopls = pkgs.gopls;
        };
        devShell = pkgs.mkShell {
          buildInputs = with self.packages.${system}; [
            go
            gotools
            gopls
          ];
        };
      }
    );
}
