package bundle

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
)

// Continents lists the valid continent filter values.
var Continents = []string{
	"Europe",
	"North America",
	"South America",
	"Asia",
	"Oceania",
	"Africa",
}

// Country is an ISO 3166-1 alpha-2 country and its continent.
type Country struct {
	Code      string
	Name      string
	Continent string
}

var countries = []Country{
	// Europe
	{"AD", "Andorra", "Europe"},
	{"AL", "Albania", "Europe"},
	{"AT", "Austria", "Europe"},
	{"AX", "Åland Islands", "Europe"},
	{"BA", "Bosnia and Herzegovina", "Europe"},
	{"BE", "Belgium", "Europe"},
	{"BG", "Bulgaria", "Europe"},
	{"BY", "Belarus", "Europe"},
	{"CH", "Switzerland", "Europe"},
	{"CY", "Cyprus", "Europe"},
	{"CZ", "Czechia", "Europe"},
	{"DE", "Germany", "Europe"},
	{"DK", "Denmark", "Europe"},
	{"EE", "Estonia", "Europe"},
	{"ES", "Spain", "Europe"},
	{"FI", "Finland", "Europe"},
	{"FO", "Faroe Islands", "Europe"},
	{"FR", "France", "Europe"},
	{"GB", "United Kingdom", "Europe"},
	{"GG", "Guernsey", "Europe"},
	{"GI", "Gibraltar", "Europe"},
	{"GR", "Greece", "Europe"},
	{"HR", "Croatia", "Europe"},
	{"HU", "Hungary", "Europe"},
	{"IE", "Ireland", "Europe"},
	{"IM", "Isle of Man", "Europe"},
	{"IS", "Iceland", "Europe"},
	{"IT", "Italy", "Europe"},
	{"JE", "Jersey", "Europe"},
	{"LI", "Liechtenstein", "Europe"},
	{"LT", "Lithuania", "Europe"},
	{"LU", "Luxembourg", "Europe"},
	{"LV", "Latvia", "Europe"},
	{"MC", "Monaco", "Europe"},
	{"MD", "Moldova", "Europe"},
	{"ME", "Montenegro", "Europe"},
	{"MK", "North Macedonia", "Europe"},
	{"MT", "Malta", "Europe"},
	{"NL", "Netherlands", "Europe"},
	{"NO", "Norway", "Europe"},
	{"PL", "Poland", "Europe"},
	{"PT", "Portugal", "Europe"},
	{"RO", "Romania", "Europe"},
	{"RS", "Serbia", "Europe"},
	{"RU", "Russia", "Europe"},
	{"SE", "Sweden", "Europe"},
	{"SI", "Slovenia", "Europe"},
	{"SJ", "Svalbard and Jan Mayen", "Europe"},
	{"SK", "Slovakia", "Europe"},
	{"SM", "San Marino", "Europe"},
	{"UA", "Ukraine", "Europe"},
	{"VA", "Vatican City", "Europe"},
	{"XK", "Kosovo", "Europe"},

	// North America
	{"AG", "Antigua and Barbuda", "North America"},
	{"AI", "Anguilla", "North America"},
	{"AW", "Aruba", "North America"},
	{"BB", "Barbados", "North America"},
	{"BL", "Saint Barthélemy", "North America"},
	{"BM", "Bermuda", "North America"},
	{"BQ", "Bonaire, Sint Eustatius and Saba", "North America"},
	{"BS", "Bahamas", "North America"},
	{"BZ", "Belize", "North America"},
	{"CA", "Canada", "North America"},
	{"CR", "Costa Rica", "North America"},
	{"CU", "Cuba", "North America"},
	{"CW", "Curaçao", "North America"},
	{"DM", "Dominica", "North America"},
	{"DO", "Dominican Republic", "North America"},
	{"GD", "Grenada", "North America"},
	{"GL", "Greenland", "North America"},
	{"GP", "Guadeloupe", "North America"},
	{"GT", "Guatemala", "North America"},
	{"HN", "Honduras", "North America"},
	{"HT", "Haiti", "North America"},
	{"JM", "Jamaica", "North America"},
	{"KN", "Saint Kitts and Nevis", "North America"},
	{"KY", "Cayman Islands", "North America"},
	{"LC", "Saint Lucia", "North America"},
	{"MF", "Saint Martin", "North America"},
	{"MQ", "Martinique", "North America"},
	{"MS", "Montserrat", "North America"},
	{"MX", "Mexico", "North America"},
	{"NI", "Nicaragua", "North America"},
	{"PA", "Panama", "North America"},
	{"PM", "Saint Pierre and Miquelon", "North America"},
	{"PR", "Puerto Rico", "North America"},
	{"SV", "El Salvador", "North America"},
	{"SX", "Sint Maarten", "North America"},
	{"TC", "Turks and Caicos Islands", "North America"},
	{"TT", "Trinidad and Tobago", "North America"},
	{"US", "United States", "North America"},
	{"VC", "Saint Vincent and the Grenadines", "North America"},
	{"VG", "British Virgin Islands", "North America"},
	{"VI", "U.S. Virgin Islands", "North America"},

	// South America
	{"AR", "Argentina", "South America"},
	{"BO", "Bolivia", "South America"},
	{"BR", "Brazil", "South America"},
	{"CL", "Chile", "South America"},
	{"CO", "Colombia", "South America"},
	{"EC", "Ecuador", "South America"},
	{"FK", "Falkland Islands", "South America"},
	{"GF", "French Guiana", "South America"},
	{"GY", "Guyana", "South America"},
	{"PE", "Peru", "South America"},
	{"PY", "Paraguay", "South America"},
	{"SR", "Suriname", "South America"},
	{"UY", "Uruguay", "South America"},
	{"VE", "Venezuela", "South America"},

	// Asia
	{"AE", "United Arab Emirates", "Asia"},
	{"AF", "Afghanistan", "Asia"},
	{"AM", "Armenia", "Asia"},
	{"AZ", "Azerbaijan", "Asia"},
	{"BD", "Bangladesh", "Asia"},
	{"BH", "Bahrain", "Asia"},
	{"BN", "Brunei", "Asia"},
	{"BT", "Bhutan", "Asia"},
	{"CC", "Cocos (Keeling) Islands", "Asia"},
	{"CN", "China", "Asia"},
	{"CX", "Christmas Island", "Asia"},
	{"GE", "Georgia", "Asia"},
	{"HK", "Hong Kong", "Asia"},
	{"ID", "Indonesia", "Asia"},
	{"IL", "Israel", "Asia"},
	{"IN", "India", "Asia"},
	{"IO", "British Indian Ocean Territory", "Asia"},
	{"IQ", "Iraq", "Asia"},
	{"IR", "Iran", "Asia"},
	{"JO", "Jordan", "Asia"},
	{"JP", "Japan", "Asia"},
	{"KG", "Kyrgyzstan", "Asia"},
	{"KH", "Cambodia", "Asia"},
	{"KP", "North Korea", "Asia"},
	{"KR", "South Korea", "Asia"},
	{"KW", "Kuwait", "Asia"},
	{"KZ", "Kazakhstan", "Asia"},
	{"LA", "Laos", "Asia"},
	{"LB", "Lebanon", "Asia"},
	{"LK", "Sri Lanka", "Asia"},
	{"MM", "Myanmar", "Asia"},
	{"MN", "Mongolia", "Asia"},
	{"MO", "Macao", "Asia"},
	{"MV", "Maldives", "Asia"},
	{"MY", "Malaysia", "Asia"},
	{"NP", "Nepal", "Asia"},
	{"OM", "Oman", "Asia"},
	{"PH", "Philippines", "Asia"},
	{"PK", "Pakistan", "Asia"},
	{"PS", "Palestine", "Asia"},
	{"QA", "Qatar", "Asia"},
	{"SA", "Saudi Arabia", "Asia"},
	{"SG", "Singapore", "Asia"},
	{"SY", "Syria", "Asia"},
	{"TH", "Thailand", "Asia"},
	{"TJ", "Tajikistan", "Asia"},
	{"TL", "Timor-Leste", "Asia"},
	{"TM", "Turkmenistan", "Asia"},
	{"TR", "Turkey", "Asia"},
	{"TW", "Taiwan", "Asia"},
	{"UZ", "Uzbekistan", "Asia"},
	{"VN", "Vietnam", "Asia"},
	{"YE", "Yemen", "Asia"},

	// Oceania
	{"AS", "American Samoa", "Oceania"},
	{"AU", "Australia", "Oceania"},
	{"CK", "Cook Islands", "Oceania"},
	{"FJ", "Fiji", "Oceania"},
	{"FM", "Micronesia", "Oceania"},
	{"GU", "Guam", "Oceania"},
	{"KI", "Kiribati", "Oceania"},
	{"MH", "Marshall Islands", "Oceania"},
	{"MP", "Northern Mariana Islands", "Oceania"},
	{"NC", "New Caledonia", "Oceania"},
	{"NF", "Norfolk Island", "Oceania"},
	{"NR", "Nauru", "Oceania"},
	{"NU", "Niue", "Oceania"},
	{"NZ", "New Zealand", "Oceania"},
	{"PF", "French Polynesia", "Oceania"},
	{"PG", "Papua New Guinea", "Oceania"},
	{"PN", "Pitcairn", "Oceania"},
	{"PW", "Palau", "Oceania"},
	{"SB", "Solomon Islands", "Oceania"},
	{"TK", "Tokelau", "Oceania"},
	{"TO", "Tonga", "Oceania"},
	{"TV", "Tuvalu", "Oceania"},
	{"UM", "United States Minor Outlying Islands", "Oceania"},
	{"VU", "Vanuatu", "Oceania"},
	{"WF", "Wallis and Futuna", "Oceania"},
	{"WS", "Samoa", "Oceania"},

	// Africa
	{"AO", "Angola", "Africa"},
	{"BF", "Burkina Faso", "Africa"},
	{"BI", "Burundi", "Africa"},
	{"BJ", "Benin", "Africa"},
	{"BW", "Botswana", "Africa"},
	{"CD", "Democratic Republic of the Congo", "Africa"},
	{"CF", "Central African Republic", "Africa"},
	{"CG", "Republic of the Congo", "Africa"},
	{"CI", "Côte d'Ivoire", "Africa"},
	{"CM", "Cameroon", "Africa"},
	{"CV", "Cabo Verde", "Africa"},
	{"DJ", "Djibouti", "Africa"},
	{"DZ", "Algeria", "Africa"},
	{"EG", "Egypt", "Africa"},
	{"EH", "Western Sahara", "Africa"},
	{"ER", "Eritrea", "Africa"},
	{"ET", "Ethiopia", "Africa"},
	{"GA", "Gabon", "Africa"},
	{"GH", "Ghana", "Africa"},
	{"GM", "Gambia", "Africa"},
	{"GN", "Guinea", "Africa"},
	{"GQ", "Equatorial Guinea", "Africa"},
	{"GW", "Guinea-Bissau", "Africa"},
	{"KE", "Kenya", "Africa"},
	{"KM", "Comoros", "Africa"},
	{"LR", "Liberia", "Africa"},
	{"LS", "Lesotho", "Africa"},
	{"LY", "Libya", "Africa"},
	{"MA", "Morocco", "Africa"},
	{"MG", "Madagascar", "Africa"},
	{"ML", "Mali", "Africa"},
	{"MR", "Mauritania", "Africa"},
	{"MU", "Mauritius", "Africa"},
	{"MW", "Malawi", "Africa"},
	{"MZ", "Mozambique", "Africa"},
	{"NA", "Namibia", "Africa"},
	{"NE", "Niger", "Africa"},
	{"NG", "Nigeria", "Africa"},
	{"RE", "Réunion", "Africa"},
	{"RW", "Rwanda", "Africa"},
	{"SC", "Seychelles", "Africa"},
	{"SD", "Sudan", "Africa"},
	{"SH", "Saint Helena", "Africa"},
	{"SL", "Sierra Leone", "Africa"},
	{"SN", "Senegal", "Africa"},
	{"SO", "Somalia", "Africa"},
	{"SS", "South Sudan", "Africa"},
	{"ST", "São Tomé and Príncipe", "Africa"},
	{"SZ", "Eswatini", "Africa"},
	{"TD", "Chad", "Africa"},
	{"TG", "Togo", "Africa"},
	{"TN", "Tunisia", "Africa"},
	{"TZ", "Tanzania", "Africa"},
	{"UG", "Uganda", "Africa"},
	{"YT", "Mayotte", "Africa"},
	{"ZA", "South Africa", "Africa"},
	{"ZM", "Zambia", "Africa"},
	{"ZW", "Zimbabwe", "Africa"},
}

