package main

//komen untuk tugas

import (
	"fmt"
	"sort"
	"strings"
)

type Calon struct {
	Nama   string
	Partai string
	Suara  int
}

type Pemilih struct {
	Nama string
}

var calons []Calon
var pemilihs []Pemilih
var votingOpen bool
var threshold int

func main() {
	// Example threshold for candidate selection
	threshold = 5
	menu()
}

func menu() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Calon")
		fmt.Println("2. Edit Calon")
		fmt.Println("3. Hapus Calon")
		fmt.Println("4. Tambah Pemilih")
		fmt.Println("5. Hapus Pemilih")
		fmt.Println("6. Mulai Pemilihan")
		fmt.Println("7. Tutup Pemilihan")
		fmt.Println("8. Pilih Calon")
		fmt.Println("9. Lihat Hasil Pemilihan")
		fmt.Println("10. Cari Data Calon")
		fmt.Println("99. Keluar")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tambahCalon()
		case 2:
			editCalon()
		case 3:
			hapusCalon()
		case 4:
			tambahPemilih()
		case 5:
			hapusPemilih()
		case 6:
			mulaiPemilihan()
		case 7:
			tutupPemilihan()
		case 8:
			pilihCalon()
		case 9:
			lihatHasilPemilihan()
		case 10:
			cariDataCalon()
		case 99:
			return
		default:
			fmt.Println("Opsi tidak valid.")
		}
	}
}

func tambahCalon() {
	var nama, partai string
	fmt.Print("Masukkan nama calon: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan partai calon: ")
	fmt.Scan(&partai)
	calons = append(calons, Calon{Nama: nama, Partai: partai, Suara: 0})
	fmt.Println("Calon berhasil ditambahkan.")
}

func editCalon() {
	var nama string
	fmt.Print("Masukkan nama calon yang ingin diubah: ")
	fmt.Scan(&nama)
	for i, calon := range calons {
		if calon.Nama == nama {
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&calons[i].Nama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&calons[i].Partai)
			fmt.Println("Calon berhasil diubah.")
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

func hapusCalon() {
	var nama string
	fmt.Print("Masukkan nama calon yang ingin dihapus: ")
	fmt.Scan(&nama)
	for i, calon := range calons {
		if calon.Nama == nama {
			calons = append(calons[:i], calons[i+1:]...)
			fmt.Println("Calon berhasil dihapus.")
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

func tambahPemilih() {
	var nama string
	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&nama)
	pemilihs = append(pemilihs, Pemilih{Nama: nama})
	fmt.Println("Pemilih berhasil ditambahkan.")
}

func hapusPemilih() {
	var nama string
	fmt.Print("Masukkan nama pemilih yang ingin dihapus: ")
	fmt.Scan(&nama)
	for i, pemilih := range pemilihs {
		if pemilih.Nama == nama {
			pemilihs = append(pemilihs[:i], pemilihs[i+1:]...)
			fmt.Println("Pemilih berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pemilih tidak ditemukan.")
}

func mulaiPemilihan() {
	votingOpen = true
	fmt.Println("Pemilihan telah dibuka.")
}

func tutupPemilihan() {
	votingOpen = false
	fmt.Println("Pemilihan telah ditutup.")
}

var suaraCalon = make(map[string][]string)

func pilihCalon() {
	if !votingOpen {
		fmt.Println("Pemilihan tidak sedang berlangsung.")
		return
	}
	var namaPemilih, namaCalon string
	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&namaPemilih)
	if !isPemilihValid(namaPemilih) {
		fmt.Println("Pemilih tidak ditemukan.")
		return
	}
	if sudahMemilih(namaPemilih) {
		fmt.Println("Pemilih sudah memberikan suara sebelumnya.")
		return
	}
	fmt.Print("Masukkan nama calon yang ingin dipilih: ")
	fmt.Scan(&namaCalon)
	for i, calon := range calons {
		if calon.Nama == namaCalon {
			calons[i].Suara++
			suaraCalon[calon.Nama] = append(suaraCalon[calon.Nama], namaPemilih)
			fmt.Println("Suara berhasil diberikan.")
			tandaiSudahMemilih(namaPemilih)
			return
		}
	}
	fmt.Println("Calon tidak ditemukan.")
}

var pemilihSudahMemilih = make(map[string]bool)

func sudahMemilih(nama string) bool {
	return pemilihSudahMemilih[nama]
}

func tandaiSudahMemilih(nama string) {
	pemilihSudahMemilih[nama] = true
}

func isPemilihValid(nama string) bool {
	for _, pemilih := range pemilihs {
		if pemilih.Nama == nama {
			return true
		}
	}
	return false
}

func lihatHasilPemilihan() {
	sort.Slice(calons, func(i, j int) bool {
		return calons[i].Suara > calons[j].Suara
	})
	fmt.Println("Hasil Pemilihan:")
	for _, calon := range calons {
		status := "Tidak Terpilih"
		if calon.Suara >= threshold {
			status = "Terpilih"
		}
		fmt.Printf("Nama: %s, Partai: %s, Suara: %d, Status: %s\n", calon.Nama, calon.Partai, calon.Suara, status)
	}
}

func cariDataCalon() {
	var partai, nama string
	fmt.Print("Masukkan nama partai (tekan Enter jika tidak ingin mencari berdasarkan partai): ")
	fmt.Scan(&partai)
	fmt.Print("Masukkan nama calon (tekan Enter jika tidak ingin mencari berdasarkan nama): ")
	fmt.Scan(&nama)

	found := false
	for _, calon := range calons {
		if (partai == "" || strings.Contains(calon.Partai, partai)) && (nama == "" || strings.Contains(calon.Nama, nama)) {
			fmt.Printf("Nama: %s, Partai: %s, Suara: %d\n", calon.Nama, calon.Partai, calon.Suara)
			if len(suaraCalon[calon.Nama]) > 0 {
				fmt.Println("Pemilih yang memilih:")
				for _, pemilih := range suaraCalon[calon.Nama] {
					fmt.Println(pemilih)
				}
			} else {
				fmt.Println("Tidak ada pemilih yang memilih calon ini.")
			}
			found = true
		}
	}

	if !found {
		fmt.Println("Calon tidak ditemukan.")
	}
}
