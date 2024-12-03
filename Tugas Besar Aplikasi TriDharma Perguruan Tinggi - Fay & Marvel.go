package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const nmax int = 2048
const maxanggota int = 4

type data_penelitian struct {
	judul, ketua, sumber_dana, prodi, fakultas string
	anggota                                    [maxanggota]string
	publikasi, produk, tahun_kegiatan          int
}
type data_abdimas struct {
	judul, ketua, sumber_dana, prodi, fakultas            string
	anggota                                               [maxanggota]string
	publikasi, produk, seminar, pelatihan, tahun_kegiatan int
}
type tabdata struct {
	penelitian            [nmax]data_penelitian
	abdimas               [nmax]data_abdimas
	npenelitian, nabdimas int
}

func main() {
	var x tabdata
	datadami_abdimas(&x)
	datadatami_penelitian(&x)
	menu(x)

}

func header1() {
	fmt.Println("-----------------------------------")
	fmt.Println("|             Aplikasi            |")
	fmt.Println("|   Tri-Darma Perguruan Tinggi 2  |")
	fmt.Println("|        By Marvel & Fayza        |")
	fmt.Println("-----------------------------------")
	fmt.Printf("\n%s, Selamat datang!\n", ucapanselamat())
}
func header2() {
	fmt.Println("\n----------------------------------")
	fmt.Println("|Pilih Data yang ingin anda akses|")
	fmt.Println("|--------------------------------|")
	fmt.Println("|1. Data penelitian              |")
	fmt.Println("|2. Data pengabdian masyarakat   |")
	fmt.Println("|3. Keluar                       |")
	fmt.Println("----------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header3() {
	fmt.Println("\n---------------------------------------")
	fmt.Println("|---------------ERROR-----------------|")
	fmt.Println("|----Masukkan pilihan yang benar !----|")
	fmt.Println("---------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header4() {
	fmt.Println("---------------------------------------")
	fmt.Println("|          1. Menambahkan data        |")
	fmt.Println("|          2. Mengubah data           |")
	fmt.Println("|          3. Menghapus data          |")
	fmt.Println("|          4. Menampilkan data        |")
	fmt.Println("|          5. Mengurutkan data        |")
	fmt.Println("|          6. kembali                 |")
	fmt.Println("---------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header5() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|               1. Urutkan dari terlama                |")
	fmt.Println("|               2. Urutkan dari terbaru                |")
	fmt.Println("|               3. Kembali                             |")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header6() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("|TERIMA KASIH TELAH MENGGUNAKAN APLIKASI KAMI !!!|")
	fmt.Println("--------------------------------------------------")
}
func header7() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|               1. Berdasarkan Tahun                  |")
	fmt.Println("|               2. Bedasarkan fakultas/prodi          |")
	fmt.Println("|               3. Semua                              |")
	fmt.Println("|               4. Kembali                            |")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header8() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|                     1. Prodi                        |")
	fmt.Println("|                     2. Fakultas                     |")
	fmt.Println("|                     3. Kembali                      |")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header9() {
	fmt.Println("-----------Data Penelitian-------------")
}
func header10() {
	fmt.Println("--------------Data Abdimas-------------")
}
func header11() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|         1. Berdasarkan Tahun                        |")
	fmt.Println("|         2. Berdasarkan Total Kegiatan per Tahun     |")
	fmt.Println("|         3. Kembali                                  |")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func header12() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println("|             1. Urutkan dari terbanyak               |")
	fmt.Println("|             2. Urutkan dari tersedikit              |")
	fmt.Println("|             3. Kembali                              |")
	fmt.Println("-------------------------------------------------------")
	fmt.Print("\nMasukkan pilihan anda: ")
}
func ucapanselamat() string {
	jam := time.Now().Hour()
	var ucapan string
	if jam < 12 {
		ucapan = "Selamat pagi"
	} else if jam < 18 {
		ucapan = "Selamat siang"
	} else {
		ucapan = "Selamat malam"
	}
	return ucapan
}

func datadatami_penelitian(x *tabdata) {
	x.penelitian[0].judul = "Olahan Limbah Kulit Pisang Sebagai Keripik Bergizi"
	x.penelitian[0].ketua = "Ali Haidar"
	x.penelitian[0].anggota[0] = "Ahmad Fahrurroza"
	x.penelitian[0].anggota[1] = "Abdul Hanafi"
	x.penelitian[0].anggota[2] = "Arafah Desianti"
	x.penelitian[0].anggota[3] = "Siti Maimunah"
	x.penelitian[0].prodi = "informatika"
	x.penelitian[0].fakultas = "FIF"
	x.penelitian[0].sumber_dana = "internal"
	x.penelitian[0].publikasi = 25000
	x.penelitian[0].produk = 50000
	x.penelitian[0].tahun_kegiatan = 2022
	x.penelitian[1].judul = "Permen Bonggol Nanas"
	x.penelitian[1].ketua = "Yusuf Mutholib"
	x.penelitian[1].anggota[0] = "Mukhlis Wahab"
	x.penelitian[1].anggota[1] = "Abdul Qadir"
	x.penelitian[1].anggota[2] = "Serina"
	x.penelitian[1].anggota[3] = "Mandraguna"
	x.penelitian[1].prodi = "sains data"
	x.penelitian[1].fakultas = "FIF"
	x.penelitian[1].sumber_dana = "internal"
	x.penelitian[1].publikasi = 20000
	x.penelitian[1].produk = 75000
	x.penelitian[1].tahun_kegiatan = 2021
	x.penelitian[2].judul = "Mesin Cuci Ramah Lingkungan"
	x.penelitian[2].ketua = "Gema Barokah"
	x.penelitian[2].anggota[0] = "Candil"
	x.penelitian[2].anggota[1] = " Dorothy Sheinafia"
	x.penelitian[2].anggota[2] = "-"
	x.penelitian[2].anggota[3] = "-"
	x.penelitian[2].prodi = "seni rupa"
	x.penelitian[2].fakultas = "FIK"
	x.penelitian[2].sumber_dana = "eksternal"
	x.penelitian[2].publikasi = 30000
	x.penelitian[2].produk = 45000
	x.penelitian[2].tahun_kegiatan = 2023
	x.penelitian[3].judul = "Pengembangan Karakter Anak Usia Dini"
	x.penelitian[3].ketua = "Elgato Munyanyo"
	x.penelitian[3].anggota[0] = "Theodore hanafi"
	x.penelitian[3].anggota[1] = "-"
	x.penelitian[3].anggota[2] = "-"
	x.penelitian[3].anggota[3] = "-"
	x.penelitian[3].prodi = "akuntansi"
	x.penelitian[3].fakultas = "FEB"
	x.penelitian[3].sumber_dana = "internal"
	x.penelitian[3].publikasi = 50000
	x.penelitian[3].produk = 50000
	x.penelitian[3].tahun_kegiatan = 2020
	x.penelitian[4].judul = "Dompet_pintar"
	x.penelitian[4].ketua = "Elgato Munyanyo"
	x.penelitian[4].anggota[0] = "Theodore hanafi"
	x.penelitian[4].anggota[1] = "Munaroh Anais"
	x.penelitian[4].anggota[2] = "Diana Alexandra"
	x.penelitian[4].anggota[3] = "-"
	x.penelitian[4].prodi = "ilmu komunikasi"
	x.penelitian[4].fakultas = "FKB"
	x.penelitian[4].sumber_dana = "internal"
	x.penelitian[4].publikasi = 15000
	x.penelitian[4].produk = 35000
	x.penelitian[4].tahun_kegiatan = 2022
	x.npenelitian = 5
}
func datadami_abdimas(y *tabdata) {
	y.abdimas[0].judul = "Pelatihan Pengelolaan SDM di Era Digital"
	y.abdimas[0].ketua = "Taufik Hidayat"
	y.abdimas[0].anggota[0] = "Rahmat Gobel"
	y.abdimas[0].anggota[1] = "Ocid"
	y.abdimas[0].anggota[2] = "Alexandra"
	y.abdimas[0].anggota[3] = "-"
	y.abdimas[0].prodi = "dkv"
	y.abdimas[0].fakultas = "FIK"
	y.abdimas[0].sumber_dana = "internal"
	y.abdimas[0].publikasi = 15000
	y.abdimas[0].produk = 35000
	y.abdimas[0].seminar = 20000
	y.abdimas[0].pelatihan = 10000
	y.abdimas[0].tahun_kegiatan = 2022
	y.abdimas[1].judul = "Pelatihan Pemasaran & Distibusi di Masa Covid-19"
	y.abdimas[1].ketua = "David Guetta"
	y.abdimas[1].anggota[0] = "Okta Prima"
	y.abdimas[1].anggota[1] = "Julian Draxler"
	y.abdimas[1].anggota[2] = "-"
	y.abdimas[1].anggota[3] = "-"
	y.abdimas[1].prodi = "sistem informasi"
	y.abdimas[1].fakultas = "FIT"
	y.abdimas[1].sumber_dana = "internal"
	y.abdimas[1].publikasi = 17500
	y.abdimas[1].produk = 70000
	y.abdimas[1].seminar = 25000
	y.abdimas[1].pelatihan = 15000
	y.abdimas[1].tahun_kegiatan = 2021
	y.abdimas[2].judul = "Emas sebagai Salah Satu Alternatif Menyimpan Uang"
	y.abdimas[2].ketua = "Theo Hernandrz"
	y.abdimas[2].anggota[0] = "Raisa"
	y.abdimas[2].anggota[1] = "Hamish Daud"
	y.abdimas[2].anggota[2] = "-"
	y.abdimas[2].anggota[3] = "-"
	y.abdimas[2].prodi = "teknik industri"
	y.abdimas[2].fakultas = "FRI"
	y.abdimas[2].sumber_dana = "eksternal"
	y.abdimas[2].publikasi = 33000
	y.abdimas[2].produk = 44000
	y.abdimas[2].seminar = 22000
	y.abdimas[2].pelatihan = 11000
	y.abdimas[2].tahun_kegiatan = 2023
	y.abdimas[3].judul = "Perancangan Aset Konten Digital untuk Kegiatan Promosi Agrowisata"
	y.abdimas[3].ketua = "Keiko"
	y.abdimas[3].anggota[0] = "Andra"
	y.abdimas[3].anggota[1] = "-"
	y.abdimas[3].anggota[2] = "-"
	y.abdimas[3].anggota[3] = "-"
	y.abdimas[3].prodi = "teknik elektro"
	y.abdimas[3].fakultas = "FTE"
	y.abdimas[3].sumber_dana = "eksternal"
	y.abdimas[3].publikasi = 50000
	y.abdimas[3].produk = 50000
	y.abdimas[3].seminar = 20000
	y.abdimas[3].pelatihan = 10000
	y.abdimas[3].tahun_kegiatan = 2020
	y.abdimas[4].judul = "Sosialisasi COVID19 melalui Media Komik untuk Anak Sekolah Dasar"
	y.abdimas[4].ketua = "Shinobu"
	y.abdimas[4].anggota[0] = "Kanao"
	y.abdimas[4].anggota[1] = "Tanjiro"
	y.abdimas[4].anggota[2] = "Nezuko"
	y.abdimas[4].anggota[3] = "Zenitsu"
	y.abdimas[4].prodi = "ilmu komunikasi"
	y.abdimas[4].fakultas = "FKB"
	y.abdimas[4].sumber_dana = "internal"
	y.abdimas[4].publikasi = 150600
	y.abdimas[4].produk = 356000
	y.abdimas[4].seminar = 210000
	y.abdimas[4].pelatihan = 120000
	y.abdimas[4].tahun_kegiatan = 2023
	y.nabdimas = 5
}
func menu(x tabdata) {
	var i, pilihan int
	header1()
	header2()
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 && pilihan != 3 {
		header3()
		header2()
		fmt.Scan(&pilihan)
		i++
	}
	if pilihan == 1 {
		choicedata_penelitian(&x)
	} else if pilihan == 2 {
		choicedata_abdimas(&x)
	} else {
		header6()
	}
}

func choicedata_penelitian(a *tabdata) {
	var i, pilihan int
	header9()
	header4()
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 && pilihan != 5 && pilihan != 6 {
		header3()
		fmt.Scan(&pilihan)
		i++
	}
	InsertionsortTahunPenelitian(a)
	if pilihan == 1 {
		add_datapenelitian(a)
	} else if pilihan == 2 {
		editdatapenelitian(a)
	} else if pilihan == 3 {
		removedata_penelitian(a)
		choicedata_penelitian(a)
	} else if pilihan == 4 {
		header7()
		fmt.Scan(&pilihan)
		for pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 {
			header3()
			fmt.Scan(&pilihan)
			i++
		}
		if pilihan == 1 {
			printdatapenelitian_tahun(a)
		} else if pilihan == 2 {
			header8()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				printprodipeneletian(a)
			} else if pilihan == 2 {
				printfakultaspeneletian(a)
			} else {
				choicedata_penelitian(a)
			}

		} else if pilihan == 3 {
			printsemua_datapenelitian(a)
		} else {
			choicedata_penelitian(a)
		}
	} else if pilihan == 5 {
		header11()
		fmt.Scan(&pilihan)
		for pilihan != 1 && pilihan != 2 && pilihan != 3 {
			header3()
			fmt.Scan(&pilihan)
			i++
		}
		if pilihan == 1 {
			header5()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				InsertionsortTahunPenelitian(a)
				printsemua_datapenelitian(a)
			} else if pilihan == 2 {
				SelectionTahunPenelitian(a)
				printsemua_datapenelitian(a)
			} else {
				choicedata_penelitian(a)
			}
		} else if pilihan == 2 {
			header12()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				sortdescendbytotalkegiatanpertahunpenelitian(a)
			} else if pilihan == 2 {
				sortascnytotalkegiatanpertahunpenelitian(a)
			} else {
				choicedata_penelitian(a)
			}
		} else {
			choicedata_penelitian(a)
		}
	} else if pilihan == 6 {
		menu(*a)
	}
}
func choicedata_abdimas(a *tabdata) {
	var i, pilihan int
	header10()
	header4()
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 && pilihan != 5 && pilihan != 6 {
		header3()
		fmt.Scan(&pilihan)
		i++
	}
	InsertionsortTahunAbdimas(a)
	if pilihan == 1 {
		add_abdimas(a)
	} else if pilihan == 2 {
		editdataabdimas(a)
	} else if pilihan == 3 {
		removedata_abdimas(a)
		choicedata_abdimas(a)
	} else if pilihan == 4 {
		header7()
		fmt.Scan(&pilihan)
		for pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 {
			header3()
			fmt.Scan(&pilihan)
			i++
		}
		if pilihan == 1 {
			printdataabdimas_tahun(a)
		} else if pilihan == 2 {
			header8()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				printprodiabdimas(a)
			} else if pilihan == 2 {
				printfakultasabdimas(a)
			} else {
				choicedata_abdimas(a)
			}

		} else if pilihan == 3 {
			printsemua_dataabdimas(a)
		} else {
			choicedata_abdimas(a)
		}
	} else if pilihan == 5 {
		header11()
		fmt.Scan(&pilihan)
		for pilihan != 1 && pilihan != 2 && pilihan != 3 {
			header3()
			fmt.Scan(&pilihan)
			i++
		}
		if pilihan == 1 {
			header5()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				InsertionsortTahunAbdimas(a)
				printsemua_dataabdimas(a)
			} else if pilihan == 2 {
				SelectionsortTahunAbdimas(a)
				printsemua_dataabdimas(a)
			} else {
				choicedata_abdimas(a)
			}
		} else if pilihan == 2 {
			header12()
			fmt.Scan(&pilihan)
			for pilihan != 1 && pilihan != 2 && pilihan != 3 {
				header3()
				fmt.Scan(&pilihan)
				i++
			}
			if pilihan == 1 {
				sortdescendbytotalkegiatanpertahunabdimas(a)
			} else if pilihan == 2 {
				sortascnytotalkegiatanpertahunabdimas(a)
			} else {
				choicedata_abdimas(a)
			}
		} else {
			choicedata_abdimas(a)
		}
	} else if pilihan == 6 {
		menu(*a)
	}
}
func add_datapenelitian(a *tabdata) {
	var i int
	fmt.Print("\nJudul: ")
	error_input_managementString(&a.penelitian[a.npenelitian].judul)
	error_input_managementString(&a.penelitian[a.npenelitian].judul)
	fmt.Print("Ketua: ")
	error_input_managementString(&a.penelitian[a.npenelitian].ketua)
	fmt.Println("{JIKA ANGGOTA TIDAK ADA, ISI DENGAN LAMBANG - }")
	for i = 0; i < maxanggota; i++ {
		fmt.Print("Anggota ", i+1, ": ")
		error_input_managementString(&a.penelitian[a.npenelitian].anggota[i])
		for a.penelitian[a.npenelitian].anggota[i] == "" {
			fmt.Println("TIDAK BOLEH KOSONG")
			fmt.Print("Anggota ", i+1, ": ")
			error_input_managementString(&a.penelitian[a.npenelitian].anggota[i])
		}
	}
	fmt.Print("Prodi: ")
	error_input_managementString(&a.penelitian[a.npenelitian].prodi)
	fmt.Println("{Gunakan huruf kapital semua. Contoh fakultas informatika = FIF}")
	fmt.Print("Fakultas: ")
	error_input_managementstringUpper(&a.penelitian[a.npenelitian].fakultas)
	fmt.Println("{Format penulisan sumber dana adalah internal atau eksternal}")
	fmt.Print("Sumber dana: ")
	error_input_managementString(&a.penelitian[a.npenelitian].sumber_dana)
	for a.penelitian[a.npenelitian].sumber_dana != "internal" && a.penelitian[a.npenelitian].sumber_dana != "eksternal" {
		fmt.Print("Sumber dana: ")
		error_input_managementString(&a.penelitian[a.npenelitian].sumber_dana)
		i++
	}
	fmt.Print("Iuran Publikasi: ")
	error_input_managementINT(&a.penelitian[a.npenelitian].publikasi)
	fmt.Print("Iuran produk: ")
	error_input_managementINT(&a.penelitian[a.npenelitian].produk)
	fmt.Print("Tahun kegiatan: ")
	error_input_managementINT(&a.penelitian[a.npenelitian].tahun_kegiatan)
	a.npenelitian++
	fmt.Println("\nData berhasil ditambahkan")
	choicedata_penelitian(a)
}

