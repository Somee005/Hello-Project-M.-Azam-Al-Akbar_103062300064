package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Pabrikan struct {
	Nama      string
	Mobil     []Mobil
	Penjualan int
}

type Mobil struct {
	Nama        string
	TahunKeluar int
	Pabrikan    string
	Penjualan   int
}

var pabrikanList = []Pabrikan{}

func tambahPabrikan(nama string) {
	pabrikanList = append(pabrikanList, Pabrikan{Nama: nama})
}

func editPabrikan(namaLama string, namaBaru string) {
	for i, p := range pabrikanList {
		if p.Nama == namaLama {
			pabrikanList[i].Nama = namaBaru
			break
		}
	}
}

func hapusPabrikan(nama string) {
	for i, p := range pabrikanList {
		if p.Nama == nama {
			pabrikanList = append(pabrikanList[:i], pabrikanList[i+1:]...)
			break
		}
	}
}

func tambahMobil(pabrikanNama, namaMobil string, tahunKeluar, penjualan int) {
	for i, p := range pabrikanList {
		if p.Nama == pabrikanNama {
			pabrikanList[i].Mobil = append(p.Mobil, Mobil{Nama: namaMobil, TahunKeluar: tahunKeluar, Pabrikan: pabrikanNama, Penjualan: penjualan})
			break
		}
	}
}

func editMobil(pabrikanNama, namaLama, namaBaru string, tahunKeluar, penjualan int) {
	for i, p := range pabrikanList {
		if p.Nama == pabrikanNama {
			for j, m := range p.Mobil {
				if m.Nama == namaLama {
					pabrikanList[i].Mobil[j] = Mobil{Nama: namaBaru, TahunKeluar: tahunKeluar, Pabrikan: pabrikanNama, Penjualan: penjualan}
					break
				}
			}
		}
	}
}

func hapusMobil(pabrikanNama, namaMobil string) {
	for i, p := range pabrikanList {
		if p.Nama == pabrikanNama {
			for j, m := range p.Mobil {
				if m.Nama == namaMobil {
					pabrikanList[i].Mobil = append(p.Mobil[:j], p.Mobil[j+1:]...)
					break
				}
			}
		}
	}
}

func cariMobilBerdasarkanPabrikan(namaPabrikan string) []Mobil {
	for _, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			return p.Mobil
		}
	}
	return nil
}

func cariMobil(namaMobil string) *Mobil {
	for _, p := range pabrikanList {
		for _, m := range p.Mobil {
			if m.Nama == namaMobil {
				return &m
			}
		}
	}
	return nil
}

func daftarPabrikanBerdasarkanJumlahProduk() []Pabrikan {
	sort.Slice(pabrikanList, func(i, j int) bool {
		return len(pabrikanList[i].Mobil) > len(pabrikanList[j].Mobil)
	})
	return pabrikanList
}

func daftarMobilTerurutBerdasarkanTahun() []Mobil {
	var semuaMobil []Mobil
	for _, p := range pabrikanList {
		semuaMobil = append(semuaMobil, p.Mobil...)
	}
	sort.Slice(semuaMobil, func(i, j int) bool {
		return semuaMobil[i].TahunKeluar < semuaMobil[j].TahunKeluar
	})
	return semuaMobil
}

func daftarMobilTerurutBerdasarkanNama() []Mobil {
	var semuaMobil []Mobil
	for _, p := range pabrikanList {
		semuaMobil = append(semuaMobil, p.Mobil...)
	}
	sort.Slice(semuaMobil, func(i, j int) bool {
		return semuaMobil[i].Nama < semuaMobil[j].Nama
	})
	return semuaMobil
}

func daftarMobilTerurutBerdasarkanPenjualan() []Mobil {
	var semuaMobil []Mobil
	for _, p := range pabrikanList {
		semuaMobil = append(semuaMobil, p.Mobil...)
	}
	sort.Slice(semuaMobil, func(i, j int) bool {
		return semuaMobil[i].Penjualan > semuaMobil[j].Penjualan
	})
	return semuaMobil
}

