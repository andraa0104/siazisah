package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	FullName  string    `json:"full_name"`
	Role      string    `json:"role"`
	MasjidID  *int      `json:"masjid_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Masjid struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Telepon   string    `json:"telepon"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PengurusMasjid struct {
	ID        int       `json:"id"`
	MasjidID  int       `json:"masjid_id"`
	Nama      string    `json:"nama"`
	Jabatan   string    `json:"jabatan"`
	Telepon   string    `json:"telepon"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PengurusZakat struct {
	ID        int       `json:"id"`
	MasjidID  int       `json:"masjid_id"`
	Nama      string    `json:"nama"`
	Jabatan   string    `json:"jabatan"`
	Telepon   string    `json:"telepon"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type KadarZakat struct {
	ID         int       `json:"id"`
	MasjidID   int       `json:"masjid_id"`
	Tahun      int       `json:"tahun"`
	Kelas      string    `json:"kelas"`
	Nominal    float64   `json:"nominal"`
	Keterangan string    `json:"keterangan"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PengaturanZakat struct {
	MasjidID           int                `json:"masjid_id"`
	Kelas1             float64            `json:"kelas1"`
	Kelas2             float64            `json:"kelas2"`
	Kelas3             float64            `json:"kelas3"`
	FitrahBerasPerJiwa float64            `json:"fitrahBerasPerJiwa"`
	FidyahPerHari      float64            `json:"fidyahPerHari"`
	FidyahBerasPerHari float64            `json:"fidyahBerasPerHari"`
	MalRates           map[string]float64 `json:"malRates"`
}

type JenisZakatTersedia struct {
	ID         int       `json:"id"`
	MasjidID   int       `json:"masjid_id"`
	JenisZakat string    `json:"jenis_zakat"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Mustahiq struct {
	ID            int       `json:"id"`
	MasjidID      int       `json:"masjid_id"`
	Nama          string    `json:"nama"`
	JenisPenerima string    `json:"jenis_penerima"`
	Alamat        string    `json:"alamat"`
	Lokasi        string    `json:"lokasi"`
	RT            string    `json:"rt"`
	Telepon       string    `json:"telepon"`
	Keterangan    string    `json:"keterangan"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Muzakki struct {
	ID        int       `json:"id"`
	MasjidID  int       `json:"masjid_id"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Telepon   string    `json:"telepon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransaksiZakat struct {
	ID                  int       `json:"id"`
	MasjidID            int       `json:"masjid_id"`
	MuzakkiID           int       `json:"muzakki_id"`
	JenisZakat          string    `json:"jenis_zakat"`
	BentukZakat         *string   `json:"bentuk_zakat"`
	JenisHarta          *string   `json:"jenis_harta"`
	NominalHarta        *float64  `json:"nominal_harta"`
	PersentaseZakat     *float64  `json:"persentase_zakat"`
	KelasZakat          *string   `json:"kelas_zakat"`
	JumlahOrang         int       `json:"jumlah_orang"`
	JumlahHariFidyah    *int      `json:"jumlah_hari_fidyah"`
	StandarBerasPerJiwa *float64  `json:"standar_beras_per_jiwa"`
	KgBerasDibayar      *float64  `json:"kg_beras_dibayar"`
	HargaBerasPerKg     *float64  `json:"harga_beras_per_kg"`
	NominalPerOrang     float64   `json:"nominal_per_orang"`
	TotalWajib          float64   `json:"total_wajib"`
	TotalDibayar        float64   `json:"total_dibayar"`
	InfaqTambahan       float64   `json:"infaq_tambahan"`
	Keterangan          string    `json:"keterangan"`
	Tahun               int       `json:"tahun"`
	TanggalBayar        string    `json:"tanggal_bayar"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	MuzakkiNama         string    `json:"muzakki_nama,omitempty"`
	MuzakkiAlamat       string    `json:"muzakki_alamat,omitempty"`
	MuzakkiTelepon      string    `json:"muzakki_telepon,omitempty"`
}

type DistribusiZakat struct {
	ID                int       `json:"id"`
	MasjidID          int       `json:"masjid_id"`
	MustahiqID        int       `json:"mustahiq_id"`
	MustahiqNama      string    `json:"mustahiq_nama,omitempty"`
	JenisPenerima     string    `json:"jenis_penerima,omitempty"`
	Nominal           float64   `json:"nominal"`
	BerasKg           float64   `json:"beras_kg"`
	Mode              string    `json:"mode,omitempty"`
	TanggalDistribusi string    `json:"tanggal_distribusi"`
	WaktuInput        string    `json:"waktu_input,omitempty"`
	Keterangan        string    `json:"keterangan"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type DistribusiInsight struct {
	TotalMustahiqAktif       int     `json:"total_mustahiq_aktif"`
	TotalFitrahBerasKg       float64 `json:"total_fitrah_beras_kg"`
	TotalFidyahBerasKg       float64 `json:"total_fidyah_beras_kg"`
	TotalFitrahUang          float64 `json:"total_fitrah_uang"`
	TotalFidyahUang          float64 `json:"total_fidyah_uang"`
	TotalZakatMalUang        float64 `json:"total_zakat_mal_uang"`
	TotalPoolBerasKg         float64 `json:"total_pool_beras_kg"`
	TotalPoolUang            float64 `json:"total_pool_uang"`
	TotalDistribusiBerasKg   float64 `json:"total_distribusi_beras_kg"`
	TotalDistribusiUang      float64 `json:"total_distribusi_uang"`
	SisaBerasKg              float64 `json:"sisa_beras_kg"`
	SisaUang                 float64 `json:"sisa_uang"`
	RekomendasiBerasPerOrang float64 `json:"rekomendasi_beras_per_orang"`
	RekomendasiUangPerOrang  float64 `json:"rekomendasi_uang_per_orang"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DashboardStats struct {
	TotalMasjid             int                      `json:"total_masjid"`
	TotalMuzakki            int                      `json:"total_muzakki"`
	TotalMustahiq           int                      `json:"total_mustahiq"`
	TotalOrangDizakati      int                      `json:"total_orang_dizakati"`
	TotalZakatFitrah        float64                  `json:"total_zakat_fitrah"`
	TotalZakatFitrahUang    float64                  `json:"total_zakat_fitrah_uang"`
	TotalZakatFitrahBerasKg float64                  `json:"total_zakat_fitrah_beras_kg"`
	TotalZakatFitrahBerasRp float64                  `json:"total_zakat_fitrah_beras_rupiah"`
	TotalZakatMal           float64                  `json:"total_zakat_mal"`
	TotalFidyah             float64                  `json:"total_fidyah"`
	TotalFidyahUang         float64                  `json:"total_fidyah_uang"`
	TotalFidyahBerasKg      float64                  `json:"total_fidyah_beras_kg"`
	TotalFidyahBerasRp      float64                  `json:"total_fidyah_beras_rupiah"`
	TotalInfaq              float64                  `json:"total_infaq"`
	LastUpdate              string                   `json:"last_update"`
	MustahiqPerJenis        []map[string]interface{} `json:"mustahiq_per_jenis"`
}
