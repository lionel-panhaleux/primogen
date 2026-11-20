package tools

import (
	"cmp"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/lionel-panhaleux/primogen/internal/bundle"
)

type Archon struct {
	GeneratedAt string
	Users       []ArchonUser
	Tournaments []ArchonTournament
	Decks       []ArchonDeck
	Leagues     []ArchonLeague

	UserByUID         map[string]*ArchonUser
	UserByVEKN        map[string]*ArchonUser
	TournamentByUID   map[string]*ArchonTournament
	TournamentByCode  map[string]*ArchonTournament
	TournamentByVEKN  map[string]*ArchonTournament
	DeckByUID         map[string]*ArchonDeck
	DecksByTournament map[string][]*ArchonDeck
	DecksByUser       map[string][]*ArchonDeck
	LeagueByUID       map[string]*ArchonLeague
}

func (a *Archon) Date() string {
	if len(a.GeneratedAt) >= 16 {
		return strings.Replace(a.GeneratedAt[:16], "T", " ", 1) + " UTC"
	}
	return a.GeneratedAt
}

func (a *Archon) VEKN(userUID string) string {
	if userUID == "" {
		return "anonymous"
	}
	if u := a.UserByUID[userUID]; u != nil && u.VEKNID != "" {
		return "VEKN " + u.VEKNID
	}
	return userUID
}

func (a *Archon) FindUser(q string) *ArchonUser {
	q = strings.TrimSpace(q)
	if q == "" {
		return nil
	}
	vekn := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(strings.ToUpper(q), "VEKN")), "#"))
	if vekn != "" {
		if u := a.UserByVEKN[vekn]; u != nil {
			return u
		}
	}
	return a.UserByUID[q]
}

func (a *Archon) FindTournament(q string) *ArchonTournament {
	q = strings.TrimSpace(q)
	if q == "" {
		return nil
	}
	if t := a.TournamentByUID[q]; t != nil {
		return t
	}
	if t := a.TournamentByCode[strings.ToLower(q)]; t != nil {
		return t
	}
	return a.TournamentByVEKN[q]
}

type ArchonUser struct {
	UID                string        `json:"uid"`
	VEKNID             string        `json:"vekn_id"`
	Country            string        `json:"country"`
	City               string        `json:"city"`
	CityGeonameID      int           `json:"city_geoname_id"`
	Roles              []string      `json:"roles"`
	CommunityLinks     []ArchonLink  `json:"community_links"`
	ConstructedOnline  *ArchonRating `json:"constructed_online"`
	ConstructedOffline *ArchonRating `json:"constructed_offline"`
	LimitedOnline      *ArchonRating `json:"limited_online"`
	LimitedOffline     *ArchonRating `json:"limited_offline"`
	Wins               []string      `json:"wins"`
}

func (u *ArchonUser) Rating(cat string) *ArchonRating {
	switch cat {
	case "constructed_offline":
		return u.ConstructedOffline
	case "constructed_online":
		return u.ConstructedOnline
	case "limited_offline":
		return u.LimitedOffline
	case "limited_online":
		return u.LimitedOnline
	}
	return nil
}

type ArchonLink struct {
	Type       string   `json:"type"`
	URL        string   `json:"url"`
	Label      string   `json:"label"`
	Languages  []string `json:"languages"`
	Country    string   `json:"country"`
	Moderation string   `json:"moderation"`
}

type ArchonRating struct {
	Total       int                 `json:"total"`
	Tournaments []ArchonRatedResult `json:"tournaments"`
}

type ArchonRatedResult struct {
	TournamentUID    string  `json:"tournament_uid"`
	TournamentName   string  `json:"tournament_name"`
	Date             string  `json:"date"`
	PlayerCount      int     `json:"player_count"`
	Rank             string  `json:"rank"`
	VP               float64 `json:"vp"`
	GW               int     `json:"gw"`
	FinalistPosition int     `json:"finalist_position"`
	Points           int     `json:"points"`
	Position         int     `json:"position"`
}

