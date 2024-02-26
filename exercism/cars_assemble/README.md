## Penjelasan

Pada paket cars, terdapat beberapa fungsi untuk melakukan perhitungan terkait produksi mobil dan biaya produksi. Berikut adalah penjelasan singkat dari setiap fungsi yang disediakan:

CalculateWorkingCarsPerHour: Fungsi ini menghitung jumlah mobil yang berhasil diproduksi setiap jam oleh lini perakitan. Fungsi ini menerima dua parameter: productionRate yang merupakan jumlah mobil yang diproduksi setiap jam (bertipe int), dan successRate yang merupakan tingkat keberhasilan produksi dalam bentuk persentase (bertipe float64). Fungsi ini mengembalikan jumlah mobil yang berhasil diproduksi setiap jam dalam bentuk float64.

CalculateWorkingCarsPerMinute: Fungsi ini menghitung jumlah mobil yang berhasil diproduksi setiap menit oleh lini perakitan. Fungsi ini memanfaatkan fungsi CalculateWorkingCarsPerHour untuk menghitung jumlah mobil yang berhasil diproduksi setiap jam, lalu mengkonversi hasilnya ke dalam jumlah mobil per menit. Fungsi ini menerima parameter yang sama dengan CalculateWorkingCarsPerHour dan mengembalikan jumlah mobil yang berhasil diproduksi setiap menit dalam bentuk int.

CalculateCost: Fungsi ini menghitung biaya produksi untuk jumlah mobil yang diberikan. Fungsi ini menerima satu parameter yaitu carsCount, yang merupakan jumlah mobil yang akan diproduksi. Fungsi ini menghitung biaya produksi berdasarkan jumlah mobil yang diberikan, dengan asumsi biaya per mobil adalah 95000 untuk setiap kelompok sepuluh mobil dan 10000 untuk mobil yang tersisa. Fungsi ini mengembalikan biaya produksi dalam bentuk uin
