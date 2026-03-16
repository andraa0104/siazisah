<template>
  <!-- CSS Links -->
  <link rel="stylesheet" href="/homepage/assets/vendor/aos/aos.css" />
  <link
    rel="stylesheet"
    href="/homepage/assets/vendor/bootstrap/css/bootstrap.min.css"
  />
  <link
    rel="stylesheet"
    href="/homepage/assets/vendor/bootstrap-icons/bootstrap-icons.css"
  />
  <link
    rel="stylesheet"
    href="/homepage/assets/vendor/boxicons/css/boxicons.min.css"
  />
  <link
    rel="stylesheet"
    href="/homepage/assets/vendor/glightbox/css/glightbox.min.css"
  />
  <link
    rel="stylesheet"
    href="/homepage/assets/vendor/swiper/swiper-bundle.min.css"
  />
  <link rel="stylesheet" href="/homepage/assets/css/style.css" />

  <!-- Header -->
  <header id="header" class="d-flex align-items-center">
    <div class="container d-flex align-items-center justify-content-between">
      <h1 class="logo">
        <router-link to="/">SI<span>AZISAH</span></router-link>
      </h1>

      <nav id="navbar" class="navbar">
        <ul>
          <li><a class="nav-link scrollto active" href="#hero">Home</a></li>
          <li><a class="nav-link scrollto" href="#about">Hadist</a></li>
          <li><router-link to="/login" class="nav-link">Login</router-link></li>
        </ul>
        <i class="bi bi-list mobile-nav-toggle"></i>
      </nav>
    </div>
  </header>

  <!-- Hero Section -->
  <section id="hero" class="d-flex align-items-center">
    <div class="container" data-aos="zoom-out" data-aos-delay="100">
      <h1>
        Selamat Datang di
        <span
          >Sistem Informasi Zakat, Infaq, Sadaqah, dan Fidyah (SIaZISah)</span
        >
      </h1>
      <h2>
        Memudahkan memantau dan merekap data zakat fitrah, infaq, shodaqoh, dan
        fidyah untuk transparansi data
      </h2>
      <div class="d-flex">
        <a href="#about" class="btn-get-started scrollto">Hadist Zakat</a>
      </div>
    </div>
  </section>

  <main id="main">
    <!-- Zakat Types Section -->
    <section id="zakat-types" class="featured-services">
      <div class="container" data-aos="fade-up">
        <div class="section-title">
          <h2>Jenis-Jenis Zakat</h2>
          <h3>
            Kadar dan Kategori <span>Zakat Fitrah, Zakat Mal, dan Fidyah</span>
          </h3>
        </div>

        <!-- Zakat Types Slider -->
        <div class="zakat-slider-wrapper">
          <div class="zakat-slider swiper-container">
            <div class="swiper-wrapper">
              <!-- Zakat Fitrah -->
              <div class="swiper-slide">
                <div class="zakat-card">
                  <div class="zakat-header zakat-fitrah">
                    <i class="bx bxs-leaf"></i>
                    <h3>Zakat Fitrah</h3>
                  </div>
                  <div class="zakat-content">
                    <p class="description">
                      Diwajibkan setiap Muslim pada akhir bulan Ramadhan sebelum
                      shalat Idul Fitri
                    </p>

                    <div class="kategori-list">
                      <div class="kategori-item">
                        <h5>Kategori Umum</h5>
                        <ul>
                          <li>
                            <strong>Beras/Gandum:</strong> 2,7 kg per jiwa
                          </li>
                          <li>
                            <strong>Uang Tunai:</strong> Rp 38.000 - Rp 68.000
                          </li>
                        </ul>
                      </div>

                      <div class="kategori-item">
                        <h5>Penerima</h5>
                        <ul>
                          <li>Fakir</li>
                          <li>Miskin</li>
                          <li>Amil</li>
                          <li>Muallaf</li>
                          <li>Riqab</li>
                          <li>Gharim</li>
                          <li>Fisabilillah</li>
                          <li>Ibnu Sabil</li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Zakat Mal -->
              <div class="swiper-slide">
                <div class="zakat-card">
                  <div class="zakat-header zakat-mal">
                    <i class="bx bxs-coin"></i>
                    <h3>Zakat Mal</h3>
                  </div>
                  <div class="zakat-content">
                    <p class="description">
                      Zakat atas harta yang halal dan berkembang, dengan
                      ketentuan nisab serta waktu penunaian yang berbeda sesuai
                      jenis hartanya
                    </p>

                    <div class="kategori-list">
                      <div class="kategori-item">
                        <h5>Jenis Harta & Kadar</h5>
                        <ul>
                          <li>
                            <strong>Emas dan perak:</strong> zakatnya 2,5%
                            setelah mencapai nisab dan berlalu satu haul
                          </li>
                          <li>
                            <strong>Harta usaha/perdagangan:</strong> meliputi
                            modal usaha, barang dagangan, stok barang, dan
                            keuntungan usaha; dikeluarkan 2,5% saat nilainya
                            mencapai nisab setara emas dan telah berlalu haul
                          </li>
                          <li>
                            <strong>Hasil pertanian:</strong> 10% jika diairi
                            oleh hujan atau aliran air alami, dan 5% jika
                            diairi dengan biaya
                          </li>
                          <li>
                            <strong>Ternak:</strong> kambing, sapi, dan unta
                            memiliki kadar zakat bertingkat sesuai jumlah
                          </li>
                          <li>
                            <strong>Rikaz:</strong> harta temuan terpendam
                            dikeluarkan zakatnya sebesar 20% saat ditemukan
                          </li>
                        </ul>
                      </div>

                      <div class="kategori-item">
                        <h5>Waktu Zakat</h5>
                        <ul>
                          <li>
                            Emas, perak, dan harta usaha ditunaikan setelah
                            mencapai nisab dan berlalu satu tahun hijriah
                          </li>
                          <li>
                            Hasil pertanian ditunaikan setiap kali panen
                          </li>
                          <li>
                            Rikaz ditunaikan saat harta ditemukan
                          </li>
                          <li>
                            Ternak mengikuti nisab dan hitungan jumlah yang
                            telah ditetapkan dalam fikih zakat
                          </li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Fidyah -->
              <div class="swiper-slide">
                <div class="zakat-card">
                  <div class="zakat-header zakat-fidyah">
                    <i class="bx bxs-heart-circle"></i>
                    <h3>Fidyah</h3>
                  </div>
                  <div class="zakat-content">
                    <p class="description">
                      Tebusan untuk yang tidak mampu menjalankan puasa atau
                      sudah meninggal masih berhutang puasa
                    </p>

                    <div class="kategori-list">
                      <div class="kategori-item">
                        <h5>Kadar Fidyah</h5>
                        <ul>
                          <li>
                            <strong>Per hari:</strong> 0,7 kg beras/makanan
                          </li>
                          <li>
                            <strong>Atau:</strong> Rp 30.000 per
                            hari
                          </li>
                          <li>
                            <strong>Atau:</strong> Memberi makan 1-2 orang untuk
                            1 hari
                          </li>
                          <li>
                            <strong>Total:</strong> Sesuai hari yang tidak
                            dipuasa
                          </li>
                        </ul>
                      </div>

                      <div class="kategori-item">
                        <h5>Penerima Fidyah</h5>
                        <ul>
                          <li>Fakir</li>
                          <li>Miskin</li>
                          <li>Amil</li>
                          <li>Muallaf</li>
                          <li>Riqab</li>
                          <li>Gharim</li>
                          <li>Fisabilillah</li>
                          <li>Ibnu Sabil</li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Pagination -->
            <div class="swiper-pagination"></div>
          </div>
        </div>

        <!-- Info Box -->
        <div class="info-box" data-aos="fade-up" data-aos-delay="200">
          <div class="row">
            <div class="col-md-4">
              <div class="info-item">
                <i class="bx bxs-check-circle"></i>
                <h5>Semua Zakat Dapat Dicatat</h5>
                <p>
                  Sistem ini membantu mencatat dan memantau semua jenis zakat
                  dengan transparan
                </p>
              </div>
            </div>
            <div class="col-md-4">
              <div class="info-item">
                <i class="bx bxs-user-check"></i>
                <h5>Terdistribusi Tepat Sasaran</h5>
                <p>
                  Memastikan harta zakat sampai ke orang-orang yang benar-benar
                  berhak menerimanya
                </p>
              </div>
            </div>
            <div class="col-md-4">
              <div class="info-item">
                <i class="bx bxs-bar-chart-alt-2"></i>
                <h5>Laporan Terperinci</h5>
                <p>
                  Laporan lengkap untuk wajib zakat dan pengelola terhadap
                  penggunaan dana zakat
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section id="about" class="about section-bg">
      <div class="container" data-aos="fade-up">
        <div class="section-title">
          <h2>Dasar Kewajiban Zakat</h2>
          <h3>Dalil Al-Qur'an dan Hadist tentang <span>Zakat</span></h3>
        </div>

        <div
          class="dasar-slider swiper-container"
          data-aos="fade-up"
          data-aos-delay="100"
        >
          <div class="swiper-wrapper">
            <div class="swiper-slide">
              <div class="dasar-card">
                <div class="dasar-title">Surat Al-Baqarah Ayat 43</div>
                <div class="dasar-arabic">
                  وَأَقِيمُوا الصَّلَاةَ وَآتُوا الزَّكَاةَ وَارْكَعُوا مَعَ
                  الرَّاكِعِينَ
                </div>
                <div class="dasar-translation">
                  Dan laksanakanlah salat, tunaikanlah zakat, dan rukuklah
                  beserta orang-orang yang rukuk.
                </div>
              </div>
            </div>

            <div class="swiper-slide">
              <div class="dasar-card">
                <div class="dasar-title">Surat At-Taubah Ayat 103</div>
                <div class="dasar-arabic">
                  خُذْ مِنْ أَمْوَالِهِمْ صَدَقَةً تُطَهِّرُهُمْ وَتُزَكِّيهِمْ
                  بِهَا وَصَلِّ عَلَيْهِمْ إِنَّ صَلَاتَكَ سَكَنٌ لَّهُمْ
                  وَاللَّهُ سَمِيعٌ عَلِيمٌ
                </div>
                <div class="dasar-translation">
                  Ambillah zakat dari sebagian harta mereka, dengan zakat itu
                  kamu membersihkan dan menyucikan mereka dan berdoalah untuk
                  mereka. Sesungguhnya doamu itu (menjadi) ketenteraman bagi
                  mereka. Allah Maha Mendengar, Maha Mengetahui.
                </div>
              </div>
            </div>

            <div class="swiper-slide">
              <div class="dasar-card">
                <div class="dasar-title">
                  Hadist Keutamaan Zakat Fitrah di Bulan Ramadhan
                </div>
                <div class="dasar-arabic">
                  فَرَضَ رَسُولُ اللَّهِ صَلَّى اللهُ عَلَيْهِ وَسَلَّمَ زَكَاةَ
                  الْفِطْرِ طُهْرَةً لِلصَّائِمِ مِنَ اللَّغْوِ وَالرَّفَثِ،
                  وَطُعْمَةً لِلْمَسَاكِينِ، فَمَنْ أَدَّاهَا قَبْلَ الصَّلَاةِ
                  فَهِيَ زَكَاةٌ مَقْبُولَةٌ، وَمَنْ أَدَّاهَا بَعْدَ الصَّلَاةِ
                  فَهِيَ صَدَقَةٌ مِنَ الصَّدَقَاتِ
                </div>
                <div class="dasar-translation">
                  Dari Ibnu Abbas RA, ia berkata, Rasulullah SAW mewajibkan
                  zakat fitrah untuk membersihkan orang yang berpuasa dari
                  perkataan sia-sia dan perkataan kotor, dan sebagai makanan
                  bagi orang-orang miskin. Barang siapa yang menunaikannya
                  sebelum shalat (Idul Fitri), berarti ini merupakan zakat yang
                  diterima, dan barang siapa yang menunaikannya setelah shalat
                  (Idul Fitri) berarti hal itu merupakan sedekah biasa. (HR. Abu
                  Daud, Ibnu Majah, dan Daru Quthni)
                </div>
              </div>
            </div>
          </div>
          <div class="swiper-pagination"></div>
        </div>
      </div>
    </section>

    <!-- Counts Section -->
    <section id="counts" class="counts">
      <div class="container" data-aos="fade-up">
        <div class="section-title">
          <h2>Data Terkini</h2>
          <h3>Statistik <span>Zakat, Infaq, Sadaqah & Fidyah</span></h3>
          <p>Data terkini pengumpulan dan penyaluran dari ketiga jenis zakat</p>
        </div>

        <!-- Tabs Navigation -->
        <div class="stats-tabs" data-aos="fade-up" data-aos-delay="100">
          <button
            v-for="tab in ['fitrah', 'mal', 'fidyah', 'mustahiq']"
            :key="tab"
            @click="
              () => {
                console.log('Clicking tab:', tab, '-> Stats:', zakatStats[tab]);
                activeTab = tab;
              }
            "
            :class="['tab-btn', { active: activeTab === tab }]"
          >
            <i
              :class="
                tab === 'fitrah'
                  ? 'bx bxs-leaf'
                  : tab === 'mal'
                    ? 'bx bxs-coin'
                    : tab === 'fidyah'
                      ? 'bx bxs-heart-circle'
                      : 'bx bxs-user-detail'
              "
            ></i>
            {{
              tab === "fitrah"
                ? "Zakat Fitrah"
                : tab === "mal"
                  ? "Zakat Mal"
                  : tab === "fidyah"
                    ? "Fidyah"
                    : "Data Mustahiq"
            }}
          </button>
        </div>

        <!-- Stats Display -->
        <div class="stats-content" data-aos="fade-up" data-aos-delay="150">
          <div v-if="isLoadingStats" class="row">
            <div class="col-lg-3 col-md-6" v-for="i in 4" :key="i">
              <SkeletonCard />
            </div>
          </div>

          <div v-else class="row">
            <template v-if="activeTab === 'fitrah'">
              <div class="col-lg-4 col-md-6">
                <div class="count-box">
                  <i class="bi bi-people-fill"></i>
                  <span>{{ currentStats.muzakki }}</span>
                  <p>Muzakki (Pembayar)</p>
                </div>
              </div>
              <div class="col-lg-4 col-md-6 mt-5 mt-md-0">
                <div class="count-box">
                  <i class="bi bi-bag-check"></i>
                  <span>{{ currentStats.beras }} kg</span>
                  <p>Beras Terkumpul</p>
                </div>
              </div>
              <div class="col-lg-4 col-md-6 mt-5 mt-lg-0">
                <div class="count-box">
                  <i class="bi bi-cash-coin"></i>
                  <span>Rp {{ formatCurrency(currentStats.uang) }}</span>
                  <p>Dana Terkumpul</p>
                </div>
              </div>
            </template>

            <template v-else-if="activeTab === 'mal'">
              <div class="col-lg-6 col-md-6">
                <div class="count-box">
                  <i class="bi bi-people-fill"></i>
                  <span>{{ currentStats.muzakki }}</span>
                  <p>Muzakki (Pembayar)</p>
                </div>
              </div>
              <div class="col-lg-6 col-md-6 mt-5 mt-md-0">
                <div class="count-box">
                  <i class="bi bi-cash-coin"></i>
                  <span>Rp {{ formatCurrency(currentStats.uang) }}</span>
                  <p>Dana Terkumpul</p>
                </div>
              </div>
            </template>

            <template v-else-if="activeTab === 'fidyah'">
              <div class="col-lg-4 col-md-6">
                <div class="count-box">
                  <i class="bi bi-people-fill"></i>
                  <span>{{ currentStats.muzakki }}</span>
                  <p>Muzakki (Pembayar)</p>
                </div>
              </div>
              <div class="col-lg-4 col-md-6 mt-5 mt-md-0">
                <div class="count-box">
                  <i class="bi bi-bag-check"></i>
                  <span>{{ currentStats.beras }} kg</span>
                  <p>Beras Terkumpul</p>
                </div>
              </div>
              <div class="col-lg-4 col-md-6 mt-5 mt-lg-0">
                <div class="count-box">
                  <i class="bi bi-cash-coin"></i>
                  <span>Rp {{ formatCurrency(currentStats.uang) }}</span>
                  <p>Dana Terkumpul</p>
                </div>
              </div>
            </template>

            <template v-else>
              <div class="col-lg-3 col-md-6 mb-4">
                <div class="count-box">
                  <i class="bi bi-heart-hands"></i>
                  <span>{{ mustahiqGlobalStats.total }}</span>
                  <p>Total Mustahiq (Global)</p>
                </div>
              </div>
              <div
                v-for="item in mustahiqJenisCards"
                :key="item.jenis_penerima"
                class="col-lg-3 col-md-6 mb-4"
              >
                <div class="count-box">
                  <i class="bi bi-person-badge"></i>
                  <span>{{ item.total }}</span>
                  <p>{{ formatJenisMustahiq(item.jenis_penerima) }}</p>
                </div>
              </div>
            </template>
          </div>
        </div>

        <!-- CTA Button -->
        <div class="cta-section" data-aos="fade-up" data-aos-delay="200">
          <router-link to="/public" class="btn-dashboard">
            <i class="bi bi-graph-up"></i>
            Akses Dashboard Publik
          </router-link>
          <p class="cta-text">
            Lihat detail lengkap data zakat, infaq, sadaqah, dan fidyah
          </p>
        </div>
      </div>
    </section>

    <!-- Testimonials Section -->
    <section id="testimonials" class="testimonials">
      <div class="container" data-aos="zoom-in">
        <div
          class="testimonials-slider swiper-container"
          data-aos="fade-up"
          data-aos-delay="100"
        >
          <div class="swiper-wrapper">
            <div class="swiper-slide">
              <div class="testimonial-item">
                <h3>Niat Zakat Untuk Diri Sendiri</h3>
                <br />
                <h1>
                  ﻧَﻮَﻳْﺖُ ﺃَﻥْ ﺃُﺧْﺮِﺝَ ﺯَﻛَﺎﺓَ ﺍﻟْﻔِﻄْﺮِ ﻋَنِّيْ ﻭَﻋَﻦْ
                  ﺟَﻤِﻴْﻊِ ﻣَﺎ ﻳَﻠْﺰَﻣُنِيْ ﻧَﻔَﻘَﺎﺗُﻬُﻢْ ﺷَﺮْﻋًﺎ ﻓَﺮْﺿًﺎ ﻟﻠﻪِ
                  ﺗَﻌَﺎﻟَﻰ
                </h1>
                <p>
                  <i class="bx bxs-quote-alt-left quote-icon-left"></i>
                  Aku niat mengeluarkan zakat fitrah untuk diriku dan seluruh
                  orang yang nafkahnya menjadi tanggunganku fardhu karena Allah
                  Taala.
                  <i class="bx bxs-quote-alt-right quote-icon-right"></i>
                </p>
              </div>
            </div>

            <div class="swiper-slide">
              <div class="testimonial-item">
                <h3>Niat Zakat Untuk Diri Sendiri dan seluruh keluarga</h3>
                <br />
                <h1>
                  ﻧَﻮَﻳْﺖُ ﺃَﻥْ ﺃُﺧْﺮِﺝَ ﺯَﻛَﺎﺓَ ﺍﻟْﻔِﻄْﺮِ ﻋَنِّيْ ﻭَﻋَﻦْ
                  ﺟَﻤِﻴْﻊِ ﻣَﺎ ﻳَﻠْﺰَﻣُنِيْ ﻧَﻔَﻘَﺎﺗُﻬُﻢْ ﺷَﺮْﻋًﺎ ﻓَﺮْﺿًﺎ ﻟﻠﻪِ
                  ﺗَﻌَﺎﻟَﻰ
                </h1>
                <p>
                  <i class="bx bxs-quote-alt-left quote-icon-left"></i>
                  Aku niat mengeluarkan zakat fitrah untuk diriku dan seluruh
                  orang yang nafkahnya menjadi tanggunganku fardhu karena Allah
                  Taala.
                  <i class="bx bxs-quote-alt-right quote-icon-right"></i>
                </p>
              </div>
            </div>
          </div>
          <div class="swiper-pagination"></div>
        </div>
      </div>
    </section>
  </main>

  <!-- Footer -->
  <footer id="footer">
    <div class="container py-4">
      <div class="copyright">
        &copy; Copyright <strong><span>SI-AZISAH</span></strong
        >. All Rights Reserved
      </div>
      <div class="credits">
        Sistem Informasi Zakat, Infaq, Sadaqah, dan Fidyah
      </div>
    </div>
  </footer>

  <div id="preloader"></div>
  <a
    href="#"
    class="back-to-top d-flex align-items-center justify-content-center"
    ><i class="bi bi-arrow-up-short"></i
  ></a>
