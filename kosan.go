package main

import "fmt"

const NKMAX int = 20
const NPMAX int = 20

type kamar struct {
	nokamar   string
	tipekamar string
	hargasewa int
	avail     string
}

type penghuni struct {
	id       string
	nama     string
	nokamar  string
	notelp   int
	alamat   string
	tglmasuk string
}

type tabkamar [NKMAX]kamar
type tabpenghuni [NPMAX]penghuni

//dummy kamar
// tabkamar[1].nokamar = 01, tabkamar[2].tipekamar = "S",

func main() {
	var a tabpenghuni
	var b tabkamar
	var jumdatK, jumdatP int
	var quit bool
	var choose int
	jumdatK = 10
	jumdatP = 0
	quit = true

	b[0].nokamar = "01"
	b[0].tipekamar = "S"
	b[0].hargasewa = 900000
	b[0].avail = "tersedia"
	b[1].nokamar = "02"
	b[1].tipekamar = "S"
	b[1].hargasewa = 900000
	b[1].avail = "tersedia"
	b[2].nokamar = "03"
	b[2].tipekamar = "S"
	b[2].hargasewa = 900000
	b[2].avail = "tersedia"
	b[3].nokamar = "04"
	b[3].tipekamar = "S"
	b[3].hargasewa = 900000
	b[3].avail = "tersedia"
	b[4].nokamar = "05"
	b[4].tipekamar = "S"
	b[4].hargasewa = 900000
	b[4].avail = "tersedia"
	b[5].nokamar = "06"
	b[5].tipekamar = "M"
	b[5].hargasewa = 1100000
	b[5].avail = "tersedia"
	b[6].nokamar = "07"
	b[6].tipekamar = "M"
	b[6].hargasewa = 1100000
	b[6].avail = "tersedia"
	b[7].nokamar = "08"
	b[7].tipekamar = "M"
	b[7].hargasewa = 1100000
	b[7].avail = "tersedia"
	b[8].nokamar = "09"
	b[8].tipekamar = "M"
	b[8].hargasewa = 1100000
	b[8].avail = "tersedia"
	b[9].nokamar = "10"
	b[9].tipekamar = "M"
	b[9].hargasewa = 1100000
	b[9].avail = "tersedia"

	for quit == true {
		fmt.Println("=============================")
		fmt.Println("       Pilih Nomor Menu      ")
		fmt.Println("=============================")
		fmt.Println("1. Tambah data penghuni")
		fmt.Println("2. Tambah data kamar")
		fmt.Println("3. Tampilkan semua penghuni")
		fmt.Println("4. Tampilkan semua kamar")
		fmt.Println("5. Edit data penghuni")
		fmt.Println("6. Edit data kamar")
		fmt.Println("7. Hapus data penghuni")
		fmt.Println("8. Hapus data kamar")
		fmt.Println("9. Hitung harga sewa kamar")
		fmt.Println("0. Exit")
		fmt.Println("=============================")
		fmt.Print("Masukan nomor: ")
		fmt.Scanln(&choose)
		if choose == 0 {
			quit = false
		} else if choose == 1 {
			tambahpenghuni(&a, &jumdatP, &b, &jumdatK)
		} else if choose == 2 {
			tambahkamar(&b, &jumdatK)
		} else if choose == 3 {
			printpenghuni(a, jumdatP)
		} else if choose == 4 {
			printkamar(b, jumdatK)
		} else if choose == 5 {
			editdatapenghuni(&a, &jumdatP, &b, &jumdatK)
		} else if choose == 6 {
			editdatakamar(&b, &jumdatK)
		} else if choose == 7 {
			hapuspenghuni(&a, &jumdatP, &b, &jumdatK)
		} else if choose == 8 {
			hapuskamar(&b, &jumdatK)
		} else if choose == 9 {
			hitungharga(b, jumdatK)
		}
	}
}

