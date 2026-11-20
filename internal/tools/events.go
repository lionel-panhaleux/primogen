package tools

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type eventsInput struct {
	Tournament string `json:"tournament,omitempty"`
	Member     string `json:"member,omitempty"`
	Rankings   bool   `json:"rankings,omitempty"`
	Category   string `json:"category,omitempty"`
	League     string `json:"league,omitempty"`
	Links      bool   `json:"links,omitempty"`
	Country    string `json:"country,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	Rank       string `json:"rank,omitempty"`
	Format     string `json:"format,omitempty"`
	State      string `json:"state,omitempty"`
	Detail     bool   `json:"detail,omitempty"`
}

var eventsSchema = &jsonschema.Schema{
	Type: "object",
	Properties: map[string]*jsonschema.Schema{
		"tournament": {Type: "string", Description: `a tournament's event code ("MU-2606"), Archon uid, VEKN event id ("13370"), or words of its name or city; or "list" (or omitted with filters) to list tournaments`},
		"member":     {Type: "string", Description: `a member's VEKN ID ("3200340") or Archon uid: their ratings, worldwide and national ranks, known wins, tournament history, published decklists and community links`},
		"rankings":   {Type: "boolean", Description: "true to list the rating rankings (constructed_offline by default, or the category given)"},
		"category":   {Type: "string", Description: "rating category for rankings: constructed_offline, constructed_online, limited_offline, limited_online"},
		"league":     {Type: "string", Description: `a league's Archon uid or words of its name, or "all" to list leagues`},
		"links":      {Type: "boolean", Description: "true to list the VTES community links (websites, Discords, forums, channels); works without Archon login"},
		"country":    {Type: "string", Description: `ISO 3166-1 alpha-2 country code ("FR", "SE") to filter tournaments, rankings, leagues or community links`},
		"from":       {Type: "string", Description: "start date lower bound (YYYY-MM-DD) for filtering tournaments"},
		"to":         {Type: "string", Description: "start date upper bound (YYYY-MM-DD) for filtering tournaments"},
		"rank":       {Type: "string", Description: `championship level for filtering tournaments: "NC" ("National Championship"), "CC" ("Continental Championship"), or "basic"`},
		"format":     {Type: "string", Description: "tournament format for filtering tournaments: Standard, V5, Limited, Storyline"},
		"state":      {Type: "string", Description: "tournament state for filtering tournaments: Planned, Registration, Waiting, Playing, Finished"},
		"detail":     {Type: "boolean", Description: "longer lists (up to 60 rows instead of 20, and every standing of a tournament)"},
	},
}

var ratingCategories = []struct {
	key, label string
}{
	{"constructed_offline", "Constructed offline"},
	{"constructed_online", "Constructed online"},
	{"limited_offline", "Limited offline"},
	{"limited_online", "Limited online"},
}

const archonLoginRequired = "Archon tournaments, standings, rankings, member records and decklists are available to logged-in VEKN members only: " +
	"tell the user to connect to https://mcp.krcg.org/mcp with Archon login (or https://mcp.krcg.org/mcp?login=1 in a client that starts OAuth on an initial 401; local installs do not hold Archon data). " +
	"Community links (call vtes_events with {\"links\": true}) and all public tools work without login."

