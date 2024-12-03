#ifndef MATAKULIAH_H_INCLUDED
#define MATAKULIAH_H_INCLUDED

#include <iostream>

#define first(L) (L.first)
#define info(P) (P)->info
#define next(P) (P)->next
#define prev(P) (P)->prev

using namespace std;

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
    //adrRelasi nextR;
};

struct listMatkul {
    adrMatkul first;
};

void createListMatKul(listMatkul &L);
adrMatkul alokasiMatkul(infotypeMatkul x);
void insertDataMatkulFIRST(listMatkul &L, adrMatkul P);
void insertDataMatkulAFTER(listMatkul &L, adrMatkul Prec, adrMatkul P);
void insertDataMatkulLAST(listMatkul &L, adrMatkul P);
void showAllDataMatkul(listMatkul L);
void editDataMatkul(string namaMatkul, listMatkul &L);
void deleteDataMatkulFIRST(listMatkul &L, adrMatkul &P);
void deleteDataMatkulAFTER(listMatkul &L, adrMatkul Prec, adrMatkul &P);
void deleteDataMatkulLAST(listMatkul &L, adrMatkul &P);
void removeDataMatkul(string namaMatkul, listMatkul &L);
adrMatkul findDataMatkul(string namaMatkul, listMatkul L);

//PROSEDUR YANG BELUM :
//Hubungkan Data Mahasiswa dengan Mata Kuliah
//Tampilkan Semua Data Mahasiswa dan Mata Kuliah
//Tampilkan Mata Kuliah yang Diambil oleh Mahasiswa Tertentu
//Hapus Hubungan Data Mahasiswa dengan Mata Kuliah
//Hitung Jumlah Mata Kuliah yang Diambil oleh Mahasiswa

#endif // MATAKULIAH_H_INCLUDED
