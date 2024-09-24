package main

import (
	"fmt"
)

type Akun struct {
	Username        string
	Password        string
	Saldo         float64
	HistoriTransaksi []string
}

func (a *Akun) TambahkanSaldo(jumlah float64) {
	a.Saldo += jumlah
	a.HistoriTransaksi = append(a.HistoriTransaksi, fmt.Sprintf("Added: %.2f", jumlah))
}

func (a *Akun) Tarik(jumlah float64) bool {
	if jumlah <= 0 {
		fmt.Println("Saldo akun harus lebih besar dari 0.")
		return false
	}
	if jumlah > a.Saldo {
		fmt.Println("Saldo tidak mencukupi!")
		return false
	}
	a.Saldo -= jumlah
	a.HistoriTransaksi = append(a.HistoriTransaksi, fmt.Sprintf("Penarikan: %.2f", jumlah))
	return true
}

func (a *Akun) TunjukkanSaldo() {
	fmt.Printf("Saldo Terkini: %.2f\n", a.Saldo)
}

func (a *Akun) TunjukkanHistori() {
	fmt.Println("Histori Transaksi:")
	for _, transaksi := range a.HistoriTransaksi {
		fmt.Println(transaksi)
	}
}

func (a *Akun) TunjukkanInfoAkun() {
	fmt.Printf("Username: %s\n", a.Username)
	fmt.Printf("Saldo Terkini: %.2f\n", a.Saldo)
}

func displayMenu(options []string) {
	fmt.Println("\nATM Menu:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func main() {
	akun := Akun{
		Username: "syana",
		Password: "2406420324",
		Saldo:  3500000,
	
	}


var username, password string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	if username == akun.Username && password == akun.Password {
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
				akun.TunjukkanInfoAkun()
				fmt.Println("--- End of Account Information ---")
			case 2:
				fmt.Println("\n--- Showing Balance ---")
				akun.TunjukkanSaldo()
				fmt.Println("--- End of Balance ---")
			case 3:
				fmt.Println("\n--- Menambahkan Saldo ---")
				var jumlah float64
				fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
				fmt.Scan(&jumlah)
				akun.TambahkanSaldo(jumlah)
				fmt.Println("--- Saldo berhasil ditambahkan ---")
			case 4:
				fmt.Println("\n--- Tarik Tunai ---")
				var jumlah float64
				fmt.Print("Masukkan jumlah yang akan ditarik: ")
				fmt.Scan(&jumlah)
				if akun.Tarik(jumlah) {
					fmt.Println("--- Penarikan tunai berhasil ---")
				}
			case 5:
				fmt.Println("\n--- Histori Transaksi ---")
				akun.TunjukkanHistori()
				fmt.Println("--- End of History ---")
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