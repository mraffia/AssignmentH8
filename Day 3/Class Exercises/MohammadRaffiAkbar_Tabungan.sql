CREATE TABLE nasabah (
    nasabah_id INT PRIMARY KEY IDENTITY (1, 1),
    nama VARCHAR (255) NOT NULL,
    alamat VARCHAR (255) NOT NULL,
    tanggal_lahir DATETIME,
    tempat_lahir VARCHAR(255),
    gender VARCHAR(255),
    nomor_identitas INT,
    nomor_hp VARCHAR(255),
    nama_ibu_kandung VARCHAR(255)
);

CREATE TABLE produk (
    produk_id INT PRIMARY KEY IDENTITY (1, 1),
    nama VARCHAR (255) NOT NULL,
    deskripsi VARCHAR (255) NOT NULL,
);

CREATE TABLE cabang (
    cabang_id INT PRIMARY KEY IDENTITY (1, 1),
    nama VARCHAR (255) NOT NULL,
    alamat VARCHAR (255) NOT NULL,
);

CREATE TABLE cabang (
    cabang_id INT PRIMARY KEY IDENTITY (1, 1),
    nama VARCHAR (255) NOT NULL,
    alamat VARCHAR (255) NOT NULL,
);

CREATE TABLE rekening (
    rekening_no INT PRIMARY KEY,
    saldo INT NOT NULL,
    tanggal DATETIME NOT NULL,
    status VARCHAR(255) NOT NULL,
    nasabah_id INT NOT NULL,
    cabang_id INT NOT NULL,
    produk_id INT NOT NULL,
    FOREIGN KEY (nasabah_id) REFERENCES nasabah (nasabah_id),
    FOREIGN KEY (cabang_id) REFERENCES cabang (cabang_id),
    FOREIGN KEY (produk_id) REFERENCES produk (produk_id),
);

CREATE TABLE riwayat (
    riwayat_id INT PRIMARY KEY IDENTITY (1, 1),
    jenis_transaksi VARCHAR(255) NOT NULL,
    jumlah INT NOT NULL,
    tanggal DATETIME NOT NULL,
    bank_tujuan VARCHAR(255) NOT NULL,
    rekening_tujuan INT NOT NULL,
    biaya_admin INT NOT NULL,
    rekening_no INT NOT NULL,
    FOREIGN KEY (rekening_no) REFERENCES rekening (rekening_no)
);

INSERT INTO produk
VALUES 
('BRITAMA', 'Tabungan'),
('SIMPEDES', 'Tabungan'),
('Simpanan Pelajar', 'Tabungan');

INSERT INTO cabang
VALUES 
('Gatsu', 'Jl. Gatot Subroto No.Kav. 9-11'),
('Cibubur', 'Jl. Alternatif Cibubur No. 19'),
('Tangmer', 'Jl. Raya Merdeka No. 108');

INSERT INTO nasabah
VALUES 
('Raffi', 'Kalibata City', '1998/08/03', 'Sumbawa Besar', 'Pria', 123456789, '081927110578', 'Martha'),
('Iqbal', 'Kalibata', '1998/08/04', 'Sumbawa', 'Pria', 123456781, '081927110579', 'Tilaar');

INSERT INTO rekening
VALUES 
(123456, 250000, '2023/08/03', 'Baik', 1, 1, 1),
(123457, 250000, '2023/08/04', 'Baik', 2, 2, 1);

INSERT INTO riwayat
VALUES 
('Setor Tunai', 50000, '2023/08/05', 'BRI', 123456, 0, 123456),
('Tarik Tunai', 50000, '2023/08/07', 'BRI', 123457, 0, 123457);

