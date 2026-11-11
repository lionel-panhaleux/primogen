#!/bin/sh
# Pins the real upstreams for the tests: the card data, vtes-rulings at its
# head, the archive cut to its ids.
set -eu
cd "$(dirname "$0")"
sha=$(curl -fsS -H 'Accept: application/vnd.github.sha' https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main)
rm -rf vtes-rulings && echo "$sha" > commit
for f in rulings groups references; do
	curl -fsS --create-dirs -o "vtes-rulings/$sha/$f.yaml" "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/$sha/rulings/$f.yaml"
done
curl -fsS -o vtes.json https://static.krcg.org/data/v4/vtes.json
curl -fsS https://usenet.krcg.org/archive.json.gz | gunzip |
	jq -c '{LastPost, Threads: [.Threads[] | {ThreadId, Messages: [.Messages[] | {Id}]}]}' | gzip -9 > archive.json.gz
