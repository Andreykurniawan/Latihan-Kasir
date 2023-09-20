package main

import (
	"fmt"
)

type Barang struct {
	Nama  string
	Harga float64
}

func main() {
	// Inisialisasi daftar barang
	daftarBarang := map[string]Barang{
		"Buku Tulis": Barang{"Buku Tulis", 3000.0},
		"Bolpoin":    Barang{"Bolpoin", 2000.0},
		"Pensil":     Barang{"Pensil", 1500.0},
	}

	// Inisialisasi keranjang belanja
	keranjang := make(map[string]int)

	fmt.Println("Selamat datang di aplikasi kasir!")

	for {
		// Menampilkan daftar barang yang tersedia
		fmt.Println("\nDaftar Barang:")
		for key, barang := range daftarBarang {
			fmt.Printf("%s - %s - Rp %.2f\n", key, barang.Nama, barang.Harga)
		}

		fmt.Print("Pilih barang (ketik 'selesai' untuk menyelesaikan transaksi): ")
		var input string
		fmt.Scanln(&input)

		if input == "selesai" {
			break
		}

		barang, found := daftarBarang[input]
		if !found {
			fmt.Println("Barang tidak ditemukan. Mohon pilih barang yang valid.")
			continue
		}

		fmt.Printf("Jumlah %s yang dibeli: ", barang.Nama)
		var jumlah int
		fmt.Scanln(&jumlah)

		if jumlah < 0 {
			fmt.Println("Jumlah barang tidak valid. Mohon masukkan jumlah yang valid.")
			continue
		}

		keranjang[input] += jumlah
	}

	// Menampilkan total belanja
	total := 0.0
	fmt.Println("\nBarang yang dibeli:")
	for key, jumlah := range keranjang {
		barang := daftarBarang[key]
		subtotal := float64(jumlah) * barang.Harga
		fmt.Printf("%s - %s - %d buah - Subtotal: Rp %.2f\n", key, barang.Nama, jumlah, subtotal)
		total += subtotal
	}

	fmt.Printf("Total biaya belanja: Rp %.2f\n", total)
}
