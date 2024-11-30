#!/bin/sh
exec bazel run -- @rules_go//go/tools/gopackagesdriver "${@}"
