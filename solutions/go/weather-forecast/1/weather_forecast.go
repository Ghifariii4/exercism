// Package weather menyediakan fungsi dan variabel untuk memprakirakan
// kondisi cuaca di berbagai kota di Goblinocus.
package weather

var (
// CurrentCondition menyimpan informasi tentang kondisi cuaca terkini.
	CurrentCondition string
// CurrentLocation menyimpan informasi tentang nama kota atau lokasi saat ini.
	CurrentLocation  string
)

// Forecast menerima nama kota dan kondisi cuaca sebagai argumen,
// lalu mengembalikan string teks ringkasan prakiraan cuaca untuk lokasi tersebut.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
