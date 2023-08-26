# (6) String – Advance Function – Pointer – Method – Struct and Interface

## String
String merupakan salah satu variabel yang penting ketika ingin membuat program. Di golang, string memiliki package sendiri yang memuat banyak fungsi untuk manipulasi string. 

## Advance Function
Macam - macam Advance function :
1. Variadic Func => digunakan ketika argument yang akan di input ke func bisa berbeda dari parameter yang dibuat. Solusinya adalah membuat temporary slice (`...`) sebagai parameter sehingga argument tetap ditulis seperti biasa.
2. Anonymous Function/Literal Function => merupakan function yang tidak punya nama yang digunakan untuk mengelompokkan sekelompok kode agar bisa di panggil berkali-kali. Closure merupakan tindak lanjut dari Anonymous function yang akan melakukan reference ketika telah menggunakan Anonymous function variabel diluar. Jadi, ketika variabel dari Anonymous function dipanggil berkali-kali, maka akan melanjutkan proses dari pemanggilan sebelumnya.
3. Defer Function => Kelompok function baik variadic atau anonymous yang didepannya ditaruh kata `defer` akan di run di akhir ketika semua perintah yang ada telah selesai di run. Defer bersifat seperti stack Last In First Out (LIFO), jika ada multiple defer maka yang defer yang terakhir di code yang akan dikeluarkan lebih dulu.

## Pointer
Pointer merupakan variabel yang dapat menyimpan memory address dari variabel yang lain. memory adress dari pointer dan memory adress dari variabel yang disimpan memory addressnya, pasti berbeda. Tipe data variabel pointer harus sama dengan tipe data variabel yang akan disimpan memory adress nya. Memory address dari sebuah variabel pasti dinamis. Variabel pointer dan variabel biasa akan saling berkaitan, sehingga jika salah satu berubah, maka yang lainnya ikut berubah.

Deklarasi variabel pointer => `var p *int = &a`

Mengambil value dari variabel yang disimpan memory addressnya oleh variabel pointer => `*<varPointer>`

Mengambil memory address dari variabel => `&<var>`

## Struct and Interface
Struct merupakan object di golang. Struct biasanya berisi properties atau function (method).

Deklarasi Struct => `type <struct-var-type> struct { field <data_type>; }`

## Method
Method itu seperti function, tapi yang membedakan adalah method memiliki identifier dan method ini merupakan function yang di attach pada sebuah type. Method membantu untuk membuat OOP di golang, membantu menghindari name conflict, serta jauh lebih mudah dibaca dan dipahami. 

Deklarasi method => `func (<receiver_structType>) <Method_Name>(<parameter>) (<return_type>) { //block statement }`