func addEvents(s *mcp.Server, l *Live) {
	mcp.AddTool(s, &mcp.Tool{
		Name:  "vtes_events",
		Title: "VTES tournaments, rankings, members and community links",
		Description: "Call before any claim about Vampire: The Eternal Struggle (VTES) tournaments, standings, player rankings, ratings, a member's wins or decklists, leagues, or community links: never state tournament results, rankings or links from memory. " +
			"With Archon login (a VEKN member signed in at https://mcp.krcg.org/mcp), returns tournaments from Archon (by event code, uid, VEKN event id, name, country, date window, championship rank, format or state) with their standings and published decklists; " +
			"rating rankings by category and country; a member's ratings, worldwide and national ranks, known wins, tournament history and published decklists (by VEKN ID; Archon's export publishes VEKN IDs, never member names); and leagues. " +
			"With links: true, returns the community links (websites, Discords, forums, channels) from Archon, available to everyone without login.",
		InputSchema: eventsSchema,
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: new(false)},
	}, func(_ context.Context, req *mcp.CallToolRequest, in eventsInput) (*mcp.CallToolResult, any, error) {
		member := MemberFrom(req)
		a := l.Archon()
		linksOnly := in.Links && in.Tournament == "" && in.Member == "" && !in.Rankings && in.Category == "" && in.League == ""
		if !linksOnly && member == "" {
			return nil, nil, fmt.Errorf("%s", archonLoginRequired)
		}
		if a == nil {
			if linksOnly {
				return nil, nil, fmt.Errorf("this server holds no Archon export (local installs do not download Archon data): connect to https://mcp.krcg.org/mcp for community links and member data")
			}
			return nil, nil, fmt.Errorf("the Archon export is not loaded on this server yet: try again shortly")
		}
		d := l.p.Load()
		var out strings.Builder
		var err error
		switch {
		case linksOnly:
			writeCommunityLinks(&out, a, in.Country, member != "", in.Detail)
		case in.Member != "":
			err = writeMember(&out, d, a, in.Member, in.Detail)
		case in.Rankings || in.Category != "":
			err = writeRankings(&out, a, in.Category, in.Country, in.Detail)
		case in.League != "":
			err = writeLeagues(&out, a, in.League, in.Country, in.Format, in.Detail)
		default:
			err = writeTournaments(&out, d, a, in, in.Detail)
		}
		if err != nil {
			return nil, nil, err
		}
		if in.Links && !linksOnly {
			out.WriteString("\n")
			writeCommunityLinks(&out, a, in.Country, member != "", in.Detail)
		}
		out.WriteString("\n## How to answer\n\n" +
			"Answer from these Archon figures, never from memory. Archon's export identifies members by VEKN ID, never by name: cite the VEKN ID and never guess a name for it.")
		if !linksOnly {
			out.WriteString(" For any published decklist above, call vtes_deck with its deck uid for its full card list and analysis.")
		}
		out.WriteString("\n\n")
		out.WriteString(ArchonDataLine(d.b, a))
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: out.String()}}}, nil, nil
	})
}

func ArchonDataLine(b *bundle.Bundle, a *Archon) string {
	base := strings.TrimSuffix(DataLine(b), ".")
	if a == nil {
		return base + "."
	}
	return fmt.Sprintf("%s; archon.vekn.net export %s.", base, a.Date())
}

func (d *data) archonArchetype(a *Archon, ad *ArchonDeck) string {
	if d.rel == nil || ad == nil {
		return ""
	}
	if ad.Winner {
		if t := a.TournamentByUID[ad.TournamentUID]; t != nil {
			if vekn := t.ExternalIDs["vekn"]; vekn != "" {
				if lbl := d.rel.Label(vekn); lbl != nil {
					return fmt.Sprintf("%s (%s)", lbl.Name, lbl.Ref)
				}
			}
		}
	}
	v := d.rel.Vector(ad.DeckCards())
	if v.Empty() {
		return ""
	}
	if best := d.rel.Archetypes(v, "", 1); len(best) > 0 && best[0].Similarity >= d.rel.Threshold {
		arch := d.rel.Class.Archetypes[best[0].Index]
		return fmt.Sprintf("%s (%s)", arch.Name, arch.Ref)
	}
	return ""
}

func (d *data) deckCounts(ad *ArchonDeck) (crypt, library int) {
	for id, n := range ad.Cards {
		if n <= 0 {
			continue
		}
		if c := d.byID[id]; c != nil && c.IsCrypt() {
			crypt += n
		} else {
			library += n
		}
	}
	return crypt, library
}

func writeCommunityLinks(out *strings.Builder, a *Archon, country string, loggedIn, detail bool) {
	country = strings.ToUpper(strings.TrimSpace(country))
	type item struct {
		link   ArchonLink
		veknID string
	}
	var items []item
	seen := map[string]bool{}
	for _, u := range a.Users {
		for _, lk := range u.CommunityLinks {
			if lk.Moderation == "hidden" || lk.URL == "" {
				continue
			}
			lk.Country = cmp.Or(lk.Country, u.Country)
			if country != "" && !strings.EqualFold(lk.Country, country) && lk.Moderation != "global" {
				continue
			}
			key := lk.URL + "|" + lk.Country
			if seen[key] {
				continue
			}
			seen[key] = true
			items = append(items, item{link: lk, veknID: u.VEKNID})
		}
	}
	modRank := func(m string) int {
		switch m {
		case "global":
			return 0
		case "national":
			return 1
		default:
			return 2
		}
	}
	slices.SortFunc(items, func(x, y item) int {
		return cmp.Or(
			cmp.Compare(modRank(x.link.Moderation), modRank(y.link.Moderation)),
			strings.Compare(x.link.Country, y.link.Country),
			strings.Compare(x.link.Type, y.link.Type),
			strings.Compare(x.link.Label, y.link.Label),
			strings.Compare(x.link.URL, y.link.URL),
		)
	})
	scope := "all countries"
	if country != "" {
		scope = country + " (and global)"
	}
	fmt.Fprintf(out, "# Community links (%d, %s)\n\n", len(items), scope)
	if len(items) == 0 {
		out.WriteString("No community links found.\n")
		return
	}
	limit := 30
	if detail {
		limit = 120
	}
	for _, it := range items[:min(len(items), limit)] {
		var meta []string
		if it.link.Type != "" {
			meta = append(meta, it.link.Type)
		}
		if it.link.Country != "" {
			meta = append(meta, it.link.Country)
		}
		if len(it.link.Languages) > 0 {
			meta = append(meta, "languages: "+strings.Join(it.link.Languages, ", "))
		}
		if it.link.Moderation != "" {
			meta = append(meta, it.link.Moderation)
		}
		if loggedIn && it.veknID != "" {
			meta = append(meta, "VEKN "+it.veknID)
		}
		label := cmp.Or(it.link.Label, it.link.URL)
		fmt.Fprintf(out, "- **%s**: %s (%s)\n", label, it.link.URL, strings.Join(meta, "; "))
	}
	if len(items) > limit {
		fmt.Fprintf(out, "\nShowing %d of %d links; pass detail: true or filter by country for more.\n", limit, len(items))
	}
}

