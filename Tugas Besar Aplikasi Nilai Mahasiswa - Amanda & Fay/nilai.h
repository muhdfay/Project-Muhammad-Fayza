#ifndef NILAI_H_INCLUDED
#define NILAI_H_INCLUDED
#include "mataKuliah.h"

#include <iostream>

#define first(L) (L.first)
//#define info(P) (P)->info
#define next(P) (P)->next
#define prev(P) (P)->prev

using namespace std;

struct nilai {
    float UTS, UAS, QUIZ, GRADE, IPK;
};

typedef struct nilai infotypeNilai;
typedef struct elmListNilai *adrNilai;

struct elmListNilai{
    infotypeNilai info;
    adrNilai next;
    adrNilai prev;
    //adrRelasi relasi;
};

struct listNilai{
    adrNilai first;
};

void createListNilai(listNilai &L);
adrNilai alokasiNilai(infotypeNilai x);
void insertDataNilaiFIRST(listNilai &L, adrNilai P);
void showAllDataNilai(listNilai L);
void deleteDataNilaiFIRST(listNilai &L, adrNilai &P);
adrNilai findDataNilai(float UAS, listNilai L);

#endif // NILAI_H_INCLUDED
