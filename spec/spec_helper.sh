# shellcheck shell=sh

root="/tmp/snekcheck_test"
bin="./result/bin/snekcheck"

create_valid_directory() {
  dir_name=$(echo "$RANDOM" | md5sum | head -c 20)
  dir_path="$1/$dir_name"

  if [ -d "$dir_path" ]; then
    create_valid_directory "$1"
  else
    mkdir --parents "$dir_path"
    echo "$dir_path"
  fi
}

create_valid_file() {
  file_name=$(echo "$RANDOM" | md5sum | head -c 20)
  file_path="$1/$file_name"

  if [ -e "$file_path" ]; then
    create_valid_file "$1"
  else
    mkdir --parents "$1"
    touch "$file_path"
    echo "$file_path"
  fi
}

create_invalid_directory() {
  dir_name=$(echo "$RANDOM" | md5sum | head -c 20)InVaLiD
  dir_path="$1/$dir_name"

  if [ -d "$dir_path" ]; then
    create_valid_directory "$1"
  else
    mkdir --parents "$dir_path"
    echo "$dir_path"
  fi
}

create_invalid_file() {
  file_name=$(echo "$RANDOM" | md5sum | head -c 20)InVaLiD
  file_path="$1/$file_name"

  if [ -e "$file_path" ]; then
    create_valid_file "$1"
  else
    mkdir --parents "$1"
    touch "$file_path"
    echo "$file_path"
  fi
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
