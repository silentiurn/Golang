#!/bin/bash
curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq '.[] | select(.id==70) | .name'
exit 0