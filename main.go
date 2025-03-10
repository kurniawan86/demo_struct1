package main

import "fmt"

// definisi struct
type Karyawan struct {
	idKaryawan      int
	namaKaryawan    string
	alamatKaryawan  string
	noTelpKaryawan  [2]string
	anggotaKaryawan Anggota
}

type Anggota struct {
	nama     string
	hubungan string
	no_telp  string
}

func main() {
	fmt.Println("program struct demo ke - 1")

	//pembuatan data karyawan pertama
	var karyawan1 = Karyawan{
		idKaryawan:     1,
		namaKaryawan:   "kurniawan",
		alamatKaryawan: "surabaya, sidosermo",
		noTelpKaryawan: [2]string{"083862118631", "031-88373845"},
		anggotaKaryawan: Anggota{
			nama:     "istri dari pak kurniawan",
			hubungan: "istri",
			no_telp:  "939393939339",
		},
	}

	//output
	fmt.Println("id karyawan \t\t\t: ", karyawan1.idKaryawan)
	fmt.Println("alamat karyawan \t\t: ", karyawan1.alamatKaryawan)
	fmt.Println("nama karyawan \t\t\t: ", karyawan1.namaKaryawan)
	fmt.Println("nomer telp HP karyawan \t\t: ", karyawan1.noTelpKaryawan[0])
	fmt.Println("nomer telp Rumah karyawan \t: ", karyawan1.noTelpKaryawan[1])
	fmt.Println("----- informasi anggota keluarga -----")
	fmt.Println("nama anggota \t\t: ", karyawan1.anggotaKaryawan.nama)
	fmt.Println("hubungan \t\t: ", karyawan1.anggotaKaryawan.hubungan)
	fmt.Println("telp anggota \t\t: ", karyawan1.anggotaKaryawan.no_telp)
}
