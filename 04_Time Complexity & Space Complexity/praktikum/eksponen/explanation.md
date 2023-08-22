# Time Complexity Program Menghitung Eksponen

Pada program ini, saya mencoba membuat 2 function untuk menyelesaikan kasus mmenghitung eksponen, yaitu fungsi pow dan pow2. Berikut analisis Time Complexity dari kedua function.

1. Function pow() :
    Saya mencoba membuat function ini menggunakan `math.Pow()`. Pada function ini, tidak terdapat perulangan `for` dan berapapun inputannya tidak akan mempengaruhi proses yang ada di dalam function karena hanya dilakukan dalam satu langkat. Maka kompleksitasnya dapat dianggap sebagai O(1).

2. Function pow2() :
    Pada function ini, prosesnya menerapkan rekursif. Pada setiap langkah rekursif, kasus dibagi setengah, dan oleh karena itu setiap langkah hanya memerlukan satu operasi perkalian dan bilangan basis (x) untuk kasus ganjil. Dalam hal ini, terdapat log(n) langkah rekursif, dan di setiap langkah ada operasi perkalian dan pengali pangkat. Oleh karena itu, kompleksitas waktu pow2() dapat dianggap sebagai O(log n).