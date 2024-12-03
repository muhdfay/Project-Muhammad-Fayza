#ifndef MAHASISWA_H_INCLUDED
#define MAHASISWA_H_INCLUDED

#include <iostream>

#define first(L) (L.first)
#define last(L) (L.last)
#define info(P) (P)->info
#define next(P) (P)->next

using namespace std;

struct dataMahasiswa {
    int NIM;
    string nama, jenisKelamin, programStudi;
};

typedef struct dataMahasiswa infotypeMhs;
typedef struct elmListMhs *adrMhs;

struct elmListMhs {
    infotypeMhs info;
    adrMhs next;
    //adrRelasi Relasi;

};

struct listMhs {
    adrMhs first;
    adrMhs last;
};

void createListMhs(listMhs &L);
infotypeMhs createInfoMhs(int NIM, string nama, string jenisKelamin, string programStudi);
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

//PROSEDUR YANG BELUM :
//Hubungkan Data Mahasiswa dengan Mata Kuliah
//Tampilkan Semua Data Mahasiswa dan Mata Kuliah
//Tampilkan Mata Kuliah yang Diambil oleh Mahasiswa Tertentu
//Hapus Hubungan Data Mahasiswa dengan Mata Kuliah
//Hitung Jumlah Mata Kuliah yang Diambil oleh Mahasiswa

#endif // MAHASISWA_H_INCLUDED
