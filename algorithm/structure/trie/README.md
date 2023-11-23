## Trie

Trie, singkatan dari "reTRIEval," adalah struktur data pohon yang digunakan untuk penyimpanan dan pencarian data, terutama dalam konteks mencari kata-kata atau string. Ini memiliki node-root dan setiap node memiliki beberapa anak node yang mewakili karakter dalam string. Ini adalah struktur data efisien untuk memeriksa keberadaan kata-kata atau string tertentu dalam kumpulan besar kata-kata.

### Struktur Node Trie

Setiap node Trie memiliki:

- **children** (anak-anak): Menyimpan referensi ke node-node anak berdasarkan karakter yang mewakili hubungan antara node-node tersebut.
- **isLeaf**: Menandakan apakah node ini adalah akhir dari suatu kata atau string.

### Operasi Utama

Operasi yang umum dilakukan pada Trie meliputi:

- **Insert**: Menambahkan kata atau string ke Trie, menciptakan node baru jika karakter tidak ada dalam struktur saat ini.
- **Find**: Memeriksa apakah suatu kata atau string ada dalam Trie.
- **Remove**: Menghapus kata atau string dari Trie.
- **Compact**: Memadatkan Trie, menghapus node-node yang tidak lagi digunakan setelah penghapusan string-string tertentu.

### Penggunaan

Implementasi Trie dapat digunakan untuk keperluan seperti:

- Otentikasi kata sandi.
- Saran kata.
- Pengecekan ejaan.
- Aplikasi berbasis kata atau string.

Dengan struktur data ini, pencarian string berbasis awalan juga menjadi efisien karena Trie menyimpan informasi secara hierarkis berdasarkan karakter awal yang identik.
