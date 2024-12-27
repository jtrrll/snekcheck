# shellcheck shell=sh

Describe "check"
  Context "with an empty directory"
    It "succeeds"
      When call "$bin" "$root"
      The status should be success
    End
  End

  Context "with no subdirectories"
    create_valid_files() { for _ in {1.."$1"}; do create_valid_file "$root"; done; }
    BeforeEach "create_valid_files 200"

    It "succeeds"
      When call "$bin" "$root"
      The status should be success
    End

    Context "and one invalid file"
      create_and_assign_invalid_file() { invalid=$(create_invalid_file "$root"); }
      BeforeEach "create_and_assign_invalid_file"

      It "fails"
        When call "$bin" "$root"
        The status should be failure
        The stderr should include "$(basename "$invalid")"
        The file "$invalid" should be exist
      End

      It "succeeds with a limited depth"
        When call "$bin" --depth 0 "$root"
        The status should be success
        The stderr should not include "$(basename "$invalid")"
        The file "$invalid" should be exist
      End
    End

    Context "and one invalid directory"
      create_and_assign_invalid_directory() { invalid=$(create_invalid_directory "$root"); }
      BeforeEach "create_and_assign_invalid_directory"

      It "fails"
        When call "$bin" "$root"
        The status should be failure
        The stderr should include "$(basename "$invalid")"
        The file "$invalid" should be exist
      End

      It "succeeds with a limited depth"
        When call "$bin" --depth 0 "$root"
        The status should be success
        The stderr should not include "$(basename "$invalid")"
        The file "$invalid" should be exist
      End
    End
  End

  Context "with many subdirectories"
    create_valid_files() { for _ in {1.."$1"}; do create_valid_file "$2"; done; }
    BeforeEach "create_valid_files 200 $root"
    create_valid_directories() {
      for _ in {1.."$1"}; do
        dir=$(create_valid_directory "$root")
        create_valid_files "$1" "$dir"
      done
    }
    BeforeEach "create_valid_directories 200"
    
    It "succeeds"
      When call "$bin" "$root"
      The status should be success
    End
  End
End
