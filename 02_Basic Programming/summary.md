# (2) Basic Programming

## Golang
Golang merupakan salah satu bahasa pemrograman yang open source yang dikembangkan oleh google untuk backend yang sedang populer saat ini. Bahasa ini dikembangkan dengan tujuan memudahkan development tapi tetap simple, efisien, dan handal. Di masa depan, golang digunakan di banyak bidang mulai dari devOps hingga data scientist.Contoh penggunaan Golang untuk development yaitu aplikasi e-commerce, social media, API, dan aplikasi CLI. 
<!-- Sejarah golang :
1. akhir 2007 - pertengahan 2008 => Pengembangan.
2. November 2009 => Dirilis ke publik.
3. Maret 2012 => Rilis versi 1.0.
4. Februari 2019 => Rilis versi 1.12.
5. September 2021 => Rilis versi 1.16. -->

Persiapan untuk memulai mempelajari Golang:
1. IDE => Vs Code, GoLand, Atom, Vim
2. Golang Plugin
3. Go Workspace :
    * bin : berisi binary
    * pkg : berisi go package
    * src : berisi go source file
4. Setting Gopath

## Variabel dan tipe data
Variabel merupakan tempat untuk menyimpan data yang memiliki nama yang unik yang akan dipanggil jika data yang disimpan oleh variabel tersebut dibutuhkan. Variabel hanya dapat menampung data dengan tipe data yang sama. Tipe data pada golang :
1. Boolean
    * True
    * False
2. Numeric
    * Integer (bilangan bulat)
        - uint8, uint16, uint32, uint64 (unsigned dengan nilai min 0)
        - int8, int16, int32, int64
        - rune (int32)
        - uint (int32/int64)
        - byte (uint8)
    * Float (desimal)
        - float32
        - float64
    * Complex
        - complex64
        - complex128
3. String

Zero value tipe data golang :
1. Boolean = `False`
2. Floats = `0.0`
3. Integer = `0`
4. String = `""`

Konstanta merupakan jenis variabel yang nilai nya tidak bisa dirubah.

## Operator, Expression, Control Structure (Branching, Looping)
Dalam golang, ada istilah operand dan operator. Operand merupakan bagian paling sederhana dari expression seperti konstanta, sedangkan operator merupakan simbol yang digunakan dalam proses operasi.

Expression digunakan dalam program untuk membuat program dapat berjalan lebih mudah.

Untuk menggunakan branching (percabangan), dapat menggunakan keyword `if, else if, else` atau `switch case`.

Looping (perulangan) menggunakan keyword `for`, `continue`, dan `break`. `continue` memaksakan perulangan berlanjut atau diskip ke perulangan selanjurnya, sedangkan `break` menghentikan paksa perulangan.

`%c` dan `%d` merupakan format verb yang memiliki fungsi tersendiri.

<!-- ### Note
Syntax pada golang :
1. `fmt` => package untuk menghandle input-output pada golang
2. Deklarasi variabel :
    * `var <nama_variabel> <tipe_data>`
    * `var <nama_variabel> <tipe_data> = <value>`
    * `var <list_nama_variabel> <tipe_data>`
    * `var <list_nama_variabel> <tipe_data> = <value1>, <value2>`
    * `<nama_variabel> := <value>` => tipe data menyesuaikan value
3. `const <nama_variabel> <tipe_data> = <value>` => Deklarasi single constanta
4. `const (`
    
    `<nama_variabel1> <tipe_data1> = <value1>`
    
    `<nama_variabel2> <tipe_data2> = <value2>`
`)` => Deklarasi multiple constanta

5. `if <kondisi1> {`
        
        `perintah1`
    
    `} else if <kondisi2> {`
       
        `perintah2`
   
    `} else {`
   
        `perintah3`
   
    `}` => Penulisan branching `if, else if, else`
6. `switch <variabel> {`

    `case 1 :`
     
        `perintah1`
        
    `case 2 :` 

        `perintah2`

    `default :` 

        `perintah3`

    `}` => Penulisan branching `switch case` -->