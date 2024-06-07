package main

import (
	"fmt"
)

// Data structures
type Calon struct {
	ID          int
	Nama        string
	Partai      string
	WNI         bool
	JumlahSuara int
}

type Pemilih struct {
	ID           int
	Nama         string
	umur         int
	SudahMemilih bool
}

// Global arrays
type tabCalon [100]Calon
type tabPemilih [100]Pemilih

func main() {
	var jumlahCalon, jumlahPemilih int
	var datacalon tabCalon
	var datapemilih tabPemilih

	// Program utama
	// fmt.Println("Selamat datang di aplikasi Pemilihan Umum Calon Legislatif")

	// Menu
	for {
		fmt.Println("\033[1;33m█▀ █ █▀ ▀█▀ █▀▀ █▀▄▀█\033[0m")
		fmt.Println("\033[1;33m▄█ █ ▄█ ░█░ ██▄ █░▀░█\033[0m")
		fmt.Println("\n\033[1;33m█▀█ █▀▀ █▀▄▀█ █ █░░ █ █░█ ▄▀█ █▄░█\033[0m")
		fmt.Println("\033[1;33m█▀▀ ██▄ █░▀░█ █ █▄▄ █ █▀█ █▀█ █░▀█\033[0m")
		fmt.Println("\n\033[1;33m█▀▀ ▄▀█ █░░ █▀█ █▄░█\033[0m")
		fmt.Println("\033[1;33m█▄▄ █▀█ █▄▄ █▄█ █░▀█\033[0m")
		fmt.Println("\n\033[1;33m█▀█ █▀▀ █▀▄▀█ █ █▀▄▀█ █▀█ █ █▄░█\033[0m")
		fmt.Println("\033[1;33m█▀▀ ██▄ █░▀░█ █ █░▀░█ █▀▀ █ █░▀█\033[0m")

		fmt.Println("\n\033[1mMenu Utama:\033[0m")

		fmt.Println("1. ‍ \033[36mTambah Data Calon\033[0m")
		fmt.Println("2. ‍‍ \033[36mTambah Data Pemilih\033[0m")
		fmt.Println("3.  \033[36mCetak Data Calon\033[0m")
		fmt.Println("4.  \033[36mCetak Data Pemilih\033[0m")
		fmt.Println("5.  \033[36mCari Data Calon Berdasarkan Partai\033[0m")
		fmt.Println("6.  \033[36mEdit Data Calon\033[0m")
		fmt.Println("7.  \033[36mEdit Data Pemilih\033[0m")
		fmt.Println("8.  \033[36mHapus Data Calon\033[0m")
		fmt.Println("9.  \033[36mHapus Data Pemilih\033[0m")
		fmt.Println("10. \033[36mPilih Calon\033[0m")
		fmt.Println("11. \033[36mTampilkan Hasil Perolehan Suara\033[0m")
		fmt.Println("\033[1m12. Keluar\033[0m")

		var pilihan int
		fmt.Print("\nPilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahCalon(&datacalon, &jumlahCalon)
		case 2:
			tambahPemilih(&datapemilih, &jumlahPemilih)
		case 3:
			cetakcalon(datacalon, jumlahCalon)
		case 4:
			cetakpemilih(datapemilih, jumlahPemilih)
		case 5:
			cetakcalonbypartai(datacalon, jumlahCalon)
		case 6:
			editcalon(&datacalon, jumlahCalon)
		case 7:
			editpemilih(&datapemilih, jumlahPemilih)
		case 8:
			DeleteCalon(&datacalon, &jumlahCalon)
		case 9:
			DeletePemilih(&datapemilih, &jumlahPemilih)
		case 10:
			pilihCalon(&datapemilih, &datacalon, jumlahPemilih, jumlahCalon)
		case 11:
			var pilihan int
			fmt.Println("\n Tampilkan Hasil Perolehan Suara: ")
			fmt.Println("1. Urutkan secara Ascending")
			fmt.Println("2. Urutkan secara Descending")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				fmt.Println("Hasil perolehan suara tertinggi (Ascending):")
				ascendingSort(&datacalon, jumlahCalon)
				cetakcalonAsort(datacalon, jumlahCalon)
			case 2:
				fmt.Println("Hasil perolehan suara terkecil (Descending):")
				descendingSort(&datacalon, jumlahCalon)
				cetakcalonAsort(datacalon, jumlahCalon)
			default:
				fmt.Println("Pilihan tidak valid.")
				return
			}
		case 12:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

/*
func selectionSort(ascending bool) {
	for i := 0; i < jumlahCalon-1; i++ {
		idx := i
		for j := i + 1; j < jumlahCalon; j++ {
			if ascending {
				if daftarCalon[j].JumlahSuara < daftarCalon[idx].JumlahSuara {
					idx = j
				}
			} else {
				if daftarCalon[j].JumlahSuara > daftarCalon[idx].JumlahSuara {
					idx = j
				}
			}
		}
		daftarCalon[i], daftarCalon[idx] = daftarCalon[idx], daftarCalon[i]
	}
}
*/

func ascendingSort(A *tabCalon, n int) {
	var i, j int
	var temp Calon
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.JumlahSuara > A[j-1].JumlahSuara {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
		i++
	}
}

func descendingSort(A *tabCalon, n int) {
	var i, j int
	var temp Calon
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.JumlahSuara < A[j-1].JumlahSuara {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
		i++
	}
}

func tambahCalon(A *tabCalon, n *int) {
	var nama, partai string
	var wni bool
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Partai Anda: ")
	fmt.Scan(&partai)
	fmt.Print("Apakah Anda WNI(true/false): ")
	fmt.Scan(&wni)
	if wni == false {
		fmt.Println("Maaf, Anda tidak memenuhi syarat.")
	} else {
		A[*n].ID = *n
		A[*n].Nama = nama
		A[*n].Partai = partai
		A[*n].WNI = wni
		A[*n].JumlahSuara = 0
		*n = *n + 1
		fmt.Println("Data calon berhasil ditambahkan.")
	}
}

func cetakcalon(A tabCalon, n int) {
	if n != 0 {
		fmt.Println("==== Data Calon ===")
		fmt.Printf("%-15s %-15s %s\n", "NamaCalon", "PartaiCalon", "Warga Negara Indonesia(true/false)")
		fmt.Println("-------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-15s %-15s %t\n", A[i].Nama, A[i].Partai, A[i].WNI)
		}
	} else {
		fmt.Println("Data calon kosong.")
	}
}

func cetakcalonAsort(A tabCalon, n int) {
	if n != 0 {
		fmt.Println("==== Data Calon ===")
		fmt.Printf("%-15s %-15s %s\n", "NamaCalon", "PartaiCalon", "JumlahSuara")
		fmt.Println("-------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-15s %-15s %v\n", A[i].Nama, A[i].Partai, A[i].JumlahSuara)
		}
	} else {
		fmt.Println("Data calon kosong.")
	}
}