</template>

<script setup>
import { onMounted, onUnmounted, ref, computed } from "vue";
import SkeletonCard from "../components/SkeletonCard.vue";

const isInitialized = ref(false);
const zakatStats = ref({
  fitrah: { muzakki: 0, beras: 0, uang: 0, mustahiq: 0 },
  mal: { muzakki: 0, beras: 0, uang: 0, mustahiq: 0 },
  fidyah: { muzakki: 0, beras: 0, uang: 0, mustahiq: 0 },
});
const mustahiqGlobalStats = ref({
  total: 0,
  perJenis: [],
});
const isLoadingStats = ref(true);
const activeTab = ref("fitrah");
const resolveApiUrl = () => {
  const rawUrl =
    import.meta.env.VITE_API_BASE_URL || "http://localhost:8082/api";
  const withoutTrailingSlash = rawUrl.replace(/\/$/, "");
  return withoutTrailingSlash.endsWith("/api")
    ? withoutTrailingSlash
    : `${withoutTrailingSlash}/api`;
};

const apiUrl = resolveApiUrl();

// Computed property for current zakat stats based on active tab
const currentStats = computed(() => {
  console.log(`Getting stats for tab: ${activeTab.value}`);
  return (
    zakatStats.value[activeTab.value] || {
      muzakki: 0,
      beras: 0,
      uang: 0,
      mustahiq: 0,
    }
  );
});