func add_abdimas(a *tabdata) {
	var i int
	fmt.Print("\nJudul: ")
	error_input_managementString(&a.abdimas[a.nabdimas].judul)
	error_input_managementString(&a.abdimas[a.nabdimas].judul)
	fmt.Print("Ketua: ")
	error_input_managementString(&a.abdimas[a.nabdimas].ketua)
	fmt.Println("{JIKA ANGGOTA TIDAK ADA, ISI DENGAN LAMBANG - }")
	for i = 0; i < maxanggota; i++ {
		fmt.Print("Anggota ", i+1, ": ")
		error_input_managementString(&a.abdimas[a.nabdimas].anggota[i])
		for a.abdimas[a.nabdimas].anggota[i] == "" {
			fmt.Println("TIDAK BOLEH KOSONG")
			fmt.Print("Anggota ", i+1, ": ")
			error_input_managementString(&a.abdimas[a.nabdimas].anggota[i])
		}
	}
	fmt.Print("Prodi: ")
	error_input_managementString(&a.abdimas[a.nabdimas].prodi)
	fmt.Println("{Gunakan huruf kapital semua. Contoh fakultas informatika = FIF}")
	fmt.Print("Fakultas: ")
	error_input_managementstringUpper(&a.abdimas[a.nabdimas].fakultas)
	fmt.Println("{Format penulisan sumber dana adalah internal atau eksternal}")
	fmt.Print("Sumber dana: ")
	error_input_managementString(&a.abdimas[a.nabdimas].sumber_dana)
	for a.abdimas[a.nabdimas].sumber_dana != "internal" && a.abdimas[a.nabdimas].sumber_dana != "eksternal" {
		fmt.Print("Sumber dana: ")
		error_input_managementString(&a.abdimas[a.nabdimas].sumber_dana)
		i++
	}
	fmt.Print("Iuran Publikasi: ")
	error_input_managementINT(&a.abdimas[a.nabdimas].publikasi)
	fmt.Print("Iuran produk: ")
	error_input_managementINT(&a.abdimas[a.nabdimas].produk)
	fmt.Print("Tahun kegiatan: ")
	error_input_managementINT(&a.abdimas[a.nabdimas].tahun_kegiatan)
	a.nabdimas++
	fmt.Println("\nData berhasil ditambahkan")
	choicedata_abdimas(a)
}
func editdatapenelitian(a *tabdata) {
	var i, j int
	var judul string
	printjuduldatapenelitian(a)
	fmt.Print("\nMasukkan judul data yang ingin diubah: ")
	error_input_managementString(&judul)
	error_input_managementString(&judul)
	//sequentialsearch
	for i = 0; i < a.npenelitian; i++ {
		if a.penelitian[i].judul == judul {
			fmt.Printf("\nData penelitian ditemukan. Silakan masukkan data baru.\n")
			fmt.Print("Judul: ")
			error_input_managementString(&a.penelitian[i].judul)
			fmt.Print("Ketua: ")
			error_input_managementString(&a.penelitian[i].ketua)
			fmt.Println("{JIKA ANGGOTA TIDAK ADA, ISI DENGAN LAMBANG - }")
			for j = 0; j < maxanggota; j++ {
				fmt.Print("Anggota ", j+1, ": ")
				error_input_managementString(&a.penelitian[i].anggota[j])
			}
			error_input_managementString(&a.penelitian[i].prodi)
			fmt.Println("{Gunakan huruf kapital semua. Contoh fakultas informatika = FIF}")
			fmt.Print("Fakultas: ")
			error_input_managementstringUpper(&a.penelitian[i].fakultas)
			fmt.Println("{Format penulisan sumber dana adalah internal atau eksternal}")
			fmt.Print("Sumber dana: ")
			error_input_managementString(&a.penelitian[i].sumber_dana)
			for a.penelitian[i].sumber_dana != "internal" && a.penelitian[i].sumber_dana != "eksternal" {
				fmt.Print("Sumber dana: ")
				error_input_managementString(&a.penelitian[i].sumber_dana)
				j++
			}
			fmt.Print("Iuran Publikasi: ")
			error_input_managementINT(&a.penelitian[i].publikasi)
			fmt.Print("Iuran produk: ")
			error_input_managementINT(&a.penelitian[i].produk)
			fmt.Print("Tahun kegiatan: ")
			error_input_managementINT(&a.penelitian[i].tahun_kegiatan)
			fmt.Println("\nData penelitian berhasil diubah.")
		}
	}
	fmt.Println("\nData tidak ditemukan")
	choicedata_penelitian(a)
}
func editdataabdimas(a *tabdata) {
	var i, j int
	var judul string
	printjuduldataabdimas(a)
	fmt.Print("\nMasukkan judul data yang ingin diubah: ")
	error_input_managementString(&judul)
	error_input_managementString(&judul)
	//sequentialsearch
	for i = 0; i < a.nabdimas; i++ {
		if a.abdimas[i].judul == judul {
			fmt.Printf("\nData penelitian ditemukan. Silakan masukkan data baru.\n")
			fmt.Print("Judul: ")
			error_input_managementString(&a.abdimas[i].judul)
			fmt.Print("Ketua: ")
			error_input_managementString(&a.abdimas[i].ketua)
			fmt.Println("{JIKA ANGGOTA TIDAK ADA, ISI DENGAN LAMBANG - }")
			for j = 0; j < maxanggota; j++ {
				fmt.Print("Anggota ", j+1, ": ")
				error_input_managementString(&a.abdimas[i].anggota[j])
			}
			error_input_managementString(&a.abdimas[i].prodi)
			fmt.Println("{Gunakan huruf kapital semua. Contoh fakultas informatika = FIF}")
			fmt.Print("Fakultas: ")
			error_input_managementstringUpper(&a.abdimas[i].fakultas)
			fmt.Println("{Format penulisan sumber dana adalah internal atau eksternal}")
			fmt.Print("Sumber dana: ")
			error_input_managementString(&a.abdimas[i].sumber_dana)
			for a.penelitian[i].sumber_dana != "internal" && a.abdimas[i].sumber_dana != "eksternal" {
				fmt.Print("Sumber dana: ")
				error_input_managementString(&a.abdimas[i].sumber_dana)
				j++
			}
			fmt.Print("Iuran Publikasi: ")
			error_input_managementINT(&a.abdimas[i].publikasi)
			fmt.Print("Iuran produk: ")
			error_input_managementINT(&a.abdimas[i].produk)
			fmt.Print("Iuran seminar: ")
			error_input_managementINT(&a.abdimas[*&a.nabdimas].seminar)
			fmt.Print("Iuran pelatihan: ")
			error_input_managementINT(&a.abdimas[*&a.nabdimas].pelatihan)
			fmt.Print("Tahun kegiatan: ")
			error_input_managementINT(&a.abdimas[*&a.nabdimas].tahun_kegiatan)
			fmt.Println("\nData penelitian berhasil diubah.")
		}
	}
	fmt.Println("\nData tidak ditemukan")
	choicedata_abdimas(a)
}
func removedata_penelitian(a *tabdata) {
	var judul string
	var i, j int
	fmt.Println("")
	printjuduldatapenelitian(a)
	fmt.Print("\nMasukkan judul data penelitian yang ingin dihapus: ")
	error_input_managementString(&judul)
	error_input_managementString(&judul)
	//sequentialsearch
	for i = 0; i < a.npenelitian; i++ {
		if a.penelitian[i].judul == judul {
			for j = i; j < a.npenelitian-1; j++ {
				a.penelitian[j] = a.penelitian[j+1]
			}
			a.npenelitian--
			fmt.Println("\nData penelitian berhasil dihapus.")
			return
		}
	}
	fmt.Println("\nData penelitian tidak ditemukan.")

}
func removedata_abdimas(a *tabdata) {
	var judul string
	var i, j int
	fmt.Println("")
	printjuduldataabdimas(a)
	fmt.Print("\nMasukkan judul data penelitian yang ingin dihapus: ")
	error_input_managementString(&judul)
	error_input_managementString(&judul)
	//sequentialsearch
	for i = 0; i < a.nabdimas; i++ {
		if a.abdimas[i].judul == judul {
			for j = i; j < a.nabdimas-1; j++ {
				a.abdimas[j] = a.abdimas[j+1]
			}
			a.nabdimas--
			fmt.Println("\nData penelitian berhasil dihapus.")
			return
		}
	}
	fmt.Println("\nData penelitian tidak ditemukan.")

}
func printjuduldatapenelitian(x *tabdata) {
	var i int
	InsertionsortTahunPenelitian(x)
	for i = 0; i < x.npenelitian; i++ {
		fmt.Println("Judul:", x.penelitian[i].judul)
	}
}
func printjuduldataabdimas(x *tabdata) {
	var i int
	InsertionsortTahunAbdimas(x)
	for i = 0; i < x.nabdimas; i++ {
		fmt.Println("Judul:", x.abdimas[i].judul)
	}
}
func printsemua_datapenelitian(x *tabdata) {
	var i, j int
	for i = 0; i < x.npenelitian; i++ {
		fmt.Println("Judul:", x.penelitian[i].judul)
		fmt.Println("Ketua:", x.penelitian[i].ketua)
		for j = 0; j < maxanggota; j++ {
			if x.penelitian[i].anggota[j] != "-" {
				fmt.Println("anggota:", x.penelitian[i].anggota[j])
			}
		}
		fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
		fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
		fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
		fmt.Println("Iuran produk:", x.penelitian[i].produk)
		fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
		fmt.Println("___________________________________________________________________________")
	}
	choicedata_penelitian(x)
}
func printsemua_dataabdimas(x *tabdata) {
	var i, j int
	for i = 0; i < x.nabdimas; i++ {
		fmt.Println("Judul:", x.abdimas[i].judul)
		fmt.Println("Ketua:", x.abdimas[i].ketua)
		for j = 0; j < maxanggota; j++ {
			if x.abdimas[i].anggota[j] != "-" {
				fmt.Println("anggota:", x.abdimas[i].anggota[j])
			}
		}
		fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
		fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
		fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
		fmt.Println("Iuran produk:", x.abdimas[i].produk)
		fmt.Println("Tahun kegiatan:", x.abdimas[i].seminar)
		fmt.Println("Iuran publikasi:", x.abdimas[i].pelatihan)
		fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
		fmt.Println("___________________________________________________________________________")
	}
	choicedata_abdimas(x)
}
func printprodipeneletian(x *tabdata) {
	var found, left, right, med, j, i int
	var check bool = false
	var prodi string
	found = -1
	left = 0
	right = x.npenelitian - 1
	fmt.Print("\nMasukkan prodi: ")
	error_input_managementString(&prodi)
	error_input_managementString(&prodi)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	descsortprodipenelitian(x)
	for left <= right && found == -1 {
		med = (left + right) / 2
		if prodi > x.penelitian[med].prodi {
			right = med - 1
		} else if prodi < x.penelitian[med].prodi {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.npenelitian; i++ {
			if x.penelitian[i].prodi == x.penelitian[found].prodi {
				fmt.Println("Judul:", x.penelitian[i].judul)
				fmt.Println("Ketua:", x.penelitian[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.penelitian[found].anggota[j] != "-" {
						fmt.Println("anggota:", x.penelitian[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
				fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
				fmt.Println("Iuran produk:", x.penelitian[i].produk)
				fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_penelitian(x)
}
func printprodiabdimas(x *tabdata) {
	var found, left, right, med, j, i int
	var check bool = false
	var prodi string
	found = -1
	left = 0
	right = x.nabdimas - 1
	fmt.Print("\nMasukkan prodi: ")
	error_input_managementString(&prodi)
	error_input_managementString(&prodi)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	descsortprodiabdimas(x)
	for left <= right && found == -1 {
		med = (left + right) / 2
		if prodi > x.abdimas[med].prodi {
			right = med - 1
		} else if prodi < x.abdimas[med].prodi {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.nabdimas; i++ {
			if x.abdimas[i].prodi == x.abdimas[found].prodi {
				fmt.Println("Judul:", x.abdimas[i].judul)
				fmt.Println("Ketua:", x.abdimas[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.abdimas[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.abdimas[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
				fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
				fmt.Println("Iuran produk:", x.abdimas[i].produk)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].seminar)
				fmt.Println("Iuran publikasi:", x.abdimas[i].pelatihan)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_abdimas(x)
}
func printfakultaspeneletian(x *tabdata) {
	var found, left, right, med, j, i int
	var check bool = false
	var fakultas string
	found = -1
	left = 0
	right = x.npenelitian - 1
	fmt.Print("\nMasukkan fakultas: ")
	error_input_managementString(&fakultas)
	error_input_managementString(&fakultas)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	descsortfakultaspenelitian(x)
	for left <= right && found == -1 {
		med = (left + right) / 2
		if fakultas > x.penelitian[med].fakultas {
			right = med - 1
		} else if fakultas < x.penelitian[med].fakultas {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.npenelitian; i++ {
			if x.penelitian[i].fakultas == x.penelitian[found].fakultas {
				fmt.Println("Judul:", x.penelitian[i].judul)
				fmt.Println("Ketua:", x.penelitian[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.penelitian[found].anggota[j] != "-" {
						fmt.Println("anggota:", x.penelitian[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
				fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
				fmt.Println("Iuran produk:", x.penelitian[i].produk)
				fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_penelitian(x)
}
func printfakultasabdimas(x *tabdata) {
	var found, left, right, med, j, i int
	var check bool = false
	var fakultas string
	found = -1
	left = 0
	right = x.nabdimas - 1
	fmt.Print("\nMasukkan fakultas: ")
	error_input_managementString(&fakultas)
	error_input_managementString(&fakultas)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	descsortfakultasabdimas(x)
	for left <= right && found == -1 {
		med = (left + right) / 2
		if fakultas > x.abdimas[med].fakultas {
			right = med - 1
		} else if fakultas < x.abdimas[med].fakultas {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.nabdimas; i++ {
			if x.abdimas[i].fakultas == x.abdimas[found].fakultas {
				fmt.Println("Judul:", x.abdimas[i].judul)
				fmt.Println("Ketua:", x.abdimas[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.abdimas[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.abdimas[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
				fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
				fmt.Println("Iuran produk:", x.abdimas[i].produk)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].seminar)
				fmt.Println("Iuran publikasi:", x.abdimas[i].pelatihan)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_abdimas(x)
}
func printdatapenelitian_tahun(x *tabdata) {
	var found, left, right, med, tahun, j, i int
	var check bool = false
	found = -1
	left = 0
	right = x.npenelitian - 1
	fmt.Println("\nMasukkan Tahun: ")
	error_input_managementINT(&tahun)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	for left <= right && found == -1 {
		med = (left + right) / 2
		if tahun > x.penelitian[med].tahun_kegiatan {
			right = med - 1
		} else if tahun < x.penelitian[med].tahun_kegiatan {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.npenelitian; i++ {
			if x.penelitian[i].tahun_kegiatan == x.penelitian[found].tahun_kegiatan {
				fmt.Println("Judul:", x.penelitian[i].judul)
				fmt.Println("Ketua:", x.penelitian[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.penelitian[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.penelitian[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
				fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
				fmt.Println("Iuran produk:", x.penelitian[i].produk)
				fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_penelitian(x)
}
func printdataabdimas_tahun(x *tabdata) {
	var found, left, right, med, tahun, j, i int
	var check bool = false
	found = -1
	left = 0
	right = x.nabdimas - 1
	fmt.Println("\nMasukkan Tahun: ")
	error_input_managementINT(&tahun)
	fmt.Println("___________________________________________________________________________")
	//binarysearch
	for left <= right && found == -1 {
		med = (left + right) / 2
		if tahun > x.abdimas[med].tahun_kegiatan {
			right = med - 1
		} else if tahun < x.abdimas[med].tahun_kegiatan {
			left = med + 1
		} else {
			found = med
			check = true
		}
	}
	if check == true {
		for i = 0; i < x.nabdimas; i++ {
			if x.abdimas[i].tahun_kegiatan == x.abdimas[found].tahun_kegiatan {
				fmt.Println("Judul:", x.abdimas[i].judul)
				fmt.Println("Ketua:", x.abdimas[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.abdimas[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.abdimas[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
				fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
				fmt.Println("Iuran produk:", x.abdimas[i].produk)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	} else {
		fmt.Println("\nData tidak ditemukan")
	}
	choicedata_abdimas(x)
}
func SelectionTahunPenelitian(T *tabdata) {
	var pass, idx, i int
	var temp data_penelitian
	pass = 1
	for pass <= T.npenelitian-1 {
		idx = pass - 1
		i = pass
		for i < T.npenelitian {
			if T.penelitian[idx].tahun_kegiatan > T.penelitian[i].tahun_kegiatan {
				idx = i
			}
			i++
		}
		temp = T.penelitian[idx]
		T.penelitian[idx] = T.penelitian[pass-1]
		T.penelitian[pass-1] = temp
		pass++
	}
}
func SelectionsortTahunAbdimas(T *tabdata) {
	var pass, idx, i int
	var temp data_abdimas
	pass = 1
	for pass <= T.nabdimas-1 {
		idx = pass - 1
		i = pass
		for i < T.nabdimas {
			if T.abdimas[idx].tahun_kegiatan > T.abdimas[i].tahun_kegiatan {
				idx = i
			}
			i++
		}
		temp = T.abdimas[idx]
		T.abdimas[idx] = T.abdimas[pass-1]
		T.abdimas[pass-1] = temp
		pass++
	}
}
func InsertionsortTahunPenelitian(T *tabdata) {
	var pass, i int
	var temp data_penelitian
	pass = 1
	for pass <= T.npenelitian-1 {
		i = pass
		temp = T.penelitian[pass]
		for i > 0 && (temp.tahun_kegiatan > T.penelitian[i-1].tahun_kegiatan) {
			T.penelitian[i] = T.penelitian[i-1]
			i--
		}
		T.penelitian[i] = temp
		pass++
	}
}
func InsertionsortTahunAbdimas(T *tabdata) {
	var pass, i int
	var temp data_abdimas
	pass = 1
	for pass <= T.nabdimas-1 {
		i = pass
		temp = T.abdimas[pass]
		for i > 0 && (temp.tahun_kegiatan > T.abdimas[i-1].tahun_kegiatan) {
			T.abdimas[i] = T.abdimas[i-1]
			i--
		}
		T.abdimas[i] = temp
		pass++
	}
}
func descsortprodipenelitian(T *tabdata) {
	var pass, i int
	var temp data_penelitian
	pass = 1
	for pass < T.npenelitian {
		i = pass
		temp = T.penelitian[pass]
		for i > 0 && (temp.prodi > T.penelitian[i-1].prodi) {
			T.penelitian[i] = T.penelitian[i-1]
			i--
		}
		T.penelitian[i] = temp
		pass++
	}
}
func descsortprodiabdimas(T *tabdata) {
	var pass, i int
	var temp data_abdimas
	pass = 1
	for pass < T.nabdimas {
		i = pass
		temp = T.abdimas[pass]
		for i > 0 && (temp.prodi > T.abdimas[i-1].prodi) {
			T.abdimas[i] = T.abdimas[i-1]
			i--
		}
		T.abdimas[i] = temp
		pass++
	}
}
func descsortfakultaspenelitian(T *tabdata) {
	var pass, i int
	var temp data_penelitian
	pass = 1
	for pass < T.npenelitian {
		i = pass
		temp = T.penelitian[pass]
		for i > 0 && (temp.fakultas > T.penelitian[i-1].fakultas) {
			T.penelitian[i] = T.penelitian[i-1]
			i--
		}
		T.penelitian[i] = temp
		pass++
	}
}
func descsortfakultasabdimas(T *tabdata) {
	var pass, i int
	var temp data_abdimas
	pass = 1
	for pass < T.nabdimas {
		i = pass
		temp = T.abdimas[pass]
		for i > 0 && (temp.fakultas > T.abdimas[i-1].fakultas) {
			T.abdimas[i] = T.abdimas[i-1]
			i--
		}
		T.abdimas[i] = temp
		pass++
	}
}
func sortdescendbytotalkegiatanpertahunpenelitian(x *tabdata) {
	var j int
	// Membuat map untuk menyimpan jumlah kegiatan per tahun
	jumlahKegiatan := make(map[int]int)

	// Menghitung jumlah kegiatan per tahun
	for i := 0; i < x.nabdimas; i++ {
		tahun := x.abdimas[i].tahun_kegiatan
		jumlahKegiatan[tahun]++
	}
	// Membuat slice untuk menyimpan hasil urutan berdasarkan jumlah kegiatan
	type JumlahKegiatan struct {
		Tahun  int
		Jumlah int
	}
	var hasil []JumlahKegiatan
	// Menambahkan data dari map ke slice
	for tahun, jumlah := range jumlahKegiatan {
		hasil = append(hasil, JumlahKegiatan{Tahun: tahun, Jumlah: jumlah})
	}
	// Mengurutkan slice berdasarkan jumlah kegiatan secara descending
	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].Jumlah > hasil[j].Jumlah
	})
	// Menampilkan semua data yang sudah terurut berdasarkan jumlah kegiatan
	for _, data := range hasil {
		tahun := data.Tahun
		// Menampilkan data berdasarkan tahun kegiatan
		for i := 0; i < x.nabdimas; i++ {
			if x.abdimas[i].tahun_kegiatan == tahun {
				fmt.Println("Judul:", x.penelitian[i].judul)
				fmt.Println("Ketua:", x.penelitian[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.penelitian[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.penelitian[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
				fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
				fmt.Println("Iuran produk:", x.penelitian[i].produk)
				fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	}
	choicedata_penelitian(x)
}
func sortascnytotalkegiatanpertahunpenelitian(x *tabdata) {
	var j int
	// Membuat map untuk menyimpan jumlah kegiatan per tahun
	jumlahKegiatan := make(map[int]int)
	// Menghitung jumlah kegiatan per tahun
	for i := 0; i < x.npenelitian; i++ {
		tahun := x.penelitian[i].tahun_kegiatan
		jumlahKegiatan[tahun]++
	}
	// Membuat slice untuk menyimpan hasil urutan berdasarkan jumlah kegiatan
	type JumlahKegiatan struct {
		Tahun  int
		Jumlah int
	}
	var hasil []JumlahKegiatan
	// Menambahkan data dari map ke slice
	for tahun, jumlah := range jumlahKegiatan {
		hasil = append(hasil, JumlahKegiatan{Tahun: tahun, Jumlah: jumlah})
	}
	// Mengurutkan slice berdasarkan jumlah kegiatan secara ascending
	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].Jumlah < hasil[j].Jumlah
	})
	// Menampilkan semua data yang sudah terurut berdasarkan jumlah kegiatan secara ascending
	for _, data := range hasil {
		tahun := data.Tahun

		// Menampilkan data berdasarkan tahun kegiatan
		for i := 0; i < x.nabdimas; i++ {
			if x.abdimas[i].tahun_kegiatan == tahun {
				fmt.Println("Judul:", x.penelitian[i].judul)
				fmt.Println("Ketua:", x.penelitian[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.penelitian[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.penelitian[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.penelitian[i].prodi, "/", x.penelitian[i].fakultas)
				fmt.Println("sumber dana:", x.penelitian[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.penelitian[i].publikasi)
				fmt.Println("Iuran produk:", x.penelitian[i].produk)
				fmt.Println("Tahun kegiatan:", x.penelitian[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	}
	choicedata_penelitian(x)
}

func sortdescendbytotalkegiatanpertahunabdimas(x *tabdata) {
	var j int
	// Membuat map untuk menyimpan jumlah kegiatan per tahun
	jumlahKegiatan := make(map[int]int)

	// Menghitung jumlah kegiatan per tahun
	for i := 0; i < x.nabdimas; i++ {
		tahun := x.abdimas[i].tahun_kegiatan
		jumlahKegiatan[tahun]++
	}
	// Membuat slice untuk menyimpan hasil urutan berdasarkan jumlah kegiatan
	type JumlahKegiatan struct {
		Tahun  int
		Jumlah int
	}
	var hasil []JumlahKegiatan
	// Menambahkan data dari map ke slice
	for tahun, jumlah := range jumlahKegiatan {
		hasil = append(hasil, JumlahKegiatan{Tahun: tahun, Jumlah: jumlah})
	}
	// Mengurutkan slice berdasarkan jumlah kegiatan secara descending
	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].Jumlah > hasil[j].Jumlah
	})
	// Menampilkan semua data yang sudah terurut berdasarkan jumlah kegiatan
	for _, data := range hasil {
		tahun := data.Tahun
		// Menampilkan data berdasarkan tahun kegiatan
		for i := 0; i < x.nabdimas; i++ {
			if x.abdimas[i].tahun_kegiatan == tahun {
				fmt.Println("Judul:", x.abdimas[i].judul)
				fmt.Println("Ketua:", x.abdimas[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.abdimas[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.abdimas[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
				fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
				fmt.Println("Iuran produk:", x.abdimas[i].produk)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].seminar)
				fmt.Println("Iuran publikasi:", x.abdimas[i].pelatihan)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	}
	choicedata_abdimas(x)
}
func sortascnytotalkegiatanpertahunabdimas(x *tabdata) {
	var j int
	// Membuat map untuk menyimpan jumlah kegiatan per tahun
	jumlahKegiatan := make(map[int]int)
	// Menghitung jumlah kegiatan per tahun
	for i := 0; i < x.nabdimas; i++ {
		tahun := x.abdimas[i].tahun_kegiatan
		jumlahKegiatan[tahun]++
	}
	// Membuat slice untuk menyimpan hasil urutan berdasarkan jumlah kegiatan
	type JumlahKegiatan struct {
		Tahun  int
		Jumlah int
	}
	var hasil []JumlahKegiatan
	// Menambahkan data dari map ke slice
	for tahun, jumlah := range jumlahKegiatan {
		hasil = append(hasil, JumlahKegiatan{Tahun: tahun, Jumlah: jumlah})
	}
	// Mengurutkan slice berdasarkan jumlah kegiatan secara ascending
	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].Jumlah < hasil[j].Jumlah
	})
	// Menampilkan semua data yang sudah terurut berdasarkan jumlah kegiatan secara ascending
	for _, data := range hasil {
		tahun := data.Tahun
		// Menampilkan data berdasarkan tahun kegiatan
		for i := 0; i < x.nabdimas; i++ {
			if x.abdimas[i].tahun_kegiatan == tahun {
				fmt.Println("Judul:", x.abdimas[i].judul)
				fmt.Println("Ketua:", x.abdimas[i].ketua)
				for j = 0; j < maxanggota; j++ {
					if x.abdimas[i].anggota[j] != "-" {
						fmt.Println("anggota:", x.abdimas[i].anggota[j])
					}
				}
				fmt.Println("Prodi/Fakultas: ", x.abdimas[i].prodi, "/", x.abdimas[i].fakultas)
				fmt.Println("sumber dana:", x.abdimas[i].sumber_dana)
				fmt.Println("Iuran publikasi:", x.abdimas[i].publikasi)
				fmt.Println("Iuran produk:", x.abdimas[i].produk)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].seminar)
				fmt.Println("Iuran publikasi:", x.abdimas[i].pelatihan)
				fmt.Println("Tahun kegiatan:", x.abdimas[i].tahun_kegiatan)
				fmt.Println("___________________________________________________________________________")
			}
		}
	}
	choicedata_abdimas(x)
}

func error_input_managementString(s *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*s = scanner.Text()
}
func error_input_managementINT(x *int) {
	var input string
	var err error
	var valid bool = false
	for !valid {
		fmt.Scan(&input)
		*x, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukkan input yang valid !!!")
		} else {
			valid = true
		}
	}
}
func error_input_managementstringUpper(s *string) {
	reader := bufio.NewReader(os.Stdin)

	var input string
	isValid := false

	for !isValid {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		isValid = true
		for _, ch := range input {
			if !unicode.IsUpper(ch) {
				isValid = false
			}
		}
		if !isValid {
			fmt.Println("Input tidak valid. Mohon masukkan huruf kapital saja.")
		}
	}
	*s = input
}
