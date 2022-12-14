package data

func TownData() []string {
	townDataList := []string{
		"İstanbul", "Adalar", "Arnavutkoy", "Ataşehir", "Avcilar", "Bagcilar", "Bahcelievler", "Bakirkoy", "Başakşehir", "Bayrampaşa",
		"Beşiktaş", "Beykoz", "Beylikduzu", "Beyoglu", "Büyükçekmece", "Çatalca", "Çekmeköy", "Esenler", "Esenyurt", "Eyüpsultan",
		"Fatih", "Gaziosmanpaşa", "Güngören", "Kadikoy", "Kagithane", "Kartal", "Küçükçekmece", "Maltepe", "Pendik", "Sancaktepe",
		"Sariyer", "Silivri", "Sultanbeyli", "Sultangazi", "Şile", "Şişli", "Tuzla", "Ümraniye", "Üsküdar", "Zeytinburnu"}

	return townDataList
}

func DistrictData() []string {
	districtDataList := []string{
		"Atakent Küçükçekmece", "Adnan Kahveci Beylikdüzü", "Zafer Bahçelievler", "Zümrütevler Maltepe", "Halkali Merkez Küçükçekmece",
		"Kayabaşi Başakşehir", "Karadeniz Gaziosmanpaşa", "İnönü Küçükçekmece", "Kocasinan Merkez Bahçelievler",
		"Soğanli Bahçelievler", "Siyavuşpaşa Bahçelievler", "Esentepe Sultangazi", "Cihangir Avcilar",
		"Cebeci Sultangazi", "Kavakpinar Pendik", "Findikli Maltepe ", "Başakşehir Başakşehir",
		"Yenişehir Pendik", "Şirinevler Bahçelievler", "Güvercintepe Başakşehir", "Aydinli Tuzla",
		"İsmetpaşa Sultangazi", "Tahtakale Avcilar", "Bariş Beylikdüzü", "Güneştepe Güngören",
		"Cumhuriyet Küçükçekmece", "Kaynarca Pendik", "Hürriyet Kartal", "Hürriyet Bahçelievler",
		"İstiklal Ümraniye", "Kazim Karabekir Gaziosmanpaşa", "Yunus Emre Sultangazi",
		"Güneşli Bağcilar", "Denizköşkler Avcilar", "Fatih Esenler",
		"Fatih Bağcilar", "Kartaltepe Bayrampaşa", "Yakuplu Beylikdüzü",
		"Kirazli Bağcilar", "Nine Hatun Esenler",

		"Bağlarbaşi Maltepe", "Cumhuriyet Bahçelievler", "Uğur Mumcu Sultangazi", "Uğur Mumcu Kartal",
		"Gümüşpala Avcilar", "Akşemsettin Eyüpsultan", "Karayollari Gaziosmanpaşa ", "Gençosman Güngören ", "Yildiztepe Bağcilar ", " Sultançiftliği Sultangazi ",
		"Yeşilpinar Eyüpsultan ", "Ambarli Avcilar ", "Hamidiye Kağithane ", "İstasyon Küçükçekmece ",
		"Kartaltepe Bakirköy ", " Göztepe Bağcilar ", "Telsiz Zeytinburnu ", "Armağanevler Ümraniye ", " Çakmak Ümraniye ", " Fevzi Çakmak Pendik ", " Sümer Zeytinburnu ", " Çinar Bağcilar ",
		"Pinar Esenyurt ", " Velibaba Pendik ", " Cumhuriyet Üsküdar ",
		"Mehterçeşme Esenyurt ", "Kazim Karabekir Esenler ",
		"Cevizli Maltepe ", "Ayazağa Sariyer ", "Göztepe Kadiköy ", " Yenimahalle Bağcilar ", " Gazi Sultangazi ", "Kemalpaşa Bağcilar ", "Güzeltepe Eyüpsultan ",
		"Yavuztürk Üsküdar ", "Kozyataği Kadiköy ", " Ünalan Üsküdar ",
		" Muratpaşa Bayrampaşa ",
		" Fevzi Çakmak Esenler ", " Güzelyali Pendik ",
		" Çobançeşme Bahçelievler ",
		" Söğütlü Çeşme Küçükçekmece ", " Anadolu Arnavutköy ", " Merdivenköy Kadiköy ",
		" Yeşilova Küçükçekmece ", " Menderes Esenler ",
		" Çeliktepe Kağithane ",
		" Zübeyde Hanim Sultangazi ", " Yeşilkent Esenyurt ", " Abdurrahmangazi Sancaktepe ", " Esenler Pendik ", " Çamçeşme Pendik ", " Bostanci Kadiköy ",
		" Mehmet Akif Çekmeköy ",
		" Atalar Kartal ", " Erenköy Kadiköy ", " Karliktepe Kartal ",
		"Sahrayi Cedit Kadiköy", " Kurtköy Pendik ", " Bulgurlu Üsküdar ", " Acibadem Kadiköy ", " Altintepe Maltepe ",
		" Altayçeşme Maltepe ", " Gültepe Küçükçekmece ", " Çağlayan Kağithane ",
		" Gürsel Kağithane ", " Cennet Küçükçekmece ",
		" Orhantepe Kartal ", " İnkilap Ümraniye ",
		" Hamidiye Çekmeköy ",
		" Çirpici Zeytinburnu ",
		" Namik Kemal Ümraniye ", " Bağlarbaşi Gaziosmanpaşa ",
		" Yavuz Selim Bağcilar ",
		" Hürriyet Gaziosmanpaşa ",
		" Kazim Karabekir Bağcilar ", " Atatürk Sancaktepe ", " Çirçir Eyüpsultan ",
		" Bağlarçeşme Esenyurt ", " Birlik Esenler ",
		" Altintepsi Bayrampaşa ",
		" Şifa Tuzla ", " Altinşehir Ümraniye ",
		" Şenlikköy Bakirköy ",
		" Orhangazi Pendik ",
		" Fatih Sancaktepe ",
		" Veliefendi Zeytinburnu ",
		" Kavakli Beylikdüzü ", " Nuripaşa Zeytinburnu ",
		" Cevizli Kartal ",
		" Necip Fazil Ümraniye ", " Karadolap Eyüpsultan ",
		" Esenevler Ümraniye ", " Esentepe Kartal ", " Yayla Tuzla ",
		" Seyyid Ömer Fatih ", " Ahmet Yesevi Pendik ",
		" Soğanlik Yeni Kartal ", " Marmara Beylikdüzü ", " Gümüşpinar Kartal ", " Fevzi Çakmak Bahçelievler ", " Talatpaşa Esenyurt ", " Site Ümraniye ", " Nurtepe Kağithane ", " Meclis Sancaktepe ", " Yildiztabya Gaziosmanpaşa ",
		" Dumlupinar Pendik ", " İncirtepe Esenyurt ",
		" Fatih Esenyurt ", " Fevzi Çakmak Küçükçekmece ", " Hürriyet Bağcilar ",
		" Sarigöl Gaziosmanpaşa ",
		" Mehmet Akif Ümraniye ", " Akşemsettin Fatih ",
		" Acibadem Üsküdar ",
		" Atatürk Ümraniye ",
		" İnönü Bağcilar ",
		" Mahmutbey Bağcilar ",
		" Feneryolu Kadiköy ",
		" İnönü Esenyurt ",
		" Şemsipaşa Gaziosmanpaşa ",
		" Emek Sancaktepe ",
		" Mimar Sinan Tuzla ",
		" Osmaniye Bakirköy ",
		" Seyitnizam Zeytinburnu ",
		" Mimar Sinan Çekmeköy ",
		" Aydintepe Tuzla ",
		" Caferağa Kadiköy ",
		" Beştelsiz Zeytinburnu ",
		" Mevlana Gaziosmanpaşa ",
		" Yeşilköy Bakirköy ",
		" Merkezefendi Zeytinburnu ",
		" Gülbahar Şişli ",
		" Yenigün Bağcilar ",
		" Ortabayir Kağithane ",
		" İnönü Sancaktepe ",
		" Fevzi Çakmak Gaziosmanpaşa ",
		" Abdurrahman Nafiz Gürman Güngören ",
		" Seyrantepe Kağithane ",
		" Suadiye Kadiköy ",
		" Çamlik Ümraniye ",
		" Kavacik Beykoz ",
		" Elmalikent Ümraniye ",
		" Şehremini Fatih ",
		" Harmantepe Kağithane ",
		" İdealtepe Maltepe ",
		" Barbaros Bağcilar ",
		" Koca Mustafapaşa Fatih ",
		" Saadetdere Esenyurt ",
		" Şirintepe Kağithane ",
		" Sarigazi Sancaktepe ",
		" Yeşiltepe Zeytinburnu ",
		" Cumhuriyet Beylikdüzü ",
		" Başibüyük Maltepe ",
		" Üniversite Avcilar ",
		" Cumhuriyet Kartal ",
		" Hadimköy Arnavutköy ",
		" Kocatepe Bayrampaşa ",
		" Namik Kemal Esenyurt ",
		" Haznedar Güngören ",
		" Bahçelievler Üsküdar ",
		" Atakent Ümraniye ",
		" Ziya Gökalp Başakşehir ",
		" Akincilar Güngören ",
		" Zuhuratbaba Bakirköy ",
		" Yenikent Esenyurt ",
		" Akpinar Sancaktepe ",
		" Adil Sultanbeyli ",
		" Büyükşehir Beylikdüzü ",
		" Çamlik Çekmeköy ",
		" Ferah Üsküdar ",
		" Mevlanakapi Fatih ",
		" Yukari Dudullu Ümraniye ",
		" Mecidiyeköy Şişli ",
		" Hürriyet Kağithane ",
		" Yenidoğan Sancaktepe ",
		" Aydinevler Maltepe ",
		" Tatlisu Ümraniye ",
		" Esentepe Eyüpsultan ",
		" Gökalp Zeytinburnu ",
		" Esenkent Ümraniye ",
		" İstasyon Tuzla ",
		" Çinar Maltepe ",
		" Kaptanpaşa Beyoğlu ",
		" Kisikli Üsküdar ",
		" Postane Tuzla ",
		" Gürpinar Beylikdüzü ",
		" Veysel Karani Sancaktepe ",
		" Sancaktepe Bağcilar ",
		" Karlitepe Gaziosmanpaşa ",
		" Kemal Türkler Sancaktepe ",
		" Fatih Büyükçekmece ",
		" Akçaburgaz Esenyurt ",
		" Esenşehir Ümraniye ",
		" Cevatpaşa Bayrampaşa ",
		" Tozkoparan Güngören ",
		" Feyzullah Maltepe ",
		" Zekeriyaköy Sariyer ",
		" Orta Kartal ",
		" Atatürk Büyükçekmece ",
		" Paşa Şişli ",
		" Yenidoğan Bayrampaşa ",
		" Arnavutköy Merkez Arnavutköy ",
		" Örnek Esenyurt ",
		" Çinardere Pendik ",
		" Düğmeciler Eyüpsultan ",
		" Pinartepe Büyükçekmece ",
		" Caddebostan Kadiköy ",
		" Molla Gürani Fatih ",
		" Mimar Sinan Sultanbeyli ",
		" Barbaros Üsküdar ",
		" İstiklal Esenyurt ",
		" Tarabya Sariyer ",
		" Güven Güngören ",
		" Yahya Kemal Kağithane ",
		" Fenerbahçe Kadiköy ",
		" Hasanpaşa Sultanbeyli ",
		" Yedikule Fatih ",
		" Fulya Şişli ",
		" İcadiye Üsküdar ",
		" Sümbül Efendi Fatih ",
		" Halide Edip Adivar Şişli ",
		" Dikilitaş Beşiktaş ",
		" Silivrikapi Fatih ",
		" Davutpaşa Esenler ",
		" Şeyhli Pendik ",
		" Aydinlar Çekmeköy ",
		" Namik Kemal Esenler ",
		" Ihlamurkuyu Ümraniye ",
		" Yakacik Çarşi Kartal ",
		" Yakacik Yeni Kartal ",
		" Girne Maltepe ",
		" İstinye Sariyer ",
		" Doğu Pendik ",
		" Çamlik Pendik ",
		" Mevlana Sancaktepe ",
		" Terazidere Bayrampaşa ",
		" Bağlar Bağcilar ",
		" Murat Çesme Büyükçekmece ",
		" Hasanpaşa Kadiköy ",
		" Bati Pendik ",
		" Yeniköy Sariyer ",
		" Konaklar Beşiktaş ",
		" İslambey Eyüpsultan ",
		" Yeni Pendik ",
		" Ortamahalle Bayrampaşa ",
		" Cumhuriyet Sultangazi ",
		" Sinanoba Büyükçekmece ",
		" Kemalpaşa Küçükçekmece ",
		" Rami Cuma Eyüpsultan ",
		" Feriköy Şişli ",
		" İsmet Paşa Bayrampaşa ",
		" Gayrettepe Beşiktaş ",
		" Cumhuriyet Esenyurt ",
		" Madenler Ümraniye ",
		" Çengelköy Üsküdar ",
		" Çavuşoğlu Kartal ",
		" Rasimpaşa Kadiköy ",
		" Yamanevler Ümraniye ",
		" Esenyali Pendik ",
		" Rami Yeni Eyüpsultan ",
		" Altunizade Üsküdar ",
		" Gültepe Kağithane ",
		" Sultan Murat Küçükçekmece ",
		" Eğitim Kadiköy ",
		" Halicioğlu Beyoğlu ",
		" Selami Ali Üsküdar ",
		" Maden Sariyer ",
		" Türkoba Büyükçekmece ",
		" Atatürk Esenyurt ",
		" Yali Maltepe ",
		" Cumhuriyet Çekmeköy ",
		" Pazariçi Gaziosmanpaşa ",
		" Aşaği Dudullu Ümraniye ",
		" Zeynep Kamil Üsküdar ",
		" Sütlüce Beyoğlu ",
		" Nisbetiye Beşiktaş ",
		" Bahçelievler Pendik ",
		" Topselvi Kartal ",
		" Kartaltepe Küçükçekmece ",
		" Eskişehir Şişli ",
		" Mimar Sinan Üsküdar ",
		" Etiler Beşiktaş ",
		" Yeşilce Kağithane ",
		" Yayalar Pendik ",
		" Telsizler Kağithane ",
		" Aksaray Fatih ",
		" Dereağzi Beylikdüzü ",
		" Karagümrük Fatih ",
		" Teşvikiye Şişli ",
		" Yayla Şişli ",
		" Sultantepe Üsküdar ",
		" Haseki Sultan Fatih ",
		" Kordonboyu Kartal ",
		" Sapan Bağlari Pendik ",
		" Ali Kuşçu Fatih ",
		" Soğuksu Beykoz ",
		" Dumlupinar Kadiköy ",
		" Bozkurt Şişli ",
		" Yenidoğan Zeytinburnu ",
		" Evliya Çelebi Tuzla ",
		" Mecidiye Beşiktaş ",
		" Örnektepe Beyoğlu ",
		" Fatih Küçükçekmece ",
		" Orhan Gazi Esenyurt ",
		" Türkali Beşiktaş ",
		" Esenkent Esenyurt ",
		" Göztepe Beykoz ",
		" Tantavi Ümraniye ",
		" Kulaksiz Beyoğlu ",
		" Alibeyköy Eyüpsultan ",
		" Ortaköy Beşiktaş ",
		" Salacak Üsküdar ",
		" Ahmediye Üsküdar ",
		" Darüşşafaka Sariyer ",
		" Emniyettepe Eyüpsultan ",
		" Küçük Piyale Beyoğlu ",
		" Silahtarağa Eyüpsultan ",
		" Cerrahpaşa Fatih ",
		" Fikirtepe Kadiköy ",
		" Emirgan Sariyer ",
		" Kirazlidere Çekmeköy ",
		" Aziz Mahmut Hüdayi Üsküdar ",
		" Yenidoğan Gaziosmanpaşa ",
		" Mimar Sinan Merkez Büyükçekmece ",
		" Selimiye Üsküdar ",
		" Çavuş Şile ",
		" Osmanağa Kadiköy ",
		" Şerifali Ümraniye ",
		" Emniyet Evleri Kağithane ",
		" Çatalmeşe Çekmeköy ",
		" Yukari Kartal ",
		" Sakizağaci Bakirköy ",
		" Zühtüpaşa Kadiköy ",
		" Maslak Sariyer ",
		" Harmandere Pendik ",
		" Meşrutiyet Şişli ",
		" Koşuyolu Kadiköy ",
		" Yeşilyurt Bakirköy ",
		" Taşdelen Çekmeköy ",
		" Acarlar Beykoz ",
		" Cumhuriyet Şişli ",
		" Ulus Beşiktaş ",
		" Esentepe Şişli ",
		" Dizdariye Büyükçekmece ",
		" Yenimahalle Bakirköy ",
		" Kumburgaz Büyükçekmece ",
		" Güzelce Büyükçekmece ",
		" Vişnezade Beşiktaş ",
		" Alemdağ Çekmeköy ",
		" Yildiz Beşiktaş ",
		" Levazim Beşiktaş ",
		" Basinköy Bakirköy ",
		" Bebek Beşiktaş ",
		" Zeytinlik Bakirköy ",
		" Cevizlik Bakirköy ",
		" Abbasağa Beşiktaş ",
		" Cami Tuzla ",
		" Baltalimani Sariyer ",
		" Mithatpaşa Eyüpsultan ",
		" Muradiye Beşiktaş ",
		" Kültür Beşiktaş ",
		" Alkent 2000 Büyükçekmece ",
		" Firuzağa Beyoğlu ",
		" İnönü Şişli ",
		" Arnavutköy Beşiktaş ",
		" Cihannüma Beşiktaş ",
		" Balmumcu Beşiktaş ",
		" Cihangir Beyoğlu ",
		" Balibey Şile ",
		" Levent Beşiktaş ",
		" Kuruçeşme Beşiktaş ",
		" Göksu Beykoz ",
		" Maltepe Zeytinburnu ",
		" Sinanpaşa Beşiktaş ",
		" Harbiye Şişli ",
		" Gümüşsuyu Beyoğlu ",
		" Ömer Avni Beyoğlu ",
		" Kandilli Üsküdar ",
		" Halaskargazi Şişli ",
		" Kumbaba Şile ",
		" Beyazit Fatih ",
	}

	return districtDataList
}
