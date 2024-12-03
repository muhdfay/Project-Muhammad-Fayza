#include "menu.h"

void header(){
    cout << "================ TUGAS BESAR STRUKTUR DATA ================" << endl;
    cout << "Kelompok \t: 9"<<endl;
    cout << "Tema \t\t: Multi Linked-List"<<endl;
    cout << "Judul \t\t: Aplikasi Nilai Mahasiswa"<<endl;
    cout << "Anggota \t: 1. Muhammad Fayza        - 1305223025"<<endl;
    cout << "\t\t  2. Amanda Kayla Putri W. - 1305223105"<<endl<<endl;
    cout << "\n" << endl;

    int i, pilihan;
    string input;
    listMhs x;
    listMatkul y;

    cout << "============================================================" << endl;
    cout << "===               APLIKASI NILAI MAHASISWA               ===" << endl;
    cout << "============================================================" << endl;
    cout << "|                      Selamat datang!                     |" << endl;
    cout << "============================================================" << endl;
    cout << ("\nPilih Data yang ingin anda akses      \n") << endl;
    cout << ("1.  Data Mahasiswa             ") << endl;
    cout << ("2.  Data Mata Kuliah           ") << endl;
    cout << ("3.  Data Nilai                 ") << endl;
    cout << ("0.  Keluar                     ") << endl;
    cout << ("\nMasukkan pilihan : ");
    cin >> pilihan;
    while (pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 0) {
		cout << ("\n---------------------------------------") << endl;
		cout << ("---             ERROR             ---") << endl;
		cout << ("--- Masukkan pilihan yang benar!! ---") << endl;
		cout << ("---------------------------------------") << endl;
		cout << ("\nMasukkan pilihan : ") << endl;
		cin >> pilihan;
		i++;
	}
    if (pilihan == 1) {
        menuMhs(x);
    } else if (pilihan == 2) {
        menuMatkul(y);
    } else {
        cout << "hehe";
    }
}

void menuMhs(listMhs x) {
    int i, pilihan;
    string input;
    adrMhs P;
    infotypeMhs z;

    cout << ("+----------------------------------------------+") << endl;
    cout << ("|        1. Menambahkan data Mahasiswa         |") << endl;
    cout << ("|        2. Mengubah data Mahasiswa            |") << endl;
    cout << ("|        3. Menghapus data Mahasiswa           |") << endl;
    cout << ("|        4. Menampilkan data Mahasiswa         |") << endl;
    cout << ("|        5. Mencari data Mahasiswa             |") << endl;
    cout << ("|        6. Kembali                            |") << endl;
    cout << ("+----------------------------------------------+") << endl;
    cout << ("\nMasukkan pilihan anda: ") << endl;
    cin >> pilihan;
    while (pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 && pilihan != 5 && pilihan != 6) {
        cout << ("\n+-------------------------------------+") << endl;
        cout << ("|---------------ERROR-----------------|") << endl;
        cout << ("|----Masukkan pilihan yang benar !----|") << endl;
        cout << ("+-------------------------------------+") << endl;
        cout << ("\nMasukkan pilihan anda: ") << endl;
        cin >> pilihan;
        i++;
    }
    if (pilihan == 1) {
        int i = 0;
        while (i < 10 && input != "tidak") {
            cout << "---------------------------------------" << endl;
            adrMhs P = alokasiMhs(z);
            if (P != NULL) {

                cout << "NIM: ";
                cin >> info(P).NIM;

                cout << "Nama: ";
                cin.ignore();
                getline(cin, info(P).nama);

                cout << "Jenis Kelamin: ";
                cin >> info(P).jenisKelamin;

                cout << "Program Studi: ";
                cin.ignore();
                getline(cin, info(P).programStudi);

                insertDataMhsFIRST(x, P);

                cout << "\n---------------------------------------" << endl;
                cout << "Lanjut menambahkan Data? (ya/tidak): ";
                cin >> input;

                while (input != "ya" && input != "tidak") {
                    cout << "Lanjut menambahkan Data? (ya/tidak): ";
                    cin >> input;
                    i++;
                }
            } else {
                cout << "Failed to allocate memory for Mahasiswa." << endl;
                break;
            }
        }
        menuMhs(x);
    } else if (pilihan == 2) {
        int NIM;
        cout << "Masukkan NIM untuk Mengubah Data: ";
        cin >> NIM;
        adrMhs found = findDataMhs(NIM, x);
        if (found != NULL) {
            editDataMhs(NIM, x);
            cout << "Data Mahasiswa berhasil diubah!" << endl;
        } else {
            cout << "Mahasiswa dengan NIM tersebut tidak ditemukan." << endl;
        }
        menuMhs(x);
    } else if (pilihan == 3) {
        int NIM;
        cout << "Masukkan NIM untuk Menghapus Data :";
        cin >> NIM;
        removeDataMhs(NIM, x);
        cout << "Data sudah di Hapus.";
        menuMhs(x);
    } else if (pilihan == 4) { //PERLU PERBAIKAN
        showAllDataMhs(x);
        menuMhs(x);
    } else if (pilihan == 5) {
        int NIM;
        cout << "Masukkan NIM untuk mencari Data :";
        cin >> NIM;
        adrMhs found = findDataMhs(NIM, x);
        if (found != NULL){
            cout << "Data Ditemukan! " << endl;
            cout << "NIM            :" << info(found).NIM << endl;
            cout << "Nama           :" << info(found).nama << endl;
            cout << "Jenis Kelamin  :" << info(found).jenisKelamin << endl;
            cout << "Jurusan        :" << info(found).programStudi << endl;
        }
        menuMhs(x);
    } else if (pilihan == 6) {
        header();
    }
}


