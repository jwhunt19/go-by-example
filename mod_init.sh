#!/bin/zsh

module_name=$1
echo "module $module_name

go 1.20" > go.mod

touch "$module_name.go"