func (a *Archon) rankInCategory(target *ArchonUser, cat, country string) (rank, total int) {
	r := target.Rating(cat)
	if r == nil || r.Total <= 0 {
		return 0, 0
	}
	higher := 0
	for i := range a.Users {
		u := &a.Users[i]
		if country != "" && !strings.EqualFold(u.Country, country) {
			continue
		}
		ur := u.Rating(cat)
		if ur == nil || ur.Total <= 0 {
			continue
		}
		total++
		if ur.Total > r.Total {
			higher++
		}
	}
	return higher + 1, total
}

func writeRankings(out *strings.Builder, a *Archon, cat, country string, detail bool) error {
	cat = cmp.Or(strings.ToLower(strings.TrimSpace(cat)), "constructed_offline")
	var label string
	for _, c := range ratingCategories {
		if c.key == cat {
			label = c.label
			break
		}
	}
	if label == "" {
		return fmt.Errorf("unknown rating category %q: choose constructed_offline, constructed_online, limited_offline or limited_online", cat)
	}
	country = strings.ToUpper(strings.TrimSpace(country))
	var rated []*ArchonUser
	for i := range a.Users {
		u := &a.Users[i]
		if country != "" && !strings.EqualFold(u.Country, country) {
			continue
		}
		if r := u.Rating(cat); r != nil && r.Total > 0 {
			rated = append(rated, u)
		}
	}
	slices.SortFunc(rated, func(x, y *ArchonUser) int {
		return cmp.Or(cmp.Compare(y.Rating(cat).Total, x.Rating(cat).Total), strings.Compare(x.VEKNID, y.VEKNID))
	})
	scope := "worldwide"
	if country != "" {
		scope = country
	}
	fmt.Fprintf(out, "# Rankings: %s (%s, %d rated members)\n\n", label, scope, len(rated))
	if len(rated) == 0 {
		out.WriteString("No rated members in this scope.\n")
		return nil
	}
	limit := 25
	if detail {
		limit = 100
	}
	rank := 1
	for i, u := range rated[:min(len(rated), limit)] {
		r := u.Rating(cat)
		if i > 0 && r.Total < rated[i-1].Rating(cat).Total {
			rank = i + 1
		}
		place := u.Country
		if u.City != "" && u.Country != "" {
			place = u.City + ", " + u.Country
		}
		fmt.Fprintf(out, "%d. **VEKN %s** (%s): **%d pts** across %d rated events; %d known wins, %d published decks\n",
			rank, u.VEKNID, cmp.Or(place, "no country"), r.Total, len(r.Tournaments), len(u.Wins), len(a.DecksByUser[u.UID]))
	}
	if len(rated) > limit {
		fmt.Fprintf(out, "\nShowing top %d of %d rated members; pass detail: true or filter by country for more.\n", limit, len(rated))
	}
	return nil
}

