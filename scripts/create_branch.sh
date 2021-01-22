#! /bin/bash
BRANCHNAME=""

RESET="\033[0m"
BOLD="\033[1m"
YELLOW="\033[36;5;11m"
read -p "$(echo -e $BOLD$YELLOW"Branch name:  "$RESET)" BRANCHNAME

git checkout -b $BRANCHNAME
