#include "relasi.h"

void createListRelasi(listRelasi &L){
    first(L) = NULL;
}

adrRelasi alokasiRelasi(adrMhs P, adrMatkul Q, adrNilai N) {
    adrRelasi R = new elmListRelasi; //3:48
    Mhs(R) = P;
    Matkul(R) = Q;
    Nilai(R) = N;
    next(R) = NULL;
    prev(R)= NULL;
    return R;
}

void insertRelasi(listRelasi &L, adrRelasi P) {
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

void deleteRelasi(listRelasi &L, adrRelasi &P) {
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

void connect(listRelasi &L, listMhs x, listMatkul y, listNilai n, adrMhs P, adrMatkul Q, adrNilai N) {
    int NIM;
    string namaMatkul;
    float UAS;
    P = findDataMhs(NIM, x);
    Q = findDataMatkul(namaMatkul, y);
    N = findDataNilai(UAS, n);
    if (P!=NULL && Q != NULL && N != NULL){
        adrRelasi R = alokasiRelasi(P, Q, N);
        insertRelasi(L, R);
    }
}

void disconnect(listRelasi &L, adrMhs P, adrMatkul Q, adrNilai N) {
    adrRelasi R = findRelasi(L, P, Q, N);
    if (P != NULL) {
        deleteRelasi(L, R);
        delete R;
        cout << "Hubungan Mahasiswa dengan Mata Kuliah berhasil dihapus." << endl;
    } else {
        cout << "Hubungan Mahasiswa dengan Mata Kuliah tidak ditemukan." << endl;
    }
}

adrRelasi findRelasi(listRelasi L, adrMhs mahasiswa, adrMatkul mataKuliah, adrNilai nilai) {
    adrRelasi P = first(L);
    while (P != NULL) {
        if (info(P).mahasiswa == mahasiswa && info(P).mataKuliah == mataKuliah) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}

void showAllDataMahasiswaMatkul(listRelasi L) {
    adrRelasi P = first(L);
    if (P == NULL) {
        cout << "Tidak ada hubungan Mahasiswa dengan Mata Kuliah." << endl;
    } else {
        cout << "Data Mahasiswa dan Mata Kuliah:" << endl;
        while (P != NULL) {
            cout << "NIM: " << info(info(P).mahasiswa).NIM << " Nama: " << info(info(P).mahasiswa).nama << endl;
            cout << "Kode Matkul: " << info(info(P).mataKuliah).kodeMatkul << " Nama Matkul: " << info(info(P).mataKuliah).namaMatkul << endl << endl;
            P = next(P);
        }
    }
}

void showMatkulMahasiswa(listRelasi L, adrMhs mahasiswa) {
    adrRelasi P = first(L);
    if (P == NULL) {
        cout << "Mahasiswa belum mengambil Mata Kuliah apapun." << endl;
    } else {
        cout << "Mata Kuliah yang diambil oleh Mahasiswa:" << endl;
        while (P != NULL) {
            if (info(info(P).mahasiswa).NIM == info(mahasiswa).NIM) {
                cout << "Kode Matkul: " << info(info(P).mataKuliah).kodeMatkul << " Nama Matkul: " << info(info(P).mataKuliah).namaMatkul << endl;
            }
            P = next(P);
        }
    }
}

int countMatkulMahasiswa(listRelasi L, adrMhs mahasiswa) {
    int count = 0;
    adrRelasi P = first(L);
    while (P != NULL) {
        if (info(info(P).mahasiswa).NIM == info(mahasiswa).NIM) {
            count++;
        }
        P = next(P);
    }
    return count;
}