func writeMember(out *strings.Builder, d *data, a *Archon, q string, detail bool) error {
	u := a.FindUser(q)
	if u == nil {
		return fmt.Errorf("no VEKN member in the Archon export has the VEKN ID or uid %q", q)
	}
	limit := 20
	if detail {
		limit = 60
	}
	fmt.Fprintf(out, "# Member: %s\n\n", a.VEKN(u.UID))
	var meta []string
	if u.Country != "" {
		if u.City != "" {
			meta = append(meta, "Location: "+u.City+", "+u.Country)
		} else {
			meta = append(meta, "Country: "+u.Country)
		}
	}
	if len(u.Roles) > 0 {
		meta = append(meta, "Roles: "+strings.Join(u.Roles, ", "))
	}
	meta = append(meta, "Archon uid: "+u.UID)
	fmt.Fprintln(out, strings.Join(meta, " | "))

	out.WriteString("\n## Ratings\n\n")
	anyRating := false
	for _, cat := range ratingCategories {
		r := u.Rating(cat.key)
		if r == nil || r.Total <= 0 {
			continue
		}
		anyRating = true
		wRank, wTotal := a.rankInCategory(u, cat.key, "")
		nat := ""
		if u.Country != "" {
			nRank, nTotal := a.rankInCategory(u, cat.key, u.Country)
			nat = fmt.Sprintf(", #%d of %d in %s", nRank, nTotal, u.Country)
		}
		fmt.Fprintf(out, "- **%s**: **%d pts** (rank #%d of %d worldwide%s, %d rated events)\n",
			cat.label, r.Total, wRank, wTotal, nat, len(r.Tournaments))
		for _, tr := range r.Tournaments {
			rk := ""
			if tr.Rank != "" {
				rk = ", " + tr.Rank
			}
			place := fmt.Sprintf("%d players", tr.PlayerCount)
			if tr.Position > 0 {
				place = fmt.Sprintf("place %d/%d", tr.Position, tr.PlayerCount)
			}
			fmt.Fprintf(out, "  - %s — %s (%s%s): %d pts (%d GW, %.1f VP, %s)\n",
				tr.Date, tr.TournamentName, tr.TournamentUID, rk, tr.Points, tr.GW, tr.VP, place)
		}
	}
	if !anyRating {
		out.WriteString("No active rating points in the current 18-month rating window.\n")
	}

	out.WriteString("\n## Known wins\n\n")
	if len(u.Wins) == 0 {
		out.WriteString("Known wins in Archon: 0.\n")
	} else {
		type winRow struct {
			uid string
			t   *ArchonTournament
		}
		rows := make([]winRow, 0, len(u.Wins))
		for _, wUID := range u.Wins {
			rows = append(rows, winRow{uid: wUID, t: a.TournamentByUID[wUID]})
		}
		slices.SortFunc(rows, func(x, y winRow) int {
			var dx, dy string
			if x.t != nil {
				dx = x.t.Start
			}
			if y.t != nil {
				dy = y.t.Start
			}
			return cmp.Or(strings.Compare(dy, dx), strings.Compare(x.uid, y.uid))
		})
		fmt.Fprintf(out, "%d known win(s):\n", len(rows))
		for _, w := range rows[:min(len(rows), limit)] {
			if w.t == nil {
				fmt.Fprintf(out, "- Tournament %s\n", w.uid)
				continue
			}
			t := w.t
			code := cmp.Or(t.EventCode, t.UID)
			deckNote := ""
			for _, ad := range a.DecksByTournament[t.UID] {
				if ad.Winner || ad.UserUID == u.UID {
					arch := d.archonArchetype(a, ad)
					if arch != "" {
						arch = ", " + arch
					}
					deckNote = fmt.Sprintf(" — deck `%s` (%s%s)", ad.UID, cmp.Or(ad.Name, "unnamed"), arch)
					break
				}
			}
			fmt.Fprintf(out, "- **%s — %s** (`%s`, %s, %s, %d players)%s\n",
				t.StartDate(), t.Name, code, cmp.Or(t.Place(), "no location"), t.Format, t.PlayerCount(), deckNote)
		}
		if len(rows) > limit {
			fmt.Fprintf(out, "\nShowing %d of %d known wins; pass detail: true for more.\n", limit, len(rows))
		}
	}

	var decks []*ArchonDeck
	seenDecks := map[string]bool{}
	for _, ad := range a.DecksByUser[u.UID] {
		seenDecks[ad.UID] = true
		decks = append(decks, ad)
	}
	if u.VEKNID != "" {
		for i := range a.Decks {
			ad := &a.Decks[i]
			if ad.Attribution.VEKNID == u.VEKNID && !seenDecks[ad.UID] {
				seenDecks[ad.UID] = true
				decks = append(decks, ad)
			}
		}
	}
	slices.SortFunc(decks, func(x, y *ArchonDeck) int {
		var dx, dy string
		if tx := a.TournamentByUID[x.TournamentUID]; tx != nil {
			dx = tx.Start
		}
		if ty := a.TournamentByUID[y.TournamentUID]; ty != nil {
			dy = ty.Start
		}
		return cmp.Or(strings.Compare(dy, dx), strings.Compare(x.UID, y.UID))
	})
	out.WriteString("\n## Published decklists\n\n")
	if len(decks) == 0 {
		out.WriteString("Published decklists in Archon: 0.\n")
	} else {
		fmt.Fprintf(out, "%d published decklist(s) (pass a deck uid to vtes_deck for the full list and analysis):\n", len(decks))
		for _, ad := range decks[:min(len(decks), limit)] {
			t := a.TournamentByUID[ad.TournamentUID]
			ev := ad.TournamentUID
			if t != nil {
				ev = fmt.Sprintf("%s — %s (`%s`, %d players)", t.StartDate(), t.Name, cmp.Or(t.EventCode, t.UID), t.PlayerCount())
			}
			role := "played"
			if ad.Winner {
				role = "winner"
			} else if ad.UserUID != u.UID && u.VEKNID != "" && ad.Attribution.VEKNID == u.VEKNID {
				role = "credited designer"
			}
			arch := d.archonArchetype(a, ad)
			if arch != "" {
				arch = ", archetype: " + arch
			}
			crypt, lib := d.deckCounts(ad)
			fmt.Fprintf(out, "- Deck `%s`: **%s** (%s, %d crypt / %d library%s) at %s\n",
				ad.UID, cmp.Or(ad.Name, "unnamed"), role, crypt, lib, arch, ev)
		}
		if len(decks) > limit {
			fmt.Fprintf(out, "\nShowing %d of %d published decklists; pass detail: true for more.\n", limit, len(decks))
		}
	}

	type playedEvent struct {
		t      *ArchonTournament
		gw, vp float64
		tp     int
		tag    string
	}
	var history []playedEvent
	for i := range a.Tournaments {
		t := &a.Tournaments[i]
		found := false
		var pe playedEvent
		pe.t = t
		for idx, s := range t.Standings {
			if s.UserUID == u.UID {
				found = true
				pe.gw, pe.vp, pe.tp = s.GW, s.VP, s.TP
				switch {
				case t.Winner == u.UID:
					pe.tag = "winner"
				case s.Finalist:
					pe.tag = "finalist"
				default:
					pe.tag = fmt.Sprintf("place #%d", idx+1)
				}
				break
			}
		}
		if !found {
			for _, p := range t.Players {
				if p.UserUID == u.UID {
					found = true
					pe.gw, pe.vp, pe.tp = float64(p.Result.GW), p.Result.VP, p.Result.TP
					switch {
					case t.Winner == u.UID:
						pe.tag = "winner"
					case p.Finalist:
						pe.tag = "finalist"
					default:
						pe.tag = p.State
					}
					break
				}
			}
		}
		if found {
			history = append(history, pe)
		}
	}
	slices.SortFunc(history, func(x, y playedEvent) int {
		return cmp.Or(strings.Compare(y.t.Start, x.t.Start), strings.Compare(x.t.UID, y.t.UID))
	})
	out.WriteString("\n## Tournament history\n\n")
	if len(history) == 0 {
		out.WriteString("No tournament appearances recorded in the Archon export.\n")
	} else {
		limit := 20
		if detail {
			limit = 60
		}
		fmt.Fprintf(out, "%d tournament appearance(s):\n", len(history))
		for _, pe := range history[:min(len(history), limit)] {
			t := pe.t
			fmt.Fprintf(out, "- %s — **%s** (`%s`, %s, %s, %d players): %s (%.0f GW, %.1f VP, %d TP)\n",
				t.StartDate(), t.Name, cmp.Or(t.EventCode, t.UID), cmp.Or(t.Place(), "no location"), t.Format, t.PlayerCount(), pe.tag, pe.gw, pe.vp, pe.tp)
		}
		if len(history) > limit {
			fmt.Fprintf(out, "\nShowing %d of %d appearances; pass detail: true for more.\n", limit, len(history))
		}
	}

	if len(u.CommunityLinks) > 0 {
		out.WriteString("\n## Community links\n\n")
		for _, lk := range u.CommunityLinks {
			if lk.Moderation == "hidden" || lk.URL == "" {
				continue
			}
			fmt.Fprintf(out, "- **%s** (%s): %s\n", cmp.Or(lk.Label, lk.URL), lk.Type, lk.URL)
		}
	}
	return nil
}

