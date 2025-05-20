package main
import "fmt"

	const NMAX = 50
	
	type data struct{
		nama, posisi, email, kota, tglnow, perstuj, noPhone string
		
		exPosisi [NMAX] string
		nExPosisi int
		exPerusaha [NMAX] string
		exKota [NMAX] string
		exTgl [NMAX] string
		Skill [NMAX] string
		nSkill int
		exPengalaman [NMAX] string
		nExPengalaman int
		
		penJur [NMAX] string
		nPenJur int
		uni [NMAX] string
		tLulus [NMAX] int
		ipk [NMAX] float64
		aktivy [NMAX] string
		nAktivy int
		
		serti [NMAX] string
		nSerti int
		penye [NMAX] string
		tSerti [NMAX] int
		
		bhs [NMAX] string
		nBhs int
		
		hobi [NMAX] string
		nHobi int
	}

func main(){
	var d data
	var p int
	var cari string
	
	menu()
	fmt.Scan(&p)
	fmt.Println()
	
	for p != 10{
		if p == 1 {
			form(&d)
		} else if p == 2 {
			cetakResume(d)
		} else if p == 3 {
			cetakSurat(d)
		} else if p == 4 {
			fmt.Print("Masukkan skill yang ingin dicari: ")
			fmt.Scan(&cari)

			if cariSkill(d, cari) {
				fmt.Printf("Skill '%s' ditemukan!\n", cari)
			} else {
				fmt.Printf("Skill '%s' tidak ditemukan.\n", cari)
			}

		} else if p == 5 {
			sortIPK(&d)
		} else if p == 6 {
			editSkill(&d)
		} else if p == 7 {
			hapusSkill(&d)
		} else if p == 8 {
			saranAI(d)
		} else if p == 9 {
			evaluasiResume(d)
		}else {
			fmt.Println("Pilihan tidak valid.")
		}
		
		menu()
		fmt.Scan(&p)	
		fmt.Println()
	}
	fmt.Println("Trimakasih telah menggunakan")
	
}

func menu(){
	fmt.Println("\n===== MENU =====")
	fmt.Println("1. Input Data")
	fmt.Println("2. Cetak Resume")
	fmt.Println("3. Cetak Surat Lamaran")
	fmt.Println("4. Cari Skill (Binary Search)")
	fmt.Println("5. Urutkan Pendidikan Berdasarkan IPK")
	fmt.Println("6. Edit Skill")
	fmt.Println("7. Hapus Skill")
	fmt.Println("8. Saran AI untuk Resume")
	fmt.Println("9. Evaluasi Resume")
	fmt.Println("10. Exit")
	fmt.Print("Silahkan pilih menu (1/2/3/4/5/6/7/8/9/10): ")
}

