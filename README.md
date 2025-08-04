# snekcheck

<!-- markdownlint-disable MD013 -->
![Version](https://img.shields.io/github/v/tag/jtrrll/snekcheck?label=version&logo=semver&sort=semver)
![CI Status](https://img.shields.io/github/actions/workflow/status/jtrrll/snekcheck/ci.yaml?branch=main&label=ci&logo=github)
![License](https://img.shields.io/github/license/jtrrll/snekcheck?label=license&logo=googledocs&logoColor=white)
<!-- markdownlint-enable MD013 -->

An opinionated filename linter that loves snake case.

![Demo](./demo.gif)

## Usage

`snekcheck` accepts a list of filenames to lint.

- To lint a filename, simply pass it to `snekcheck`.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck <filename>
   ```
   <!-- markdownlint-enable MD013 -->

- To lint several filenames, provide them in a list.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck <filename> <filename> <dirname> ...
   ```
   <!-- markdownlint-enable MD013 -->

- To recursively lint directories and apply filters,
use tools like `find` to provide arguments to `snekcheck`.

   <!-- markdownlint-disable MD013 -->
   ```sh
   find . -exec snekcheck {} +
   ```
   <!-- markdownlint-enable MD013 -->

### Flags

`snekcheck`'s behavior can be modified with various flags.

- To print a help message, specify the `--help` flag.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --help
   ```
   <!-- markdownlint-enable MD013 -->

- To automatically rename invalid filenames, specify the `--fix` flag.
*Be careful*, as the renaming strategy may not produce the results you want.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --fix <filename> ...
   ```
   <!-- markdownlint-enable MD013 -->

- To print inspected filenames, specify the `--verbose` flag.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --verbose <filename> ...
   ```
   <!-- markdownlint-enable MD013 -->

## Build From Source

1. Install [Nix](https://zero-to-nix.com/start/install/)
2. Run the build command:

   <!-- markdownlint-disable MD013 -->
   ```sh
   nix build github:jtrrll/snekcheck
   ```
   <!-- markdownlint-enable MD013 -->

3. Find the `snekcheck` binary at `result/bin/snekcheck`
