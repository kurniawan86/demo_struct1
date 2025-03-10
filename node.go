package main

// definisi struct
type Karyawan struct {
	IdKaryawan      int
	NamaKaryawan    string
	AlamatKaryawan  string
	NoTelpKaryawan  [2]string
	AnggotaKaryawan Anggota
}

type Anggota struct {
	Nama     string
	Hubungan string
	No_telp  string
}
