#ifndef MLL_H_INCLUDED
#define MLL_H_INCLUDED

#include <iostream>
using namespace std;

#include <vector>

#ifdef _WIN32
#include <bits/stdc++.h>
#endif // _WIN32

#define first(L) (L.first)
#define last(L) (L.last)
#define info(P) (P)->info
#define next(P) (P)->next
#define prev(P) (P)->prev

// MARK : - Mahasiswa
struct dataMahasiswa {
    int NIM;
    string nama, jenisKelamin, programStudi;
    float ipk;
};

typedef struct dataMahasiswa infotypeMhs;
typedef struct elmListMhs *adrMhs;

struct elmListMhs {
    infotypeMhs info;
    adrMhs next;
    adrRelasi Relasi;
};

struct listMhs {
    adrMhs first;
    adrMhs last;
};

// MARK : - Mata Kuliah
struct mataKuliah {
    int kodeMatkul, sks;
    string namaMatkul;
};

typedef struct mataKuliah infotypeMatkul;
typedef struct elmListMatkul *adrMatkul;

struct elmListMatkul {
    infotypeMatkul info;
    adrMatkul next;
    adrMatkul prev;
};

struct listMatkul {
    adrMatkul first;
};

// MARK : - Nilai
struct nilai {
    int UTS, UAS, quiz, grade;
};

typedef struct nilai infotypeNilai;
typedef struct elmListNilai *adrNilai;

struct elmListNilai{
    infotypeNilai info;
    adrNilai next;
    adrNilai prev;
};

struct listNilai{
    adrNilai first;
};

// MARK : - Relasi
struct relasi {
    adrMhs mahasiswa;
    adrMatkul mataKuliah;
    adrNilai nilai;
};

typedef struct relasi infotypeRelasi;
typedef struct elmListRelasi *adrRelasi;

struct elmListRelasi {
    infotypeRelasi info;
    adrRelasi next;
};

struct listRelasi {
    adrRelasi first;
};

// MARK : - Function/Procedure Mahasiswa
void createListMhs(listMhs &L);
adrMhs alokasiMhs(infotypeMhs x);
void insertDataMhsFIRST(listMhs &L, adrMhs P);
void insertDataMhsAFTER(listMhs &L, adrMhs Prec, adrMhs P);
void insertDataMhsLAST(listMhs &L, adrMhs P);
void showAllDataMhs(listMhs L);
void editDataMhs(int NIM, listMhs &L);
void deleteDataMhsFIRST(listMhs &L, adrMhs &P);
void deleteDataMhsAFTER(listMhs &L, adrMhs Prec, adrMhs &P);
void deleteDataMhsLAST(listMhs &L, adrMhs &P);
void removeDataMhs(int NIM, listMhs &L);
adrMhs findDataMhs(int NIM, listMhs L);


// MARK : - Function/Procedure Mata Kuliah
void createListMatKul(listMatkul &L);
adrMatkul alokasiMatkul(infotypeMatkul x);
void insertDataMatkul(listMatkul &L, adrMatkul P); // menggunakan insert first
void showAllDataMatkul(listMatkul L);
void deleteDataMatkul(listMatkul &L, adrMatkul &P); // menggunakan delete first
adrMatkul findDataMatkul(string namaMatkul, listMatkul L);

// MARK : - Function/Procedure Nilai
void createListNilai(listNilai &L);
adrNilai alokasiNilai(infotypeNilai x);
void insertDataNilai(listNilai &L, adrNilai P); // menggunakan insert first
void showAllDataNilai(listNilai L);
void deleteDataNilai(listNilai &L, adrNilai &P); // menggunakan delete first
//adrNilai findDataNilai(string nilaiMatkul, listNilai L);

// MARK : - Function/Procedure Relasi
adrRelasi createRelasi(adrMhs mahasiswa, adrMatkul mataKuliah, adrNilai nilai);
void insertRelasi(listRelasi &L, adrRelasi P);
void connectMahasiswaMatkul(listRelasi &L, adrMhs mahasiswa, adrMatkul mataKuliah, adrNilai nilai);
void deleteRelasi(listRelasi &L, adrRelasi &P);
adrRelasi findRelasi(listRelasi L, adrMhs mahasiswa, adrMatkul mataKuliah);
void showAllDataMahasiswaMatkul(listRelasi L);
void showMatkulMahasiswa(listRelasi L, adrMhs mahasiswa);
void deleteHubunganMahasiswaMatkul(listRelasi &L, adrMhs mahasiswa, adrMatkul mataKuliah);
int countMatkulMahasiswa(listRelasi L, adrMhs mahasiswa);

// MARK : - Menu
string menu();
string menuMahasiswa();
string menuMatkul();
string menuNilai();

// MARK : - etc
void clearScreen();

// MARK : - Table
