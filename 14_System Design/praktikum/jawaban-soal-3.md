## Soal

Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;` Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
- Redis
- Neo4j
- Cassandra

## Jawaban 

**Redis**

Redis yang saya tau bukan jenis DBMS yang dirancang untuk pengambilan data yang kompleks. Jadi query yang ada hanya untuk menampilkan data dari users berdasarkan key atau id nya.

`HGETALL user:1`

`HGETALL user:2`

`HGETALL user:3`

`--- dan seterusnya`

**Neo4j**

`MATCH (u:User)`

`RETURN u;`

Neo4j akan mencocokan node yang memiliki label user dan mengembalikannya dalam bentuk hasil query grafik.

**Cassandra**

Cassandra memiliki query select yang serupa dengan SQL dalam kasus ini.

`SELECT * FROM users;`