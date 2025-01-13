# shellcheck shell=sh

Describe "fix"
  Context "with an empty directory"
    It "succeeds"
      When call "$bin" --fix "$root"
      The status should be success
      The output should include "1 valid filenames"
      The output should include "0 filenames changed"
      The length of error should equal 0
    End
  End

  Context "with no subdirectories"
    create_valid_files() { for _ in $(seq 1 "$1"); do create_valid_file "$root"; done; }
    BeforeEach "create_valid_files 20"

    It "succeeds"
      When call "$bin" --fix "$root"
      The status should be success
      The output should include "21 valid filenames"
      The output should include "0 filenames changed"
      The length of error should equal 0
    End

    Context "and one invalid file"
      create_and_assign_invalid_file() { invalid=$(create_invalid_file "$root"); }
      BeforeEach "create_and_assign_invalid_file"

      It "succeeds given the root"
        When call "$bin" --fix "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the root twice"
        When call "$bin" --fix "$root" "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the invalid file"
        When call "$bin" --fix "$invalid"
        The status should be success
        The output should include "0 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the invalid file twice"
        When call "$bin" --fix "$invalid" "$invalid"
        The status should be success
        The output should include "0 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the root and the invalid file"
        When call "$bin" --fix "$root" "$invalid"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End
    End

    Context "and one invalid directory"
      create_and_assign_invalid_directory() { invalid=$(create_invalid_directory "$root"); }
      BeforeEach "create_and_assign_invalid_directory"

      It "succeeds given the root"
        When call "$bin" --fix "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the root twice"
        When call "$bin" --fix "$root" "$root"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the invalid directory"
        When call "$bin" --fix "$invalid"
        The status should be success
        The output should include "0 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the invalid file directory"
        When call "$bin" --fix "$invalid" "$invalid"
        The status should be success
        The output should include "0 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
      End

      It "succeeds given the root and the invalid directory"
        When call "$bin" --fix "$root" "$invalid"
        The status should be success
        The output should include "1 valid filenames"
        The output should include "1 filenames changed"
        The length of error should equal 0
        The file "$invalid" should not be exist
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
      When call "$bin" --fix "$root"
      The status should be success
      The output should include "1 valid filenames"
      The output should include "0 filenames changed"
      The length of error should equal 0
    End
  End
End