type ArchonTournament struct {
	UID                 string            `json:"uid"`
	Name                string            `json:"name"`
	Format              string            `json:"format"`
	Rank                string            `json:"rank"`
	Online              bool              `json:"online"`
	Start               string            `json:"start"`
	Finish              string            `json:"finish"`
	Timezone            string            `json:"timezone"`
	Country             string            `json:"country"`
	City                string            `json:"city"`
	CityGeonameID       int               `json:"city_geoname_id"`
	LeagueUID           string            `json:"league_uid"`
	State               string            `json:"state"`
	OrganizersUIDs      []string          `json:"organizers_uids"`
	Venue               string            `json:"venue"`
	VenueURL            string            `json:"venue_url"`
	Address             string            `json:"address"`
	MapURL              string            `json:"map_url"`
	RegistrationURL     string            `json:"registration_url"`
	Proxies             bool              `json:"proxies"`
	Multideck           bool              `json:"multideck"`
	DecklistRequired    bool              `json:"decklist_required"`
	Description         string            `json:"description"`
	StandingsMode       string            `json:"standings_mode"`
	DecklistsMode       string            `json:"decklists_mode"`
	MaxRounds           int               `json:"max_rounds"`
	MaxPlayers          int               `json:"max_players"`
	OpenRounds          bool              `json:"open_rounds"`
	SelfOrganizedRounds bool              `json:"self_organized_rounds"`
	ExternalIDs         map[string]string `json:"external_ids"`
	EventCode           string            `json:"event_code"`
	Players             []ArchonPlayer    `json:"players"`
	Rounds              [][]ArchonTable   `json:"rounds"`
	Finals              *ArchonFinals     `json:"finals"`
	Winner              string            `json:"winner"`
	Standings           []ArchonStanding  `json:"standings"`
	ReportedPlayerCount int               `json:"reported_player_count"`
}

func (t *ArchonTournament) StartDate() string {
	if len(t.Start) >= 10 {
		return t.Start[:10]
	}
	return t.Start
}

func (t *ArchonTournament) PlayerCount() int {
	return max(len(t.Players), len(t.Standings), t.ReportedPlayerCount)
}

func (t *ArchonTournament) Place() string {
	if t.Online {
		return "Online"
	}
	if t.City != "" && t.Country != "" {
		return t.City + ", " + t.Country
	}
	return cmp.Or(t.City, t.Country)
}

type ArchonScore struct {
	GW int     `json:"gw"`
	VP float64 `json:"vp"`
	TP int     `json:"tp"`
}

type ArchonPlayer struct {
	UserUID      string      `json:"user_uid"`
	State        string      `json:"state"`
	Toss         int         `json:"toss"`
	Result       ArchonScore `json:"result"`
	Finalist     bool        `json:"finalist"`
	NonCompeting bool        `json:"non_competing"`
}

type ArchonSeat struct {
	PlayerUID string      `json:"player_uid"`
	Result    ArchonScore `json:"result"`
	JudgeUID  string      `json:"judge_uid"`
}

type ArchonTable struct {
	Seating []ArchonSeat `json:"seating"`
	State   string       `json:"state"`
}

type ArchonFinals struct {
	Seating   []ArchonSeat `json:"seating"`
	State     string       `json:"state"`
	SeedOrder []string     `json:"seed_order"`
}

type ArchonStanding struct {
	UserUID      string  `json:"user_uid"`
	GW           float64 `json:"gw"`
	VP           float64 `json:"vp"`
	TP           int     `json:"tp"`
	Toss         int     `json:"toss"`
	Finalist     bool    `json:"finalist"`
	Disqualified bool    `json:"disqualified"`
	NonCompeting bool    `json:"non_competing"`
	Withdrawn    bool    `json:"withdrawn"`
}

type ArchonDeck struct {
	UID           string            `json:"uid"`
	TournamentUID string            `json:"tournament_uid"`
	UserUID       string            `json:"user_uid"`
	Round         *int              `json:"round"`
	Name          string            `json:"name"`
	Comments      string            `json:"comments"`
	Cards         map[int]int       `json:"cards"`
	Attribution   ArchonAttribution `json:"attribution"`
	Winner        bool              `json:"winner"`
}

func (ad *ArchonDeck) DeckCards() []bundle.DeckCard {
	out := make([]bundle.DeckCard, 0, len(ad.Cards))
	for id, count := range ad.Cards {
		if count > 0 {
			out = append(out, bundle.DeckCard{ID: id, Count: count})
		}
	}
	return out
}

type ArchonAttribution struct {
	Kind   string `json:"kind"`
	VEKNID string `json:"vekn_id"`
}

