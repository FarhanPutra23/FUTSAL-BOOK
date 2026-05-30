package main

import (
	"fmt"
)

const NMAX = 100

// ==========================================
// DEKLARASI STRUCT & ARRAY
// ==========================================
type lapangan struct {
	IDlapang, harga int
	nama            string
}
type dataPenyewa struct {
	IDpenyewa   int
	Nama, NoTlp string
}
type jadwalDANtransaksi struct {
	IDtransaksi, IDlapang, IDpenyewa int
	bulan, tanggal, jamMulai         int
	durasi, total                    int
}
type jadwalKosong struct {
	jam, harga int
	namaLap    string
}
type arrLapangan [NMAX]lapangan
type arrPenyewa [NMAX]dataPenyewa
type arrTransaksi [NMAX]jadwalDANtransaksi

// Array jadwal bisa lebih besar karena 1 lapangan bisa punya banyak jam kosong
type arrJadwal [NMAX * 24]jadwalKosong

// ==========================================
// FUNGSI BANTUAN
// ==========================================
func bersihkanLayar() {
	fmt.Print("\033[H\033[2J")
}
func tahanLayar() {
	var tahan string
	fmt.Print("\nKetik apa saja lalu Enter untuk kembali...")
	fmt.Scan(&tahan)
}

// ==========================================
// PROGRAM UTAMA
// ==========================================
func main() {
	var dataL arrLapangan
	var nLap int = 0

	var dataP arrPenyewa
	var nPenyewa int = 0

	var dataT arrTransaksi
	var nTrans int = 0

	var pilihan int
	aplikasiJalan := true

	for aplikasiJalan {
		bersihkanLayar()
		fmt.Println("==========================================")
		fmt.Println("        Penyewaan Lapangan Futsal         ")
		fmt.Println("==========================================")
		fmt.Println("1. Kelola Data Lapangan")
		fmt.Println("2. Kelola Data Penyewa")
		fmt.Println("3. Catat Transaksi Sewa")
		fmt.Println("4. Cari Data Penyewa (Sequential & Binary)")
		fmt.Println("5. Urutkan Jadwal Kosong (Selection & Insertion)")
		fmt.Println("6. Statistik Pendapatan & Jam Terlaris")
		fmt.Println("7. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Pilih menu (1-7): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuLapangan(&dataL, &nLap)
		} else if pilihan == 2 {
			menuPenyewa(&dataP, &nPenyewa)
		} else if pilihan == 3 {
			tambahTransaksi(&dataT, &nTrans, dataP, nPenyewa, dataL, nLap)
		} else if pilihan == 4 {
			cariPenyewa(&dataP, nPenyewa)
		} else if pilihan == 5 {
			urutJadwalKosong(dataL, nLap, dataT, nTrans)
		} else if pilihan == 6 {
			tampilStatistik(dataT, nTrans)
		} else if pilihan == 7 {
			aplikasiJalan = false
			bersihkanLayar()
			fmt.Println("Terima kasih telah menggunakan Futsal-Book!")
		} else {
			fmt.Println("\nPilihan tidak valid!")
			tahanLayar()
		}
	}
}

// ==========================================
// FUNGSI MENU CRUD (Disatukan Biar Rapi)
// ==========================================
func menuLapangan(L *arrLapangan, n *int) {
	bersihkanLayar()
	fmt.Println("--- KELOLA DATA LAPANGAN ---")
	if *n == 0 {
		fmt.Println("Data lapangan masih kosong.")
	} else {
		i := 0
		for i < *n {
			fmt.Printf("%d. ID: %d | Nama: %s | Harga/Jam: Rp%d\n", i+1, L[i].IDlapang, L[i].nama, L[i].harga)
			i++
		}
	}
	fmt.Println("\nAksi:")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih aksi (1-4): ")

	var aksi int
	fmt.Scan(&aksi)

	if aksi == 1 {
		if *n < NMAX {
			fmt.Print("\nID Lapangan Baru: ")
			fmt.Scan(&L[*n].IDlapang)
			fmt.Print("Nama Lapangan: ")
			fmt.Scan(&L[*n].nama)
			fmt.Print("Harga Sewa Per Jam: Rp ")
			fmt.Scan(&L[*n].harga)
			*n = *n + 1
			fmt.Println("Data berhasil ditambahkan!")
		} else {
			fmt.Println("Kapasitas penuh!")
		}
		tahanLayar()
	} else if aksi == 2 {
		var cariID int
		fmt.Print("\nMasukkan ID Lapangan yang diubah: ")
		fmt.Scan(&cariID)
		ketemu := false
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if L[j].IDlapang == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&L[idx].nama)
			fmt.Print("Harga Baru: Rp ")
			fmt.Scan(&L[idx].harga)
			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	} else if aksi == 3 {
		var cariID int
		fmt.Print("\nMasukkan ID Lapangan yang dihapus: ")
		fmt.Scan(&cariID)
		ketemu := false
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if L[j].IDlapang == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			k := idx
			for k < *n-1 {
				L[k] = L[k+1]
				k++
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	}
}
func menuPenyewa(P *arrPenyewa, n *int) {
	bersihkanLayar()
	fmt.Println("--- KELOLA DATA PENYEWA ---")
	if *n == 0 {
		fmt.Println("Data penyewa masih kosong.")
	} else {
		i := 0
		for i < *n {
			fmt.Printf("%d. ID: %d | Nama: %s | No.Tlp: %s\n", i+1, P[i].IDpenyewa, P[i].Nama, P[i].NoTlp)
			i++
		}
	}
	fmt.Println("\nAksi:")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih aksi (1-4): ")

	var aksi int
	fmt.Scan(&aksi)

	if aksi == 1 {
		if *n < NMAX {
			fmt.Print("\nID Penyewa Baru: ")
			fmt.Scan(&P[*n].IDpenyewa)
			fmt.Print("Nama Penyewa (Tanpa Spasi): ")
			fmt.Scan(&P[*n].Nama)
			fmt.Print("Nomor Telepon: ")
			fmt.Scan(&P[*n].NoTlp)
			*n = *n + 1
			fmt.Println("Data berhasil ditambahkan!")
		} else {
			fmt.Println("Kapasitas penuh!")
		}
		tahanLayar()
	} else if aksi == 2 {
		var cariID int
		fmt.Print("\nMasukkan ID Penyewa yang diubah: ")
		fmt.Scan(&cariID)
		ketemu := false
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if P[j].IDpenyewa == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&P[idx].Nama)
			fmt.Print("No Tlp Baru: ")
			fmt.Scan(&P[idx].NoTlp)
			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	} else if aksi == 3 {
		var cariID int
		fmt.Print("\nMasukkan ID Penyewa yang dihapus: ")
		fmt.Scan(&cariID)
		ketemu := false
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if P[j].IDpenyewa == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			k := idx
			for k < *n-1 {
				P[k] = P[k+1]
				k++
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	}
}

