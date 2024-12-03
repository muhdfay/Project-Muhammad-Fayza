#include "mataKuliah.h"

void createListMatKul(listMatkul &L) {
    first(L) = NULL;
}

adrMatkul alokasiMatkul(infotypeMatkul x) {
    adrMatkul P = new elmListMatkul;
    if (P != NULL) {
        info(P) = x;
        next(P) = NULL;
        prev(P) = NULL;
    }
    return P;
}

void insertDataMatkulFIRST(listMatkul &L, adrMatkul P) {
    if (first(L) == NULL) {
        first(L) = P;
        next(P) = P;
        prev(P) = P;
    } else {
        next(P) = first(L);
        prev(P) = prev(first(L));
        next(prev(P)) = P;
        prev(first(L)) = P;
        first(L) = P;
    }
}

void insertDataMatkulAFTER(listMatkul &L, adrMatkul Prec, adrMatkul P) {
    if (first(L) == NULL) {
        cout << "List Mata Kuliah masih kosong, gunakan insertDataMatkulFIRST." << endl;
    } else {
        next(P) = next(Prec);
        prev(P) = Prec;
        prev(next(Prec)) = P;
        next(Prec) = P;

        if (Prec == first(L)) {
            first(L) = P;
        }
    }
}

void insertDataMatkulLAST(listMatkul &L, adrMatkul P) {
    adrMatkul Prec;
    if (first(L) == NULL){
        insertDataMatkulFIRST(L, P);
    } else {
        insertDataMatkulAFTER(L, Prec, P);
    }
}


void showAllDataMatkul(listMatkul L) {
    adrMatkul P = first(L);
    if (P == NULL) {
        cout << "Tidak ada data mata kuliah." << endl;
    } else {
        cout << "Data Mata Kuliah:" << endl;
        while (P != first(L)) {
            cout << "Kode Matkul: " << info(P).kodeMatkul << endl;
            cout << "Nama Matkul: " << info(P).namaMatkul << endl;
            cout << "SKS: " << info(P).sks << endl << endl;
            P = next(P);
        }
    }
}

void editDataMatkul(string namaMatkul, listMatkul &L) {
    adrMatkul P = findDataMatkul(namaMatkul, L);
    if (P != NULL) {
        cout << "Edit Data Mata Kuliah '" << info(P).namaMatkul << "':" << endl;

        cout << "Masukkan Kode Mata Kuliah baru: ";
        cin >> info(P).kodeMatkul;

        cout << "Masukkan Nama Mata Kuliah baru: ";
        cin.ignore();
        getline(cin, info(P).namaMatkul);

        cout << "Masukkan SKS baru: ";
        cin >> info(P).sks;

        cout << "Data Mata Kuliah berhasil diubah." << endl;
    } else {
        cout << "Mata Kuliah dengan nama '" << namaMatkul << "' tidak ditemukan." << endl;
    }
}

void deleteDataMatkulFIRST(listMatkul &L, adrMatkul &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mata kuliah yang dapat dihapus." << endl;
    } else {
        P = first(L);
        if (next(P) == P) {
            first(L) = NULL;
        } else {
            next(prev(P)) = next(P);
            prev(next(P)) = prev(P);
        }
        prev(P) = NULL;
        next(P) = NULL;
        cout << "Data Mata Kuliah pertama berhasil dihapus." << endl;
    }
}

void deleteDataMatkulAFTER(listMatkul &L, adrMatkul Prec, adrMatkul &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mata kuliah yang dapat dihapus." << endl;
    } else {
        if (Prec == first(L)) {
            cout << "Prec berada pada elemen terakhir, tidak bisa menghapus elemen setelahnya." << endl;
        } else {
            P = next(Prec);
            next(Prec) = next(P);
            prev(next(P)) = Prec;

            if (P == first(L)) {
                first(L) = Prec;
            }

            delete P;
            cout << "Data Mata Kuliah setelah elemen Prec berhasil dihapus." << endl;
        }
    }
}

void deleteDataMatkulLAST(listMatkul &L, adrMatkul &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mata kuliah yang dapat dihapus." << endl;
    } else {
        P = prev(first(L));

        if (first(L) == next(P)) {
            first(L) = NULL;
        } else {
            prev(first(L)) = prev(P);
            next(prev(P)) = first(L);
        }

        delete P;
        cout << "Data Mata Kuliah terakhir berhasil dihapus." << endl;
    }
}

void removeDataMatkul(string namaMatkul, listMatkul &L){
    adrMatkul P = findDataMatkul(namaMatkul, L);
    if (P != NULL) {
        if (P == first(L)) {
            deleteDataMatkulFIRST(L, P);
         } else {
            adrMatkul Q = first(L);
            while (next(Q) != P) {
                Q = next(Q);
            }
            deleteDataMatkulAFTER(L, Q, P);
        }
        cout << "Data Mata Kuliah " << namaMatkul << " berhasil dihapus." << endl;
    } else {
        cout << "Data Mata Kuliah " << namaMatkul << " tidak ditemukan." << endl;
    }
}

adrMatkul findDataMatkul(string namaMatkul, listMatkul L) {
    adrMatkul P = first(L);
    while (P != first(L)) {
        if (info(P).namaMatkul == namaMatkul) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}
