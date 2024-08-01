{
  description = "movie-matcher - A back-end test for potential Generate engineers";

  inputs = {
    devenv.url = "github:cachix/devenv";
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = {
    devenv,
    flake-utils,
    nixpkgs,
    ...
  } @ inputs:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells = {
        default = devenv.lib.mkShell {
          inherit inputs pkgs;
          modules = [
            ({pkgs, ...}: {
              enterShell = ''
                printf "movie-matcher\n" | ${pkgs.lolcat}/bin/lolcat
                printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
              '';
              languages = {
                go.enable = true;
                javascript = {
                  enable = true;
                  npm = {
                    enable = true;
                    install.enable = true;
                  };
                };
                nix.enable = true;
                typescript.enable = true;
              };
              packages = with pkgs; [
                commitizen
                docker
                gnumake
                golangci-lint
                postgresql
                wget
              ];
              pre-commit = {
                default_stages = ["pre-push"];
                hooks = {
                  actionlint.enable = true;
                  alejandra.enable = true;
                  check-added-large-files.enable = true;
                  check-yaml.enable = true;
                  deadnix.enable = true;
                  end-of-file-fixer.enable = true;
                  flake-checker.enable = true;
                  gofmt.enable = true;
                  golangci-lint.enable = true;
                  govet.enable = true;
                  markdownlint.enable = true;
                  mixed-line-endings.enable = true;
                  nil.enable = true;
                  no-commit-to-branch = {
                    enable = true;
                    stages = ["pre-commit"];
                  };
                  statix.enable = true;
                };
              };
            })
          ];
        };
      };
      formatter = pkgs.alejandra;
    });
}
