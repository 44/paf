paf vimgrep|grep|find|tag|auto -- fzf-args|@fzf-args-file
- transforms input interpreted as specified
- launches fzf with fzf-args
- sets environment variable allowing to refer to parent paf process from child paf processes (for preview)
- configures preview (and nth)
- transforms output back
- use IDs in mem for each line / maybe only folder / hash


grep: - grep, rg with -n option

long/long/path/file.txt:10:    some text with ansi coloring

file.txt:10 l/l/p some text with ansi coloring
<magenta>:<green> <grey> <original>

vimgrep: rg --vimgrep

long/long/path/file.txt:10:12: some text with ansi coloring

CONFIG:
in .git/paf.ini

keep Product
keep Test
max-length 32
align yes
dir-length 2
reverse yes
camel-case
snake-case
kebab-case
preserve-numbers
alias EM=EventMachine

Win10Convert - W11C
W11
W10
PR