type ArchonLeague struct {
	UID                  string             `json:"uid"`
	Name                 string             `json:"name"`
	Kind                 string             `json:"kind"`
	StandingsMode        string             `json:"standings_mode"`
	Format               string             `json:"format"`
	Country              string             `json:"country"`
	Start                string             `json:"start"`
	Finish               string             `json:"finish"`
	Description          string             `json:"description"`
	OrganizersUIDs       []string           `json:"organizers_uids"`
	ParentUID            string             `json:"parent_uid"`
	OpenToCountryPrinces bool               `json:"open_to_country_princes"`
	Final                *ArchonLeagueFinal `json:"final"`
}

type ArchonLeagueFinal struct {
	SeedOrder []string     `json:"seed_order"`
	Seating   []ArchonSeat `json:"seating"`
	Winner    string       `json:"winner"`
}

func ReadArchon(r io.Reader) (*Archon, error) {
	z, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	dec := json.NewDecoder(z)
	var head struct {
		Type        string `json:"type"`
		Version     int    `json:"version"`
		GeneratedAt string `json:"generated_at"`
	}
	if err := dec.Decode(&head); err != nil {
		return nil, err
	}
	if head.Type != "header" || head.Version != 2 || len(head.GeneratedAt) < len(time.DateOnly) {
		return nil, fmt.Errorf("no v2 export header")
	}
	a := &Archon{
		GeneratedAt:       head.GeneratedAt,
		UserByUID:         map[string]*ArchonUser{},
		UserByVEKN:        map[string]*ArchonUser{},
		TournamentByUID:   map[string]*ArchonTournament{},
		TournamentByCode:  map[string]*ArchonTournament{},
		TournamentByVEKN:  map[string]*ArchonTournament{},
		DeckByUID:         map[string]*ArchonDeck{},
		DecksByTournament: map[string][]*ArchonDeck{},
		DecksByUser:       map[string][]*ArchonDeck{},
		LeagueByUID:       map[string]*ArchonLeague{},
	}
	rows := 0
	for dec.More() {
		var line struct {
			Type  string          `json:"type"`
			Count int             `json:"count"`
			Data  json.RawMessage `json:"data"`
		}
		if err := dec.Decode(&line); err != nil {
			return nil, err
		}
		if line.Type == "eof" {
			if line.Count != rows || dec.More() {
				return nil, fmt.Errorf("export eof count %d after %d rows", line.Count, rows)
			}
			for i := range a.Users {
				u := &a.Users[i]
				a.UserByUID[u.UID] = u
				if u.VEKNID != "" {
					a.UserByVEKN[u.VEKNID] = u
				}
			}
			for i := range a.Tournaments {
				t := &a.Tournaments[i]
				a.TournamentByUID[t.UID] = t
				if t.EventCode != "" {
					a.TournamentByCode[strings.ToLower(t.EventCode)] = t
				}
				if vekn := t.ExternalIDs["vekn"]; vekn != "" {
					a.TournamentByVEKN[vekn] = t
				}
			}
			for i := range a.Decks {
				d := &a.Decks[i]
				a.DeckByUID[d.UID] = d
				if d.TournamentUID != "" {
					a.DecksByTournament[d.TournamentUID] = append(a.DecksByTournament[d.TournamentUID], d)
				}
				if d.UserUID != "" {
					a.DecksByUser[d.UserUID] = append(a.DecksByUser[d.UserUID], d)
				}
			}
			for i := range a.Leagues {
				l := &a.Leagues[i]
				a.LeagueByUID[l.UID] = l
			}
			return a, nil
		}
		rows++
		switch line.Type {
		case "user":
			var u ArchonUser
			err = json.Unmarshal(line.Data, &u)
			a.Users = append(a.Users, u)
		case "tournament":
			var t ArchonTournament
			err = json.Unmarshal(line.Data, &t)
			a.Tournaments = append(a.Tournaments, t)
		case "deck":
			var d ArchonDeck
			err = json.Unmarshal(line.Data, &d)
			a.Decks = append(a.Decks, d)
		case "league":
			var l ArchonLeague
			err = json.Unmarshal(line.Data, &l)
			a.Leagues = append(a.Leagues, l)
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("truncated export: no eof")
}

func (l *Live) SetArchon(a *Archon) { l.archon.Store(a) }
func (l *Live) Archon() *Archon     { return l.archon.Load() }

func MemberFrom(req *mcp.CallToolRequest) string {
	if req != nil && req.Extra != nil && req.Extra.TokenInfo != nil {
		return req.Extra.TokenInfo.UserID
	}
	return ""
}
