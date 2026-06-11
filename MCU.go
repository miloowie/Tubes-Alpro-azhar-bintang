package main

import "fmt"

const NMAX int = 100

type Pasien struct {
	id    int
	nama  string
	umur  int
	paket string
}

var dataPasien [NMAX]Pasien
var jumlahPasien int = 0

type LaporanPaket struct {
	nama  string
	count int
	harga int
	total int
}

type Paket struct {
	id        int
	namaPaket string
	harga     int
}

var dataPaket [NMAX]Paket
var jumlahPaket int = 0

type Hasil struct {
	namaPasien   string
	paket        string
	tanggal      string
	status       string
	tekanan      int
	nadi         int
	suhu         float64
	gula         int
	kolesterol   int
	hdl          int
	ldl          int
	trigliserida int
	asamUrat     float64
	ekg          string
	hba1c        float64
	usgPerut     string
	kondisiMata  string
}

var dataHasil [NMAX]Hasil
var jumlahHasil int = 0

func InputDataPasien(dataPasien *[NMAX]Pasien, jumlahPasien *int) {
	var jumlahInput, i int
	var isPenuh bool
	isPenuh = false

	fmt.Print("Masukkan jumlah pasien yang ingin diinput: ")
	fmt.Scan(&jumlahInput)

	fmt.Println("\nMasukkan data dengan format: ID Nama Umur (pisahkan dengan spasi)")
	fmt.Println("Contoh: 1001 Budi 67")
	for i = 0; i < jumlahInput && !isPenuh; i++ {
		if *jumlahPasien < NMAX {
			fmt.Scan(&dataPasien[*jumlahPasien].id, &dataPasien[*jumlahPasien].nama, &dataPasien[*jumlahPasien].umur)
			*jumlahPasien++
		} else {
			fmt.Println("Peringatan, kapasitas data pasien sudah penuh! Sisa data tidak dapat dimasukkan.")
			isPenuh = true
		}
	}
	fmt.Println("Data berhasil ditambahkan.")
}

func CetakDataPasien(dataPasien [NMAX]Pasien, jumlah int) {
	var i int
	for i = 0; i < jumlah; i++ {
		fmt.Printf("%2d. ID: %-3d | Nama: %-10s | Umur: %d\n", i+1, dataPasien[i].id, dataPasien[i].nama, dataPasien[i].umur)
	}
	fmt.Println()
}

func CariDataPasienSeq(dataPasien [NMAX]Pasien, jumlah int, namaCari string) Pasien {
	var i, ketemu int
	var hasil Pasien

	ketemu = 0
	hasil = Pasien{-1, "Tidak Ditemukan", -1, ""}

	for i = 0; i < jumlah && ketemu == 0; i++ {
		if dataPasien[i].nama == namaCari {
			hasil = dataPasien[i]
			ketemu = 1
		}
	}
	return hasil
}