// ==========================================
// TRANSAKSI SEWA (Poin b)
// ==========================================
func tambahTransaksi(T *arrTransaksi, nT *int, P arrPenyewa, nP int, L arrLapangan, nL int) {
	bersihkanLayar()
	fmt.Println("--- CATAT TRANSAKSI SEWA ---")
	if *nT < NMAX {
		if nP == 0 || nL == 0 {
			fmt.Println("Data Lapangan dan Penyewa harus ada minimal 1!")
		} else {
			fmt.Print("ID Transaksi: ")
			fmt.Scan(&T[*nT].IDtransaksi)
			fmt.Print("ID Lapangan yang disewa: ")
			fmt.Scan(&T[*nT].IDlapang)
			fmt.Print("ID Penyewa: ")
			fmt.Scan(&T[*nT].IDpenyewa)
			fmt.Print("Bulan (1-12): ")
			fmt.Scan(&T[*nT].bulan)
			fmt.Print("Tanggal (1-31): ")
			fmt.Scan(&T[*nT].tanggal)
			fmt.Print("Jam Mulai (contoh 14 untuk 14:00): ")
			fmt.Scan(&T[*nT].jamMulai)
			fmt.Print("Durasi Sewa (Jam): ")
			fmt.Scan(&T[*nT].durasi)

			// Cek Harga Lapangan
			hargaPerJam := 0
			ketemuLap := false
			i := 0
			for i < nL && !ketemuLap {
				if L[i].IDlapang == T[*nT].IDlapang {
					hargaPerJam = L[i].harga
					ketemuLap = true
				}
				i++
			}

			if ketemuLap {
				T[*nT].total = T[*nT].durasi * hargaPerJam
				*nT = *nT + 1
				fmt.Printf("\nTransaksi Berhasil! Total Harga: Rp%d\n", T[*nT-1].total)
			} else {
				fmt.Println("ID Lapangan tidak ditemukan, transaksi gagal.")
			}
		}
	} else {
		fmt.Println("Kapasitas transaksi penuh!")
	}
	tahanLayar()
}

