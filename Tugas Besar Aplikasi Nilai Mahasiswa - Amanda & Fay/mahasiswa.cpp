#include "mahasiswa.h"
#include "menu.h"

void createListMhs(listMhs &L) {
    first(L) = NULL;
    last(L) = NULL;
}

infotypeMhs createInfoMhs(int NIM, string nama, string jenisKelamin, string programStudi){
    infotypeMhs X;
    X.NIM = NIM;
    X.nama = nama;
    X.jenisKelamin = jenisKelamin;
    X.programStudi = programStudi;
    return X;
}

adrMhs alokasiMhs(infotypeMhs x) {
    adrMhs P = new elmListMhs;
    if (P != NULL) {
        info(P) = x;
        next(P) = NULL;
    }
    return P;
}

void insertDataMhsFIRST(listMhs &L, adrMhs P) {
    if (first(L) == NULL) {
        first(L) = P;
        last(L) = P;
    } else {
        next(P) = first(L);
        first(L) = P;
    }
}

void insertDataMhsAFTER(listMhs &L, adrMhs Prec, adrMhs P) {
    if (first(L) == NULL) {
        first(L) = P;
        last(L) = P;
    } else {
        if (Prec != NULL) {
            next(P) = next(Prec);
            next(Prec) = P;
            if (Prec == last(L)) {
                last(L) = P;
            }
        } else {
            next(P) = first(L);
            first(L) = P;
        }
    }
}

void insertDataMhsLAST(listMhs &L, adrMhs P){
    if (first(L) == NULL){
        insertDataMhsFIRST(L, P);
    } else {
        insertDataMhsAFTER(L, last(L), P);
    }
}

//PERLU PERBAIKAN
void showAllDataMhs(listMhs L) {
    adrMhs p = first(L);
    if (p == NULL) {
        cout << "Tidak ada data mahasiswa.";
    } else {
        while (p != NULL) {
            cout << "NIM: " << info(p).NIM << endl;
            cout << "Nama: " << info(p).nama << endl;
            cout << "Jenis Kelamin: " << info(p).jenisKelamin << endl;;
            cout << "Program Studi: " << info(p).programStudi << endl;;
            p = next(p);
        }
    }
}

void editDataMhs(int NIM, listMhs &L) {
    adrMhs P = findDataMhs(NIM, L);

    if (P != NULL) {
        cout << "Data Mahasiswa dengan NIM " << NIM << " ditemukan. Silakan masukkan data baru:" << endl;
        infotypeMhs updatedData;
        cout << "NIM: ";
        cin >> updatedData.NIM;
        cout << "Nama: ";
        cin.ignore();
        getline(cin, updatedData.nama);
        cout << "Jenis Kelamin: ";
        cin >> updatedData.jenisKelamin;
        cout << "Program Studi: ";
        cin.ignore();
        getline(cin, updatedData.programStudi);
        info(P) = updatedData;
        cout << "Data Mahasiswa dengan NIM " << NIM << " berhasil diupdate." << endl;
    } else {
        cout << "Data Mahasiswa dengan NIM " << NIM << " tidak ditemukan." << endl;
    }
}


void deleteDataMhsFIRST(listMhs &L, adrMhs &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mahasiswa yang dapat dihapus." << endl;
    } else {
        P = first(L);
        if (next(P) == NULL) {
            first(L) = NULL;
            last(L) = NULL;
        } else {
            first(L) = next(P);
            next(P) = NULL;
        }
    }
}

void deleteDataMhsAFTER(listMhs &L, adrMhs Prec, adrMhs &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mahasiswa yang dapat dihapus." << endl;
    } else {
        if (Prec != NULL) {
            P = next(Prec);
            if (P != NULL) {
                next(Prec) = next(P);
                next(P) = NULL;
                if (P == last(L)) {
                    last(L) = Prec;
                }
            } else {
                cout << "Tidak dapat menghapus setelah node terakhir." << endl;
            }
        } else {
            cout << "Tidak dapat menghapus setelah node NULL." << endl;
        }
    }
}

void deleteDataMhsLAST(listMhs &L, adrMhs &P) {
    if (first(L) == NULL) {
        cout << "Tidak ada data mahasiswa yang dapat dihapus." << endl;
    } else {
        P = last(L);
        if (first(L) == P) {
            first(L) = NULL;
            last(L) = NULL;
        } else {
            adrMhs Q = first(L);
            while (next(Q) != P) {
                Q = next(Q);
            }
            last(L) = Q;
            next(Q) = NULL;
        }
    }
}

void removeDataMhs(int NIM, listMhs &L){
    adrMhs P = findDataMhs(NIM, L);
    if (P != NULL) {
        if (P == first(L)) {
            deleteDataMhsFIRST(L, P);
        } else if (P == last(L)) {
            deleteDataMhsLAST(L, P);
         } else {
            adrMhs Q = first(L);
            while (next(Q) != P) {
                Q = next(Q);
            }
            deleteDataMhsAFTER(L, Q, P);
        }
        cout << "Data mahasiswa dengan NIM " << NIM << " berhasil dihapus." << endl;
    } else {
        cout << "Data mahasiswa dengan NIM " << NIM << " tidak ditemukan." << endl;
    }
}


adrMhs findDataMhs(int NIM, listMhs L) {
    adrMhs P = first(L);
    while (P != NULL) {
        if (info(P).NIM == NIM) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}
