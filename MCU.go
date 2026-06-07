package main

import "fmt"

const NMAX int = 100

type Pasien struct {
	id   int
	nama string
	umur int
}

var dataPasien [NMAX]Pasien
var jumlahPasien int = 0

type Paket struct {
	id        int
	namaPaket string
	harga     int
}

var dataPaket [NMAX]Paket
var jumlahPaket int = 0

type Hasil struct {
	namaPasien string
	idPaket    int
	tanggal    string
	tekanan    int
	gula       int
	kolesterol int
	status     string
}

var dataHasil [NMAX]Hasil
var jumlahHasil int = 0

func InputDataPasien(dataPasien *[NMAX]Pasien, jumlahPasien *int) {
	var jumlahInput, i int

	// Meminta jumlah data yang ingin dimasukkan
	fmt.Print("Masukkan jumlah pasien yang ingin diinput: ")
	fmt.Scan(&jumlahInput)

	// Memberi tahu user format inputannya
	fmt.Println("\nMasukkan data dengan format: ID Nama Umur (pisahkan dengan spasi)")

	for i = 0; i < jumlahInput; i++ {
		// Cek apakah array belum penuh
		if *jumlahPasien < NMAX {
			fmt.Scan(&dataPasien[*jumlahPasien].id, &dataPasien[*jumlahPasien].nama, &dataPasien[*jumlahPasien].umur)

			*jumlahPasien++
		} else {
			fmt.Println("Peringatan, kapasitas data pasien sudah penuh! Sisa data tidak dapat dimasukkan.")
			break // Keluar dari loop jika array sudah mencapai NMAX
		}
	}
	fmt.Println("Proses penambahan data pasien selesai.")
}

func CetakDataPasien(dataPasien [NMAX]Pasien, jumlah int) {
	var i int
	for i = 0; i < jumlah; i++ {
		// Mengganti idx menjadi i, dan menghapus " " di akhir
		fmt.Printf("%2d. ID: %-3d | Nama: %-10s | Umur: %d\n", i+1, dataPasien[i].id, dataPasien[i].nama, dataPasien[i].umur)
	}
	fmt.Println()
}

func CariDataPasien(dataPasien [NMAX]Pasien, jumlah int, namaCari string) Pasien {
	var i int

	// Looping untuk mengecek setiap pasien satu per satu
	for i = 0; i < jumlah; i++ {
		// Jika nama di data sama dengan nama yang dicari
		if dataPasien[i].nama == namaCari {
			return dataPasien[i] // Kembalikan data pasien tersebut
		}
	}

	// Jika looping selesai tapi nama tidak ditemukan
	return Pasien{-1, "Tidak Ditemukan", -1}
}

func BinarySearchPasien(data [NMAX]Pasien, jumlah int, namaCari string) int {
	var kiri int = 0
	var kanan int = jumlah - 1
	var tengah int

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah].nama == namaCari {
			return tengah // Mengembalikan INDEKS/POSISI data jika ditemukan
		} else if data[tengah].nama < namaCari {
			kiri = tengah + 1 // Geser pencarian ke kanan
		} else {
			kanan = tengah - 1 // Geser pencarian ke kiri
		}
	}
	return -1 // Mengembalikan -1 jika nama tidak ditemukan
}

