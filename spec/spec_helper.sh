# shellcheck shell=sh

root="/tmp/snekcheck_test"
bin="./result/bin/snekcheck"

create_valid_directory() {
  dir_name=$(tr -dc a-z0-9 </dev/urandom | head -c 13)
  mkdir --parent "$1"/"$dir_name"
  echo "$1"/"$dir_name"
}
create_valid_file() {
  file_name=$(tr -dc a-z0-9 </dev/urandom | head -c 13)
  mkdir --parent "$1"
  touch "$1"/"$file_name"
  echo "$1"/"$file_name"

  mkdir --parent "$1"
  filename=$(mktemp --tmpdir="$1" --quiet XXXvalidXXX)
  lowercase=$(echo "$filename" | tr '[:upper:]' '[:lower:]')
  mv "$filename" "$lowercase"
  echo "$lowercase"
}
create_invalid_directory() {
  mkdir --parent "$1"
  mktemp --directory --tmpdir="$1" --quiet XXXInVaLiDXXX
}
create_invalid_file() {
  mkdir --parent "$1"
  mktemp --tmpdir="$1" --quiet XXXInVaLiDXXX
}

spec_helper_precheck() {
  minimum_version '0.28.1'
}

spec_helper_configure() {
  create_root() { mkdir --parent "$root"; }
  before_each "create_root"

  delete_root() { rm --recursive "$root"; }
  after_each "delete_root"
}