func writeLeagues(out *strings.Builder, a *Archon, q, country, format string, detail bool) error {
	q = strings.TrimSpace(q)
	country = strings.ToUpper(strings.TrimSpace(country))
	format = strings.TrimSpace(format)
	var matches []*ArchonLeague
	if lg := a.LeagueByUID[q]; lg != nil {
		matches = []*ArchonLeague{lg}
	} else {
		folded := bundle.Fold(q)
		all := strings.EqualFold(q, "all")
		for i := range a.Leagues {
			lg := &a.Leagues[i]
			if country != "" && !strings.EqualFold(lg.Country, country) {
				continue
			}
			if format != "" && !strings.EqualFold(lg.Format, format) {
				continue
			}
			if all || strings.Contains(bundle.Fold(lg.Name), folded) {
				matches = append(matches, lg)
			}
		}
	}
	if len(matches) == 0 {
		return fmt.Errorf("no league in the Archon export matches %q", q)
	}
	slices.SortFunc(matches, func(x, y *ArchonLeague) int {
		return cmp.Or(strings.Compare(y.Start, x.Start), strings.Compare(x.Name, y.Name))
	})
	limit := 20
	if detail {
		limit = 60
	}
	if len(matches) == 1 && !strings.EqualFold(q, "all") {
		lg := matches[0]
		fmt.Fprintf(out, "# League: %s\n\n", lg.Name)
		fmt.Fprintf(out, "Archon uid: `%s` | Kind: %s | Scoring: %s | Format: %s | Country: %s | Dates: %s to %s\n",
			lg.UID, cmp.Or(lg.Kind, "League"), cmp.Or(lg.StandingsMode, "RTP"), cmp.Or(lg.Format, "any"), cmp.Or(lg.Country, "worldwide"), cmp.Or(lg.Start, "?"), cmp.Or(lg.Finish, "ongoing"))
		if len(lg.OrganizersUIDs) > 0 {
			orgs := make([]string, 0, len(lg.OrganizersUIDs))
			for _, o := range lg.OrganizersUIDs {
				orgs = append(orgs, a.VEKN(o))
			}
			fmt.Fprintf(out, "Organizers: %s\n", strings.Join(orgs, ", "))
		}
		if lg.Description != "" {
			fmt.Fprintf(out, "\n%s\n", strings.TrimSpace(lg.Description))
		}
		if lg.Final != nil && lg.Final.Winner != "" {
			fmt.Fprintf(out, "\nLeague final winner: **%s**\n", a.VEKN(lg.Final.Winner))
		}
		var events []*ArchonTournament
		for i := range a.Tournaments {
			if a.Tournaments[i].LeagueUID == lg.UID {
				events = append(events, &a.Tournaments[i])
			}
		}
		slices.SortFunc(events, func(x, y *ArchonTournament) int {
			return strings.Compare(y.Start, x.Start)
		})
		fmt.Fprintf(out, "\n## League tournaments (%d)\n\n", len(events))
		for _, t := range events[:min(len(events), limit)] {
			win := ""
			if t.Winner != "" {
				win = ", winner: " + a.VEKN(t.Winner)
			}
			fmt.Fprintf(out, "- %s — **%s** (`%s`, %s, %s, %d players%s)\n",
				t.StartDate(), t.Name, cmp.Or(t.EventCode, t.UID), cmp.Or(t.Place(), "no location"), t.State, t.PlayerCount(), win)
		}
		if len(events) > limit {
			fmt.Fprintf(out, "\nShowing %d of %d league tournaments; pass detail: true for more.\n", limit, len(events))
		}
		return nil
	}
	fmt.Fprintf(out, "# Leagues (%d)\n\n", len(matches))
	for _, lg := range matches[:min(len(matches), limit)] {
		events := 0
		for i := range a.Tournaments {
			if a.Tournaments[i].LeagueUID == lg.UID {
				events++
			}
		}
		fmt.Fprintf(out, "- **%s** (`%s`, %s, %s, %s, %s to %s, %d tournaments)\n",
			lg.Name, lg.UID, cmp.Or(lg.Kind, "League"), cmp.Or(lg.StandingsMode, "RTP"), cmp.Or(lg.Country, "worldwide"), cmp.Or(lg.Start, "?"), cmp.Or(lg.Finish, "ongoing"), events)
	}
	if len(matches) > limit {
		fmt.Fprintf(out, "\nShowing %d of %d leagues; pass detail: true or filter by country or format for more.\n", limit, len(matches))
	}
	return nil
}