func tambahpenghuni(p *tabpenghuni, np *int, k *tabkamar, nk *int) {
	var valid, vld bool
	var no_kamar string
	var i, j int

	fmt.Println("==========================")
	fmt.Println("   Tambah Data Penghuni   ")
	fmt.Println("==========================")
	valid = true
	vld = true
	for *np <= NPMAX && valid {
		fmt.Print("Masukan ID: ")
		fmt.Scanln(&p[*np].id)
		fmt.Print("Masukan nama: ")
		fmt.Scanln(&p[*np].nama)
		fmt.Print("Masukan no telpon: ")
		fmt.Scanln(&p[*np].notelp)
		fmt.Print("Masukan alamat: ")
		fmt.Scanln(&p[*np].alamat)
		fmt.Print("Masukan tanggal masuk: ")
		fmt.Scanln(&p[*np].tglmasuk)
		fmt.Print("Masukan no kamar: ")
		fmt.Scanln(&no_kamar)
		for i = 0; i < *np && vld; i++ {
			if p[i].nokamar == no_kamar {
				fmt.Println("kamar sudah terisi")
				vld = false
			}
		}
		if vld != false {
			p[*np].nokamar = no_kamar
			for j = 0; j < *nk; j++ {
				if (*k)[j].nokamar == p[*np].nokamar {
					(*k)[j].avail = "terisi"
				}
			}
			fmt.Println("Data penghuni berhasil ditambahkan")
			*np++
		}
		valid = false
	}
}

func printpenghuni(p tabpenghuni, np int) {
	var i, j int
	var temp penghuni

	for i = 0; i < np-1; i++ {
		for j = i + 1; j < np; j++ {
			if p[i].id > p[j].id {
				temp = p[i]
				p[i] = p[j]
				p[j] = temp
			}
		}
	}

	fmt.Printf("%-5s %-10s %-10s %-10s %-15s %-15s %-10s\n", "No.", "No ID", "Nama", "No kamar", "No telp", "Alamat", "Tgl masuk")
	for i = 0; i < np; i++ {
		fmt.Printf("%-5d %-10s %-10s %-10s %-15d %-15s %-10s\n", i+1, p[i].id, p[i].nama, p[i].nokamar, p[i].notelp, p[i].alamat, p[i].tglmasuk)
	}
}

func tambahkamar(k *tabkamar, nk *int) {
	var valid bool

	valid = true
	for *nk < NKMAX && valid {
		fmt.Print("Masukan nomor kamar: ")
		fmt.Scanln(&k[*nk].nokamar)
		fmt.Print("Masukan tipe kamar: ")
		fmt.Scanln(&k[*nk].tipekamar)
		fmt.Print("Masukan harga sewa /bulan: ")
		fmt.Scanln(&k[*nk].hargasewa)
		fmt.Print("Masukan ketersediaan kamar: ")
		fmt.Scanln(&k[*nk].avail)
		valid = false
	}
	*nk++
	fmt.Println("Data kamar berhasil ditambahkan")
}

func printkamar(k tabkamar, nk int) {
	var j, m, l int
	var temp kamar
	for j = 0; j < nk-1; j++ {
		for m = j + 1; m < nk; m++ {
			if k[j].avail > k[m].avail {
				temp = k[j]
				k[j] = k[m]
				k[m] = temp
			}
		}
	}

	for l = 0; l < nk; l++ {
		fmt.Println("No kamar: ", k[l].nokamar, "|", "Tipe kamar: ", k[l].tipekamar, "|", "Harga sewa: ", k[l].hargasewa, "|", "Ketersediaan: ", k[l].avail)
	}
}

func editdatapenghuni(p *tabpenghuni, np *int, k *tabkamar, nk *int) {
	var i, j, l, n, m int
	var no_kamar, noid string
	var valid, vld, valdi bool
	valid = true
	valdi = true
	vld = true
	fmt.Println("==========================")
	fmt.Println("   Edit Data Penghuni   ")
	fmt.Println("==========================")
	fmt.Print("No ID : ")
	fmt.Scanln(&noid)
	for i = 0; i < *np && valid; i++ {
		if p[i].id == noid {
			for m = 0; m < *nk && valdi; m++ {
				if k[m].nokamar == p[i].nokamar {
					n = m
					valdi = false
				}
			}
			fmt.Print("Masukan ID: ")
			fmt.Scanln(&p[i].id)
			fmt.Print("Masukan nama: ")
			fmt.Scanln(&p[i].nama)
			fmt.Print("Masukan no telpon: ")
			fmt.Scanln(&p[i].notelp)
			fmt.Print("Masukan alamat: ")
			fmt.Scanln(&p[i].alamat)
			fmt.Print("Masukan tanggal masuk: ")
			fmt.Scanln(&p[i].tglmasuk)
			fmt.Print("Masukan no kamar: ")
			fmt.Scanln(&no_kamar)
			for j = 0; j < *np && vld; j++ {
				if p[j].nokamar == no_kamar {
					fmt.Println("kamar sudah terisi, gagal mengedit kamar")
					vld = false
				}
			}
			if vld != false {
				p[i].nokamar = no_kamar
				for l = 0; l < *nk; l++ {
					if (*k)[l].nokamar == p[i].nokamar {
						(*k)[l].avail = "terisi"
					}

				}
				k[n].avail = "tersedia"
				fmt.Println("Data penghuni berhasil diedit")
			}
			valid = false
		}
	}
	if valid == true {
		fmt.Println("ID penghuni tidak ditemukan")
	}
}