func InsertionSortAscendPasien(data *[NMAX]Pasien, jumlah int) {
	var i, j int
	for i = 1; i < jumlah; i++ {
		key := data[i]
		j = i - 1
		for j >= 0 && data[j].id > key.id {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func HapusDataPasien(dataPasien *[NMAX]Pasien, jumlahPasien *int) {
	var namaTarget string
	var i, j int
	var ketemu bool = false

	fmt.Print("\nMasukkan Nama Pasien yang ingin dihapus: ")
	fmt.Scan(&namaTarget)

	for i = 0; i < *jumlahPasien; i++ {
		if dataPasien[i].nama == namaTarget {
			ketemu = true

			// Geser semua data di kanannya ke arah kiri
			for j = i; j < *jumlahPasien-1; j++ {
				dataPasien[j] = dataPasien[j+1]
			}

			// Kurangi total jumlah pasien karena ada 1 yang dihapus
			*jumlahPasien--

			fmt.Println("Data pasien berhasil dihapus dari sistem!")
			break
		}
	}

	if ketemu == false {
		fmt.Println("Maaf, data pasien dengan Nama tersebut tidak ditemukan.")
	}
}

func InputDataCheckup(dataHasil *[NMAX]Hasil, jumlahHasil *int, dataPasien [NMAX]Pasien, jumlahPasien int) {
	var idTarget string
	var posisi Pasien

	fmt.Print("Masukkan nama Pasien yang akan di-checkup: ")
	fmt.Scan(&idTarget)

	posisi = CariDataPasien(dataPasien, jumlahPasien, idTarget)

	if posisi.id != -1 {
		// Basis Kasus / Kasus Berhenti: Jika ID ketemu, jalankan input data
		if *jumlahHasil < NMAX {
			dataHasil[*jumlahHasil].namaPasien = idTarget
			fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
			fmt.Scan(&dataHasil[*jumlahHasil].tanggal)
			fmt.Print("Masukkan Tekanan Darah: ")
			fmt.Scan(&dataHasil[*jumlahHasil].tekanan)
			fmt.Print("Masukkan Gula Darah: ")
			fmt.Scan(&dataHasil[*jumlahHasil].gula)
			fmt.Print("Masukkan Kolesterol: ")
			fmt.Scan(&dataHasil[*jumlahHasil].kolesterol)

			if dataHasil[*jumlahHasil].gula > 140 || dataHasil[*jumlahHasil].kolesterol > 200 {
				dataHasil[*jumlahHasil].status = "Perlu Tindakan"
			} else {
				dataHasil[*jumlahHasil].status = "Normal"
			}

			*jumlahHasil++
			fmt.Println("Hasil check-up berhasil disimpan.")
		} else {
			fmt.Println("Data hasil sudah penuh.")
		}
	} else {
		// REKURSIF: Jika tidak ketemu, panggil fungsi ini lagi dari awal
		fmt.Println("ID Pasien tidak ditemukan! Silakan coba lagi.")
		InputDataCheckup(dataHasil, jumlahHasil, dataPasien, jumlahPasien)
	}
}

func CetakHasilCheckup(dataHasil [NMAX]Hasil, jumlah int, idx int) {
	// Basis kasus: Berhenti jika indeks sudah mencapai jumlah data
	if idx >= jumlah {
		return
	}

	// Mencetak data secara sekuensial
	fmt.Printf("\n--- Hasil Check-up ke-%d ---\n", idx+1)
	fmt.Printf("Nama Pasien  : %s\n", dataHasil[idx].namaPasien)
	fmt.Printf("Tanggal    : %s\n", dataHasil[idx].tanggal)
	fmt.Printf("Tensi      : %d\n", dataHasil[idx].tekanan)
	fmt.Printf("Gula Darah : %d\n", dataHasil[idx].gula)
	fmt.Printf("Kolesterol : %d\n", dataHasil[idx].kolesterol)
	fmt.Printf("Status     : %s\n", dataHasil[idx].status)

	// Panggil fungsi ini lagi untuk indeks selanjutnya
	CetakHasilCheckup(dataHasil, jumlah, idx+1)
}

func main() {
	var menu int
	var jalan bool = true // Variabel penanda untuk perulangan
	var namaCari string
	var pasienDitemukan Pasien

	// Perulangan akan terus berjalan selama 'jalan' bernilai true
	for jalan {
		fmt.Println("\n=== SISTEM MEDICAL CHECK-UP ===")
		fmt.Println("1. Kelola Data Pasien")
		fmt.Println("2. Tambah Data Check-up")
		fmt.Println("3. Cari Data Check-up")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&menu)

		switch menu {
		case 1: // Anggap ini adalah pilihan "Kelola Data Pasien" di Menu Utama
			var subJalan bool = true
			var menuPasien int

			// Perulangan khusus untuk Sub-Menu
			for subJalan {
				fmt.Println("\n=== MENU KELOLA DATA PASIEN ===")
				fmt.Println("1. Tambah Data Pasien")
				fmt.Println("2. Lihat Daftar Pasien")
				fmt.Println("3. Cari Data Pasien")
				fmt.Println("4. Hapus Data Pasien")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilih sub-menu: ")
				fmt.Scan(&menuPasien)

				switch menuPasien {
				case 1:
					InputDataPasien(&dataPasien, &jumlahPasien)
				case 2:
					fmt.Println("\n--- DAFTAR PASIEN ---")
					if jumlahPasien == 0 {
						fmt.Println("Belum ada data pasien.")
					} else {
						InsertionSortAscendPasien(&dataPasien, jumlahPasien)
						CetakDataPasien(dataPasien, jumlahPasien)
					}
				case 3:
					// 1. Minta input nama yang mau dicari
					fmt.Print("Masukkan nama Pasien yang dicari: ")
					fmt.Scan(&namaCari)

					// 2. Simpan hasil kembalian fungsinya ke dalam variabel
					// Catatan: Angka 0 di belakang dihapus karena kita sudah memakai versi for loop
					pasienDitemukan = CariDataPasien(dataPasien, jumlahPasien, namaCari)

					// 3. Cek apakah ID-nya bukan -1 (artinya ketemu)
					if pasienDitemukan.id != -1 {
						fmt.Println("\n--- DATA DITEMUKAN ---")
						fmt.Printf("ID   : %d\n", pasienDitemukan.id)
						fmt.Printf("Nama : %s\n", pasienDitemukan.nama)
						fmt.Printf("Umur : %d\n", pasienDitemukan.umur)
					} else {
						fmt.Println("Maaf, data pasien dengan nama tersebut tidak ditemukan.")
					}
				case 4:
					HapusDataPasien(&dataPasien, &jumlahPasien)
				case 5:
					fmt.Println("Kembali ke Menu Utama...")
					subJalan = false // Mematikan perulangan sub-menu agar keluar ke main switch
				default:
					fmt.Println("Pilihan sub-menu tidak tersedia.")
				}
			}
		case 2:
			InputDataCheckup(&dataHasil, &jumlahHasil, dataPasien, jumlahPasien)
		case 3:
			// 1. Minta input nama Pasien
			fmt.Print("Masukkan nama Pasien untuk melihat riwayat check-up: ")
			fmt.Scan(&namaCari)

			// 2. URUTKAN DATA PASIEN (Wajib sebelum Binary Search)
			InsertionSortAscendPasien(&dataPasien, jumlahPasien)

			// 3. CARI PASIEN MENGGUNAKAN BINARY SEARCH
			var indeksPasien int = BinarySearchPasien(dataPasien, jumlahPasien, namaCari)

			// 4. Cek apakah hasil Binary Search menemukan pasien (bukan -1)
			if indeksPasien != -1 {
				// Ambil data pasien dari hasil indeks binary search
				pasienDitemukan = dataPasien[indeksPasien]
				fmt.Printf("\n--- RIWAYAT CHECK-UP PASIEN: %s (ID: %d) ---\n", pasienDitemukan.nama, pasienDitemukan.id)

				var adaRiwayat bool = false
				var i int = 0

				// 5. Tampilkan semua riwayat check-up dari dataHasil
				for i < jumlahHasil {
					if dataHasil[i].namaPasien == namaCari {
						fmt.Printf("- Tanggal: %s | Tensi: %d | Gula: %d | Kolesterol: %d | Status: %s\n",
							dataHasil[i].tanggal, dataHasil[i].tekanan, dataHasil[i].gula, dataHasil[i].kolesterol, dataHasil[i].status)
						adaRiwayat = true
					}
					i++
				}

				if !adaRiwayat {
					fmt.Println("Pasien ini belum memiliki riwayat pemeriksaan.")
				}
			} else {
				fmt.Println("Maaf, data pasien dengan ID tersebut tidak ditemukan.")
			}
		case 4:
			fmt.Println("Terima kasih!")
			jalan = false // Mengubah nilai menjadi false agar perulangan for berhenti
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
