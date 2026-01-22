#!/bin/bash
curl -s "https://01.tomorrow-school.ai/assets/superhero/all.json" | jq ".[] | select(.id == $HERO_ID) | .connections.relatives"| sed 's/"//g'
exit 0