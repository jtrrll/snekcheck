#!/usr/bin/env nu

cd $env.SOURCE_ROOT
go test --cover ./...