func editdatakamar(k *tabkamar, nk *int) {
	var i int
	var no_kamar string
	var valid bool

	valid = true
	fmt.Println("==========================")
	fmt.Println("      Edit Data Kamar     ")
	fmt.Println("==========================")
	fmt.Print("No kamar : ")
	fmt.Scanln(&no_kamar)
	for i = 0; i < *nk && valid; i++ {
		if k[i].nokamar == no_kamar {
			fmt.Print("Masukan nomor kamar: ")
			fmt.Scanln(&k[i].nokamar)
			fmt.Print("Masukan tipe kamar: ")
			fmt.Scanln(&k[i].tipekamar)
			fmt.Print("Masukan harga sewa /bulan: ")
			fmt.Scanln(&k[i].hargasewa)
			fmt.Print("Masukan ketersediaan kamar: ")
			fmt.Scanln(&k[i].avail)
			valid = false
			fmt.Println("Data kamar berhasil diedit")
		}
	}
	if valid == true {
		fmt.Println("Data kamar tidak ditemukan")
	}
}

func hapuspenghuni(p *tabpenghuni, np *int, k *tabkamar, nk *int) {
	var i, j, m, l int
	var noid string
	var valid bool = true

	m = 0
	fmt.Println("==========================")
	fmt.Println("   Hapus Data Penghuni   ")
	fmt.Println("==========================")
	fmt.Print("No ID : ")
	fmt.Scanln(&noid)
	// for j = 0; j < *nk; j++ {}
	for i = 0; i < *np && valid; i++ {
		if p[i].id == noid {
			m = i
			valid = false
		}
	}
	if valid == false {
		for l = 0; l < *nk; l++ {
			if k[l].nokamar == p[m].nokamar {
				k[l].avail = "tersedia"
			}
		}
		for j = m; j < *np; j++ {
			p[j] = p[j+1]
		}
		*np--
	} else if valid == true {
		fmt.Println("Data penghuni tidak ditemukan")
	}
}

func hapuskamar(k *tabkamar, nk *int) {
	var i, j, m int
	var valid bool = true
	var no_kamar string

	m = 0
	fmt.Println("=======================")
	fmt.Println("   Hapus Data Kamar   ")
	fmt.Println("=======================")
	fmt.Print("No Kamar : ")
	fmt.Scanln(&no_kamar)
	for i = 0; i < *nk && valid; i++ {
		if k[i].nokamar == no_kamar {
			m = i
			valid = false
		}
	}
	if valid == false {
		for j = m; j < *nk-1; j++ {
			k[j] = k[j+1]
		}
	} else if valid == true {
		fmt.Println("Data kamar tidak ditemukan")
	}
}

func hitungharga(k tabkamar, nk int) {
	var i, m, total int
	var valid bool
	var no_kamar string
	var brpbulan int

	total = 0
	m = 0
	valid = true
	fmt.Println("=====================")
	fmt.Println("  Hitung harga sewa  ")
	fmt.Println("=====================")
	fmt.Print("Nomor kamar: ")
	fmt.Scanln(&no_kamar)
	fmt.Print("Durasi waktu (/bulan): ")
	fmt.Scanln(&brpbulan)
	for i = 0; i < nk && valid; i++ {
		if k[i].nokamar == no_kamar {
			m = i
			valid = false
		}
	}
	if valid == false {
		total = k[m].hargasewa * brpbulan
		fmt.Println("Rp.", total)
	} else if valid == true {
		fmt.Println("Nomor kamar tidak ditemukan")
	}
}
