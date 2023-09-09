# (12) Join - Union - Agregasi - Subquery - Function

## Join
Join merupakan klausa untuk menggabungkan record dari dua atau lebih tabel. Jenis - Jenis Join :
- Inner => return baris-baris dari tabel-tabel yang memenuhi syarat
- Left => return seluruh data dari tabel disebelah kiri yang memenuhi kondisi dan sebagian data dari tabel kanan yang memiliki relasi dengan tabel kiri.
- Right => return seluruh data dari tabel disebelah kanan yang memenuhi kondisi dan sebagian data dari tabel kanan yang memiliki relasi dengan tabel kanan.

## Union
Union adalah klausa yang sama seperti Join, namun field yang dipanggil harus sama.

## Agregasi
Agregasi merupakan fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal. Jenis-jenis fungsi Agregasi :
- Min => mencari data terkecil dari sebuah field 
- Max => mencari data terbesar dari sebuah field
- Sum => menjumlah seluruh data dari sebuah field
- Avg => mencari nilai rata-rata dari keseluruhan data di sebuah field
- Count => mencari banyaknya data pada sebuah field
- Having => menyeleksi data berdasarkan kriteria yang merupakan agregat

## Subquery
Subquery adalah query query di dalam query SQL lain. Digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Subquery dapat digunakan dengan select, insert, update, dan delete bersama dengan operator =, <, >, in, between, dll. 

## Function
Function merupakan kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya. Contoh Query yang digunakan :
- DELIMETER => penanda awal dan akhir dari sebuah function.
- CREATE FUNCTION => header dari function
- RETURNS => menentukan tipe data yang akan di return-kan oleh function.
- DETERMINISTIC/NOT DETERMINISTIC => membatasi pengguna dari function, DETERMINISTIC berarti hanya dapat digunakan oleh pembuatnya.
- BEGIN - END => body dari function. semua query di tulis di dalamnya.