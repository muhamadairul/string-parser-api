package parser

import "strings"

// Capitals maps Indonesian provincial capital cities to their province names.
var Capitals = map[string]string{
	"BANDA ACEH":     "ACEH",
	"MEDAN":          "SUMATERA UTARA",
	"PADANG":         "SUMATERA BARAT",
	"PEKANBARU":      "RIAU",
	"JAMBI":          "JAMBI",
	"PALEMBANG":      "SUMATERA SELATAN",
	"BENGKULU":       "BENGKULU",
	"BANDAR LAMPUNG": "LAMPUNG",
	"JAKARTA":        "DKI JAKARTA",
	"SERANG":         "BANTEN",
	"BANDUNG":        "JAWA BARAT",
	"SEMARANG":       "JAWA TENGAH",
	"YOGYAKARTA":     "DIY",
	"SURABAYA":       "JAWA TIMUR",
	"DENPASAR":       "BALI",
	"MATARAM":        "NTB",
	"KUPANG":         "NTT",
	"PONTIANAK":      "KALIMANTAN BARAT",
	"PALANGKARAYA":   "KALIMANTAN TENGAH",
	"BANJARMASIN":    "KALIMANTAN SELATAN",
	"SAMARINDA":      "KALIMANTAN TIMUR",
	"TANJUNG SELOR":  "KALIMANTAN UTARA",
	"MANADO":         "SULAWESI UTARA",
	"PALU":           "SULAWESI TENGAH",
	"MAKASSAR":       "SULAWESI SELATAN",
	"KENDARI":        "SULAWESI TENGGARA",
	"GORONTALO":      "GORONTALO",
	"MAMUJU":         "SULAWESI BARAT",
	"AMBON":          "MALUKU",
	"SOFIFI":         "MALUKU UTARA",
	"JAYAPURA":       "PAPUA",
}

// EnrichCity appends province name if city is a provincial capital.
func EnrichCity(city string) string {
	key := strings.ToUpper(strings.TrimSpace(city))
	if prov, ok := Capitals[key]; ok {
		return key + " " + prov
	}
	return strings.ToUpper(key)
}
