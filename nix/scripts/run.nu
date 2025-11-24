#!/usr/bin/env nu

cd $env.SOURCE_ROOT
go mod tidy
nix run $env.PROJECT_ROOT -- ...$rest
