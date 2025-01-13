# shellcheck shell=sh

Describe "help"
  It "shows the help message given --help"
    When call "$bin" --help
    The status should be success
    The error should include "Usage"
  End

  It "shows the help message given -h"
    When call "$bin" -h
    The status should be success
    The error should include "Usage"
  End

  It "shows the help message given nothing"
    When call "$bin"
    The status should be success
    The error should include "Usage"
  End
End
