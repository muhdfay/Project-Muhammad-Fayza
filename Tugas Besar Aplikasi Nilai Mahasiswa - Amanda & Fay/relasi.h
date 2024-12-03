#ifndef RELASI_H_INCLUDED
#define RELASI_H_INCLUDED

#include "mahasiswa.h"
#include "matakuliah.h"
#include "nilai.h"
#include<iostream>
using namespace std;

#define first(L) (L.first)
#define info(P) (P)->info
#define next(P) (P)->next
#define prev(P) (P)->prev
#define Mhs(P) (P)->mahasiswa
#define Matkul(P) (P)->mataKuliah
#define Nilai(P) (P)->nilai

typedef struct relasi infotypeRelasi;
typedef struct elmListRelasi *adrRelasi;

struct relasi {
    adrMhs mahasiswa;
    adrMatkul mataKuliah;
    adrNilai nilai;
};

struct elmListRelasi {
    infotypeRelasi info;
    adrRelasi next;
    adrRelasi prev;
    adrMatkul mataKuliah;
    adrMhs mahasiswa;
    adrNilai nilai;
};

struct listRelasi {
    adrRelasi first;
};

void createListRelasi(listRelasi &L);
adrRelasi alokasiRelasi(adrMhs P, adrMatkul Q, adrNilai N);
void insertRelasi(listRelasi &L, adrRelasi P);
void deleteRelasi(listRelasi &L, adrRelasi &P);
void connect(listRelasi &L, listMhs x, listMatkul y, listNilai n, adrMhs P, adrMatkul Q, adrNilai N);
void disconnect(listRelasi &L, adrMhs P, adrMatkul Q, adrNilai N);
adrRelasi findRelasi(listRelasi L, adrMhs mahasiswa, adrMatkul mataKuliah, adrNilai nilai);
void showAllDataMahasiswaMatkul(listRelasi L);
void showMatkulMahasiswa(listRelasi L, adrMhs mahasiswa);
int countMatkulMahasiswa(listRelasi L, adrMhs mahasiswa);

#endif // RELASI_H_INCLUDED
