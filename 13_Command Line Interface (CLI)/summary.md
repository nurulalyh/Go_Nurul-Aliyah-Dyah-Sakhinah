# (13) Configuration Management and CLI

## Introduction Hook
OS memiliki 2 jenis, yaitu GUI dan CLI. GUI merupakan tampilan yang memiliki interface. Sedangkan CLI hanya berupa teks. Command Line lebih cepat dan powerfull karena hanya perlu mengetikkan perintah, sedangkan GUI lebih lambat karena GUI perlu waktu untuk meload tampilan setiap kali di klik. Alasan menggunakan CLI :
1. Bisa memanajemen lebih cepat.
2. Bisa melakukan controlling access (sebagai user atau admin) ke OS.
3. Bisa menyimpan beberapa perintah di dalam sebuah file .sh sehingga hanya perlu dijalankan.
4. Lebih memudahkan karena dapat memberikan informasi untuk membantu menyelesaikan masalah troubleshooting.

## Directory Command
1. pwd => See Current Directory
2. ls => List directory, -a => see hidden file, -l => see the detail, -R => see folder and subdir
3. mkdir => Lake directory
4. cd => Go to directory
5. rm => Delete file, rmdir => Delete empty directory, rm -r => delete folder dan isi
6. cp => Copy file into directory
7. mv => move file into directory
8. ln => make link, -s => softlink

## Files Command
1. create file => touch
2. view file => head, cat, tail, less
3. editor => vim, nano
4. permission => chown, chmod
5. different => diff

## Network Command
1. ping => tes connection IP
2. ssh => connetct ke remote server
3. netstat => mengetahui apa yang terjadi di jaringan kita
4. nmap => analisis port
5. ip addr (ifconfig) => menampilkan list ip, mengubah ip dan subnet
6. wget => mengunduh file dari url
7. curl => mengunduh file dari url
8. telnet => untuk remote server dengan keamanan rendah
9. netcat => monitoring jaringan

## Utility Command
1. man => untuk informasi tentang perintah yang dibingungkan
2. env => environment yang bisa menggunakan variabel
3. echo => untuk menampilkan text
4. date => untuk menampilkan waktu
5. which => mengetahui lokasi file
6. watch => untuk memonitoring
7. sudo => untuk masuk ke mode admin
8. history => untuk menampilkan history Command CLI
9. grep => mencari di dalam file bagian yang mengandung kata yang dicari
10. locate => untuk pencarian ke dalam file

## File SH
Shell merupakan program berisi function sebagai jembatan antaran user dan kernel. Shell Script adalah bahasa pemrograman yang di compiled berdasarkan shell commands. Digunakan untuk mengautomasi.