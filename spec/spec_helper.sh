# shellcheck shell=sh

root="/tmp/snekcheck_test"
bin="./result/bin/snekcheck"

create_valid_directory() {
  dir_name=$(echo "$RANDOM" | md5sum | head -c 20)
  mkdir --parents "$1"/"$dir_name"
  echo "$1"/"$dir_name"
}
create_valid_file() {
  file_name=$(echo "$RANDOM" | md5sum | head -c 20)
  mkdir --parents "$1"
  touch "$1"/"$file_name"
  echo "$1"/"$file_name"
}
create_invalid_directory() {
  mkdir --parents "$1"
  mktemp --directory --tmpdir="$1" --quiet XXXInVaLiDXXX
}
create_invalid_file() {
  mkdir --parents "$1"
  mktemp --tmpdir="$1" --quiet XXXInVaLiDXXX
}

spec_helper_precheck() {
  minimum_version '0.28.1'
}

spec_helper_configure() {
  create_root() { mkdir --parent "$root"; }
  delete_root() { rm --force --recursive "$root"; }

  before_all "delete_root"
  before_each "create_root"
  after_each "delete_root"
}
