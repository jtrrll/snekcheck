# shellcheck shell=sh

Describe "check"
  Context "with an empty directory"
    It "succeeds"
      When call "$bin" "$root"
      The status should be success
      The output should include "1 valid filenames"
      The output should include "0 invalid filenames"
      The length of error should equal 0
    End
  End

  Context "with no subdirectories"
    create_valid_files() { for _ in $(seq 1 "$1"); do create_valid_file "$root"; done; }
    BeforeEach "create_valid_files 20"

    It "succeeds"
      When call "$bin" "$root"
      The status should be success
      The output should include "21 valid filenames"
      The output should include "0 invalid filenames"
      The length of error should equal 0
    End

    Context "and one invalid file"
      BeforeEach "create_invalid_file $root"

      It "fails"
        When call "$bin" "$root"
        The status should be failure
        The output should include "21 valid filenames"
        The output should include "1 invalid filenames"
        The error should include "invalid filenames found"
      End

      It "succeeds with a limited depth"
        When call "$bin" --depth 0 "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "0 invalid filenames"
        The length of error should equal 0
      End
    End

    Context "and one invalid directory"
      BeforeEach "create_invalid_file $root"

      It "fails"
        When call "$bin" "$root"
        The status should be failure
        The output should include "21 valid filenames"
        The output should include "1 invalid filenames"
        The error should include "invalid filenames found"
      End

      It "succeeds with a limited depth"
        When call "$bin" --depth 0 "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "0 invalid filenames"
        The length of error should equal 0
      End
    End
  End

  Context "with many subdirectories"
    create_valid_files() { for _ in $(seq 1 "$1"); do create_valid_file "$2"; done; }
    BeforeEach "create_valid_files 20 $root"
    create_valid_directories() {
      for _ in $(seq 1 "$1"); do
        dir=$(create_valid_directory "$root")
        create_valid_files "$1" "$dir"
      done
    }
    BeforeEach "create_valid_directories 20"
    
    It "succeeds"
      When call "$bin" "$root"
      The status should be success
      The output should include "441 valid filenames"
      The output should include "0 invalid filenames"
      The length of error should equal 0
    End
  End
End
