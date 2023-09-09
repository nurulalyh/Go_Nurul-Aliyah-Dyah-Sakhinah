# (10) Clean Code

Clean code adalah istilah untuk kode yang mudah dibaca, dipahami, dan diubah oleh programmer lain. Clean code harus diterapkan karena 3 hal yaitu :
1. Work Collaboration => Semua orang harus bisa membaca yang kita kerjakan
2. Feature Development => Fitur yang dikembangkan diketahui fungsinya sehingga lebih mudah dikembangkan
3. Faster Development => Development menjadi lebih cepat karena clean code cepat dimengerti fungsinya sehingga tidak perlu mempelajari/membedah kode dari awal.

Karakteristik clean code :
1. Mudah dipahami
2. Mudah dieja dan dicari
3. Singkat namun mendeskripsikan konteks
4. Konsisten
5. Hindari penambahan konteks yang tidak perlu
6. Komentar
7. Good function
8. Gunakan Konvensi
9. Formatting :
    - lebar baris code 80 - 120 Karakter.
    - satu class 300-500 baris.
    - Baris code yang berhubungan saling berdekatan.
    - Dekatkan fungsi dengan pemanggilnya.
    - Perhatikan indentasi.
    - Menggunakan prettier atau formatter.

Clean Code Principle :
1. KISS (Keep It So Simple) => jangan membuat func yang memiliki terlalu banyak tugas berbeda-beda. Lebih baik sebuah func hanya memiliki satu tugas spesifik. Tips :
- Fungsi atau class harus kecil.
- Fungsi dibuat untuk melakukan satu tugas saja.
- Jangan gunakan terlalu banyak argumen pada fungsi.
- Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil dan jumlahnya minimal.
2. DRY (Don't Repeat Yourself) => jangan duplikasi kode, lebih baik buatkan function khusus agar dapat digunakan berulang-ulang.

Refactoring merupakan proses restrukturisasi kode dengan mengubah struktur internal tanpa mengubah perilaku eskternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor. Teknik refactoring :
- Membuat sebuah abstraksi.
- Memecah kode dengan fungsi/class.
- Perbaiki penamaan dan lokasi kode.
- Deteksi kode yang memiliki duplikat.