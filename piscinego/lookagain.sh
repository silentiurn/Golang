#!/bin/bash
find . -type f -name "*.sh" 2>/dev/null \
  | rev | cut -d'/' -f1 | rev \
  | cut -d'.' -f1 \
  | sort -r

  exit 0