const defaultMustahiqJenis = [
  "amil",
  "fakir",
  "miskin",
  "mualaf",
  "riqab",
  "gharim",
  "fisabilillah",
  "ibnu sabil",
];

const normalizeJenisMustahiq = (jenis) =>
  String(jenis || "")
    .trim()
    .toLowerCase();

const mustahiqJenisCards = computed(() => {
  const incoming = Array.isArray(mustahiqGlobalStats.value.perJenis)
    ? mustahiqGlobalStats.value.perJenis
    : [];

  const jenisMap = new Map();
  incoming.forEach((item) => {
    const key = normalizeJenisMustahiq(item.jenis_penerima);
    if (!key) return;
    jenisMap.set(key, Number(item.total || 0));
  });

  const defaults = defaultMustahiqJenis.map((jenis) => ({
    jenis_penerima: jenis,
    total: Number(jenisMap.get(jenis) || 0),
  }));

  const extras = [];
  jenisMap.forEach((total, jenis) => {
    if (!defaultMustahiqJenis.includes(jenis)) {
      extras.push({ jenis_penerima: jenis, total });
    }
  });

  return [...defaults, ...extras];
});

const formatJenisMustahiq = (jenis) => {
  const value = String(jenis || "").trim();
  if (!value) return "Lainnya";
  return value
    .split(" ")
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1).toLowerCase())
    .join(" ");
};