func form(d *data){
	var input string
	
	fmt.Println("Untuk Spasi gunakan '_'")
	fmt.Println()
	
	fmt.Print("Tanggal Sekarang: ")
	fmt.Scan(&d.tglnow)
	fmt.Println()

	fmt.Print("Nama Lengkap: ")
	fmt.Scan(&d.nama)
	fmt.Println()

	fmt.Print("Posisi yang diinginkan: ")
	fmt.Scan(&d.posisi)
	fmt.Println()
	
	fmt.Print("Perusahaan Tujuan: ")
	fmt.Scan(&d.perstuj)
	fmt.Println()
	
	fmt.Print("Email: ")
	fmt.Scan(&d.email)
	fmt.Println()
	
	fmt.Print("Kota Asal dan negara: <kota, negara>: ")
	fmt.Scan(&d.kota)
	fmt.Println()

	fmt.Print("No Telephone: ")
	fmt.Scan(&d.noPhone)
	fmt.Println()

	fmt.Println("Masukan Pengalaman Kerja: <Posisi 1 (enter), Perusahaan, Kota, Bulan_YYYY, Pengalaman>")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nExPengalaman
			d.exPosisi[i] = input
			fmt.Scan(&d.exPerusaha[i], &d.exKota[i], &d.exTgl[i], &d.exPengalaman[i])
			d.nExPengalaman++
			d.nExPosisi++
		}
	}
	fmt.Println()
	
	fmt.Println("Masukan Histori Pendidikan: <Program Studi (enter), Universitas, Tahun Lulus, IPK, Aktivitas/kepanitiaan>")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nPenJur
			d.penJur[i] = input
			fmt.Scan(&d.uni[i], &d.tLulus[i], &d.ipk[i], &d.aktivy[i])
			d.nPenJur++
			d.nAktivy++
		}
	}
	fmt.Println()
	
	fmt.Println("Masukan Skill:")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nSkill
			d.Skill[i] = input
			d.nSkill++
		}
	}
	fmt.Println()
	
	fmt.Println("Masukan Sertifikat: <Nama Sertifikasi (enter), Penyelenggara, YYYY>")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nSerti
			d.serti[i] = input
			fmt.Scan(&d.penye[i], &d.tSerti[i])
			d.nSerti++
		}
	}
	fmt.Println()
	
	fmt.Println("Masukan Kemampuan Bahasa:")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nBhs
			d.bhs[i] = input
			d.nBhs++
		}
	}
	fmt.Println()

	fmt.Println("Masukan Hobi:")
	fmt.Println("Ketik '.' untuk selesai")
	input = ""
	for input != "." {
		fmt.Scan(&input)
		if input != "." {
			i := d.nHobi
			d.hobi[i] = input
			d.nHobi++
		}
	}
	fmt.Println()
}

