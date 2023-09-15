# (14) System Design

## Diagram
Diagram adalah kumpulan simbol yang merepresentasikan informasi yang ingin disampaikan yaitu system Design. Jenis - jenis software Design :
1. Flowchart => merepresentasikan 1 proses, memberi gambaran alur sistem ke user
2. Use case => Ringkasan dari sistem (isi sistem dan interaksi sistem dengan aktor)
3. ERD => Tipe flowchart yang menggambarkan relasi dari entity, object yang berhubungan dengan sistem
4. HLA => Gambaran tertinggi mengenai teknologi yang digunakan pada projek.

## Distributed System
Distributed System adalah cara agar sistem kita dapat berjalan di multi komputer/server. Karakteristik Distributed system :
1. Scalability
    
    Kapabilitas dari sistem, proses, atau jaringan untuk mengelola lonjakan permintaan. Jika sistem mencapai limit, dapat dilakukan vertical scale dengan mengupgrade teknologi dan hardware yang digunakan. Atau dapat melakukan horizontal scale dengan membeli server yang baru.
2. Reliability
    
    Kehandalan dari sistem, fungsi-fungsi yang harus dapat dilakukan sistem. Dilakukan dengan cara menggunakan banyak server sehingga ketika salah satu server bermasalah, server lainnya dapat mem-backup.
3. Availability

    Ketersediaan sistem, bisa atau tidak sistem diakses dalam 24 jam oleh user. 
4. Efficiency

    Sistem tetap efisien dan tidak memiliki waktu tunggu yang lama. Dilakukan dengan meningkatkan code, bandwith, hingga server.
5. Serviceability of Manageability

    Sistem yang dapat di maintenance, dengan menerapkan clean code yang baik.

## Job Queue and Microservices
Job Queue membuat user agar tetap bisa mengeksplorasi sistem walau permintaannya sedang dalam proses sehingga user tidak hanya menunggu. Dilakukan dengan mengalihkan permintaan tersebut ke Job Queue. Load balancing merupakan component penting di dalam sistem yang memiliki multi server. Load balancing akan membagi trafic ke server yang memiliki space sehingga tidak ada sistem yang terlalu terbebani. Peletakan load balancer :
1. antara user dan web
2. antara web dan server
3. antara server dan database

Monolithic yaitu sistem yang memiliki banyak service dan diletakkan di satu server. Bisa dilakukan ketika permintaan tidak begitu banyak perhari. Microservices merupakan service yang dipecah ke dalam beberapa server.

## SQL and No SQL
SQL adalah relational database yang memiliki entity yang saling berhubungan dan terstruktur di awal. No SQL itu sebaliknya yaitu tidak terstruktur di awal, tidak relasional dan bersifat dinamis. Keuntungan relational DB :
1. Dirancang untuk segala keperluan
2. Memiliki standar yang jelas
3. Memiliki banyak tool 

Prinsip dalam relational database :
1. Atomicity => transaksi terjadi semua atau tidak sama sekali
2. Consistency => data merupakan data valid yang ditentukan berdasarkan aturan tertentu
3. Isolation => Pada saat terjadi concurrent, transaksi dijalankan seakan-akan secara sekuensial
4. Durability => Data dijamin tetap tersimpan

No SQL (Not Only SQL) menyediakan mekanisme yang lebih fleksibel dibandingan dengan RDBMS. Dibuat untuk menghindari :
1. Effort pada sifat transaksi ACID
2. Kompleksitas SQL
3. Design schema di depan
4. Transactions (ditangani oleh aplikasi)

Kelebihan No SQL adalah schema less dan fast development. Digunakan ketika membutuhkan schema fklesibel, ACID tidak diperlukan, data logging (json), data sementara (cache), stc. Jangan digunakan ketika berhubungan dengan data transaksi, data finansial, business critical, SCID - compliant.

## Chache and Indexing
Cache adalah penyimpanan sementara disimpan di dalam RAM, sehingga untuk input yang berulang-ulang tidak perlu melakukan proses keseluruhan karena output dapat ditaruh di cache dan akan langsung ditampilkan ketika menerima input yang sama. Fitur indexing pada database akan melakukan sorting indexing sehingga data terurut. Resiko mengaktifkan fitur ini adalah akan lama create dan update data karena diperhitungkan posisi index dari data. Tapi proses select akan berjalan lebih cepat. Ketika fitur ini diaktifkan, maka akan ada data baru yang masuk yang digunakan menyimpan index.