func tambahPemilih(A *tabPemilih, n *int) {
	var nama string
	var umur int
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Umur Anda: ")
	fmt.Scan(&umur)
	if umur < 17 {
		fmt.Println("Maaf, Anda tidak memenuhi syarat.")
	} else {
		A[*n].ID = *n
		A[*n].Nama = nama
		A[*n].umur = umur
		A[*n].SudahMemilih = false
		*n = *n + 1
		fmt.Println("Data Pemilih berhasil ditambahkan.")
	}
}

func cetakpemilih(A tabPemilih, n int) {
	if n != 0 {
		fmt.Println("==== Data Pemilih ===")
		fmt.Printf("%-15s %-15s %s\n", "NamaPemilih", "UmurPemilih", "IsVoted?(true/false)")
		fmt.Println("-------------------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-15s %-15v %v\n", A[i].Nama, A[i].umur, A[i].SudahMemilih)
		}
	} else {
		fmt.Println("Data Pemilih kosong.")
	}
}

func cetakcalonbypartai(A tabCalon, n int) {
	var partai string
	var count int
	fmt.Print("Masukkan nama partai: ")
	fmt.Scan(&partai)
	fmt.Printf("==== Data Calon Dari Partai: %s ===\n", partai)
	fmt.Printf("%-15s %-15s %s\n", "NamaCalon", "PartaiCalon", "Warga Negara Indonesia\n")
	fmt.Println("-------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		index := searchcalonbypartai(A, i, partai)
		if index == -1 {
			count++
		} else {
			fmt.Printf("%-15s %-15s %t\n", A[index].Nama, A[index].Partai, A[index].WNI)
		}
		if count == n {
			fmt.Println("Data Calon tidak ditemukan.")
		}
	}
}

func searchcalonbypartai(A tabCalon, i int, partai string) int {
	if A[i].Partai == partai {
		return i
	}
	return -1
}

func cariCalon(A tabCalon, nama string, n int) int {
	for i := 0; i < n; i++ {
		if A[i].Nama == nama {
			return i
		}
	}
	return -1
}

