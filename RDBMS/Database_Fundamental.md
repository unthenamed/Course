
### INTRODUCTION (Chapter 1)
---
> Database adalah tempat menimpan data secara terstruktur dan terorganisir.
> Database menimpan data berbentuk table. Dan didalam database yang sama bisa terdapat banyak table.

Tabel memiliki banyak bagian,
* Filed (Kolom)
* Record (Baris)

### INSTALASI POSTGRESSQL (Chapter 2)
---

### TIPE DATA
---
Tipe data di database yang paling banyak digunakan

| TIPE DATA | PENGGUNAAN |
| :------: | :------: |
| VARCHAR(n) | Digunakan untuk menyimpan text alphanumeric sepanjang "n" |
| INTEGER | Digunakan untuk menyimpan Angka |
| DATE | Digunakan untuk menyimpan tanggal |
| Content 4 | Content 4 |
| Content 5 | Content 5 |

-
### CONSTRAINT
---
Constraint adalah batasan atau aturan yang di terapkan pada table untuk menjaga konsistensi data.

| CONSTRAIN | PENJELASAN |
| :------: | :------: |
| NOT NULL | Memasitikan bahwa kolom tidak boleh memiliki nilai NULL |
| UNIQUE | Memastikan  bahwa semua nilai di kolom berbeda |
| PRIMARY KEY | Suatu nilai yang di gunakan untuk mengidentifikasi suatu baris dalam table |
| FOREIGN KEY | Atribut dalam table yang merujuk ke Primary Key dari table lain |




-
### SQL INTRODUCTION (Chapter 3)
---
Apa itu SQL ? 
Sql merupakan kepanjangan Dari Structured Query Language.
Sql merupakan bahasa yang di gunakan untuk mengirim perintah ke RDBMS.
Sql Adalah bahasa yang mudah karna hanya berisi intruksi untuk menyimpan, mengubah, menghapus atau mengambil data melalui RDBMS.


```
    Terminal psql
    - Show DB              namadb=# \l 
    - Connect DB           namadb=# \c 
    - Show table           namadb=# \d 
    - Show structur table  namadb=# \d namedb
```
Berikut adalah grup perintah dalam sql :


#### DDL (Data Definition Language) (Chapter 4)
DDL merupakan perintah SQL yang di gunakan untuk mendefinisikan skema database.
Seperti membuat, dan memodifikasi truktur objek database.
DDL Mendefinisikan struktur table dan barisnnya.
Dan objek yang di maksud bisa berupa Database itu sendiri atau table nya.

**CREATE**

> Digunakan untuk membuat Database atau table

``` sql
-- Create DB
CREATE DATABASE namedb ; 

-- Create Table
CREATE TABLE namedb (table1 TIPEDATA CONSTRAINT, table2 TIPEDATA CONSTRAINT, table3 TIEDATA);
```

**ALTER (Chapter 6)**

Digunakan untuk Mengubah struktur Database atau table
Perintah ALTER digunakan untuk dapat melakukan atau mengubah suatu struktur pada table yang sebelumnya pernah di buat.
Jenis perintah DDL ALTER :

1. ALTER RENAME 
> Digunakan untuk mengubah nama pada DataBase, Table, atau Kolom


``` sql
-- Mengubah nama Database
ALTER DATABASE nama_database RENAME TO nama_baru_database;

-- Mengubah nama Table
ALTER TABLE nama_table RENAME TO nama_baru_table;

-- Mengubah nama Kolom	
ALTER TABLE nama_table RENAME COLUMN nama_kolom
TO nama_kolom_baru;
```
            
2. ALTER ADD
> Digunakan untuk menambahkan kolom/constraint baru.

``` sql
-- Menambahkan kolom baru pada table
ALTER TABLE nama_table ADD COLUMN nama_kolom tipe_data;

-- Menambahkan constraint
ALTER TABLE nama_table ADD CONSTRAINT nama_label_constraint
TIPE_CONSTRAINTNYA (nama_kolom);
```

3. ALTER DROP	
> Digunakan untuk Menghapus sebuab kolom/constraint.

``` sql
-- Menghapus Kolom
ALTER TABLE nama_table DROP COLUMN nama_kolom;

-- Menghapus Constraint
ALTER TABLE nama_table DROP CONSTRAINT nama_label_constraint;
```
            
4. ALTER MODIFY
> Digunakan untuk melalukan perubahan pada nama/tipedata dalam database.

- Mengubah Tipe Data
``` sql
//ALTER MODIFY COLUMN DATA TYPE
ALTER TABLE nama_table ALTER COLUMN nama_column TYPE tipe_data_baru(panjang_data);
```
>  - Perntah modify ini bisa digunakan walaupun kolom sudah memiliki record dengan syarat untuk perubahan ke tipe data VARCHAR panjang data yang di tuliskan tidak kurang dari panjang data record.
> - Dan tipe data baru harus menyesuaikan isi record dari data sebelumnya, misalnya record dari kolom VARCHAR di ubah ke INTEGER maka akan terjadi error

