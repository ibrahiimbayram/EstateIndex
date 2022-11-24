# Go ile Docker, Selenium, ElasticSearch ve Kibana

Go ile Emlak endeksi projesi geliştirdim. Docker, Selenium, Elasticsearch ve Kibana teknolojilerini kullandım. 
Projenin amacı, zingat.com üzerinde bulunan İstanbul'un İlçe ve Semt emlak endeksi verilerini anlamlı hale getirerek Kibana üzerinde görselleştirmek.

Özetle mimari yapı aşağıdaki görseldeki gibidir.

![EstateIndexDiagram](https://user-images.githubusercontent.com/47785669/203158047-af4f13e7-0ab8-494d-ad9a-6cad308bb3a3.png)

# Projeye ait görseller
![step1](https://user-images.githubusercontent.com/47785669/203446807-00ce5a2f-3237-4cad-96ac-02461488bb2f.png)
![step2](https://user-images.githubusercontent.com/47785669/203446810-5f0d58b9-7f14-488f-8ecf-13c9b7e877bc.png)
![step3](https://user-images.githubusercontent.com/47785669/203446834-f00a6d8a-872b-4801-b826-ccfbf50fad0f.png)


# Kurulum

1. `Docker Desktop`'ı yükleyin .
2. Docker Compose yapılandırmasını alın

```
git clone https://github.com/ibrahiimbayram/EstateIndex
```


3. Kurulum Ortamı

* Windows

```
cd EstateIndex
docker-compose up -d
```


* Linux & MacOS

```
cd EstateIndex
docker-compose up -d
```

4. Kibana'ya export.ndjson dosyamızı import edelim.

```
curl -X POST "http://localhost:5601/api/saved_objects/_import" -H "kbn-xsrf: true" --form file=@export.ndjson
```

# VS Code görünümü ve Tanımlar

![image](https://user-images.githubusercontent.com/47785669/203442193-5ff1edfe-d0c4-466a-bee4-43b70bfb3e82.png)

* Goservices = Selenium'u yönetir.Elasticsearch'e veri akışını sağlar.

* Selenium = Web sitesi üzerinden veri toplar.

* Elasticsearch = Veritabanı görevi sağlar.

* Kibana =  Verilerin analizini yapıp anlamlı hale getirerek görselleştirmeyi sağlar.

# Geliştirilmesi gereken noktalar

* Dinamik bir yapı kurarak taskların kontrolü yapılabilir.

* Eşzamanlı olarak çalışan Selenium containerların performans geliştirmeleri yapılabilir.  
