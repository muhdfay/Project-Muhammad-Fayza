#include "nilai.h"
#include "mataKuliah.h"

void createListNilai(listNilai &L) {
    first(L) = NULL;
}

adrNilai alokasiNilai(infotypeNilai x) {
    adrNilai P = new elmListNilai;
    info(P) = x;
    next(P) = NULL;
    prev(P) = NULL;
    return P;
}

void insertDataNilaiFIRST(listNilai &L, adrNilai P) {
    if (first(L) == NULL) {
        first(L) = P;
    } else {
        next(P) = first(L);
        prev(first(L)) = P;
        first(L) = P;
    }
}

void showAllDataNilai(listNilai L) {
    adrNilai P = first(L);
    while (P != NULL) {
        cout << "UTS: " << info(P).UTS << ", UAS: " << info(P).UAS
             << ", Quiz: " << info(P).QUIZ << ", Grade: " << info(P).GRADE << endl;
        P = next(P);
    }
}

void deleteDataNilaiFIRST(listNilai &L, adrNilai &P) {
    if (first(L) == P) {
        if (next(P) != NULL) {
            first(L) = next(P);
            prev(first(L)) = NULL;
        } else {
            first(L) = NULL;
        }
        delete P;
    } else {
        adrNilai Q = first(L);
        while (next(Q) != P) {
            Q = next(Q);
        }
        next(Q) = next(P);
        if (next(P) != NULL) {
            prev(next(P)) = Q;
        }
        delete P;
    }
}

adrNilai findDataNilai(float UAS, listNilai L) {
    adrNilai P = first(L);
    while (P != NULL) {
        if (info(P).UAS == UAS) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}