var countryAliases = map[string]string{
	"usa":                                    "US",
	"u.s.a.":                                 "US",
	"u.s.":                                   "US",
	"united states of america":               "US",
	"uk":                                     "GB",
	"u.k.":                                   "GB",
	"great britain":                          "GB",
	"britain":                                "GB",
	"england":                                "GB",
	"scotland":                               "GB",
	"wales":                                  "GB",
	"northern ireland":                       "GB",
	"czech republic":                         "CZ",
	"the netherlands":                        "NL",
	"holland":                                "NL",
	"russian federation":                     "RU",
	"bosnia":                                 "BA",
	"macedonia":                              "MK",
	"moldova, republic of":                   "MD",
	"holy see":                               "VA",
	"korea, republic of":                     "KR",
	"republic of korea":                      "KR",
	"korea, democratic people's republic of": "KP",
	"taiwan, province of china":              "TW",
	"viet nam":                               "VN",
	"iran, islamic republic of":              "IR",
	"syrian arab republic":                   "SY",
	"lao people's democratic republic":       "LA",
	"brunei darussalam":                      "BN",
	"burma":                                  "MM",
	"macau":                                  "MO",
	"state of palestine":                     "PS",
	"east timor":                             "TL",
	"turkiye":                                "TR",
	"uae":                                    "AE",
	"venezuela, bolivarian republic of":      "VE",
	"bolivia, plurinational state of":        "BO",
	"tanzania, united republic of":           "TZ",
	"ivory coast":                            "CI",
	"dr congo":                               "CD",
	"drc":                                    "CD",
	"congo":                                  "CG",
	"swaziland":                              "SZ",
	"cape verde":                             "CV",
}

