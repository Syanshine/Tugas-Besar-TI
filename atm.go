package main

import (
	"fmt"
)

func main() {
	
	var username string = "syana"
	var password string = "2406420324"
	var saldo float64 = 3500000
	var historiTransaksi []string

	var inputUsername, inputPassword string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&inputUsername)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&inputPassword)

	if inputUsername == username && inputPassword == password {
		pilihanMenu := []string{
			"Lihat Informasi Akun",
			"Lihat Saldo",
			"Tambahkan Saldo",
			"Tarik Saldo",
			"Histori Transaksi",
			"Keluar dari Program",
		}

		var pilihan int
		for {
			displayMenu(pilihanMenu)
			fmt.Print("Masukkan pilihan menu: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Println("\n--- Informasi Akun ---")
				fmt.Printf("Username: %s\n", username)
				fmt.Printf("Saldo Terkini: %.2f\n", saldo)
			case 2:
				fmt.Println("\n--- Menampilkan Saldo ---")
				fmt.Printf("Saldo Terkini: %.2f\n", saldo)
			case 3:
				fmt.Println("\n--- Menambahkan Saldo ---")
				var jumlah float64
				fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
				fmt.Scan(&jumlah)
				if jumlah <= 0 {
					fmt.Println("Penambahan saldo harus lebih besar dari 0.")
				} else {
				saldo += jumlah
				historiTransaksi = append(historiTransaksi, fmt.Sprintf("Penambahan: %.2f", jumlah))
				fmt.Println("--- Saldo berhasil ditambahkan ---")
				}
			case 4:
				fmt.Println("\n--- Tarik Tunai ---")
				var jumlah float64
				fmt.Print("Masukkan jumlah yang akan ditarik: ")
				fmt.Scan(&jumlah)
				if jumlah <= 0 {
					fmt.Println("Jumlah penarikan harus lebih besar dari 0.")
				} else if jumlah > saldo {
					fmt.Println("Saldo tidak mencukupi!")
				} else {
					saldo -= jumlah
					historiTransaksi = append(historiTransaksi, fmt.Sprintf("Penarikan: %.2f", jumlah))
					fmt.Println("--- Penarikan tunai berhasil ---")
				}
			case 5:
				fmt.Println("\n--- Histori Transaksi ---")
				fmt.Println("Histori Transaksi:")
				for _, transaksi := range historiTransaksi {
					fmt.Println(transaksi)
				}
				fmt.Println("--- Histori transaksi berhasil ditampilkan ---")
			case 6:
				fmt.Println("Terima kasih telah menggunakan ATM!")
				return
			default:
				fmt.Println("Pilihan tidak valid. Mohon coba lagi.")
			}
		}
	} else {
		fmt.Println("Username atau Password salah, mohon coba lagi!")
	}
}

func displayMenu(options []string) {
	fmt.Println("\nATM Menu:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}



		
