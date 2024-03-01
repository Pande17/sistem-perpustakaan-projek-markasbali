package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Sitem Manajemen Perpustakaan

type buku struct {
	kode      string
	judul     string
	pengarang string
	penerbit  string
	halaman   int
	tahun     int
}

var perpustakaan []buku

func main() {
	pilihan := 0

	//menu pilihan
	fmt.Println("\n===== Sistem Manajemen Perpusatakaan =====")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Daftar Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")

	fmt.Print("Masukkan Pilihan : ")
	_, err := fmt.Scanln(&pilihan)
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
	}
	switch pilihan {
	case 1:
		tambahBuku()
	case 2:
		daftarBuku()
	case 3:
		editBuku()
	case 4:
		hapusBuku()
	case 5:
		fmt.Println("\nTerima Kasih Sudah Berkunjung")
		os.Exit(0)
	}

	main()
}

func tambahBuku() {
	//deklarasi variabel
	kodeBuku := bufio.NewReader(os.Stdin)
	judulBuku := bufio.NewReader(os.Stdin)
	pengarangBuku := bufio.NewReader(os.Stdin)
	penerbitBuku := bufio.NewReader(os.Stdin)
	halamanBuku := 0
	tahunBuku := 0

	fmt.Println("\n===== Tambah Buku Baru =====")

	//kode buku
	fmt.Print("Masukkan Kode Buku Baru : ")
	kodeBukuBaru, err := kodeBuku.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}
	kodeBukuBaru = strings.TrimSpace(kodeBukuBaru)

	//judul buku
	fmt.Print("Masukkan Judul Buku Baru : ")
	judulBukuBaru, err := judulBuku.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}
	judulBukuBaru = strings.TrimSpace(judulBukuBaru)

	//pengarang buku
	fmt.Print("Masukkan Pengarang Buku Baru : ")
	pengarangBukuBaru, err := pengarangBuku.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}
	pengarangBukuBaru = strings.TrimSpace(pengarangBukuBaru)

	//penerbit buku
	fmt.Print("Masukkan Penerbit Buku Baru : ")
	penerbitBukuBaru, err := penerbitBuku.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}
	penerbitBukuBaru = strings.TrimSpace(penerbitBukuBaru)

	//jumlah halaman buku
	fmt.Print("Masukkan Jumlah Halaman Buku Baru : ")
	_,err = fmt.Scanln(&halamanBuku)
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}

	//tahun terbit buku
	fmt.Print("Masukkan Tahun Terbit Buku Baru : ")
	_,err = fmt.Scanln(&tahunBuku)
	if err != nil {
		fmt.Println("Terjadi Error: ", err)
		return
	}


	//pemasukan data ke dalam struct
	perpustakaan = append(perpustakaan, buku {
		kode 		: kodeBukuBaru,
		judul		: judulBukuBaru,
		pengarang 	: pengarangBukuBaru,
		penerbit	: penerbitBukuBaru,
		halaman		: halamanBuku,
		tahun		: tahunBuku,
	})
	fmt.Println("Berhasil Menambah Buku Ke Perpustakaan!")
}

func daftarBuku() {
	fmt.Println("\n===== Daftar Buku =====")
	if len(perpustakaan) == 0 {
		fmt.Println("Tidak Ada Buku Di Dalam Perpustakaan")
		return
	}
	for i, buku := range perpustakaan {
		fmt.Printf("%d. Kode Buku : %s \n- Judul : %s \n- Pengarang : %s \n- Penerbit : %s \n- Jumlah Halaman : %d \n- Tahun Rilis : %d \n\n",
			i+1, buku.kode, buku.judul, buku.pengarang, buku.penerbit, buku.halaman, buku.tahun)
	}
}

func editBuku() {
	var kode string

	fmt.Println("\n===== Edit Data Buku =====")
	fmt.Print("Masukkan Kode Buku : ")
	fmt.Scanln(&kode)

	for i, buku := range perpustakaan {
		if buku.kode == kode {
			inputanUser := bufio.NewReader(os.Stdin)

			fmt.Println("Masukkan Data Baru : ")
			fmt.Print("Judul Buku : ")
			perpustakaan[i].judul, _ = inputanUser.ReadString('\n')

			fmt.Print("Pengarang Buku : ")
			perpustakaan[i].pengarang, _ = inputanUser.ReadString('\n')

			fmt.Print("Penerbit Buku : ")
			perpustakaan[i].penerbit, _ = inputanUser.ReadString('\n')

			fmt.Print("Jumlah Halaman Buku : ")
			fmt.Scanln(&perpustakaan[i].halaman)

			fmt.Print("Tahun Terbit Buku : ")
			fmt.Scanln(&perpustakaan[i].tahun)
			fmt.Println("\nData Buku Berhasil Diedit")
			fmt.Println("Tekan 'Enter' Untuk Melanjutkan...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}else{
			fmt.Println("Terjadi Error Dalam Pemilihan Kode Buku atau Kode Buku Tidak Ada!")
			fmt.Println("Tekan 'Enter' Untuk Melanjutkan...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func hapusBuku() {
	var kodeInput string
	fmt.Println("\n===== Hapus Buku =====")
	daftarBuku()

	fmt.Print("\nMasukkan Kode Buku Yang Ingin Dihapus : ")
	fmt.Scanln(&kodeInput)

	for i,j := range perpustakaan{
		if j.kode == kodeInput{
			perpustakaan = append(
				perpustakaan[:i],
				perpustakaan[i+1:]...
		)
		fmt.Println("Buku Berhasil Dihapus Dari Perpustakaan")
	}else{
		fmt.Println("Terjadi Error Dalam Pemilihan Kode Buku atau Kode Buku Tidak Ada!")
	}
	}
	fmt.Println("Tekan 'Enter' Untuk Melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
