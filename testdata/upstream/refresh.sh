#!/bin/sh
# Pins the real upstreams for the tests: the card data, the TWDA cut to the
# archetype labels' window (their params.since) and three decks the tests
# name, decks of three deckbuilders, vtes-rulings at its head, archon-vibe's rules documents
# at its head with the ids archon.vekn.net gives them, zero-vtes's rules-qa and rulings ledger, and the archive cut to
# its ids and titles, with the text of the posts the rulings and the advanced
# rules cite, of the threads the calibration corpus's settled-by names and of
# the threads the tests name.
# Needs curl, jq, node and npm (mise install).
set -eu
cd "$(dirname "$0")"
sha=$(curl -fsS -H 'Accept: application/vnd.github.sha' https://api.github.com/repos/vtes-biased/vtes-rulings/commits/main)
rm -rf vtes-rulings && echo "$sha" > commit
for f in rulings groups references; do
	curl -fsS --create-dirs -o "vtes-rulings/$sha/$f.yaml" "https://raw.githubusercontent.com/vtes-biased/vtes-rulings/$sha/rulings/$f.yaml"
done
curl -fsS -o vtes.json https://static.krcg.org/data/vtes.json
curl -fsS -o expansions.json https://static.krcg.org/data/v5/expansions.json
since=$(jq -r .params.since ../../knowledge/classification.json)
curl -fsS https://static.krcg.org/data/twda.json |
	jq -c --arg since "$since" '[to_entries[] | select(.value.event.date >= $since or (.key | IN("2k6munduko", "12434", "zink2k1")))] | from_entries' > twda.json
# the deckbuilders' APIs, each under its host: a TWDA deck as VDB's and VTES
# Decks's answer it, under an id the TWDA lacks, which the tests read from them;
# the graded Banu list from VDB; an Amaranth deck, the answer for a deck
# Amaranth lacks, and its card catalog
curl -fsS --create-dirs -o vdb.im/api/deck/notwda https://vdb.im/api/deck/12947
curl -fsS -o vdb.im/api/deck/7dcfd9fcb https://vdb.im/api/deck/7dcfd9fcb
curl -fsS --create-dirs -o api.vtesdecks.com/1.0/decks/notwda https://api.vtesdecks.com/1.0/decks/tournament-12947
for id in 4d3aa426-70da-44b7-8cb7-92377a1a0dbd notpinned; do
	curl -fsS --create-dirs -o "amaranth.vtes.co.nz/api/deck/$id" "https://amaranth.vtes.co.nz/api/deck?id=$id"
done
curl -fsS -o amaranth.vtes.co.nz/api/cards https://amaranth.vtes.co.nz/api/cards
# the last commits touching archon-vibe's documents and its renderer
last() { curl -fsS "https://api.github.com/repos/vtes-biased/archon-vibe/commits?per_page=1&path=$1" | jq -r '.[0].sha'; }
archon=$(last frontend/src/lib/help-content)
renderer=$(last frontend/src/lib/markdown.ts)
rm -rf archon-vibe && echo "$archon" > archon-commit && echo "$renderer" > archon-renderer
raw="https://raw.githubusercontent.com/vtes-biased/archon-vibe/$archon/frontend"
for f in vtes-rules tournament-rules judges-guide; do
	curl -fsS --create-dirs -o "archon-vibe/$archon/$f.md" "$raw/src/lib/help-content/$f.md"
done
# the ids archon.vekn.net gives them, by its renderer of that commit; the
# sanitizer it runs after does not touch ids
render=$(mktemp -d)
curl -fsS "https://raw.githubusercontent.com/vtes-biased/archon-vibe/$renderer/frontend/src/lib/markdown.ts" |
	sed 's|^import DOMPurify from "dompurify";|const DOMPurify = { sanitize: (h: string, _o?: unknown) => h };|' > "$render/markdown.ts"
cat > "$render/ids.ts" <<'TS'
import { readFileSync } from "node:fs";
import { basename } from "node:path";
import { renderDocument } from "./markdown.ts";
const ids: Record<string, string[]> = {};
for (const path of process.argv.slice(2)) {
  ids[basename(path, ".md")] = [...renderDocument(readFileSync(path, "utf8")).matchAll(/ id="([^"]+)"/g)].map((m) => m[1]!);
}
console.log(JSON.stringify(ids));
TS
marked=$(curl -fsS "https://raw.githubusercontent.com/vtes-biased/archon-vibe/$renderer/frontend/package.json" | jq -r '.dependencies.marked')
(cd "$render" && npm init -y >/dev/null && npm pkg set type=module && npm install --silent --no-audit --no-fund "marked@$marked")
node "$render/ids.ts" archon-vibe/"$archon"/*.md | jq . > "archon-vibe/$archon/ids.json"
rm -rf "$render"
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
whole=$({
	grep -h '^settled-by:' ../../calibration/corpus/*.md | grep -oE 'usenet\.krcg\.org/t/[^/#]+/' |
		sed -E 's|usenet\.krcg\.org/t/([^/]+)/|\1|'
	printf '%s\n' z_kvG6T-DTQ w6jOOcMIUbM
} | jq -Rn '[inputs]')
curl -fsS https://usenet.krcg.org/archive.json.gz | gunzip |
	jq -c --argjson cited "$cited" --argjson whole "$whole" '{LastPost, Threads: [.Threads[] | .ThreadId as $t | ($cited[$t] // []) as $c |
		{ThreadId, Title, Messages: [.Messages | to_entries[] | .key as $i | .value as $m |
			if ($whole | index([$t])) or ($c | index("m\($i)")) or ($m.Id != null and ($c | index($m.Id)))
			then $m | {Id, Author, Date, Body} else {Id: $m.Id} end]}]}' | gzip -9 > archive.json.gz
