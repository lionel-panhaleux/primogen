#!/bin/sh
# Pins the real upstreams for the tests: the card data, vtes-rulings at its
# head, vekn.net's rulebook, tournament rules and judges' guide, zero-vtes's
# rules-qa and rulings ledger, and the archive cut to its ids and titles, with
# the text of the posts the rulings and the advanced rules cite.
set -eu
cd "$(dirname "$0")"
sha=$(curl -fsS -H 'Accept: application/vnd.github.sha' https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main)
rm -rf vtes-rulings && echo "$sha" > commit
for f in rulings groups references; do
	curl -fsS --create-dirs -o "vtes-rulings/$sha/$f.yaml" "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/$sha/rulings/$f.yaml"
done
curl -fsS -o vtes.json https://static.krcg.org/data/v4/vtes.json
for page in rulebook tournament-rules judges-guide; do
	curl -fsS -A primogen --create-dirs -o "vekn/$page.html" "https://www.vekn.net/$page"
done
for f in docs/rules-qa.md data/rulings-v5.csv; do
	curl -fsS --create-dirs -o "zero-vtes/$f" "https://raw.githubusercontent.com/lionel-panhaleux/zero-vtes/main/$f"
done
# the cited posts, as {"<thread>": ["m<n>" or "<forum post>", ...]}
cited=$(
	{
		grep -ohE 'https://usenet\.krcg\.org/t/[^/]+/#[A-Za-z0-9_-]+' "vtes-rulings/$sha/references.yaml" ../../knowledge/rules/advanced-rules.md |
			sed -E 's|https://usenet\.krcg\.org/t/([^/]+)/#(.*)|\1 \2|'
		grep -ohE 'https://www\.vekn\.net/forum/[^ ]*/[0-9]+-[^/# ]*#[0-9]+' "vtes-rulings/$sha/references.yaml" |
			sed -E 's|.*/([0-9]+)-[^/#]*#([0-9]+)$|vekn-\1 \2|'
	} | jq -Rn '[inputs | split(" ")] | group_by(.[0]) | map({key: .[0][0], value: map(.[1])}) | from_entries'
)
curl -fsS https://usenet.krcg.org/archive.json.gz | gunzip |
	jq -c --argjson cited "$cited" '{LastPost, Threads: [.Threads[] | .ThreadId as $t | ($cited[$t] // []) as $c |
		{ThreadId, Title, Messages: [.Messages | to_entries[] | .key as $i | .value as $m |
			if ($c | index("m\($i)")) or ($m.Id != null and ($c | index($m.Id)))
			then $m | {Id, Author, Date, Body} else {Id: $m.Id} end]}]}' | gzip -9 > archive.json.gz
