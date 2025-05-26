package main

import "fmt"

const NMAX int = 10

type UserData struct {
	Pengalaman   [NMAX]string
	Keterampilan [NMAX]string
	Pendidikan   [NMAX]string
	IdxPengalaman   int
	IdxKeterampilan int
	IdxPendidikan   int
}

type tabData UserData

var pekerjaan = [5]string{
	"AI_Engineer",
	"Backend_Developer",
	"Data_Analyst",
	"Frontend_Developer",
	"UI/UX_Designer",
}
var gaji = [5]int{12000000, 10000000, 9000000, 9500000, 11000000} // paralel

func main() {
	var data tabData
	for {
		menu()
		var pilih int
		fmt.Scan(&pilih)
		fmt.Scanln() 
		fmt.Println()

		switch pilih {
		case 1:
			tambahData(&data)
		case 2:
			ubahData(&data)
		case 3:
			hapusData(&data)
		case 4:
			cetakResume(data)
		case 5:
			cetakSurat(data)
		case 6:
			saranResume(data)
		case 7:
			saranSurat(data)
		case 8:
			cariPekerjaanSequential()
		case 9:
			cariPekerjaanBinary()
		case 10:
			selectionSort()
		case 11:
			insertionSort()
		case 12:
			evaluasiResume(data)
		case 0:
			fmt.Println("Trimakasih telah menggunakan")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menu(){
	fmt.Println("\n=== MENU UTAMA ===")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Cetak Resume")
	fmt.Println("5. Cetak Surat Lamaran")
	fmt.Println("6. Saran AI Resume")
	fmt.Println("7. Saran AI Surat Lamaran")
	fmt.Println("8. Cari Pekerjaan (Sequential)")
	fmt.Println("9. Cari Pekerjaan (Binary)")
	fmt.Println("10. Urutkan Pekerjaan (Selection Sort - Nama)")
	fmt.Println("11. Urutkan Pekerjaan (Insertion Sort - Gaji)")
	fmt.Println("12. Evaluasi Resume")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahData(A *tabData) {
	var input string
	
	fmt.Println("Masukkan Pengalaman Kerja: (. untuk selesai)")
	input = " "
	for input != "."{
		fmt.Scan(&input)
		if input != "."{
			A.Pengalaman[A.IdxPengalaman] = input
			A.IdxPengalaman++
		}
	}
	
	fmt.Println("Masukkan Keterampilan: (. untuk selesai)")
	input = " "
	for input != "."{
		fmt.Scan(&input)
		if input != "."{
			A.Keterampilan[A.IdxKeterampilan] = input
			A.IdxKeterampilan++
		}
	}
	
	fmt.Println("Masukkan Pendidikan: (. untuk selesai)")
	input = " "
	for input != "."{
		fmt.Scan(&input)
		if input != "."{
			A.Pendidikan[A.IdxPendidikan] = input
			A.IdxPendidikan++
		}
	}

	fmt.Println("Data berhasil ditambahkan.")
}

func ubahData(A *tabData) {
	var kategori int
	fmt.Println("\n== Ubah Data ==")
	fmt.Println("1. Pengalaman")
	fmt.Println("2. Keterampilan")
	fmt.Println("3. Pendidikan")
	fmt.Print("Pilih kategori yang ingin diubah: ")
	fmt.Scan(&kategori)

	var idx int
	var input string

	switch kategori {
	case 1:
		if A.IdxPengalaman == 0 {
			fmt.Println("Belum ada data pengalaman.")
			return
		}
		
		fmt.Println("Daftar Pengalaman:")
		for i := 0; i < A.IdxPengalaman; i++ {
			fmt.Printf("[%d] %s\n", i, A.Pengalaman[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin diubah: ")
		fmt.Scan(&idx)
		
		if idx >= 0 && idx < A.IdxPengalaman {
			fmt.Print("Masukkan pengalaman baru: ")
			fmt.Scan(&input)
			A.Pengalaman[idx] = input
			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
		
	case 2:
		if A.IdxKeterampilan == 0 {
			fmt.Println("Belum ada data keterampilan.")
			return
		}
		
		fmt.Println("Daftar Keterampilan:")
		for i := 0; i < A.IdxKeterampilan; i++ {
			fmt.Printf("[%d] %s\n", i, A.Keterampilan[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin diubah: ")
		fmt.Scan(&idx)
		
		if idx >= 0 && idx < A.IdxKeterampilan {
			fmt.Print("Masukkan keterampilan baru: ")
			fmt.Scan(&input)
			A.Keterampilan[idx] = input
			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
		
	case 3:
		if A.IdxPendidikan == 0 {
			fmt.Println("Belum ada data pendidikan.")
			return
		}
		
		fmt.Println("Daftar Pendidikan:")
		for i := 0; i < A.IdxPendidikan; i++ {
			fmt.Printf("[%d] %s\n", i, A.Pendidikan[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < A.IdxPendidikan {
			fmt.Print("Masukkan pendidikan baru: ")
			fmt.Scan(&input)
			A.Pendidikan[idx] = input
			fmt.Println("Data berhasil diubah.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
		
	default:
		fmt.Println("Kategori tidak valid.")
	}
}

func hapusData(A *tabData) {
	var kategori int
	fmt.Println("\n== Hapus Data ==")
	fmt.Println("1. Pengalaman")
	fmt.Println("2. Keterampilan")
	fmt.Println("3. Pendidikan")
	fmt.Print("Pilih kategori yang ingin dihapus: ")
	fmt.Scan(&kategori)

	var idx int

	switch kategori {
	case 1:
		if A.IdxPengalaman == 0 {
			fmt.Println("Belum ada data pengalaman.")
			return
		}
		
		fmt.Println("Daftar Pengalaman:")
		for i := 0; i < A.IdxPengalaman; i++ {
			fmt.Printf("[%d] %s\n", i, A.Pengalaman[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin dihapus: ")
		fmt.Scan(&idx)
		
		if idx >= 0 && idx < A.IdxPengalaman {
			for i := idx; i < A.IdxPengalaman-1; i++ {
				A.Pengalaman[i] = A.Pengalaman[i+1]
			}
			A.IdxPengalaman--
			fmt.Println("Data berhasil dihapus.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
	case 2:
		if A.IdxKeterampilan == 0 {
			fmt.Println("Belum ada data keterampilan.")
			return
		}
		
		fmt.Println("Daftar Keterampilan:")
		for i := 0; i < A.IdxKeterampilan; i++ {
			fmt.Printf("[%d] %s\n", i, A.Keterampilan[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < A.IdxKeterampilan {
			for i := idx; i < A.IdxKeterampilan-1; i++ {
				A.Keterampilan[i] = A.Keterampilan[i+1]
			}
			A.IdxKeterampilan--
			fmt.Println("Data berhasil dihapus.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
		
	case 3:
		if A.IdxPendidikan == 0 {
			fmt.Println("Belum ada data pendidikan.")
			return
		}
		
		fmt.Println("Daftar Pendidikan:")
		for i := 0; i < A.IdxPendidikan; i++ {
			fmt.Printf("[%d] %s\n", i, A.Pendidikan[i])
		}
		
		fmt.Print("Pilih indeks data yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx >= 0 && idx < A.IdxPendidikan {
			for i := idx; i < A.IdxPendidikan-1; i++ {
				A.Pendidikan[i] = A.Pendidikan[i+1]
			}
			A.IdxPendidikan--
			fmt.Println("Data berhasil dihapus.")
		} else {
			fmt.Println("Indeks tidak valid.")
		}
		
	default:
		fmt.Println("Kategori tidak valid.")
	}
}

func cetakResume(A tabData) {
	fmt.Println("\n=== RESUME ===")

	fmt.Println("\n> Pengalaman Kerja:")
	if A.IdxPengalaman == 0 {
		fmt.Println("  (Belum ada data)")
	} else {
		for i := 0; i < A.IdxPengalaman; i++ {
			fmt.Printf("  - %s\n", A.Pengalaman[i])
		}
	}

	fmt.Println("\n> Keterampilan:")
	if A.IdxKeterampilan == 0 {
		fmt.Println("  (Belum ada data)")
	} else {
		for i := 0; i < A.IdxKeterampilan; i++ {
			fmt.Printf("  - %s\n", A.Keterampilan[i])
		}
	}

	fmt.Println("\n> Pendidikan:")
	if A.IdxPendidikan == 0 {
		fmt.Println("  (Belum ada data)")
	} else {
		for i := 0; i < A.IdxPendidikan; i++ {
			fmt.Printf("  - %s\n", A.Pendidikan[i])
		}
	}
}

func cetakSurat(A tabData) {
	var nama string
	
	fmt.Println("\n=== SURAT LAMARAN KERJA ===")
	fmt.Println("Kepada Yth. HRD Perusahaan")
	fmt.Println("Dengan hormat,\n")

	if A.IdxPengalaman > 0 {
		fmt.Print("Saya memiliki pengalaman sebagai ")
		for i := 0; i < A.IdxPengalaman; i++ {
			fmt.Print(A.Pengalaman[i])
			if i == A.IdxPengalaman-2 {
				fmt.Print(" dan ")
			} else if i < A.IdxPengalaman-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println(".")
	}

	if A.IdxKeterampilan > 0 {
		fmt.Print("Saya memiliki keahlian dalam ")
		for i := 0; i < A.IdxKeterampilan; i++ {
			fmt.Print(A.Keterampilan[i])
			if i == A.IdxKeterampilan-2 {
				fmt.Print(" dan ")
			} else if i < A.IdxKeterampilan-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println(".")
	}

	if A.IdxPendidikan > 0 {
		fmt.Printf("Saya merupakan lulusan dari %s.\n", A.Pendidikan[0])
	}

	fmt.Println("\nSaya berharap dapat bergabung dan berkontribusi di perusahaan Bapak/Ibu.")
	fmt.Println("Atas perhatian dan kesempatannya, saya ucapkan terima kasih.\n")
	fmt.Println("Hormat saya,")
	fmt.Scan(&nama)
}

func saranResume(A tabData) {
	fmt.Println("\n=== SARAN AI UNTUK RESUME ===")

	if A.IdxPengalaman == 0 {
		fmt.Println("- Tambahkan pengalaman kerja, meskipun itu magang atau freelance.")
	} else if A.IdxPengalaman < 2 {
		fmt.Println("- Coba tambahkan 1–2 pengalaman lagi agar lebih meyakinkan.")
	}

	if A.IdxKeterampilan == 0 {
		fmt.Println("- Tambahkan keterampilan seperti software, bahasa pemrograman, atau tools.")
	} else if A.IdxKeterampilan < 2 {
		fmt.Println("- Coba tambahkan lebih banyak skill teknis atau soft skill.")
	}

	if A.IdxPendidikan == 0 {
		fmt.Println("- Pendidikan belum diisi. Minimal tulis pendidikan terakhir.")
	} else if A.Pendidikan[0] == "SMA" || A.Pendidikan[0] == "SMK" {
		fmt.Println("- Pertimbangkan lanjut kuliah untuk peluang kerja lebih luas.")
	}

	if A.IdxPengalaman >= 2 && A.IdxKeterampilan >= 2 && A.IdxPendidikan > 0 {
		fmt.Println("- Resume Anda sudah cukup kuat! Pertimbangkan untuk mulai melamar posisi sesuai minat.")
	}
}

func saranSurat(A tabData) {
	fmt.Println("\n=== SARAN AI UNTUK SURAT LAMARAN ===")

	if A.IdxPengalaman == 0 {
		fmt.Println("- Surat Anda belum menyebutkan pengalaman kerja. Tambahkan agar lebih meyakinkan.")
	}

	if A.IdxKeterampilan == 0 {
		fmt.Println("- Anda belum mencantumkan keterampilan apa pun. Tambahkan untuk menunjukkan keunggulan.")
	}

	if A.IdxPendidikan == 0 {
		fmt.Println("- Pendidikan belum disebutkan. Cantumkan agar HR bisa menilai latar belakang Anda.")
	}

	if A.IdxPengalaman > 0 && A.IdxKeterampilan > 0 && A.IdxPendidikan > 0 {
		fmt.Println("- Struktur surat lamaran sudah baik. Jangan lupa gunakan bahasa formal dan sopan.")
	}

	if A.IdxPengalaman+A.IdxKeterampilan+A.IdxPendidikan < 3 {
		fmt.Println("- Surat Anda masih sangat singkat. Tambahkan konten agar terlihat serius dan profesional.")
	}
}

func cariPekerjaanSequential() {
	const N = 5
	var nama string
	var found bool = false
	var i int

	fmt.Print("Masukkan nama pekerjaan yang dicari: ")
	fmt.Scanln(&nama)

	for i = 0; i < N; i++ {
		if pekerjaan[i] == nama {
			fmt.Printf("Pekerjaan ditemukan: %s dengan gaji Rp%d\n", pekerjaan[i], gaji[i])
			found = true
			return
		}
	}

	if !found {
		fmt.Println("Pekerjaan tidak ditemukan.")
	}
}

func cariPekerjaanBinary() {
	const N = 5
	var nama string
	var found bool = false
	var i, j, idx int
	var left, right, mid int

	fmt.Print("Masukkan nama pekerjaan yang dicari (Binary Search): ")
	fmt.Scanln(&nama)

	for i = 0; i < N-1; i++ {
		idx = i
		for j = i + 1; j < N; j++ {
			if pekerjaan[j] < pekerjaan[idx] {
				idx = j
			}
		}
		pekerjaan[i], pekerjaan[idx] = pekerjaan[idx], pekerjaan[i]
		gaji[i], gaji[idx] = gaji[idx], gaji[i]
	}

	left = 0
	right = N - 1

	for left <= right {
		mid = (left + right) / 2
		if pekerjaan[mid] == nama {
			fmt.Printf("Pekerjaan ditemukan: %s dengan gaji Rp%d\n", pekerjaan[mid], gaji[mid])
			found = true
			return
		} else if nama < pekerjaan[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		fmt.Println("Pekerjaan tidak ditemukan.")
	}
}

func selectionSort() {
	const N = 5
	var i, j, idx int

	for i = 0; i < N-1; i++ {
		idx = i
		for j = i + 1; j < N; j++ {
			if pekerjaan[j] < pekerjaan[idx] {
				idx = j
			}
		}

		pekerjaan[i], pekerjaan[idx] = pekerjaan[idx], pekerjaan[i]
		gaji[i], gaji[idx] = gaji[idx], gaji[i]
	}

	fmt.Println("\nDaftar Pekerjaan setelah diurutkan berdasarkan nama (A-Z):")
	i = 0
	for i < N {
		fmt.Printf("- %s (Rp%d)\n", pekerjaan[i], gaji[i])
		i++
	}
}

func insertionSort() {
	const N = 5
	var i, j, tempGaji int
	var tempNama string

	for i = 1; i < N; i++ {
		tempGaji = gaji[i]
		tempNama = pekerjaan[i]
		j = i - 1

		for j >= 0 && gaji[j] > tempGaji {
			gaji[j+1] = gaji[j]
			pekerjaan[j+1] = pekerjaan[j]
			j--
		}

		gaji[j+1] = tempGaji
		pekerjaan[j+1] = tempNama
	}

	fmt.Println("\nDaftar Pekerjaan setelah diurutkan berdasarkan gaji (termurah-termahal):")
	i = 0
	for i < N {
		fmt.Printf("- %s (Rp%d)\n", pekerjaan[i], gaji[i])
		i++
	}
}

func evaluasiResume(A tabData) {
	fmt.Println("\n=== EVALUASI RESUME ===")

	var skor, nilaiPengalaman, nilaiKeterampilan, nilaiPendidikan int

	if A.IdxPengalaman >= 3 {
		nilaiPengalaman = 30
	} else if A.IdxPengalaman == 2 {
		nilaiPengalaman = 20
	} else if A.IdxPengalaman == 1 {
		nilaiPengalaman = 10
	} else {
		nilaiPengalaman = 0
	}

	if A.IdxKeterampilan >= 3 {
		nilaiKeterampilan = 30
	} else if A.IdxKeterampilan == 2 {
		nilaiKeterampilan = 20
	} else if A.IdxKeterampilan == 1 {
		nilaiKeterampilan = 10
	} else {
		nilaiKeterampilan = 0
	}

	if A.IdxPendidikan >= 1 {
		nilaiPendidikan = 40
	} else {
		nilaiPendidikan = 0
	}

	skor = nilaiPengalaman + nilaiKeterampilan + nilaiPendidikan

	fmt.Printf("- Pengalaman: %d poin\n", nilaiPengalaman)
	fmt.Printf("- Keterampilan: %d poin\n", nilaiKeterampilan)
	fmt.Printf("- Pendidikan: %d poin\n", nilaiPendidikan)
	fmt.Printf("=> Total Skor Resume: %d/100\n", skor)

	if skor == 100 {
		fmt.Println("Resume sangat siap! Langsung kirim ke perusahaan impianmu.")
	} else if skor >= 70 {
		fmt.Println("Resume cukup baik. Bisa ditingkatkan sedikit lagi.")
	} else if skor >= 40 {
		fmt.Println("Resume masih kurang. Tambahkan pengalaman atau skill.")
	} else {
		fmt.Println("Resume belum layak. Segera lengkapi dulu datanya.")
	}
}