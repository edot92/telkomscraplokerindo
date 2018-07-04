"""
Suppose you have some texts of news and know their categories.
You want to train a system with this pre-categorized/pre-classified 
texts. So, you have better call this data your training set.
"""
from naiveBayesClassifier import tokenizer
from naiveBayesClassifier.trainer import Trainer
from naiveBayesClassifier.classifier import Classifier

newsTrainer = Trainer(tokenizer)

# You need to train the system passing each text one by one to the trainer module.
newsSet = [
    {'text': 'CV. Asia Afrika Dental yang bergerak di bidang pembuatan gigi palsu dan klinik dokter gigi sedang membutuhkan perawat yang berkompeten di bidangnyaTanggung Jawab Pekerjaan :– Menyiapkan alat alat– Menyiapkan bahan bahan– Assistant Dokter gigi– RontgenSyarat Pengalaman :Pengalaman minimal 1 tahun sebagai perawatKeahlian :–Kualifikasi :Pria/Wanita maksimal 25 tahun, Jujur, Cekatan, Bertanggung Jawab, dapat bekerja dibawah deadlineDiutamakan domisil di BandungTunjangan :BPJS, Tunjangan harian, Tunjangan Hari RayaInsentif :Uang Lembur, BonusGaji : 1.500.000 – 3.000.000Waktu Bekerja :Jam 8 s/d 17, Senin sampai Sabtu', 'category': 'dokter'},
    {'text': 'Sebagai GURU TIK (komputer). – menguasai aplikasi office dan coreldraw atau photoshop Penempatan SENTRA PRIMER CAKUNG, PSR MINGGU DAN CIANJUR WEB Admin : Menguasai script php dan desain web penempatan : Cibubur Jakarta Timur Teknisi Komputer: – menguasai troubleshooting Penempatan Bintaro Tanggung Jawab Pekerjaan : Sebagai guru komputer dan web admin serta teknisi komputer Syarat Pengalaman : Pengalaman 1-2 tahun Keahlian : Guru TIK: Menguasai aplikasi office dan coreldraw, photoshop Teknisi komputer: Mnguasai troubleshooting Web Admin: Mnguasai script php dn desain web Kualifikasi : Pria/Wanita Tunjangan : uang harian Insentif : uang lembur', 'category': 'guru'},
]
for news in newsSet:
    newsTrainer.train(news['text'], news['category'])

# When you have sufficient trained data, you are almost done and can start to use
# a classifier.
newsClassifier = Classifier(newsTrainer.data, tokenizer)

# Now you have a classifier which can give a try to classifiy text of news whose
# category is unknown, yet.
classification = newsClassifier.classify("Yayasan Pendidikan Harapan Nusantara Denpasar membutuhkan guru Agama Islam dengan pendidikan minimal S1 berpenampilan sopan, lamaran segera di bawa ke Yayasan Pendidikan Harapan Nusantara dengan alamat Jl. Cargo Sari III No. 3 DenpasarTanggung Jawab Pekerjaan :Mengajar Agama Islam di Yayasan Pendidikan Harapan NusantaraSyarat Pengalaman :Pengalaman minimal 1 tahun mengajarKeahlian :Bisa mengajar, dan menguasai Microsoft officeKualifikasi :Pria /Wanita, maksimal 35 tahun, jujur dan telitiTunjangan :Uang transport dan honor mengajarWaktu Bekerja :Jam 8 s.d.13 Senin s.d. Jumat")

# the classification variable holds the detected categories sorted
print(classification)
