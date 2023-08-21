# (4) Time Complexity & Space Complexity

## Time Complexity
Time Complexity membahas bagaimana cara kita menghitung waktu eksekusi sebuah algoritma atau program. Hal yang harus diperhatikan dalam menghitung Time Complexity :
1. Tentukan seberapa banyak operasi dominan/primitif yang dilakukan seperti penjumlahan dan perkalian.
2. Dari banyak kode program, harus ditentukan bagian yang paling dominan.
3. Operasi dominan dilihat dari jumlah eksekusinya.

Jenis dari Time Complexity :
1. Constant Time - O(1) => Berapapun inputnya, perasi yang ada di algoritma hanya dieksekusi 1 kali. Dengan kata lain, operasi tidak terpengaruh inputan. 
2. Linear Time - O(n) => Ketika dimasukkan input, operasi yang dilakukan sebanding dengan input tersebut. Time Complexity ini, merupakan yang paling normal dan mudah. Selain O(n) ada juga O(n+m) yang serupa konsepnya, namun karena inputnya ada 2, maka operasinya dilakukan sebanyak jumlah dari kedua input.
3. Logarithmic Time - O(log n) => Merupakan yang paling efektif setelah O(1) jika ingin melakukan sebuat problem solving yang berkaitan dengan kecepatan eksekusi. Rumus dari time Complexity ini adalah 2^x = n. Digunakan untuk binary search (divide and conquer) dan sorting. 
4. Quadratic Time - O(n^2) => Paling banyak digunakan saat melakukan problem solving. Rumusnya adalah 0.5 * (n * (n+1)) = 0.5 * n^2 +0.5 * n. Diterapkan di proses nested loop. Digunakan saat mencari permutasi, kombinasi, dll. Merupakan time Complexity yang tidak efektif. Prinsip ini dapat digunakan ketika tidak ada lagi time Complexity yang dapat digunakan.
5. Exponential Time - O(2^n)
6. Factorial Time - O(n!)

Hasil Analisa grafik : 
1. Constant Time = Berapapun inputannya, time Complexity akan tetap atau tidak terpengaruh dengan inputannya.
2. Logarithmic Time = time Complexity cenderung cepat karena inputannya akan terus dibagi 2 setiap iterasi.
3. Linear Time = Time Complexity nya menyesuaikan atau sebanding dengan inputannya.
4. Quadratic Time = Bahkan jika inputannya sedikit, Time Complexity nya akan langsung dikalikan sehingga melonjak naik, maka dari itu sangat tidak efektif. 

Cara Mengetahui berapa lama yang dibutuhkan untuk mengeksekusi program dalam satuan waktu. Dari referensi saat ini, komputer yang digunakan dalam satu detik mampu melakukan operasi sebanyak 10^8^ atau 100.000.000 data.

## Space Complexity
Space Complexity berbicara mengenai jumlah memori yang digunakan dalam sebuah algoritma. Space Complexity menggunakan O notation. Memori berhubungan dengan variabel, maka perhatikan banyaknya variabel yang dibutuhkan dalam algoritma tersebut. Contoh :
1. Pada Algoritma dengan Time Complexity O(1), membutuhkan 2 variabel sehingga space Complexity nya menjadi O(2) atau bersifat konstan.
2. Pada Algoritma dengan Time Complexity O(n), misalnya dibuat sebuah slice, dilakukan append sebanyak input variabel n. jadi dalam memori akan terdapat n data sehingga space Complexity akan sama dengan O(n).

Harapannya, ketika membuat algoritma gunakan time Complexity Constant atau Logarithmic dan space Complexity nya adalah Constant sehingga kecepatan operasi yang dijalankan cukup cepat sehingga user secara ux akan langsung terlihat hasilnya.
