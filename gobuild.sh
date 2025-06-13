#!/usr/bin/env bash
type -t go >/dev/null 2>&1 || {
    echo "  REQUIREs : go"
    exit 1
}
[[ -f $1 ]] || {
    echo "  USAGE : bash ${BASH_SOURCE##*/}" \$any.go
    exit 2
}

go build -o "${1%*.go}" $1

unset to
[[ -d $GOBIN ]] && to=$GOBIN
[[ -d $HOME/.bin ]] && to=$HOME/.bin
[[ -d $GOPATH/bin ]] && to=$GOPATH/bin
[[ -d $to ]] || to='.'

install "${1%*.go}" $to/

[[ $to == '.' ]] ||
    find . -type f -size +500k -exec rm "{}" \+