// Load vendor scripts dynamically
const loadScript = (src) => {
  return new Promise((resolve) => {
    const script = document.createElement("script");
    script.src = src;
    script.onload = () => {
      console.log("✓ Loaded:", src);
      resolve(true);
    };
    script.onerror = (err) => {
      console.error("✗ Failed to load:", src, err);
      resolve(false); // Don't reject - continue anyway
    };
    document.body.appendChild(script);
  });
};
// Fetch zakat statistics
const fetchZakatStats = async () => {
  try {
    isLoadingStats.value = true;
    const startTime = performance.now();
    console.log("Fetching zakat stats from:", apiUrl);

    const dashboardRes = await fetch(`${apiUrl}/public/dashboard`);
    const dashboardData = dashboardRes.ok ? await dashboardRes.json() : null;

    if (dashboardData?.data) {
      zakatStats.value.fitrah = {
        muzakki: Number(dashboardData.data.total_muzakki || 0),
        beras: Number(dashboardData.data.total_zakat_fitrah_beras_kg || 0),
        uang: Number(dashboardData.data.total_zakat_fitrah_uang || 0),
        mustahiq: Number(dashboardData.data.total_mustahiq || 0),
      };

      zakatStats.value.mal = {
        muzakki: 0,
        beras: 0,
        uang: Number(dashboardData.data.total_zakat_mal || 0),
        mustahiq: Number(dashboardData.data.total_mustahiq || 0),
      };

      zakatStats.value.fidyah = {
        muzakki: 0,
        beras: Number(dashboardData.data.total_fidyah_beras_kg || 0),
        uang: Number(dashboardData.data.total_fidyah_uang || 0),
        mustahiq: Number(dashboardData.data.total_mustahiq || 0),
      };

      mustahiqGlobalStats.value = {
        total: Number(dashboardData.data.total_mustahiq || 0),
        perJenis: Array.isArray(dashboardData.data.mustahiq_per_jenis)
          ? dashboardData.data.mustahiq_per_jenis
          : [],
      };
    } else {
      console.warn("⚠ Dashboard response has no data property");
    }

    const endTime = performance.now();
    console.log("Final stats:", zakatStats.value);
    console.log(`Fetching completed in ${(endTime - startTime).toFixed(2)}ms`);
  } catch (error) {
    console.error("Error fetching zakat stats:", error);
  } finally {
    isLoadingStats.value = false;
  }
};

