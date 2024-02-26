## Penjelasan

Instruksi tersebut memberikan detail tentang quest logic untuk game RPG yang sedang dikembangkan. Di dalam game ini, karakter utamanya adalah Annalyn, seorang gadis pemberani yang memiliki anjing peliharaan yang setia. Saat sedang mencari buah beri di hutan, sahabat baik Annalyn diculik. Annalyn akan mencoba menemukan dan membebaskan sahabatnya tersebut, opsional membawa anjingnya dalam pencarian ini.

Setelah mengikuti jejak sahabatnya, Annalyn menemukan perkemahan di mana sahabatnya ditawan. Ternyata ada dua penculik: seorang ksatria perkasa dan seorang penembak jitu yang licik.

Setelah menemukan penculik, Annalyn mempertimbangkan tindakan apa yang dapat dia lakukan:

Fast Attack: Serangan cepat dapat dilakukan jika ksatria tersebut sedang tidur, karena dia membutuhkan waktu untuk memakai baju besinya, sehingga dia akan rentan.
Spy: Kelompok tersebut bisa diamati jika setidaknya satu di antara mereka terjaga. Jika tidak, memata-matai akan sia-sia.
Signal Prisoner: Tahanan bisa diberi isyarat menggunakan suara burung jika tahanan tersebut terjaga dan penembak jitu tersebut tertidur, karena penembak jitu dilatih dalam isyarat burung sehingga mereka bisa menyadari pesan tersebut.
Free Prisoner: Annalyn dapat mencoba menyelinap ke perkemahan untuk membebaskan tahanan. Ini adalah hal yang berisiko dan hanya bisa berhasil dengan dua cara:
Jika Annalyn membawa anjing peliharaannya, dia bisa membebaskan tahanan jika penembak jitu sedang tertidur. Ksatria tersebut takut pada anjing dan penembak jitu tidak akan punya waktu untuk bersiap sebelum Annalyn dan tahanan bisa melarikan diri.
Jika Annalyn tidak membawa anjingnya, maka Annalyn dan tahanan harus sangat berhati-hati! Annalyn bisa membebaskan tahanan jika tahanan tersebut terjaga dan ksatria serta penembak jitu keduanya tertidur, tetapi jika tahanan tersebut tertidur mereka tidak bisa diselamatkan: tahanan tersebut akan terkejut dengan kehadiran tiba-tiba Annalyn dan membangunkan ksatria dan penembak jitu.
Dari penjelasan di atas, kami memiliki beberapa instruksi:

CanFastAttack: Implementasikan logika untuk menentukan apakah serangan cepat dapat dilakukan berdasarkan apakah ksatria tersebut terjaga atau tidak.
CanSpy: Implementasikan logika untuk menentukan apakah mata-mataan bisa dilakukan berdasarkan apakah ksatria, penembak jitu, atau tahanan tersebut terjaga.
CanSignalPrisoner: Implementasikan logika untuk menentukan apakah isyarat kepada tahanan bisa dilakukan berdasarkan apakah penembak jitu tersebut terjaga dan tahanan tersebut terjaga.
CanFreePrisoner: Implementasikan logika untuk menentukan apakah membebaskan tahanan bisa dilakukan berdasarkan status ksatria, penembak jitu, dan tahanan serta kehadiran anjing peliharaan Annalyn.