- SET NOT NULL
``` sql
//ALTER MODIFY SET NOT NULL  (Menambahkan Not Null)
ALTER TABLE nama_table ALTER COLUMN nama_column SET NOT NULL;
```


**DROP**
> Digunakan untuk menghapus Database atau Tabel.

``` sql
-- Drop DB
DROP DATABASE namedb;

-- Drop Table
DROP TABLE namedb; 
```
                     
**RENAME**
> Digunakan untuk mengganti nama
    
**TRUNCATE**
> Digunakan untuk menghapus/ mengosongkan semua record/data yang tersimpan pada sebuah table
   

#### DML (Data Manipulation Language) (Chapter 5)
   DMl merupakan perintah SQL yang di gunakan untuk memanipulasi data yang tersimpan dalam database.
   
**INSERT**
>Digunakan untuk memasukan data kedalam table

``` sql
INSERT INTO table_name (column1,column2,column3,...columnN)
VALUE (value1,value2,value3,...valueN);
```
>- Untuk kolom bertipedata VARCHAR, DATE(YYYY-MM-DD) harus menggunakan tanda petik satu ('n')
>- Jika meninputkan value yang sama pada kolom yang berconstrain PRIMARY KEY maka akan error.
>- Jika pada salah satu kolom berconstrain NOT NULL tidak di sertakan atau tidak di isi maka akan error.
>- Untuk yang tidak berconstrain NOT NULL bisa di kosongkan dengan tidak memyertakan nama kolom dan values.

**UPDATE**
> Digunakan untuk memperbarui data yang ada di dalam table

``` sql
// syntax
UPDATE table_name
SET column1=value1, column2=value2, ...columnN=valueN
WHERE [condition];
```

``` sql

/* memperbarui semua baris data
yang ada didalam kolom.
atau tanpa WHERE.*/

UPDATE table_name SET column=value;

/* memperbarui value kolom di
          salah satu baris.*/

UPDATE table_name SET column=value
WHERE column=value; --sebagai pencari baris
 
          /* memperbarui beberapa value kolom
          pada salah satu baris	*/

UPDATE table_name SET column1=value, column2=value
WHERE column=value --sebagai pencari baris
```
                                               
**DELETE** 
> Digunakam untuk menghapus data dari table

``` sql
// syntax
DELETE FROM table_name
WHERE [condition];
```

``` sql
            /* Delete jika tidak di beri WHERE
            maka data yang ada table akan
            terhapus semua, */
DELETE FROM table_name; 

            /* Menghapus salah satu baris */
DELETE FROM table_name
WHERE column=value;	
```

#### DQL (DATA Query Language) (Chapter 7)
---
DQL merupakan perintah SQL untuk menjalankan query atau menampilkan data ynh tersimpan dalam database.

```
[SELECT] [ORDER BY]
[SELECT] [WHERE]
[SELECT] [WHERE] [BETWEEN]
[SELECT] [WHERE] [LIKE]
```

***SELECT***
> Digunakan untuk mengambil / menampilkan data dari database.

``` sql
// syntax
SELECT column1, column2, columN
FROM table_name;

// menampilkan semua kolom
SELECT * FROM table_name;
```

- Operasi Aritmatika

``` sql
SELECT column1, (2024 - column2)
FROM table_name;

SELECT column1, ((2024 - column2)*2)
FROM table_name;
```

- Alias 

``` sql
SELECT column1 AS 'nama_alias' , column2 AS 'nama_alias'
FROM table_name;

SELECT column1, (2024 - column2) AS 'alias_name1' , ((2024 - column2)*2) AS 'Alias_name2'
FROM table_name;
```

**SELECT yang dikombinasikan ORDER BY**
> Digunakan untuk mengurutkan data saat di tampilkan sesuai urutan kolom tertentu

``` sql
SELECT column1, column2, columN
FROM table_name;
ORDER BY columnN [ASC|DESC];

SELECT column1, column2, columN
FROM table_name;
ORDER BY column1 [ASC|DESC], columnN [ASC|DESC];
```
> - ASC (Ascending) diurutkan dari nilai yang terkecil sampai yang terbesar.
> - DESC (Descending) diurutkan dari nilai yang terbesar sampai yanh terkecil.

**SELECT yang di kombinasikan WHERE CLAUSE (Chapter 8)**

* WHERE
> Digunakan untuk menampilkan data yang sesuau dengan kondisi yang di berikan.
    