// Format currency to Rupiah
const formatCurrency = (value) => {
  if (!value || value === 0) return "0";
  if (value >= 1000000) {
    return (value / 1000000).toFixed(1) + " Jt";
  } else if (value >= 1000) {
    return (value / 1000).toFixed(1) + " Rb";
  }
  return value.toString();
};
let swiper = null;

const initSwiper = () => {
  try {
    const Swiper = window.Swiper;
    if (!Swiper) {
      console.warn("Swiper not loaded");
      return;
    }

    // Testimonials slider
    new Swiper(".testimonials-slider", {
      speed: 600,
      loop: true,
      autoplay: {
        delay: 5000,
        disableOnInteraction: false,
      },
      slidesPerView: "auto",
      pagination: {
        el: ".testimonials-slider .swiper-pagination",
        type: "bullets",
        clickable: true,
      },
    });
    console.log("✓ Testimonials slider initialized");

    // Dasar kewajiban zakat slider
    new Swiper(".dasar-slider", {
      speed: 600,
      loop: true,
      autoplay: {
        delay: 6000,
        disableOnInteraction: false,
      },
      slidesPerView: 1,
      pagination: {
        el: ".dasar-slider .swiper-pagination",
        type: "bullets",
        clickable: true,
      },
    });
    console.log("✓ Dasar kewajiban zakat slider initialized");

    // Zakat types slider
    const zakatSliderEl = document.querySelector(".zakat-slider");
    const zakatPaginationEl = document.querySelector(
      ".zakat-slider .swiper-pagination",
    );

    console.log("Zakat slider element:", zakatSliderEl);
    console.log("Zakat pagination element:", zakatPaginationEl);

    if (zakatSliderEl) {
      new Swiper(".zakat-slider", {
        speed: 600,
        loop: true,
        autoplay: {
          delay: 6000,
          disableOnInteraction: false,
        },
        slidesPerView: 1,
        pagination: {
          el: ".zakat-slider .swiper-pagination",
          type: "bullets",
          clickable: true,
        },
        breakpoints: {
          768: {
            slidesPerView: 1,
          },
        },
      });
      console.log("✓ Zakat types slider initialized");
    } else {
      console.error("Zakat slider element not found!");
    }
  } catch (error) {
    console.error("Swiper init error:", error);
  }
};