func editSkill(d *data) {
	var i, idx int
	
	if d.nSkill == 0 {
		fmt.Println("Belum ada skill yang bisa diedit.")
		return
	}
	fmt.Println("Daftar Skill:")
	for i = 0; i < d.nSkill; i++ {
		fmt.Printf("%d. %s\n", i+1, d.Skill[i])
	}
	
	fmt.Print("Pilih nomor skill yang ingin diedit: ")
	fmt.Scan(&idx)

	if idx < 1 || idx > d.nSkill {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var newSkill string
	fmt.Print("Masukkan skill baru: ")
	fmt.Scan(&newSkill)
	d.Skill[idx-1] = newSkill
	fmt.Println("Skill berhasil diubah.")
}

func hapusSkill(d *data) {
	var i, idx int
	
	if d.nSkill == 0 {
		fmt.Println("Belum ada skill yang bisa dihapus.")
		return
	}
	fmt.Println("Daftar Skill:")
	for i = 0; i < d.nSkill; i++ {
		fmt.Printf("%d. %s\n", i+1, d.Skill[i])
	}
	
	fmt.Print("Pilih nomor skill yang ingin dihapus: ")
	fmt.Scan(&idx)

	if idx < 1 || idx > d.nSkill {
		fmt.Println("Nomor tidak valid.")
		return
	}

	// Geser data ke kiri
	for i = idx - 1; i < d.nSkill-1; i++ {
		d.Skill[i] = d.Skill[i+1]
	}
	d.nSkill--
	fmt.Println("Skill berhasil dihapus.")
}

func saranAI(d data) {
	fmt.Println("\n--- Saran AI untuk Resume ---")

	if d.nSkill < 3 {
		fmt.Println("• Tambahkan lebih banyak skill. Idealnya 3–5 skill utama.")
	}
	if d.nExPengalaman == 0 {
		fmt.Println("• Tambahkan pengalaman kerja, magang, atau proyek freelance.")
	}
	if d.nSerti == 0 {
		fmt.Println("• Sertifikasi akan meningkatkan nilai kamu di mata HR.")
	}
	if d.nPenJur == 0 {
		fmt.Println("• Data pendidikan kosong, segera lengkapi!")
	}
	if hitungRataIPK(d) < 3.0 {
		fmt.Println("• IPK di bawah 3.0, sebaiknya tonjolkan skill & proyek real.")
	}
	if d.nBhs == 0 {
		fmt.Println("• Tambahkan bahasa asing untuk memperluas peluang.")
	}
	if d.nHobi == 0 {
		fmt.Println("• Menambahkan hobi membuatmu lebih menarik secara personal.")
	}

	fmt.Println("Saran AI selesai ditampilkan.")
}

func hitungRataIPK(d data) float64 {
	var total float64
	for i := 0; i < d.nPenJur; i++ {
		total += d.ipk[i]
	}
	if d.nPenJur == 0 {
		return 0
	}
	return total / float64(d.nPenJur)
}

func evaluasiResume(d data) {
	skor := 0

	if d.nSkill >= 5 {
		skor += 30
	} else if d.nSkill >= 3 {
		skor += 20
	} else {
		skor += 10
	}

	if d.nExPengalaman >= 2 {
		skor += 30
	} else if d.nExPengalaman == 1 {
		skor += 15
	}

	rata := hitungRataIPK(d)
	if rata >= 3.5 {
		skor += 20
	} else if rata >= 3.0 {
		skor += 10
	}

	if d.nSerti > 0 {
		skor += 10
	}

	if skor > 90 {
		fmt.Println("Resume sangat kuat (Excellent)")
	} else if skor > 70 {
		fmt.Println("Resume cukup baik (Good)")
	} else {
		fmt.Println("Resume perlu ditingkatkan (Needs Improvement)")
	}

	fmt.Printf("Skor total: %d/100\n", skor)
}

func urutkanSkill(d *data) {
	var i, j int
	var key string
	
	for i = 1; i < d.nSkill; i++ {
		key = d.Skill[i]
		j = i - 1
		for j >= 0 && d.Skill[j] > key {
			d.Skill[j+1] = d.Skill[j]
			j--
		}
		d.Skill[j+1] = key
	}
}

func cariSkill(d data, cari string) bool {
	urutkanSkill(&d)
	var kiri, kanan, tengah int
	kiri = 0 
	kanan = d.nSkill-1
	
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if d.Skill[tengah] == cari {
			return true
		} else if d.Skill[tengah] < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return false
}

func sortIPK(d *data) {
	var i, idx, j, tempLulus int
	var tempIPK float64
	var tempJurusan, tempUniv, tempAktivitas string
	var pilih string

	fmt.Print("Pilih urutan IPK (naik/turun): ")
	fmt.Scan(&pilih)

	for i = 0; i < d.nPenJur-1; i++ {
		idx = i
		for j = i + 1; j < d.nPenJur; j++ {
			if pilih == "turun"{
				if d.ipk[j] > d.ipk[idx] {
					idx = j
				}
			} else{
				if d.ipk[j] < d.ipk[idx] {
					idx = j
				}
			}
		}
		
		if idx != i {
			tempIPK = d.ipk[i]
			d.ipk[i] = d.ipk[idx]
			d.ipk[idx] = tempIPK

			tempJurusan = d.penJur[i]
			d.penJur[i] = d.penJur[idx]
			d.penJur[idx] = tempJurusan
			
			tempUniv = d.uni[i]
			d.uni[i] = d.uni[idx]
			d.uni[idx] = tempUniv

			tempLulus = d.tLulus[i]
			d.tLulus[i] = d.tLulus[idx]
			d.tLulus[idx] = tempLulus

			tempAktivitas = d.aktivy[i]
			d.aktivy[i] = d.aktivy[idx]
			d.aktivy[idx] = tempAktivitas

		}
	}
	
	for i = 0; i < d.nPenJur; i++ {
		fmt.Printf("	%s — %s | %d\n", d.penJur[i], d.uni[i], d.tLulus[i])
		fmt.Printf("	IPK: %.2f | Aktivitas/kepanitiaan relevan: %s\n", d.ipk[i], d.aktivy[i])
	}
	fmt.Println("Data pendidikan berhasil diurutkan berdasarkan IPK (tinggi ke rendah).")
}

func cetakResume(d data) {
	fmt.Println("===============================================================")
	fmt.Printf("	%s | %s\n", d.nama, d.posisi)
	fmt.Printf("	%s · %s · %s\n", d.email, d.noPhone, d.kota)
	fmt.Println("===============================================================")

	fmt.Println("	RINGKASAN PROFIL")
	if d.nExPengalaman > 0 {
		fmt.Printf("	%s dengan pengalaman dalam %s. \n", d.exPosisi[0], d.exPengalaman[0])
	} else {
		fmt.Println("	(Belum ada pengalaman kerja)")
	}
	if d.nSkill > 1 {
		fmt.Printf("	Kuat di %s, %s. \n", d.Skill[0], d.Skill[1])
	}

	if d.nExPengalaman > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	PENGALAMAN KERJA")
		for i := 0; i < d.nExPengalaman; i++ {
			fmt.Printf("	%s — %s · %s | %s\n", d.exPosisi[i], d.exPerusaha[i], d.exKota[i], d.exTgl[i])
			fmt.Printf("	• %s \n", d.exPengalaman[i])
		}
	}

	if d.nPenJur > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	PENDIDIKAN")
		for i := 0; i < d.nPenJur; i++ {
			fmt.Printf("	%s — %s | %d\n", d.penJur[i], d.uni[i], d.tLulus[i])
			fmt.Printf("	IPK: %.2f | Aktivitas/kepanitiaan relevan: %s\n", d.ipk[i], d.aktivy[i])
		}
	}

	if d.nSkill > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	SKILL UTAMA")
		for i := 0; i < d.nSkill; i++ {
			fmt.Printf("	• %s\n", d.Skill[i])
		}
	}

	if d.nSerti > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	SERTIFIKASI & KURSUS")
		for i := 0; i < d.nSerti; i++ {
			fmt.Printf("	• %s, %s %d\n", d.serti[i], d.penye[i], d.tSerti[i])
		}
	}

	if d.nBhs > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	BAHASA")
		for i := 0; i < d.nBhs; i++ {
			fmt.Printf("	• %s ", d.bhs[i])
		}
		fmt.Println()
	}

	if d.nHobi > 0 {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("	MINAT")
		for i := 0; i < d.nHobi; i++ {
			fmt.Printf("	%s", d.hobi[i])
			if i < d.nHobi-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}

	fmt.Println("===============================================================")
}