// ==========================================
// SEARCHING (Poin c)
// ==========================================
func cariPenyewa(P *arrPenyewa, n int) {
	bersihkanLayar()
	fmt.Println("--- CARI PENYEWA ---")
	fmt.Println("1. Berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Berdasarkan No HP (Binary Search)")
	fmt.Print("Pilih (1/2): ")

	var pilCari int
	fmt.Scan(&pilCari)

	if pilCari == 1 {
		var cariNama string
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&cariNama)

		ketemu := false
		i := 0
		for i < n && !ketemu {
			if P[i].Nama == cariNama {
				fmt.Printf("\nDitemukan! ID: %d | Nama: %s | Tlp: %s\n", P[i].IDpenyewa, P[i].Nama, P[i].NoTlp)
				ketemu = true
			}
			i++
		}
		if !ketemu {
			fmt.Println("Nama tidak ditemukan.")
		}

	} else if pilCari == 2 {
		var cariNo string
		fmt.Print("Masukkan No HP: ")
		fmt.Scan(&cariNo)

		// Insertion Sort Ascending berdasarkan NoTlp untuk syarat Binary Search
		i := 1
		for i < n {
			temp := P[i]
			j := i - 1
			for j >= 0 && P[j].NoTlp > temp.NoTlp {
				P[j+1] = P[j]
				j--
			}
			P[j+1] = temp
			i++
		}

		// Binary Search
		kiri := 0
		kanan := n - 1
		ketemu := false

		for kiri <= kanan && !ketemu {
			tengah := (kiri + kanan) / 2
			if P[tengah].NoTlp == cariNo {
				fmt.Printf("\nDitemukan! ID: %d | Nama: %s | Tlp: %s\n", P[tengah].IDpenyewa, P[tengah].Nama, P[tengah].NoTlp)
				ketemu = true
			} else if P[tengah].NoTlp < cariNo {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
		if !ketemu {
			fmt.Println("No HP tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	tahanLayar()
}

// ==========================================
// SORTING JADWAL KOSONG (Poin d)
// ==========================================
func urutJadwalKosong(L arrLapangan, nLap int, T arrTransaksi, nTrans int) {
	bersihkanLayar()
	fmt.Println("--- URUTKAN JADWAL KOSONG ---")

	var cariBulan, cariTanggal int
	fmt.Print("Cek jadwal kosong untuk Bulan (1-12): ")
	fmt.Scan(&cariBulan)
	fmt.Print("Cek jadwal kosong untuk Tanggal (1-31): ")
	fmt.Scan(&cariTanggal)

	var J arrJadwal
	nJadwal := 0

	// Buat jadwal untuk semua lapangan, dari jam 08:00 sd 22:00
	iLap := 0
	for iLap < nLap {
		jam := 8
		for jam <= 22 {
			dipakai := false
			// Cek apakah jam ini bentrok dengan transaksi
			iT := 0
			for iT < nTrans && !dipakai {
				if T[iT].bulan == cariBulan && T[iT].tanggal == cariTanggal && T[iT].IDlapang == L[iLap].IDlapang {
					if jam >= T[iT].jamMulai && jam < (T[iT].jamMulai+T[iT].durasi) {
						dipakai = true
					}
				}
				iT++
			}

			if !dipakai {
				J[nJadwal].jam = jam
				J[nJadwal].namaLap = L[iLap].nama
				J[nJadwal].harga = L[iLap].harga
				nJadwal++
			}
			jam++
		}
		iLap++
	}

	fmt.Println("\nPilih Metode Pengurutan:")
	fmt.Println("1. Berdasarkan Jam Mulai (Selection Sort - Ascending)")
	fmt.Println("2. Berdasarkan Harga (Insertion Sort - Ascending)")
	fmt.Print("Pilih (1/2): ")

	var pilUrut int
	fmt.Scan(&pilUrut)

	if pilUrut == 1 {
		// Selection Sort berdasarkan Jam
		i := 0
		for i < nJadwal-1 {
			idxMin := i
			j := i + 1
			for j < nJadwal {
				if J[j].jam < J[idxMin].jam {
					idxMin = j
				}
				j++
			}
			temp := J[i]
			J[i] = J[idxMin]
			J[idxMin] = temp
			i++
		}
	} else if pilUrut == 2 {
		// Insertion Sort berdasarkan Harga
		i := 1
		for i < nJadwal {
			temp := J[i]
			j := i - 1
			for j >= 0 && J[j].harga > temp.harga {
				J[j+1] = J[j]
				j--
			}
			J[j+1] = temp
			i++
		}
	}
	fmt.Println("\n--- HASIL JADWAL KOSONG ---")
	i := 0
	for i < nJadwal {
		fmt.Printf("- Pukul %02d:00 | %s | Rp%d\n", J[i].jam, J[i].namaLap, J[i].harga)
		i++
	}
	tahanLayar()
}

// ==========================================
// STATISTIK (Poin e)
// ==========================================
func tampilStatistik(T arrTransaksi, n int) {
	bersihkanLayar()
	fmt.Println("--- STATISTIK PENDAPATAN ---")

	var cariBulan int
	fmt.Print("Masukkan Bulan (1-12) untuk melihat statistik: ")
	fmt.Scan(&cariBulan)

	totalPendapatan := 0
	var frekJam [24]int // Array untuk melacak jam berapa saja yang dipesan

	i := 0
	for i < n {
		if T[i].bulan == cariBulan {
			totalPendapatan += T[i].total

			jam := T[i].jamMulai
			k := 0
			for k < T[i].durasi {
				if (jam + k) < 24 {
					frekJam[jam+k]++
				}
				k++
			}
		}
		i++
	}
	jamTerlaris := 0
	maxOrder := 0
	i = 0
	for i < 24 {
		if frekJam[i] > maxOrder {
			maxOrder = frekJam[i]
			jamTerlaris = i
		}
		i++
	}
	fmt.Println("\n--- HASIL STATISTIK ---")
	fmt.Printf("Total Pendapatan Bulan %d : Rp %d\n", cariBulan, totalPendapatan)
	if maxOrder > 0 {
		fmt.Printf("Jam Paling Laris Bulan %d : Pukul %02d:00 (Disewa sebanyak %d jam)\n", cariBulan, jamTerlaris, maxOrder)
	} else {
		fmt.Println("Belum ada data transaksi di bulan ini.")
	}
	tahanLayar()
}