func matchRank(tRank, q string) bool {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return true
	}
	switch q {
	case "nc", "national", "national championship":
		return strings.EqualFold(tRank, "National Championship")
	case "cc", "continental", "continental championship":
		return strings.EqualFold(tRank, "Continental Championship")
	case "basic", "standard":
		return tRank == ""
	}
	return strings.Contains(strings.ToLower(tRank), q)
}

func writeTournaments(out *strings.Builder, d *data, a *Archon, in eventsInput, detail bool) error {
	q := strings.TrimSpace(in.Tournament)
	if q != "" && !strings.EqualFold(q, "list") {
		if exact := a.FindTournament(q); exact != nil {
			writeTournamentDossier(out, d, a, exact, detail)
			return nil
		}
	}
	folded := ""
	if q != "" && !strings.EqualFold(q, "list") {
		folded = bundle.Fold(q)
	}
	country := strings.ToUpper(strings.TrimSpace(in.Country))
	format := strings.TrimSpace(in.Format)
	state := strings.TrimSpace(in.State)
	var matches []*ArchonTournament
	for i := range a.Tournaments {
		t := &a.Tournaments[i]
		if country != "" && !strings.EqualFold(t.Country, country) {
			continue
		}
		if in.From != "" && t.StartDate() < in.From {
			continue
		}
		if in.To != "" && (t.StartDate() == "" || t.StartDate() > in.To) {
			continue
		}
		if !matchRank(t.Rank, in.Rank) {
			continue
		}
		if format != "" && !strings.EqualFold(t.Format, format) {
			continue
		}
		if state != "" && !strings.EqualFold(t.State, state) {
			continue
		}
		if folded != "" && !strings.Contains(bundle.Fold(t.Name), folded) && !strings.Contains(bundle.Fold(t.City), folded) && !strings.Contains(strings.ToLower(t.EventCode), folded) {
			continue
		}
		matches = append(matches, t)
	}
	if len(matches) == 0 {
		return fmt.Errorf("no tournament in the Archon export matches the query")
	}
	slices.SortFunc(matches, func(x, y *ArchonTournament) int {
		return cmp.Or(strings.Compare(y.Start, x.Start), strings.Compare(x.UID, y.UID))
	})
	if len(matches) == 1 && folded != "" {
		writeTournamentDossier(out, d, a, matches[0], detail)
		return nil
	}
	fmt.Fprintf(out, "# Tournaments (%d matching)\n\n", len(matches))
	limit := 20
	if detail {
		limit = 60
	}
	for _, t := range matches[:min(len(matches), limit)] {
		rk := ""
		if t.Rank != "" {
			rk = ", " + t.Rank
		}
		win := ""
		if t.Winner != "" {
			win = ", winner: " + a.VEKN(t.Winner)
		}
		decks := len(a.DecksByTournament[t.UID])
		deckInfo := ""
		if decks > 0 {
			deckInfo = fmt.Sprintf(", %d published deck(s)", decks)
		}
		fmt.Fprintf(out, "- **%s — %s** (`%s`, %s, %s%s, %s, %d players%s%s)\n",
			t.StartDate(), t.Name, cmp.Or(t.EventCode, t.UID), cmp.Or(t.Place(), "no location"), t.Format, rk, t.State, t.PlayerCount(), win, deckInfo)
	}
	if len(matches) > limit {
		fmt.Fprintf(out, "\nShowing %d of %d tournaments; pass detail: true or narrow by country, from/to dates, rank, format or state.\n", limit, len(matches))
	}
	return nil
}

