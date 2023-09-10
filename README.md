# Belajar-Golang-Dependency-Injection

Repo ini Lanjutan dari repo <a href="https://github.com/abrarnaim015/belajar-golang-rastful-api">Belajar-golang-restful-api</a>

Link Pembelajaran

- Dependency Injection <a href="https://www.youtube.com/watch?v=zwiDrYBGRtg&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=15&ab_channel=ProgrammerZamanNow">Programmer Zaman Now</a>

- Restfull API <a href="https://www.youtube.com/watch?v=bJ2ZFt9D0uI&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=14&ab_channel=ProgrammerZamanNow">Programmer Zaman Now</a>

Link Documentasi

- <a href="https://github.com/go-sql-driver/mysql">Driver MySQL</a>
- <a href="https://github.com/julienschmidt/httprouter">HTTP Router</a>
- <a href="https://github.com/go-playground/validator">Validation</a>
- <a href="https://github.com/stretchr/testify">Testify</a>

- <a href="https://github.com/google/wire">Google Wire</a>

> Setup

Driver MySQL

```golang
go get -u github.com/go-sql-driver/mysql
```

HTTP Router

```golang
go get github.com/julienschmidt/httprouter
```

Validator

```golang
go get github.com/go-playground/validator
```

Testing

```golang
go get github.com/stretchr/testify
```

Google Wire

```golang
go get github.com/google/wire
```

Istalation `Wire`

```golang
go install github.com/google/wire/cmd/wire@latest
```

Check `Wire` instalation Success

```golang
wire help
```

untuk menjalankan `google wire` dan meninject cukup masuk ke file injectornya lalu

```golang
wire
```

jika ingin dari luar dan filenya ada di dalam folder bisa menggunakan

```golang
wire gen {nama module}/{folder}
```

maka file `wire_gen.go` akan dibuat secara otomatis