const initAOS = () => {
  try {
    const AOS = window.AOS;
    if (!AOS) {
      console.warn("AOS not loaded");
      return;
    }

    AOS.init({
      duration: 1000,
      easing: "ease-in-out",
      once: true,
      mirror: false,
    });
    console.log("✓ AOS initialized");
  } catch (error) {
    console.error("AOS init error:", error);
  }
};

const initGLightbox = () => {
  try {
    const GLightbox = window.GLightbox;
    if (!GLightbox) {
      console.warn("GLightbox not loaded");
      return;
    }

    GLightbox({
      selector: ".glightbox",
    });
    console.log("✓ GLightbox initialized");
  } catch (error) {
    console.error("GLightbox init error:", error);
  }
};

const handleNavClick = (e) => {
  const navLinks = document.querySelectorAll("#navbar a");
  navLinks.forEach((link) => {
    link.classList.remove("active");
  });
  if (e.target.classList.contains("nav-link")) {
    e.target.classList.add("active");
  }
};

const handleMobileMenuToggle = () => {
  const navbar = document.getElementById("navbar");
  if (navbar) {
    navbar.classList.toggle("navbar-mobile");
  }
};

const scrollWithOffset = (e) => {
  if (
    e.target.classList.contains("scrollto") &&
    e.target.getAttribute("href")?.startsWith("#")
  ) {
    e.preventDefault();
    const target = document.querySelector(e.target.getAttribute("href"));
    if (target) {
      target.scrollIntoView({ behavior: "smooth" });
    }
  }
};

