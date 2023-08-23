# (5) Data Structure

## Array
Array merupakan struktur data yang berisi sekelompok elemen dengan tipe data yang sama. Indeks di array dimulai dari 0 dan ukuran dari array tidak dapat berubah-rubah, hanya mengikuti apa yang telah dideklarasikan. 

Cara deklarasi Array => `var <namaArray> [<panjangArray>] <tipeData>` 

Jika Array yang telah dideklarasikan value nya tidak terisi penuh, maka ruang yang kosong akan otomatis terisi oleh nilai default dari tipe data yang digunakan array.

## Slice 
Slice mirip seperti array, tapi slice ukurannya dinamis atau dapat berubah-berubah. Slice juga dapat dibuat reference dari data-data yang dimiliki array sehingga slice yang seperti ini tergolong **Reference Type**. Slice ini memiliki pointer, pointernya ngambil dari value yang diambil dari array. Pada kasus ini, perubahan value pada array akan mempengaruhi slice.

Selain pointer, slice juga memiliki length (len) atau panjang slice dan capacity (cap) atau kapasitas maximum slice. Karena slice bersifat dinamis, maka len akan berubah-rubah sesuai jumlah elemen dalam slice, sedangkan cap akan berubah 2x lipat dari cap sebelumnya ketika ditambah elemen baru diluar kapasitas slice sebelumnya.

Cara deklarasi slice => `var <namaSlice> []<tipeData>` dan `var <namaSlice> = []<tipeData>{<value>}`
Cara deklarasi Map menggunakan `make()` => `var <namaSlice> = make([]<tipeData>,5,10)`

## Map
Map merupakan struktur data yang memiliki key (bersifat unik) dan value.

Cara deklarasi Map => var harga = `map[string]int{"siomay":1000, "batagor":2000}`
Cara deklarasi Map menggunakan `make()` => `var harga = make(map[string]int)`

## Function
Function merupakan sekumpulan kode dengan fungsi tertentu yang dipanggil dengan nama tertentu yang dibuat di luar function main dan ketika akan dijalankan, hanya perlu memanggil nama dari function. Bertujuan untuk membuat clean code dan lebih modular.