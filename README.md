# snekcheck

<!-- markdownlint-disable MD013 -->
![GitHub Actions CI Status](https://img.shields.io/github/actions/workflow/status/jtrrll/snekcheck/ci.yaml?branch=main&logo=github&label=CI)
![License](https://img.shields.io/github/license/jtrrll/snekcheck?label=License)
<!-- markdownlint-enable MD013 -->

An opinionated filename linter that loves snake case.
![Demo](./demo.gif)

## Basic Usage

- To lint several filenames, provide them in a list.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck <filename> <filename> ...
   ```
   <!-- markdownlint-enable MD013 -->

- To lint an entire directory, or several, directory names can be provided instead.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck <dirname> <dirname> ...
   ```
   <!-- markdownlint-enable MD013 -->

- Several filenames and directory names can be provided simultaneously.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck <filename> <dirname> ...
   ```
   <!-- markdownlint-enable MD013 -->

## Advanced Usage

- To print a help message, specify the --help flag.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --help
   ```
   <!-- markdownlint-enable MD013 -->

- To automatically rename invalid filenames, specify the `--fix` flag.
Be careful, as the renaming strategy may not produce the results you want.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --fix <filename> ...
   ```
   <!-- markdownlint-enable MD013 -->

- To limit directory traversal, specify a depth limit.
This can be useful for only checking top-level names or
for improving performance on large file systems.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --depth 1 <dirname> ...
   ```
   <!-- markdownlint-enable MD013 -->

- To print inspected filenames, specify the `--verbose` flag.

   <!-- markdownlint-disable MD013 -->
   ```sh
   snekcheck --verbose <filename> ...
   ```
   <!-- markdownlint-enable MD013 -->