onMounted(async () => {
  console.log("Home.vue mounted - starting script loading...");

  // Hide preloader immediately
  const preloader = document.getElementById("preloader");
  if (preloader) {
    preloader.style.display = "none";
  }

  // Fetch zakat stats from database
  fetchZakatStats();

  // Load vendor scripts in parallel with Promise.allSettled (don't fail if any fail)
  try {
    await Promise.allSettled([
      loadScript(
        "/homepage/assets/vendor/bootstrap/js/bootstrap.bundle.min.js",
      ),
      loadScript("/homepage/assets/vendor/aos/aos.js"),
      loadScript("/homepage/assets/vendor/glightbox/js/glightbox.min.js"),
      loadScript("/homepage/assets/vendor/swiper/swiper-bundle.min.js"),
    ]);

    // Initialize components - they check if libraries exist
    setTimeout(() => {
      initAOS();
      initGLightbox();
      initSwiper();
      isInitialized.value = true;
    }, 500);
  } catch (error) {
    console.error("Script loading error:", error);
    isInitialized.value = true; // Show content anyway
  }

  // Add event listeners
  document.querySelectorAll(".mobile-nav-toggle").forEach((toggle) => {
    toggle.addEventListener("click", handleMobileMenuToggle);
  });

  document.querySelectorAll("#navbar a").forEach((link) => {
    link.addEventListener("click", handleNavClick);
    link.addEventListener("click", scrollWithOffset);
  });

  // Back to top button
  const backToTopBtn = document.querySelector(".back-to-top");
  if (backToTopBtn) {
    window.addEventListener("scroll", () => {
      if (window.scrollY > 100) {
        backToTopBtn.classList.add("active");
      } else {
        backToTopBtn.classList.remove("active");
      }
    });

    backToTopBtn.addEventListener("click", (e) => {
      e.preventDefault();
      window.scrollTo({ top: 0, behavior: "smooth" });
    });
  }
});

