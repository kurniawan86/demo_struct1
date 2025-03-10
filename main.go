package main

import (
	"bufio"
	"fmt"
	"os"
)

func insert() Karyawan {
	scanner := bufio.NewScanner(os.Stdin)

	//input karyawan
	fmt.Print("masukkan id karyawan :")
	var id int
	fmt.Scanln(&id)

	fmt.Print("masukkan nama karyawan : ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("masukkan alamat karyawan: ")
	scanner.Scan()
	alamat := scanner.Text()

	fmt.Print("Masukkan No Telp HP Karyawan: ")
	scanner.Scan()
	noHP := scanner.Text()

	fmt.Print("Masukkan No Telp Rumah Karyawan: ")
	scanner.Scan()
	noRumah := scanner.Text()

	// Input Anggota Keluarga
	fmt.Println("\n--- Data Anggota Keluarga ---")
	fmt.Print("Masukkan Nama Anggota Keluarga: ")
	scanner.Scan()
	namaAnggota := scanner.Text()

	fmt.Print("Masukkan Hubungan dengan Karyawan: ")
	scanner.Scan()
	hubungan := scanner.Text()

	fmt.Print("Masukkan No Telp Anggota: ")
	scanner.Scan()
	noTelpAnggota := scanner.Text()

	return Karyawan{
		IdKaryawan:     id,
		NamaKaryawan:   nama,
		AlamatKaryawan: alamat,
		NoTelpKaryawan: [2]string{noHP, noRumah},
		AnggotaKaryawan: Anggota{
			Nama:     namaAnggota,
			Hubungan: hubungan,
			No_telp:  noTelpAnggota,
		},
	}

}

// Fungsi untuk menampilkan data Karyawan
func display(emp Karyawan) {
	fmt.Println("\n==== Data Karyawan ====")
	fmt.Println("ID Karyawan:", emp.IdKaryawan)
	fmt.Println("Nama Karyawan:", emp.NamaKaryawan)
	fmt.Println("Alamat Karyawan:", emp.AlamatKaryawan)
	fmt.Println("Nomer Telp HP Karyawan:", emp.NoTelpKaryawan[0])
	fmt.Println("Nomer Telp Rumah Karyawan:", emp.NoTelpKaryawan[1])

	fmt.Println("\n==== Informasi Anggota Keluarga ====")
	fmt.Println("Nama Anggota:", emp.AnggotaKaryawan.Nama)
	fmt.Println("Hubungan:", emp.AnggotaKaryawan.Hubungan)
	fmt.Println("Telp Anggota:", emp.AnggotaKaryawan.No_telp)
}

func main() {
	fmt.Println("Program Input Data Karyawan")
	fmt.Println("--------------------------")

	// Memasukkan data
	karyawan := insert()

	// Menampilkan data yang sudah dimasukkan
	display(karyawan)
}
