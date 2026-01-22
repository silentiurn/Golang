#!/usr/bin/env bash

KEY_INTERVIEW=$(
  grep -R "SEE INTERVIEW" streets/* \
  | grep -o '[0-9]\+' \
  | head -n1
)
export KEY_INTERVIEW
echo "$KEY_INTERVIEW"
cat interviews/*"$KEY_INTERVIEW"*
echo "$MAIN_SUSPECT"