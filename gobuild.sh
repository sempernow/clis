#!/usr/bin/env bash
type -t go >/dev/null 2>&1 || {
    echo "  REQUIREs : go"
    exit 1
}
[[ -f $1 ]] || {
    echo "  USAGE : ${BASH_SOURCE##*/}" \$name.go
    exit 2
}

go build -o ${1%.go} $1
