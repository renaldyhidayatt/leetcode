package main

import (
	"fmt"
	"strings"
)

func longgesCommon(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func loggest() {
	flower := []string{"flower", "flow", "flight"}
	fmt.Println(longgesCommon(flower))
}

/*
Ini adalah contoh kode untuk sebuah fungsi bernama longgestCommon yang menerima sebuah array string sebagai argumen dan mengembalikan sebuah string yang merupakan prefiks yang paling panjang yang dimiliki oleh semua string dalam array tersebut. Fungsi ini pertama-tama mengecek apakah array yang diberikan kosong. Jika iya, maka fungsi akan mengembalikan string kosong. Jika tidak, fungsi akan menetapkan prefiks awal sebagai elemen pertama dari array tersebut. Kemudian, fungsi akan mengulangi sebanyak jumlah elemen dalam array, kecuali elemen pertama. Setiap iterasi, fungsi akan mengecek apakah elemen yang sedang dibandingkan memiliki prefiks yang sama dengan prefiks awal. Jika tidak, prefiks awal akan dipangkas hingga panjangnya menjadi satu karakter lebih pendek dari sebelumnya. Jika setelah dipangkas panjangnya menjadi 0, maka fungsi akan mengembalikan string kosong. Setelah semua elemen dalam array dibandingkan, fungsi akan mengembalikan prefiks awal yang telah dipangkas sesuai dengan kondisi di atas.
*/