func writeTournamentDossier(out *strings.Builder, d *data, a *Archon, t *ArchonTournament, detail bool) {
	fmt.Fprintf(out, "# Tournament: %s\n\n", t.Name)
	var ids []string
	if t.EventCode != "" {
		ids = append(ids, "Code: `"+t.EventCode+"`")
	}
	ids = append(ids, "Archon uid: `"+t.UID+"`")
	if vekn := t.ExternalIDs["vekn"]; vekn != "" {
		ids = append(ids, "VEKN event: `"+vekn+"`")
	}
	fmt.Fprintln(out, strings.Join(ids, " | "))
	rk := cmp.Or(t.Rank, "Basic")
	fmt.Fprintf(out, "State: %s | Format: %s (%s) | Start: %s (%s) | Location: %s | Players: %d | Rounds: %d | Decklists mode: %s\n",
		t.State, t.Format, rk, cmp.Or(t.Start, "?"), cmp.Or(t.Timezone, "UTC"), cmp.Or(t.Place(), "unspecified"), t.PlayerCount(), max(len(t.Rounds), t.MaxRounds), cmp.Or(t.DecklistsMode, "Winner"))
	if t.Venue != "" || t.Address != "" {
		fmt.Fprintf(out, "Venue: %s\n", strings.TrimSpace(t.Venue+" "+t.Address))
	}
	if len(t.OrganizersUIDs) > 0 {
		orgs := make([]string, 0, len(t.OrganizersUIDs))
		for _, o := range t.OrganizersUIDs {
			orgs = append(orgs, a.VEKN(o))
		}
		fmt.Fprintf(out, "Organizers: %s\n", strings.Join(orgs, ", "))
	}
	if t.Winner != "" {
		fmt.Fprintf(out, "Winner: **%s**\n", a.VEKN(t.Winner))
	}

	decksByUser := map[string]*ArchonDeck{}
	for _, ad := range a.DecksByTournament[t.UID] {
		if ad.UserUID != "" {
			decksByUser[ad.UserUID] = ad
		}
	}

	if t.Finals != nil && len(t.Finals.Seating) > 0 {
		out.WriteString("\n## Finals table\n\n")
		for i, seat := range t.Finals.Seating {
			tag := ""
			if t.Winner != "" && seat.PlayerUID == t.Winner {
				tag = " — **Winner**"
			}
			fmt.Fprintf(out, "%d. %s: %d GW, %.1f VP%s\n", i+1, a.VEKN(seat.PlayerUID), seat.Result.GW, seat.Result.VP, tag)
		}
	}

	out.WriteString("\n## Standings\n\n")
	type row struct {
		userUID  string
		gw, vp   float64
		tp       int
		finalist bool
		dq       bool
	}
	var rows []row
	for _, s := range t.Standings {
		rows = append(rows, row{userUID: s.UserUID, gw: s.GW, vp: s.VP, tp: s.TP, finalist: s.Finalist, dq: s.Disqualified})
	}
	if len(rows) == 0 {
		for _, p := range t.Players {
			rows = append(rows, row{userUID: p.UserUID, gw: float64(p.Result.GW), vp: p.Result.VP, tp: p.Result.TP, finalist: p.Finalist, dq: p.State == "Disqualified"})
		}
		slices.SortFunc(rows, func(x, y row) int {
			wx, wy := x.userUID != "" && x.userUID == t.Winner, y.userUID != "" && y.userUID == t.Winner
			if wx != wy {
				if wx {
					return -1
				}
				return 1
			}
			if x.finalist != y.finalist {
				if x.finalist {
					return -1
				}
				return 1
			}
			return cmp.Or(cmp.Compare(y.gw, x.gw), cmp.Compare(y.vp, x.vp), cmp.Compare(y.tp, x.tp))
		})
	}
	if len(rows) == 0 {
		out.WriteString("No standings recorded.\n")
	} else {
		limit := 25
		if detail {
			limit = len(rows)
		}
		for i, r := range rows[:min(len(rows), limit)] {
			var tags []string
			if t.Winner != "" && r.userUID == t.Winner {
				tags = append(tags, "winner")
			} else if r.finalist {
				tags = append(tags, "finalist")
			}
			if r.dq {
				tags = append(tags, "DQ")
			}
			tagStr := ""
			if len(tags) > 0 {
				tagStr = " [" + strings.Join(tags, ", ") + "]"
			}
			deckStr := ""
			if ad := decksByUser[r.userUID]; ad != nil {
				arch := d.archonArchetype(a, ad)
				if arch != "" {
					arch = ", " + arch
				}
				deckStr = fmt.Sprintf(" — deck `%s` (%s%s)", ad.UID, cmp.Or(ad.Name, "unnamed"), arch)
			}
			fmt.Fprintf(out, "%d. **%s**%s: %.0f GW, %.1f VP, %d TP%s\n",
				i+1, a.VEKN(r.userUID), tagStr, r.gw, r.vp, r.tp, deckStr)
		}
		if len(rows) > limit {
			fmt.Fprintf(out, "\nShowing top %d of %d players; pass detail: true for all standings.\n", limit, len(rows))
		}
	}

	decks := a.DecksByTournament[t.UID]
	out.WriteString("\n## Published decklists\n\n")
	if len(decks) == 0 {
		out.WriteString("No published decklists for this tournament.\n")
	} else {
		deckLimit := 25
		if detail {
			deckLimit = len(decks)
		}
		fmt.Fprintf(out, "%d published decklist(s) (call vtes_deck with a deck uid for its full card list and analysis):\n", len(decks))
		for _, ad := range decks[:min(len(decks), deckLimit)] {
			role := a.VEKN(ad.UserUID)
			if ad.Winner {
				role += ", winner"
			}
			arch := d.archonArchetype(a, ad)
			if arch != "" {
				arch = ", archetype: " + arch
			}
			crypt, lib := d.deckCounts(ad)
			fmt.Fprintf(out, "- Deck `%s`: **%s** (%s, %d crypt / %d library%s)\n",
				ad.UID, cmp.Or(ad.Name, "unnamed"), role, crypt, lib, arch)
		}
		if len(decks) > deckLimit {
			fmt.Fprintf(out, "\nShowing %d of %d published decklists; pass detail: true for all decklists.\n", deckLimit, len(decks))
		}
	}
}