func BinarySearchPasien(data [NMAX]Pasien, jumlah int, namaCari string) int {
	var kiri int = 0
	var kanan int = jumlah - 1
	var tengah int

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah].nama == namaCari {
			return tengah
		} else if data[tengah].nama < namaCari {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func InsertionSortAscendPasienByID(data *[NMAX]Pasien, jumlah int) {
	var i, j int
	var key Pasien

	for i = 1; i < jumlah; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].id > key.id {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func InsertionSortDescendPasienByNama(data *[NMAX]Pasien, jumlah int) {
	var i, j int
	var key Pasien
	for i = 1; i < jumlah; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].nama < key.nama {
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
	for i = 0; i < *jumlahPasien && !ketemu; i++ {
		if dataPasien[i].nama == namaTarget {
			ketemu = true

			for j = i; j < *jumlahPasien-1; j++ {
				dataPasien[j] = dataPasien[j+1]
			}
			*jumlahPasien--
			fmt.Println("Data pasien berhasil dihapus dari sistem!")
		}
	}

	if ketemu == false {
		fmt.Println("Maaf, data pasien dengan Nama tersebut tidak ditemukan.")
	}
}

func TampilkanDaftarPaket() {
	fmt.Print(`
================================================================================
                         DAFTAR PAKET MEDICAL CHECK-UP                          
================================================================================

1. Paket EPIC (Basic & Essential Check) - Rp 500.000
   "Pemeriksaan dasar, biar gak beban tim dan gak nyangkut di badak selamanya."
   
   Paket ini ditujukan untuk pemeriksaan kesehatan dasar (skrining awal) guna 
   memastikan fungsi organ vital dasar berjalan dengan baik. Cocok untuk usia 
   muda atau yang baru pertama kali MCU.
   
   Isi Pemeriksaan:
   • Pemeriksaan Fisik & Tanda Vital: 
     Tekanan darah, denyut nadi, indeks massa tubuh (IMB), dan suhu tubuh.
   • Laboratorium Dasar:
     - Darah Rutin (Hematologi): Cek hemoglobin, leukosit, trombosit (skrining anemia/infeksi).
     - Urine Rutin (Urinalisis): Cek fungsi ginjal dasar & mendeteksi infeksi saluran kemih.
     - Profil Gula Darah: Gula Darah Puasa (skrining awal diabetes).
     - Profil Lemak Dasar: Kolesterol Total.
   • Konsultasi Dokter: Evaluasi hasil pemeriksaan dokter umum.

--------------------------------------------------------------------------------

2. Paket LEGEND (Intermediate & Standard Check) - Rp 1.000.000
   "Satu langkah menuju top global. Butuh mekanik mumpuni dan organ tubuh yang responsif."
   
   Paket ini lebih lengkap dan mendalam. Selain memeriksa organ dasar, paket ini 
   mulai memantau fungsi organ dalam seperti hati, ginjal secara spesifik, dan 
   profil lemak yang lebih detail. Cocok untuk mereka yang memiliki aktivitas 
   tinggi atau sering begadang.
   
   Isi Pemeriksaan:
   • Semua isi di Paket EPIC ditambah dengan:
   • Profil Lemak Lengkap: 
     Kolesterol Total, HDL, LDL, dan Trigliserisda (untuk memantau risiko 
     stroke/jantung akibat keseringan makan gorengan saat mabar).
   • Fungsi Ginjal Lengkap: Ureum, Kreatinin, dan Asam Urat.
   • Fungsi Hati (Liver): SGOT dan SGPT (penting untuk yang sering begadang push rank).
   • Pemeriksaan Penunjang:
     - EKG (Elektrokardiografi): Rekam jantung dasar untuk melihat kelistrikan jantung.
     - Rontgen Dada (Thorax): Melihat kondisi paru-paru dan ukuran jantung.
   • Konsultasi Dokter: Dokter Umum + Konsultasi Gizi Dasar.

--------------------------------------------------------------------------------

3. Paket MYTHIC (Advanced & Comprehensive Check) - Rp 2.000.000
   "Kesehatan Glory, Mekanik Sempurna. Perlindungan total luar dan dalam."
   
   Ini adalah paket paling premium dan komprehensif. Pemeriksaannya menyeluruh 
   (eksekutif) untuk mendeteksi dini penyakit kronis, memeriksa fungsi jantung 
   saat bekerja berat, serta fungsi metabolisme tubuh secara total.
   
   Isi Pemeriksaan:
   • Semua isi di Paket LEGEND ditambah dengan:
   • Skrining Diabetes Lanjutan: HbA1c (melihat rata-rata gula darah 3 bulan terakhir).
   • Fungsi Hati Lengkap tambahan: Bilirubin Total, Protein Total, Albumen, Globulin.
   • Pemeriksaan Penunjang Advanced:
     - Treadmill Test: Menguji ketahanan jantung saat melakukan aktivitas fisik berat.
     - USG Abdomen (Perut): Melihat kondisi hati, empedu, pankreas, limpa, dan ginjal 
       secara visual (mendeteksi fatty liver atau batu ginjal).
     - Pemeriksaan Mata (Visus & Tonometri): Skrining kesehatan mata (penting banget 
       buat gamer yang matanya sering terpapar blue light).
   • Konsultasi Dokter: Dokter Spesialis Penyakit Dalam (Sp.PD).

================================================================================
`)
}

func PilihPaket(dataPasien *[NMAX]Pasien, jumlahPasien int) {
	var nama string
	var pilihan int
	var i int
	var ketemu bool

	fmt.Print("\nMasukkan nama pasien yang akan memilih paket: ")
	fmt.Scan(&nama)

	ketemu = false
	i = 0

	for i < jumlahPasien && !ketemu {
		if dataPasien[i].nama == nama {
			ketemu = true
			TampilkanDaftarPaket()

			fmt.Printf("\n--- PILIH PAKET UNTUK PASIEN: %s ---\n", dataPasien[i].nama)
			fmt.Print("Masukkan pilihan paket Anda (1/2/3): ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				dataPasien[i].paket = "1"
				fmt.Printf("=> Pendaftaran berhasil! Pasien %s didaftarkan ke Paket EPIC.\n", dataPasien[i].nama)
			case 2:
				dataPasien[i].paket = "2"
				fmt.Printf("=> Pendaftaran berhasil! Pasien %s didaftarkan ke Paket LEGEND.\n", dataPasien[i].nama)
			case 3:
				dataPasien[i].paket = "3"
				fmt.Printf("=> Pendaftaran berhasil! Pasien %s didaftarkan ke Paket MYTHIC.\n", dataPasien[i].nama)
			default:
				fmt.Println("Pilihan paket tidak tersedia. Silakan masukkan angka 1, 2, atau 3.")
			}
		}
		i++
	}
	if !ketemu {
		fmt.Println("Maaf, data pasien tidak ditemukan. Silakan daftarkan di Menu Kelola Data Pasien terlebih dahulu.")
	}
}

func InputDataCheckup(dataHasil *[NMAX]Hasil, jumlahHasil *int, dataPasien [NMAX]Pasien, jumlahPasien int) {
	var namaTarget string
	var i int
	var ketemu bool
	var pasienDipilih Pasien
	var namaPaket string

	fmt.Print("\nMasukkan nama Pasien yang akan di-checkup: ")
	fmt.Scan(&namaTarget)

	ketemu = false
	i = 0

	for i < jumlahPasien && !ketemu {
		if dataPasien[i].nama == namaTarget {
			ketemu = true
			pasienDipilih = dataPasien[i]
		}
		i++
	}

	if ketemu {
		if pasienDipilih.paket == "" {
			fmt.Println("\n[!] Peringatan: Pasien ini belum memilih paket MCU!")
			fmt.Println("Silakan pilih paket terlebih dahulu di Menu Kelola Data Check-up (Opsi 1).")
			return
		}

		if *jumlahHasil < NMAX {
			dataHasil[*jumlahHasil].namaPasien = pasienDipilih.nama
			dataHasil[*jumlahHasil].paket = pasienDipilih.paket

			if pasienDipilih.paket == "1" {
				namaPaket = "EPIC"
			} else if pasienDipilih.paket == "2" {
				namaPaket = "LEGEND"
			} else if pasienDipilih.paket == "3" {
				namaPaket = "MYTHIC"
			}

			fmt.Printf("\n==================================================")
			fmt.Printf("\n  INPUT HASIL MCU - PASIEN: %s [Paket %s]  ", pasienDipilih.nama, namaPaket)
			fmt.Printf("\n==================================================\n")

			fmt.Print("Masukkan Tanggal Periksa (DD-MM-YYYY): ")
			fmt.Scan(&dataHasil[*jumlahHasil].tanggal)

			if pasienDipilih.paket == "1" || pasienDipilih.paket == "2" || pasienDipilih.paket == "3" {
				fmt.Println("\n[ Form Pemeriksaan Fisik & Lab Dasar - EPIC ]")
				fmt.Print("-> Tekanan Darah (Tensi) : ")
				fmt.Scan(&dataHasil[*jumlahHasil].tekanan)
				fmt.Print("-> Denyut Nadi / menit   : ")
				fmt.Scan(&dataHasil[*jumlahHasil].nadi)
				fmt.Print("-> Suhu Tubuh (Celsius)  : ")
				fmt.Scan(&dataHasil[*jumlahHasil].suhu)
				fmt.Print("-> Gula Darah Puasa      : ")
				fmt.Scan(&dataHasil[*jumlahHasil].gula)
				fmt.Print("-> Kolesterol Total      : ")
				fmt.Scan(&dataHasil[*jumlahHasil].kolesterol)
			}

			if pasienDipilih.paket == "2" || pasienDipilih.paket == "3" {
				fmt.Println("\n[ Form Organ Dalam & Jantung - LEGEND ]")
				fmt.Print("-> Kolesterol HDL        : ")
				fmt.Scan(&dataHasil[*jumlahHasil].hdl)
				fmt.Print("-> Kolesterol LDL        : ")
				fmt.Scan(&dataHasil[*jumlahHasil].ldl)
				fmt.Print("-> Trigliserida          : ")
				fmt.Scan(&dataHasil[*jumlahHasil].trigliserida)
				fmt.Print("-> Asam Urat             : ")
				fmt.Scan(&dataHasil[*jumlahHasil].asamUrat)
				fmt.Print("-> Hasil EKG Jantung     : ")
				fmt.Scan(&dataHasil[*jumlahHasil].ekg)
			}

			if pasienDipilih.paket == "3" {
				fmt.Println("\n[ Form Advanced Executive - MYTHIC ]")
				fmt.Print("-> HbA1c (Rata-rata Gula): ")
				fmt.Scan(&dataHasil[*jumlahHasil].hba1c)
				fmt.Print("-> Hasil USG Abdomen     : ")
				fmt.Scan(&dataHasil[*jumlahHasil].usgPerut)
				fmt.Print("-> Kondisi Mata Gamer    : ")
				fmt.Scan(&dataHasil[*jumlahHasil].kondisiMata)
			}

			if dataHasil[*jumlahHasil].gula > 140 || dataHasil[*jumlahHasil].kolesterol > 200 {
				dataHasil[*jumlahHasil].status = "Perlu Tindakan (Gak Aman, Kurangi Begadang!)"
			} else {
				dataHasil[*jumlahHasil].status = "Normal (Aman)"
			}

			*jumlahHasil++
			fmt.Println("\n[+] Hasil check-up pasien berhasil disimpan ke sistem.")
		} else {
			fmt.Println("\n[!] Error: Kapasitas memori data hasil sudah penuh.")
		}
	} else {
		fmt.Println("\n[!] Nama Pasien tidak terdaftar di sistem! Silakan coba lagi.")
		InputDataCheckup(dataHasil, jumlahHasil, dataPasien, jumlahPasien)
	}
}

func CetakHasilCheckup(h Hasil) {
	var namaPaket string

	if h.paket == "1" {
		namaPaket = "EPIC"
	} else if h.paket == "2" {
		namaPaket = "LEGEND"
	} else if h.paket == "3" {
		namaPaket = "MYTHIC"
	}

	fmt.Printf("\n==================================================\n")
	fmt.Printf("               HASIL MEDICAL CHECK-UP             \n")
	fmt.Printf("==================================================\n")
	fmt.Printf("Nama Pasien     : %s\n", h.namaPasien)
	fmt.Printf("Tanggal Periksa : %s\n", h.tanggal)
	fmt.Printf("Paket MCU       : %s\n", namaPaket)
	fmt.Printf("Status Akhir    : %s\n", h.status)
	fmt.Printf("--------------------------------------------------\n")

	if h.paket == "1" || h.paket == "2" || h.paket == "3" {
		fmt.Println("[ Hasil Pemeriksaan Fisik & Lab Dasar ]")
		fmt.Printf("-> Tekanan Darah (Tensi) : %d mmHg\n", h.tekanan)
		fmt.Printf("-> Denyut Nadi           : %d /menit\n", h.nadi)
		fmt.Printf("-> Suhu Tubuh            : %.1f Celcius\n", h.suhu)
		fmt.Printf("-> Gula Darah Puasa      : %d mg/dL\n", h.gula)
		fmt.Printf("-> Kolesterol Total      : %d mg/dL\n", h.kolesterol)
	}

	if h.paket == "2" || h.paket == "3" {
		fmt.Println("\n[ Hasil Organ Dalam & Jantung ]")
		fmt.Printf("-> Kolesterol HDL        : %d mg/dL\n", h.hdl)
		fmt.Printf("-> Kolesterol LDL        : %d mg/dL\n", h.ldl)
		fmt.Printf("-> Trigliserida          : %d mg/dL\n", h.trigliserida)
		fmt.Printf("-> Asam Urat             : %.1f mg/dL\n", h.asamUrat)
		fmt.Printf("-> Hasil EKG Jantung     : %s\n", h.ekg)
	}

	if h.paket == "3" {
		fmt.Println("\n[ Hasil Advanced Executive]")
		fmt.Printf("-> HbA1c (Rata-rata Gula): %.1f %%\n", h.hba1c)
		fmt.Printf("-> Hasil USG Abdomen     : %s\n", h.usgPerut)
		fmt.Printf("-> Kondisi Mata Gamer    : %s\n", h.kondisiMata)
	}
	fmt.Printf("==================================================\n")
}

func CariRiwayatPerTahun(dataHasil [NMAX]Hasil, jumlahHasil int, dataPasien [NMAX]Pasien, jumlahPasien int) {
	var tahunCari string
	var i, j int
	var namaPaket string
	var idPasien, usiaPasien int
	var ketemuPasien bool
	var adaData bool

	fmt.Print("Masukkan tahun pemeriksaan yang ingin dicari (YYYY): ")
	fmt.Scan(&tahunCari)

	fmt.Println("\n=======================================================================")
	fmt.Println("ID  | Nama       | Usia | Paket Yang Dipilih | Tanggal Pemeriksaan")
	fmt.Println("=======================================================================")

	adaData = false

	for i = 0; i < jumlahHasil; i++ {
		if len(dataHasil[i].tanggal) == 10 && dataHasil[i].tanggal[6:] == tahunCari {
			adaData = true
			idPasien = -1
			usiaPasien = -1
			ketemuPasien = false
			j = 0

			for j < jumlahPasien && !ketemuPasien {
				if dataPasien[j].nama == dataHasil[i].namaPasien {
					idPasien = dataPasien[j].id
					usiaPasien = dataPasien[j].umur
					ketemuPasien = true // Mengubah boolean agar loop berhenti otomatis
				}
				j++
			}

			if dataHasil[i].paket == "1" {
				namaPaket = "EPIC"
			} else if dataHasil[i].paket == "2" {
				namaPaket = "LEGEND"
			} else if dataHasil[i].paket == "3" {
				namaPaket = "MYTHIC"
			} else {
				namaPaket = "Tanpa Paket"
			}

			fmt.Printf("%-3d | %-10s | %-4d | %-18s | %s\n",
				idPasien,
				dataHasil[i].namaPasien,
				usiaPasien,
				namaPaket,
				dataHasil[i].tanggal,
			)
		}
	}

	if !adaData {
		fmt.Println("   Tidak ada riwayat pemeriksaan pada tahun tersebut.")
	}
	fmt.Println("=======================================================================")
}

func Pemasukan(dataPasien [NMAX]Pasien, jumlahPasien int) {
	var i, j int
	var countEpic, countLegend, countMythic int
	var grandTotal int

	for i = 0; i < jumlahPasien; i++ {
		if dataPasien[i].paket == "1" {
			countEpic++
		} else if dataPasien[i].paket == "2" {
			countLegend++
		} else if dataPasien[i].paket == "3" {
			countMythic++
		}
	}

	var listPaket [3]LaporanPaket
	listPaket[0] = LaporanPaket{"Paket EPIC  ", countEpic, 500000, countEpic * 500000}
	listPaket[1] = LaporanPaket{"Paket LEGEND", countLegend, 1000000, countLegend * 1000000}
	listPaket[2] = LaporanPaket{"Paket MYTHIC", countMythic, 2000000, countMythic * 2000000}

	for i = 0; i < 2; i++ {
		var maxIdx int = i
		for j = i + 1; j < 3; j++ {
			if listPaket[j].count > listPaket[maxIdx].count {
				maxIdx = j
			}
		}
		var temp LaporanPaket = listPaket[i]
		listPaket[i] = listPaket[maxIdx]
		listPaket[maxIdx] = temp
	}

	grandTotal = listPaket[0].total + listPaket[1].total + listPaket[2].total

	fmt.Println("\n==================================================")
	fmt.Println("          LAPORAN PEMASUKAN PAKET MCU          ")
	fmt.Println("==================================================")
	for i = 0; i < 3; i++ {
		fmt.Printf("%d. %s : %2d pasien x Rp %9d = Rp %d\n",
			i+1,
			listPaket[i].nama,
			listPaket[i].count,
			listPaket[i].harga,
			listPaket[i].total,
		)
	}
	fmt.Println("--------------------------------------------------")
	fmt.Printf("TOTAL PEMASUKAN                      = Rp %d\n", grandTotal)
	fmt.Println("==================================================")
}

func main() {
	var menu int
	var jalan bool = true
	var namaCari string
	var pasienDitemukan Pasien

	for jalan {
		fmt.Println("\n=== SISTEM MEDICAL CHECK-UP ===")
		fmt.Println("1. Kelola Data Pasien")
		fmt.Println("2. Checkup & Riwayat Pasien")
		fmt.Println("3. Laporan Pemasukan")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&menu)

		switch menu {
		case 1:
			var subJalan bool = true
			var menuPasien int

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
						InsertionSortAscendPasienByID(&dataPasien, jumlahPasien)
						CetakDataPasien(dataPasien, jumlahPasien)
					}
				case 3:
					fmt.Print("Masukkan nama Pasien yang ingin dicari: ")
					fmt.Scan(&namaCari)
					InsertionSortDescendPasienByNama(&dataPasien, jumlahPasien)
					var indeksPasien int
					indeksPasien = BinarySearchPasien(dataPasien, jumlahPasien, namaCari)

					if indeksPasien != -1 {
						pasienDitemukan = dataPasien[indeksPasien]
						fmt.Println("\n==================================================")
						fmt.Println("               DATA PROFIL PASIEN                 ")
						fmt.Println("==================================================")
						fmt.Printf("ID Pasien   : %d\n", pasienDitemukan.id)
						fmt.Printf("Nama        : %s\n", pasienDitemukan.nama)
						fmt.Printf("Umur        : %d Tahun\n", pasienDitemukan.umur)
						fmt.Println("==================================================")
					} else {
						fmt.Println("Maaf, data pasien dengan nama tersebut tidak ditemukan.")
					}
				case 4:
					HapusDataPasien(&dataPasien, &jumlahPasien)
				case 5:
					fmt.Println("Kembali ke Menu Utama...")
					subJalan = false
				default:
					fmt.Println("Pilihan sub-menu tidak tersedia.")
				}
			}
		case 2:
			var subJalanCheckup bool = true
			var menuCheckup int
			for subJalanCheckup {
				fmt.Println("\n=== MENU KELOLA DATA CHECK-UP ===")
				fmt.Println("1. Pilih Paket MCU")
				fmt.Println("2. Tambah Hasil Check-up")
				fmt.Println("3. Cari Riwayat Check-up Pasien")
				fmt.Println("4. Cari Riwayat Check-up per Tahun")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilih sub-menu: ")
				fmt.Scan(&menuCheckup)

				switch menuCheckup {
				case 1:
					PilihPaket(&dataPasien, jumlahPasien)
				case 2:
					InputDataCheckup(&dataHasil, &jumlahHasil, dataPasien, jumlahPasien)
				case 3:
					fmt.Print("Masukkan nama Pasien untuk melihat riwayat check-up: ")
					fmt.Scan(&namaCari)

					pasienDitemukan = CariDataPasienSeq(dataPasien, jumlahPasien, namaCari)

					if pasienDitemukan.id != -1 {
						fmt.Printf("\n--- RIWAYAT CHECK-UP PASIEN: %s (ID: %d) ---\n", pasienDitemukan.nama, pasienDitemukan.id)

						var adaRiwayat bool = false
						var i int = 0

						for i < jumlahHasil {
							if dataHasil[i].namaPasien == namaCari {
								CetakHasilCheckup(dataHasil[i])
								adaRiwayat = true
							}
							i++
						}

						if !adaRiwayat {
							fmt.Println("Pasien ini belum memiliki riwayat pemeriksaan.")
						}
					} else {
						fmt.Println("Maaf, data pasien dengan nama tersebut tidak ditemukan.")
					}
				case 4:
					CariRiwayatPerTahun(dataHasil, jumlahHasil, dataPasien, jumlahPasien)
				case 5:
					fmt.Println("Kembali ke Menu Utama...")
					subJalanCheckup = false
				default:
					fmt.Println("Pilihan sub-menu tidak tersedia.")
				}
			}
		case 3:
			Pemasukan(dataPasien, jumlahPasien)
		case 4:
			fmt.Println("Terima kasih!")
			jalan = false
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
