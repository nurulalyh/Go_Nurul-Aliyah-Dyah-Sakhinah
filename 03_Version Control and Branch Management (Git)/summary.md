# (3) Version Control and Branch Management (Git)

1. Versioning adalah mengatur versi dari source code program.
2. **Projek** yang dikerjakan sebaiknya jika ingin **dilakukan perubahan**, akan lebih baik jika **tidak mengubah langsung** pada **file asli** tapi dengan melakakukan **perubahan pada salinannya** agar dapat melihat perubahan apa saja yang dilakukan. Tujuannya, untuk mengetrack letak kesalahan, apa yang diubah, dan siapa yang mengubah, agar lebih praktis dalam memanajemen projek.
3. Tools :
    * Version Control System (VCS) => Now
    * Source Control Manager (SCM) => Past
    * Revision Control System (RCS) => Past
4. GIT merupakan salah satu VCS yang populer yang digunakan developer untuk mengembangkan software secara bersama-sama. GIT terdistribusi bukan tersentralisasi, yang berarti GIT menghubungkan banyak device tapi tidak terpusat pada server. GIT dibuat oleh **Linus Torvalds** pada 2005. 
5. GIT bisa melakukan semacam undo dengan menggunakan reset sehingga dapat kembali ke poin tertentu yang disebut commit.
6. Readme.md merupakan dokumentasi dari sebuah repository. 
7. .gitignore merupakan file yang dapat digunakan untuk membatasi file yang akan di publikasikan di github. 
8. License digunakan untuk membatasi akses dari viewers pada sebuah repository.
9. Cara kerja di dalam GIT :
    * dibagi 3 area di dalam local directory :
        - working directory => file source code
        - staging area => area sebelum commit, yaitu saat melakukan git add.
        - repository => saat melakukan commit. commit message harus singkat namun jelas mengenai perubahan.
10. Command GIT :
    * git init = untuk menginisialisasi directory baru di git
    * git add remote origin = untuk menghubungkan directory git dengan repository github
    * git add = untuk mengirim file dari local ke staging area
    * git commit -m "" = untuk memasukkan pesan commit
    * git push origin master/main = untuk mengupload file yang di staging area ke repository github tergantung branch yang dipilih.
    * git status = untuk mengecek status dari working tree dan branch yang digunakan.
    * git diff = untuk mengetahui perubahan dalam projek
    * git stash = merupakan penyimpanan sementara.
    * git log = untuk melihat alamat dari perubahan apa saja yang dilakukan
    * git checkout = untuk membuat branch baru + langsung pindah branch + projek terakhir dari branch lain akan ikut ke branch baru. bisa juga pindah ke branch yang sudah ada.
    * git reset = ada 2 jenis, yaitu --soft dan --hard. --soft perubahan yang ingin di undo akan tetap terdeteksi di staging area, sedangkan --hard tidak akan terdeteksi tapi tetap ada di local.
    * git fetch = untuk mengambil data yang ada di server kita tapi tidak sampe masuk di local.
    * git pull = untuk menarik perubahan dari branch yang dipilih dan dimasukkan ke directory local/branch kita.
    * git branch <nama_branch> = untuk membuat brance baru
    * git branch --list = untuk melihat list branch yang dilocal
    * git branch -a = untuk melihat list branch dari origin
    * git branch -D <nama_branch> = untuk menghapus branch
    * git merge = untuk mengambil perubahan dari branch lain
11. Branches untuk mengisolasi bagian projek dari perubahan-perubahan yang dilakukan oleh orang lain.
12. Pull request yaitu minta develoment untuk merge fitur yang kita telah buat ke projek utama.
13. Workflow Collaboration merupakan alur dalam menggunakan git. 
14. Tips :
    * Branch master jangan diubah-ubah, hanya menjadi tempat untuk menerima pull request yang sudah di tes.
    * Brach master hanya mengambil pull request dari branch development, sedangkan branch development hanya mengambil pull request untuk branch fitur dan master.
    * Jangan edit branch development, lebih baik buat brach baru.
    * setiap fitur, kalau sudah selesai, minta pull request ke branch development.
    * setelah branch development fix, baru minta pull request ke branch master.