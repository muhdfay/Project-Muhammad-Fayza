#include "menu.h"
#include "mahasiswa.h"

// MARK : - Menu
string menu() {
    //clearScreen();
    string inputUser;
    cout << "===== SELAMAT DATANG DI APLIKASI NILAI MAHASISWA =====" << endl;
    cout << "1. Data Mahasiswa" << endl;
    cout << "2. Data Mata Kuliah" << endl;
    cout << "3. Data Nilai" << endl;
    cout << "0. Keluar" << endl;
    cout << "\n Masukkan pilihan : " << endl;
    cout << "Masukkan pilihan: ";
    cin >> inputUser;
    cout << endl;

    return inputUser;
}

string menuMahasiswa(listMhs &M){
    //clearScreen();
    int i, pilihan;
    string input;
    listMhs x;
    listMatkul y;
    adrMhs P;
    adrMatkul Q;
    infotypeMhs z;
    infotypeMatkul w;
    cout << "=============== DATA MAHASISWA ==============="<<endl;
    //showAllDataMhs(M);

    string inputUser;
    cout << "===== SILAHKAN PILIH OPSI =====" << endl;
    cout << "1. Menambahkan data Mahasiswa" << endl;
    cout << "2. Mengubah data Mahasiswa" << endl;
    cout << "3. Menghapus data Mahasiswa" << endl;
    cout << "4. Menampilkan data Mahasiswa" << endl;
    cout << "5. Mencari data Mahasiswa" << endl;
    cout << "0. Back" << endl;
    cout << "Masukkan pilihan: ";
    cin >> inputUser;

    if (inputUser == "0") {

    }else if (inputUser == "1"){
        int i = 0;
        string input = "ya";
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
    }else if (inputUser == "2"){
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
    }else if (inputUser == "3"){
        int NIM;
        cout << "Masukkan NIM untuk Menghapus Data :";
        cin >> NIM;
        removeDataMhs(NIM, x);
        cout << "Data sudah di Hapus.";
    }else if (inputUser == "4"){
        cout << "Data Mahasiswa :";
        showAllDataMhs(x);
    } else if (pilihan == 5) {
        int NIM;
        cin >> NIM;
        adrMhs found = findDataMhs(NIM, x);
        if (found != NULL){
            cout << "Data Ditemukan!" << endl;
            cout << "NIM :" << info(found).NIM << endl;
            cout << "Nama :" << info(found).nama << endl;
            cout << "Jenis Kelamin :" << info(found).jenisKelamin << endl;
            cout << "Jurusan :" << info(found).programStudi << endl;
        }
    }
}