var (
	countryByCode   map[string]Country
	countryByName   map[string]string
	continentByFold map[string]string
)

func init() {
	countryByCode = make(map[string]Country, len(countries))
	countryByName = make(map[string]string, len(countries)+len(countryAliases))
	for _, c := range countries {
		countryByCode[c.Code] = c
		countryByName[Fold(c.Name)] = c.Code
	}
	for alias, code := range countryAliases {
		countryByName[Fold(alias)] = code
	}
	continentByFold = make(map[string]string, len(Continents))
	for _, c := range Continents {
		continentByFold[Fold(c)] = c
	}
}

// LookupCountry resolves an ISO 3166-1 alpha-2 code or a country name/alias.
func LookupCountry(q string) (Country, bool) {
	s := strings.TrimSpace(q)
	if s == "" {
		return Country{}, false
	}
	if len(s) == 2 {
		if c, ok := countryByCode[strings.ToUpper(s)]; ok {
			return c, true
		}
	}
	if code, ok := countryByName[Fold(s)]; ok {
		return countryByCode[code], true
	}
	return Country{}, false
}

// LookupContinent resolves a continent name to its canonical form.
func LookupContinent(q string) (string, bool) {
	c, ok := continentByFold[Fold(q)]
	return c, ok
}

// CountryContinent returns the continent for an ISO 3166-1 alpha-2 code, or "".
func CountryContinent(code string) string {
	return countryByCode[strings.ToUpper(strings.TrimSpace(code))].Continent
}