``` sql
SELECT column1, column2, columnN
FROM table_name
WHERE condition;
-- WHERE condition [AND | OR]  condition;
```
``` sql
-- contoh
SELECT * FROM mst_student WHERE address = 'Jakarta';
```

`condition` disini bisa berupa comparsion operator

| Operator | Description |
| :------: | :------: |
| < | Kurang dari |
| > | lebih dari  |
| <= | kurang dari atau samadengan |
| >= | lebihdari atau samadengan  |
| = | sama dengan |
| <> or != | tidak samadengan |

``` sql
-- contoh
SELECT * FROM mst_student WHERE entry_year > 2018;
```

`condition` juga bisa menggunakan Gerbang Logika [AND|OR]

Logika AND

| X | Y | X And Y |
| :------: | :------: | :------: |
| false | false | false |
| false | true | false |
| true | false | false |
| true | true | true |

Logika OR

| X | Y | X OR Y |
| :------: | :------: | :------: |
| false | false | false |
| false | true | true |
| true | false | true |
| true | true | true |

``` sql
-- contoh

/* Menampilkan siswa yang hanya mendaftar dari tahun 2018 dan berdomisili di Jakarta */
SELECT * FROM mst_student WHERE entry_year >= 2018 AND address = 'Jakarta';

/* Menampilkan siswa hanya yang berdomisili di jakata atau di siduarjo */
SELECT * FROM mst_student WHERE address = 'Jakarta' OR address 'Sidoarjo'

/* Menampilkan siswa yang mendaftar dengan rentang tahun 2015 sampai 2018 */
SELECT * FROM mst_student WHERE entrily_year >= 2015 OR entry_year <= 2018;
```
Penggunaan clousa `BETWEEN` untuk query yang sama
```
column_name BETWEEN condition AND condition
syntax ini sama dengan :
column_name >= value AND column_name <= value
```
``` sql
-- contoh :
/* Menampilkan siswa yang mendaftar dengan rentang tahun 2015 sampai 2018 */

SELECT * FROM mst_student WHERE entry_year >= 2015 OR entry_year <= 2018;
-- sama dengan :
SELECT * FROM mst_student WHERE entry_year 2015 AND 2018 ;
```

* LIKE
> Digunakan untuk mencocokan nilai dengan pola tertenty menggunakan `wildcards` yang tipe datanya berupa text.
>
> WILDCARD adalah simbol yang digunakan utuk menggantikan atau mewakili suatu atau lebih karakter
>
> LIKE menggunakan `Case Sensitif`
Ada 2 simbol yang bisa di gunakan dalam LIKE operator :

| SIMBOL | DIGUNAKAN |
| :------: | :------: |
| Persen (%) | Simbol persen mewakili nol, satu, atau beberapa angka atau karakter |
| Underscore (_) | Simbol garis bawah mewakili satu angka atau karakter |
``` sql
-- Syntax
SELECT * FROM table_name WHERE column LIKE '%XXXX%';
SELECT * FROM table_name WHERE column LIKE '_XXXX_';
```
``` sql
-- Contoh :

/* Menampilkan siswa yang nama awalannya menggunakan hurup K */
SELECT * FROM mst_student WHERE name LIKE 'K%';

- Contoh lain :
SELECT * FROM mst_student WHERE name LIKE '%ali%';
SELECT * FROM mst_student WHERE name LIKE '_enji';
```

> LIKE menggunakan sensitif case, dan jika kita ingin menggunakan LIKE tanpa sensitif case gunakan clouse `ILIKE`
* ILIKE
``` sql
SELECT * FROM mst_student WHERE name ILIKE 'K%';

/* Maka semua nama siawa yang mempunyai awalan hurup K akan di tampilkan baik nama yang undercase atau uppercase */
```
* NOT LIKE
> Digunakan untuk menampilkan selain data yang di sebutkan, atau kebalikannya.
``` sql
SELECT * FROM mst_student WHERE name NOT LIKE 'K%';

/* Maka hanya sisa dengan nama selain K yang akan di tampilkan */
```
* IN
> Digunakan untuk mengecek apakah data tersebut ada dalam suatu list data. 

``` sql
-- Syntax

SELECT column1, column2, columnN
FROM table_name
WHERE column_name
IN ('value','value2','value3');
```
``` sql
-- Contoh:

SELECT * FROM mst_student WHERE address IN ( 'Jakarta','Sidoarjo','Malang');

-- Atau sama saja jika kita menggunakan OR

SELECT * FROM mst_student WHERE address = 'Jakarta' OR address = 'Sidoarjo' OR address = 'Malang' ;
```
#### DCL (Data Control Language)
DCL merupakan perintah SQL yang berhubungan dengan hak akses dan kontrol database
****GRANT****
****REVOKE****