# shellcheck shell=sh

Describe "help"
  It "shows the help message"
    When call "$bin" --help
    The status should be success
    The stderr should include "Usage"
  End
End