func cetakSurat(d data){
    fmt.Printf("%s, %s\n", d.kota, d.tglnow)
    fmt.Println()

    fmt.Println("Kepada Yth.")
    fmt.Printf("Tim Rekrutmen %s\n", d.perstuj)
    fmt.Println()

    fmt.Printf("Perihal: Lamaran Posisi %s\n", d.posisi)
    fmt.Println()

    fmt.Println("Salam,")
    fmt.Println()

    fmt.Printf("Perkenalkan, saya %s.\n", d.nama)
    fmt.Printf("Saya mengetahui lowongan %s di Internet dan tertarik karena selaras dengan skill dan rencana pengembangan karier yang sedang aku jalani.\n", d.posisi)//butuh Input
    fmt.Println()

	if d.nExPengalaman > 0 {
		fmt.Printf("Selama beberapa tahun terakhir, saya banyak berkutat di %s.\n", d.exPosisi[0])

		fmt.Println("Beberapa pencapaian yang relevan:")
		max := d.nExPengalaman
		if max > 3 {
			max = 3
		}
		for i := 0; i < max; i++ {
			fmt.Printf("• %s\n", d.exPengalaman[i])
		}
	} else {
		fmt.Println("Selama beberapa tahun terakhir, saya terus mengembangkan keahlian yang relevan dengan posisi ini.")
		fmt.Println("(belum ada pencapaian tercatat)")
	}
	fmt.Println()

    fmt.Printf("Pengalaman tersebut membentuk skill inti di %s,\n", d.Skill[0])
    fmt.Printf("yang saya yakin bisa membantu %s meraih target dengan cara yang lebih cepat dan efisien.\n", d.perstuj)
    fmt.Println()

    fmt.Println("Besar harapan saya untuk berdiskusi lebih lanjut.")
    fmt.Printf("CV dan portofolio terlampir—silakan hubungi saya di %s atau %s kapan pun.\n", d.noPhone, d.email)
    fmt.Println("Terima kasih atas waktunya!")
    fmt.Println()

    fmt.Println("Hormat saya,")
    fmt.Println()
    fmt.Printf("%s\n", d.nama)
    fmt.Println()
}