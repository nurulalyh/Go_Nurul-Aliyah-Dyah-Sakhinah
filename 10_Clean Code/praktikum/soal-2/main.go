package main

import "fmt"

//Camel case untuk semua penamaan
type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

//Menukar posisi kedua method agar lebih mudah dibaca
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam = m.KecepatanPerJam + kecepatanBaru
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()

	//Menambahkan fmt.Printf agar dapat output terlihat
	fmt.Printf("Kecepatan mobil cepat: %d km/jam\n", mobilCepat.KecepatanPerJam)
	fmt.Printf("Kecepatan mobil lamban: %d km/jam\n", mobilLamban.KecepatanPerJam)
}