onUnmounted(() => {
  if (swiper) {
    swiper.destroy();
  }
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,600,600i,700,700i|Roboto:300,300i,400,400i,500,500i,600,600i,700,700i|Poppins:300,300i,400,400i,500,500i,600,600i,700,700i");

:global(#preloader) {
  display: none !important;
}

/* Zakat Cards Styling */
.zakat-slider-wrapper {
  width: 100%;
  padding: 30px 0;
}

.zakat-slider {
  width: 100%;
  padding: 0;
  position: relative;
  overflow: hidden;
}

.zakat-slider .swiper-wrapper {
  display: flex;
  width: 100%;
}

.zakat-slider .swiper-slide {
  width: 100% !important;
  height: auto;
  flex-shrink: 0;
  padding: 0 15px;
}

.zakat-card {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: all 0.3s ease;
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}

.zakat-card:hover {
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  transform: translateY(-5px);
}

.zakat-header {
  padding: 30px 20px;
  text-align: center;
  color: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 120px;
}

.zakat-header i {
  font-size: 48px;
  margin-bottom: 15px;
}

.zakat-header h3 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.zakat-fitrah {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.zakat-mal {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.zakat-fidyah {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.zakat-content {
  padding: 25px;
}

.zakat-content .description {
  font-size: 14px;
  color: #555;
  margin-bottom: 20px;
  line-height: 1.6;
  font-style: italic;
}

.kategori-list {
  display: grid;
  grid-template-columns: 1fr;
  gap: 20px;
}

.kategori-item {
  background: #f8f9fa;
  padding: 15px;
  border-radius: 8px;
  border-left: 4px solid #667eea;
}

.zakat-mal .kategori-item {
  border-left-color: #f5576c;
}

.zakat-fidyah .kategori-item {
  border-left-color: #00f2fe;
}

.kategori-item h5 {
  margin: 0 0 12px 0;
  font-size: 14px;
  font-weight: 600;
  color: #333;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.kategori-item ul {
  margin: 0;
  padding: 0 0 0 20px;
}

.kategori-item li {
  margin: 8px 0;
  font-size: 13px;
  color: #666;
  line-height: 1.5;
}

.kategori-item strong {
  color: #333;
  font-weight: 600;
}

/* Info Box */
.info-box {
  margin-top: 40px;
  padding: 30px 0;
}

.info-item {
  text-align: center;
  padding: 20px;
  transition: all 0.3s ease;
}

.info-item:hover {
  transform: translateY(-5px);
}

.info-item i {
  font-size: 40px;
  color: #667eea;
  margin-bottom: 15px;
}

.info-item h5 {
  margin: 15px 0;
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.info-item p {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
}

/* Swiper Pagination */
:global(.swiper-slide) {
  display: flex !important;
}

:global(.zakat-slider .swiper-pagination) {
  position: relative;
  margin-top: 20px;
}

:global(.zakat-slider .swiper-pagination-bullet) {
  background: #667eea;
  opacity: 0.6;
}

:global(.zakat-slider .swiper-pagination-bullet-active) {
  background: #667eea;
  opacity: 1;
}

:global(.testimonials-slider .swiper-pagination-bullet) {
  background: #667eea;
  opacity: 0.6;
}

:global(.testimonials-slider .swiper-pagination-bullet-active) {
  background: #667eea;
  opacity: 1;
}

/* Dasar Kewajiban Zakat Slider */
.dasar-slider {
  margin-top: 10px;
}

.dasar-slider :global(.swiper-slide) {
  align-items: stretch;
  justify-content: center;
}

.dasar-card {
  background: white;
  border-radius: 16px;
  padding: 28px 32px;
  max-width: 900px;
  width: 100%;
  box-shadow: 0 20px 50px rgba(14, 21, 47, 0.08);
  border: 1px solid rgba(102, 126, 234, 0.12);
  text-align: center;
}

.dasar-title {
  font-size: 18px;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 14px;
}

.dasar-arabic {
  font-family: "Traditional Arabic", serif;
  font-size: 26px;
  line-height: 1.7;
  color: #111827;
  margin-bottom: 16px;
}

.dasar-translation {
  font-size: 14px;
  line-height: 1.7;
  color: #4b5563;
  font-style: italic;
}

:global(.dasar-slider .swiper-pagination) {
  position: relative;
  margin-top: 20px;
}

:global(.dasar-slider .swiper-pagination-bullet) {
  background: #667eea;
  opacity: 0.5;
}

:global(.dasar-slider .swiper-pagination-bullet-active) {
  background: #667eea;
  opacity: 1;
}

/* Responsive */
@media (max-width: 768px) {
  .zakat-header {
    min-height: 100px;
  }

  .zakat-header i {
    font-size: 36px;
  }

  .kategori-list {
    grid-template-columns: 1fr;
  }

  .info-item {
    margin-bottom: 15px;
  }
}

/* Stats Tabs */
.stats-tabs {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-bottom: 40px;
  flex-wrap: wrap;
}

.tab-btn {
  background: #f8f9fa;
  border: 2px solid #e9ecef;
  color: #666;
  padding: 12px 24px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
}

.tab-btn i {
  font-size: 18px;
}

.tab-btn:hover {
  background: #e9ecef;
  border-color: #667eea;
  color: #667eea;
}

.tab-btn.active {
  background: #667eea;
  border-color: #667eea;
  color: white;
}

/* Stats Content */
.stats-content {
  margin-bottom: 40px;
}

/* CTA Section */
.cta-section {
  text-align: center;
  padding: 40px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  margin-top: 30px;
}

.btn-dashboard {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  background: white;
  color: #667eea;
  padding: 14px 32px;
  border-radius: 8px;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.3s ease;
  font-size: 16px;
}

.btn-dashboard:hover {
  background: #f8f9fa;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.4);
  color: #667eea;
}

.btn-dashboard i {
  font-size: 18px;
}

.cta-text {
  margin-top: 15px;
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
  margin: 15px 0 0 0;
}
</style>