func editcalon(A *tabCalon, n int) {
	var nama string
	fmt.Print("Nama calon yang akan diedit: ")
	fmt.Scan(&nama)
	index := cariCalon(*A, nama, n)
	if index != -1 {
		var namaBaru, partaiBaru string
		var wniBaru bool
		fmt.Print("Nama baru: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Partai baru: ")
		fmt.Scan(&partaiBaru)
		fmt.Print("Apakah WNI (true/false): ")
		fmt.Scan(&wniBaru)
		if wniBaru == false {
			fmt.Println("Data calon tidak memenuhi syarat")
		} else {
			A[index].ID = index
			A[index].Nama = namaBaru
			A[index].Partai = partaiBaru
			A[index].WNI = wniBaru
			fmt.Println("Data calon berhasil diedit.")
		}
	} else {
		fmt.Println("Calon tidak ditemukan.")
	}
}

func cariPemilih(A tabPemilih, nama string, n int) int {
	for i := 0; i < n; i++ {
		if A[i].Nama == nama {
			return i
		}
	}
	return -1
}

func editpemilih(A *tabPemilih, n int) {
	var nama string
	fmt.Print("Nama pemilih yang akan diedit: ")
	fmt.Scan(&nama)
	index := cariPemilih(*A, nama, n)
	if index != -1 {
		var namaBaru string
		var umurBaru int
		fmt.Print("Nama baru: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Umur baru: ")
		fmt.Scan(&umurBaru)
		if umurBaru < 17 {
			fmt.Println("Data pemilih tidak memenuhi syarat")
		} else {
			A[index].ID = index
			A[index].Nama = namaBaru
			A[index].umur = umurBaru
			fmt.Println("Data pemilih berhasil diedit.")
		}

	} else {
		fmt.Println("Pemilih tidak ditemukan.")
	}
}

func hapusCalon(A *tabCalon, index int, n *int) {
	for i := index; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func DeleteCalon(A *tabCalon, n *int) {
	var nama string
	fmt.Print("Nama calon yang akan dihapus: ")
	fmt.Scan(&nama)
	index := cariCalon(*A, nama, *n)
	if index != -1 {
		hapusCalon(A, index, n)
		fmt.Println("Calon Berhasil dihapus")
	} else {
		fmt.Println("Calon tidak ditemukan")
	}
}

func hapusPemilih(A *tabPemilih, index int, n *int) {
	for i := index; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func DeletePemilih(A *tabPemilih, n *int) {
	var nama string
	fmt.Print("Nama Pemilih yang akan dihapus: ")
	fmt.Scan(&nama)
	index := cariPemilih(*A, nama, *n)
	if index != -1 {
		hapusPemilih(A, index, n)
		fmt.Println("Calon Berhasil dihapus")
	} else {
		fmt.Println("Calon tidak ditemukan")
	}
}

// Pilih Calon
func pilihCalon(A *tabPemilih, B *tabCalon, jPemilih int, jCalon int) {
	var namaPemilih, namaCalon string
	fmt.Print("Nama pemilih: ")
	fmt.Scan(&namaPemilih)
	pemilihIndex := cariPemilih(*A, namaPemilih, jPemilih)
	if pemilihIndex == -1 {
		fmt.Println("Pemilih tidak ditemukan.")
		return
	}
	if A[pemilihIndex].SudahMemilih {
		fmt.Println("Pemilih sudah melakukan pemilihan sebelumnya.")
		return
	}

	fmt.Print("Nama calon yang dipilih: ")
	fmt.Scan(&namaCalon)
	calonIndex := cariCalon(*B, namaCalon, jCalon)
	if calonIndex == -1 {
		fmt.Println("Calon tidak ditemukan.")
	} else {
		B[calonIndex].JumlahSuara++
		A[pemilihIndex].SudahMemilih = true
		fmt.Println("Pemilihan berhasil.")
	}
}

// Tampilkan Hasil
/*
func tampilkanHasil() {
	var pilihan int
	fmt.Println("\nTampilkan Hasil Perolehan Suara:")
	fmt.Println("1. Urutkan secara Ascending")
	fmt.Println("2. Urutkan secara Descending")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		selectionSort(true)
		fmt.Println("Hasil perolehan suara (Ascending):")
	case 2:
		selectionSort(false)
		fmt.Println("Hasil perolehan suara (Descending):")
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	for i := 0; i < jumlahCalon; i++ {
		fmt.Printf("Calon: %s, Partai: %s, Jumlah Suara: %d\n", daftarCalon[i].Nama, daftarCalon[i].Partai, daftarCalon[i].JumlahSuara)
	}
}
*/
