# (17) ORM and Code Structure (MVC)

ORM (Object Relational Mapping) merupakan tool yang membantu untuk konversi data dari data table ke object (struct). Kelebihan ORM :
- Less repetitive query
- Automatically fetch data into ready to use object
- Simple way if you want to screening data before store it in database
- Some have feature cache query

Kelemahan ORM :
- Add a layer in code and cost the overhead process
- Load un-necesary relationship data
- Complex raw query can be longy to write with ORM (>10 table joins)
- Specific SQL function related to one vendor may not supported or no specific function.

Library ORM yang digunakan di golang adalah GORM. GORM bisa di pairing dengan framework echo.

Database migration merupakan cara untuk mengupdate database ke versi terbaru atau mengupdate strukturnya secara simple dan ter tracking. Bisa juga untuk mengembalikan ke versi sebelumnya. Alasan menggunakan db migration :
- database update simplicity
- database rollback simplicity
- track changes on database Structure
- Database Structure history writed on a code
- always compatible with the application changes.

Ketika menggunakan auto migrate untuk membuat DB, secara otomatis GORM akan membuat nama tabel sesuai nama struct ditambah 's' dan semua nama field akan berubah menjadi huruf kecil. Tapi tipe data yang digunakan adalah tipe data terbesar.

gorm.model merupakan fungsional dari gorm yang akan menyediakan field id, created_at, update_at, dan delete_at saat membuat table db dari kode program.

Refactoring adalah cara untuk memecah kode program menjadi bagian-bagian kecil.

MVC (Model View Controller) adalah salah satu pattern/struktur yang terdiri dari Model (berisi modelling database), View (berisi kode untuk membuat tampilan program), Controller (berisi logic business).