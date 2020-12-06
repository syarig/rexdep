#!/bin/zsh
#rexdep --pattern 'import +(\S+)' --module '^export default +(\S+)|export' --start '^import +["{]' --end '^\}$"' --format dot $(find ./packages/admin/src -name '*.tsx')
#rexdep --pattern 'import +(\S+)' --module '^export default +(\S+)|export' --start '^import +["{]' --end '^\}$"' --format dot $(find ./packages/admin/src -name '*.tsx')
#rexdep --pattern '^import ([^{}\*]\S+) from' --module '^export default +(\S+)' --format dot $(find ./packages/admin/src -name '*.tsx')
#rexdep --pattern '^import ([^{}\*]\S+) from' --module '^export default +(\S+)' --format dot $(find ./packages/admin/src -name '*.tsx') | dot -Tpng -o test.png
#rexdep --pattern '^import ([^{}\*]\S+) from' --module '^export default +(\S+)' --format json $(find ./packages/admin/src -name '*.tsx')
#rexdep --pattern "from '+\S+\/(\S+)'|'(\S+)'$" --module '^(export default|export) +(\S+)' --format json $(find ./packages/admin/src -name '*.tsx')

./rexdep \
--pattern '^import ([^{}\*]\S+) from' --module '^export default +(\S+)' \
--format dot $(find /Users/arito/ghq/github.com/syuukai85/frontend/src -name '*.tsx')