void menuMatkul(listMatkul y) {
    int i, pilihan;
    string input;
    adrMatkul Q;
    infotypeMatkul w;

    cout << ("+-------------------------------------------------+") << endl;
    cout << ("|        1. Menambahkan data Mata Kuliah         |") << endl;
    cout << ("|        2. Mengubah data Mata Kuliah            |") << endl;
    cout << ("|        3. Menghapus data Mata Kuliah           |") << endl;
    cout << ("|        4. Menampilkan data Mata Kuliah         |") << endl;
    cout << ("|        5. Mencari data Mata Kuliah             |") << endl;
    cout << ("|        6. Kembali                              |") << endl;
    cout << ("+------------------------------------------------+") << endl;
    cout << ("\nMasukkan pilihan anda: ") << endl;
    cin >> pilihan;
    while (pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 && pilihan != 5 && pilihan != 6) {
        cout << ("\n+-------------------------------------+") << endl;
        cout << ("|---------------ERROR-----------------|") << endl;
        cout << ("|----Masukkan pilihan yang benar !----|") << endl;
        cout << ("+-------------------------------------+") << endl;
        cout << ("\nMasukkan pilihan anda: ") << endl;
        cin >> pilihan;
        i++;
    }
    if (pilihan == 1) {
    string input;
    do {
        cout << "---------------------------------------" << endl;
        adrMatkul Q = alokasiMatkul(w);
        if (Q != NULL) {
            cout << "Kode Mata Kuliah: ";
            cin >> info(Q).kodeMatkul;

            cout << "SKS: ";
            cin >> info(Q).sks;

            cout << "Nama Mata Kuliah: ";
            cin.ignore();
            getline(cin, info(Q).namaMatkul);

            insertDataMatkulFIRST(y, Q);

            cout << "\n---------------------------------------" << endl;
            cout << "Lanjut menambahkan Data? (ya/tidak): ";
            cin >> input;
        } else {
            cout << "Failed to allocate memory for Mata Kuliah." << endl;
            break;
            }
        } while (input == "ya");
        menuMatkul(y);
    } else if (pilihan == 2) {
        string namaMatkul;
        cout << "Masukkan Nama Mata Kuliah untuk Mengubah Data: ";
        cin >> namaMatkul;
        adrMatkul found = findDataMatkul(namaMatkul, y);
        if (found != NULL) {
            editDataMatkul(namaMatkul, y);
            cout << "Data Mahasiswa berhasil diubah!" << endl;
        } else {
            cout << "Mahasiswa dengan NIM tersebut tidak ditemukan." << endl;
        }
        menuMatkul(y);
    } else if (pilihan == 3) {
        string namaMatkul;
        cout << "Masukkan Nama Mata Kuliah untuk Menghapus Data :";
        cin >> namaMatkul;
        removeDataMatkul(namaMatkul, y);
        cout << "Data sudah di Hapus.";
        menuMatkul(y);
    } else if (pilihan == 4) {
        //PERBAIKAN!!!!!!!!!
        menuMatkul(y);
    } else if (pilihan == 5) {
        string namaMatkul;
        cout << "Masukkan Nama Mata Kuliah untuk mencari Data :";
        cin >> namaMatkul;
        adrMatkul found = findDataMatkul(namaMatkul, y);
        if (found != NULL){
            cout << "Data Ditemukan!    " << endl;
            cout << "Kode Mata Kuliah : " << info(found).kodeMatkul << endl;
            cout << "SKS              : " << info(found).sks << endl;
            cout << "Nama Mata Kuliah : " << info(found).namaMatkul << endl;
        }
        menuMatkul(y);
    } else if (pilihan == 6) {
        header();
    }
}
