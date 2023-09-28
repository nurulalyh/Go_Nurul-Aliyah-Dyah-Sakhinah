# (19) Unit Testing

Testing adalah proses analisis sebuah sistem untuk membandingkan kondisi fitur saat ini dengan kondisi fitur yang di inginkan end user. Bertujuan untuk mengetahui apakah sistem bekerja dengan baik dan tidak memiliki bug. Alasan melakukan testing :
- Preventing regression
- Confidence level in refactoring
- Improve code design
- Code documentation
- Scaling the team

Level Testing :
- UI Test => dilakukan melalui user interface aplikasinya, dengan mencoba flow aplikasi.
- Integration Test => Test terhadap modul atau sub sistem secara spesifik.
- Unit Test => Test terhadap unit terkecil dari aplikasi.

Framework dalam unit testing merupakan library yang dapat digunakan untuk melakukan unit testing. Framework dibutuhkan karena bisa lebih fokus dalam pembuatan testing yang efisien. 

Structure dalam unit testing adalah strategi yang di implementasikan dalam menaruh file testing dalam projek. Pattern yang sering digunakan :
- Centralize => menaruh semua file testing di dalam satu folder. kelebihannya, direktori akan lebih bersih.
- Together => menaruh setiap file testing di dalam folder yang sama dengan file produksinya. kelebihannya, akan lebih mudah mencari test file dari file asliya.

Komponen penting dalam unit testing :
- Test file => file yang berisi testing dan kumpulan test Suites
- Test Suites => berisi koleksi test cases
- Test Fixtures => proses/script untuk memastikan environment yang digunakan tetap konsisten
- Test Cases => kumpulan kondisi/skenario yang akan dijalankan, berisi sebuah input dan expected output dari program.

Runner adalah aplikasi yang digunakan untuk menjalankan file testing. Fitur yang disediakan runner :
- Tool that run the best file
- Use watch mode
- Choose a runner that canrun fastest

Mocking adalah objek tiruan yang dibuat untuk menyerupai hasil dari sistem luar yang digunakan.

Test case harus seindependen mungkin karena hanya ingin fokus melakukan test pada sistem kita saja, tidak perlu mengikutsertakan sistem luar. Mocking lah salah satu upaya agar hal ini terwujud.

Hal yang tidak disarankan pada saat testing :
- 3rd party modules
- Database
- 3rd party api or other external system
- Object class that you have test in other place

Testing coverage adalah alat ukur yang menunjukkan source code yang sudah di execute saat menjalankan test. Programmer harus memastikan test yang dibuat sudah cukup dan memiliki segala logika yang mungkin. coverage dapat dilihat pada saat file test dijalankan. Bagian yang dianalisa coverage Tool adalah function, statement, branch, dan lines. Report dalam format CLI, XML, HTML, LCOV. 

Tahapan untuk melakukan testing :
1. Membuat test file baru
    - diharuskan membuat nama file dengan format libarary_test.go
    - Lokasi file di same folder, same package atau same folder, different package
2. Tulis test function
    - Nama func : Test & Capitalized word
    - Harus memiliki paramaeter t *testing.T
3. Pahami apa yang ditulis di file testing
    - Memiliki lebih dari 1 test function
    - Mengikuti aturan testing di golang

Test scenario adalah gambaran umum terhadap apa yang akan di tes dan disebut juga sebagai haigh level test case. Test case adalah kumpulan langkah-langkah uji positif dan negatif dalam sebuah test scenario. Test case memiliki pre-conditions, steps, expected result, status and actual result. Test scenario membahas apa yang akan di tes, sedangkan test case adalah bagaimana cara melakukan tesnya.