// PlaceCountry extracts the ISO 3166-1 alpha-2 country code from a TWDA Place.
// Online events return ("", true); unmapped places return ("", false).
// Two-letter ISO codes are never matched so state/province codes (CA, DE, RJ)
// cannot misidentify a country.
func PlaceCountry(place string) (code string, online bool) {
	trimmed := strings.TrimSpace(place)
	if trimmed == "" {
		return "", false
	}
	if Fold(trimmed) == "online" {
		return "", true
	}
	last := trimmed
	if i := strings.LastIndex(trimmed, ","); i >= 0 {
		last = strings.TrimSpace(trimmed[i+1:])
	}
	folded := Fold(last)
	if c, ok := countryByName[folded]; ok {
		return c, false
	}
	words := strings.Fields(folded)
	for k := min(len(words), 4); k >= 1; k-- {
		if c, ok := countryByName[strings.Join(words[len(words)-k:], " ")]; ok {
			return c, false
		}
	}
	return "", false
}

func (b *Bundle) unmappedPlaces() []string {
	type count struct {
		n     int
		first string
	}
	seen := map[string]*count{}
	for _, d := range b.Decks {
		if d.Date < b.Classification.Since {
			continue
		}
		code, online := PlaceCountry(d.Place)
		if online || code != "" {
			continue
		}
		c := seen[d.Place]
		if c == nil {
			c = &count{first: d.ID}
			seen[d.Place] = c
		}
		c.n++
	}
	if len(seen) == 0 {
		return nil
	}
	places := slices.SortedFunc(maps.Keys(seen), func(a, b string) int {
		return cmp.Or(seen[b].n-seen[a].n, strings.Compare(a, b))
	})
	out := make([]string, len(places))
	for i, p := range places {
		c := seen[p]
		out[i] = fmt.Sprintf("%q (%d decks, e.g. %s)", p, c.n, c.first)
	}
	return out
}
