# (9) Concurrent Programming

Untuk melakukan Big Search website, dapat menggunakan metode sequential, parallel, atau Concurrent.
1. sequential program => menyelesaikan semua task secara berurutan.
2. parallel program => menyelesaikan semua task secara bersamaan, dengan syarat harus ada multicore.
3. Concurrent program => menyelesaikan task secara bersamaan dan secara indepent, tidak membutuhkan multicore.

Fitur dalam golang :
1. Goroutines merupakan teknologi dalam golang untuk menerapkan Concurrent. Func main merupakan Concurrent utama di golang. Goroutines merupakan sebuah Func yang ditentukan akan berjalan Concurrent dibanding Func yang lain. 
2. Channel merupakan fitur golang untuk tempat Concurrent bernaung, melakukan komunikasi dan sinkronisasi antara Concurrent/goroutines 2 task yang sedang berlangsung. 
3. Select untuk mengontrol jalannya Concurrent di channel - channel yang ada.
4. Gomaxprocs digunakan untuk mengontrol jumlah core yang digunakan program.

Race Condition merupakan salah satu kelemahan dari Concurrent Programming. Terjadi karena 2 Concurrent atau lebih tidak sinkron saat mengakses variabel yang sama. Untuk menghandle Race Condition, salah satu goroutines harus mengalah dan bisa mengakses kembali variabel setelah satu goroutines telah selesai. Cara fix data race :
1. Blocking with waitgroups => melakukan lock variabel pada goroutines yang lain hingga, goroutines yang berjalan selesai.
2. Channel Blocking => sinkronisasi antar goroutines.
3. Mutex => untuk waiting dan locking goroutines.