func daftarTop3PabrikanBerdasarkanPenjualan() []Pabrikan {
	sort.Slice(pabrikanList, func(i, j int) bool {
		return pabrikanList[i].Penjualan > pabrikanList[j].Penjualan
	})
	if len(pabrikanList) > 3 {
		return pabrikanList[:3]
	}
	return pabrikanList
}

func daftarTop3MobilBerdasarkanPenjualan() []Mobil {
	var semuaMobil []Mobil
	for _, p := range pabrikanList {
		semuaMobil = append(semuaMobil, p.Mobil...)
	}
	sort.Slice(semuaMobil, func(i, j int) bool {
		return semuaMobil[i].Penjualan > semuaMobil[j].Penjualan
	})
	if len(semuaMobil) > 3 {
		return semuaMobil[:3]
	}
	return semuaMobil
}

func tampilkanSemuaData() {
	fmt.Println("\nDaftar Semua Pabrikan dan Mobil:")
	for _, p := range pabrikanList {
		fmt.Printf("Pabrikan: %s\n", p.Nama)
		for _, m := range p.Mobil {
			fmt.Printf("\tMobil: %s, Tahun Keluar: %d, Penjualan: %d\n", m.Nama, m.TahunKeluar, m.Penjualan)
		}
	}
}

func editDataPabrikan() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Edit Data Pabrikan ===")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Edit Nama Pabrikan")
		fmt.Println("3. Hapus Pabrikan")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Masukkan pilihan Anda: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan nama pabrikan baru: ")
			scanner.Scan()
			namaPabrikan := scanner.Text()
			tambahPabrikan(namaPabrikan)
		case "2":
			fmt.Print("Masukkan nama pabrikan lama: ")
			scanner.Scan()
			namaLama := scanner.Text()

			fmt.Print("Masukkan nama pabrikan baru: ")
			scanner.Scan()
			namaBaru := scanner.Text()

			editPabrikan(namaLama, namaBaru)
		case "3":
			fmt.Print("Masukkan nama pabrikan yang akan dihapus: ")
			scanner.Scan()
			namaPabrikan := scanner.Text()
			hapusPabrikan(namaPabrikan)
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func editDataMobil() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Edit Data Mobil ===")
		fmt.Println("1. Tambah Mobil")
		fmt.Println("2. Edit Data Mobil")
		fmt.Println("3. Hapus Mobil")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Masukkan pilihan Anda: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan nama pabrikan: ")
			scanner.Scan()
			pabrikanNama := scanner.Text()

			fmt.Print("Masukkan nama mobil: ")
			scanner.Scan()
			namaMobil := scanner.Text()

			fmt.Print("Masukkan tahun keluar: ")
			scanner.Scan()
			tahunKeluar, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukkan penjualan: ")
			scanner.Scan()
			penjualan, _ := strconv.Atoi(scanner.Text())

			tambahMobil(pabrikanNama, namaMobil, tahunKeluar, penjualan)
		case "2":
			fmt.Print("Masukkan nama pabrikan: ")
			scanner.Scan()
			pabrikanNama := scanner.Text()

			fmt.Print("Masukkan nama mobil lama: ")
			scanner.Scan()
			namaLama := scanner.Text()

			fmt.Print("Masukkan nama mobil baru: ")
			scanner.Scan()
			namaBaru := scanner.Text()

			fmt.Print("Masukkan tahun keluar baru: ")
			scanner.Scan()
			tahunKeluarBaru, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukkan penjualan baru: ")
			scanner.Scan()
			penjualanBaru, _ := strconv.Atoi(scanner.Text())

			editMobil(pabrikanNama, namaLama, namaBaru, tahunKeluarBaru, penjualanBaru)
		case "3":
			fmt.Print("Masukkan nama pabrikan: ")
			scanner.Scan()
			pabrikanNama := scanner.Text()

			fmt.Print("Masukkan nama mobil yang akan dihapus: ")
			scanner.Scan()
			namaMobil := scanner.Text()

			hapusMobil(pabrikanNama, namaMobil)
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuCari() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Menu Cari ===")
		fmt.Println("1. Cari Mobil Berdasarkan Nama")
		fmt.Println("2. Cari Mobil Berdasarkan Pabrikan")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Masukkan pilihan Anda: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			fmt.Print("Masukkan nama mobil yang dicari: ")
			scanner.Scan()
			namaMobil := scanner.Text()
			mobil := cariMobil(namaMobil)
			if mobil != nil {
				fmt.Printf("Mobil: %s, Tahun Keluar: %d, Pabrikan: %s, Penjualan: %d\n", mobil.Nama, mobil.TahunKeluar, mobil.Pabrikan, mobil.Penjualan)
			} else {
				fmt.Println("Mobil tidak ditemukan.")
			}
		case "2":
			fmt.Print("Masukkan nama pabrikan: ")
			scanner.Scan()
			namaPabrikan := scanner.Text()
			mobilList := cariMobilBerdasarkanPabrikan(namaPabrikan)
			if len(mobilList) > 0 {
				fmt.Println("Daftar Mobil:")
				for _, m := range mobilList {
					fmt.Printf("\tMobil: %s, Tahun Keluar: %d, Penjualan: %d\n", m.Nama, m.TahunKeluar, m.Penjualan)
				}
			} else {
				fmt.Println("Tidak ada mobil yang ditemukan untuk pabrikan ini.")
			}
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func main() {

	tambahPabrikan("Toyota")
	tambahPabrikan("Honda")

	tambahMobil("Toyota", "Camry", 2020, 100)
	tambahMobil("Toyota", "Corolla", 2019, 150)
	tambahMobil("Honda", "Civic", 2021, 200)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Menu Pilihan ===")
		fmt.Println("1. Cari")
		fmt.Println("2. Edit Data Pabrikan")
		fmt.Println("3. Edit Data Mobil")
		fmt.Println("4. Daftar Pabrikan Berdasarkan Jumlah Produk")
		fmt.Println("5. Daftar Mobil Terurut Berdasarkan Tahun")
		fmt.Println("6. Daftar Mobil Terurut Berdasarkan Nama")
		fmt.Println("7. Top 3 Pabrikan Berdasarkan Penjualan")
		fmt.Println("8. Top 3 Mobil Berdasarkan Penjualan")
		fmt.Println("9. Tampilkan Semua Data")
		fmt.Println("10. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			menuCari()
		case "2":
			editDataPabrikan()
		case "3":
			editDataMobil()
		case "4":
			fmt.Println("Daftar Pabrikan Berdasarkan Jumlah Produk:")
			for _, p := range daftarPabrikanBerdasarkanJumlahProduk() {
				fmt.Printf("Pabrikan: %s, Jumlah Mobil: %d\n", p.Nama, len(p.Mobil))
			}
		case "5":
			fmt.Println("\nDaftar Mobil Terurut Berdasarkan Tahun:")
			for _, m := range daftarMobilTerurutBerdasarkanTahun() {
				fmt.Printf("Mobil: %s, Tahun: %d, Pabrikan: %s\n", m.Nama, m.TahunKeluar, m.Pabrikan)
			}
		case "6":
			fmt.Println("\nDaftar Mobil Terurut Berdasarkan Nama:")
			for _, m := range daftarMobilTerurutBerdasarkanNama() {
				fmt.Printf("Mobil: %s, Tahun: %d, Pabrikan: %s\n", m.Nama, m.TahunKeluar, m.Pabrikan)
			}
		case "7":
			fmt.Println("\nTop 3 Pabrikan Berdasarkan Penjualan:")
			for _, p := range daftarTop3PabrikanBerdasarkanPenjualan() {
				fmt.Printf("Pabrikan: %s, Penjualan: %d\n", p.Nama, p.Penjualan)
			}
		case "8":
			fmt.Println("\nTop 3 Mobil Berdasarkan Penjualan:")
			for _, m := range daftarTop3MobilBerdasarkanPenjualan() {
				fmt.Printf("Mobil: %s, Penjualan: %d, Pabrikan: %s\n", m.Nama, m.Penjualan, m.Pabrikan)
			}
		case "9":
			tampilkanSemuaData()
		case "10":
			fmt.Println("Terima kasih telah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// halo saya mengedit baris ini