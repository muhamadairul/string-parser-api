package parser

import "strings"

// Parse extracts Name, Age, and City from an input string.
//
// Rules (ketat):
//   - Parsing kanan ke kiri (right-to-left)
//   - No regex, no replace
//   - Max 5 variabel aktif: i, state, city, age, name
//   - Handle age suffix: TAHUN, THN, TH (attached atau dengan spasi)
//
// Contoh:
//
//	"CUT MINI 28 BANDA ACEH"     → name="CUT MINI", age="28", city="BANDA ACEH"
//	"BUDI 35THN SURABAYA"        → name="BUDI", age="35", city="SURABAYA"
//	"SITI 22TH JAKARTA"          → name="SITI", age="22", city="JAKARTA"
//	"AHMAD 40 TAHUN MAKASSAR"    → name="AHMAD", age="40", city="MAKASSAR"
func Parse(input string) (name, age, city string) {
	// Pre-process: normalkan "angka TAHUN/THN/TH" → "angkaTAHUN/THN/TH"
	// supaya loop utama tidak perlu menangani spasi antara angka dan suffix.
	// Dilakukan dengan satu loop forward sederhana — tidak menggunakan regex/replace.
	input = normalizeSuffix(input)

	// --- LOOP UTAMA: kanan ke kiri, tepat 5 variabel ---
	i := len(input) - 1 // var 1
	state := "city"     // var 2
	city = ""           // var 3
	age = ""            // var 4
	name = ""           // var 5

	for i >= 0 {
		c := input[i]

		if state == "city" {
			if c >= '0' && c <= '9' {
				// Ketemu digit pertama dari kanan → mulai state age
				state = "age"
				// Jangan i-- dulu, biarkan char ini diproses ulang di iterasi berikut
				// Tapi kita sudah switch state, jadi langsung fall-through ke age handler
				age = string(c) + age
			} else {
				city = string(c) + city
			}
		} else if state == "age" {
			if c == ' ' {
				// Spasi setelah angka age → selesai age, mulai name
				state = "name"
			} else if c >= '0' && c <= '9' {
				age = string(c) + age
			}
			// huruf (T, H, N, A, U) di-skip otomatis — tidak ada else branch
		} else {
			// state == "name": ambil semua sisa karakter
			name = string(c) + name
		}

		i--
	}

	city = strings.TrimSpace(city)
	name = strings.TrimSpace(name)
	return
}

// normalizeSuffix menggabungkan format "angka TAHUN/THN/TH" → "angkaTAHUN/THN/TH"
// supaya loop utama Parse() tidak perlu menangani spasi antara angka dan suffix.
//
// Cara kerja: scan forward, bila menemukan urutan "<digit(s)> TAHUN|THN|TH <spasi|end>",
// buang spasi di antara angka dan suffix.
// Tidak menggunakan regex atau replace.
func normalizeSuffix(s string) string {
	suffixes := []string{"TAHUN", "THN", "TH"}
	upper := strings.ToUpper(s)

	for _, sfx := range suffixes {
		sfxLen := len(sfx)
		sLen := len(upper)
		for j := 0; j < sLen-sfxLen; j++ {
			// Cek apakah posisi j adalah awal dari suffix
			if upper[j:j+sfxLen] != sfx {
				continue
			}
			// Cek apakah suffix diakhiri spasi atau end-of-string
			afterSfx := j + sfxLen
			if afterSfx < sLen && upper[afterSfx] != ' ' {
				continue
			}
			// Cek apakah sebelum suffix ada spasi + digit
			// Cari ke kiri: skip spasi, lalu harus ada digit
			k := j - 1
			if k < 0 || upper[k] != ' ' {
				continue
			}
			// Ada spasi sebelum suffix — cek apakah sebelum spasi itu ada digit
			k--
			if k < 0 || upper[k] < '0' || upper[k] > '9' {
				continue
			}
			// Kondisi terpenuhi: hapus spasi di posisi j-1
			s = s[:j-1] + s[j:]
			upper = upper[:j-1] + upper[j:]
			break
		}
	}
	return s
}
