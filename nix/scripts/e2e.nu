#!/usr/bin/env nu

cd $env.SPEC_ROOT
go test -count 1 ./...
