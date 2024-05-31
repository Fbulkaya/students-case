# Student Scheduler API

Bu proje, öğrencilerin ders programlarını oluşturmasına yardımcı olmak için bir API sunar.

## Gereksinimler

- Go 1.20 ve üzeri
- Sqlite3

## Kurulum

1. Projeyi klonlayın

```bash
git clone git@github.com:Fbulkaya/Student_scheduler.git
```

2. Projeyi çalıştırın

```bash
cd Student_scheduler
go run .
```

## API

### Ogrenci

#### Ogrenci Ekleme

```bash
curl -X POST http://localhost:8080/students -d '{"name": "John Doe", "email": "jhon@example.com"}'
```

#### Ogrenci Guncelleme

```bash
curl -X PUT http://localhost:8080/students/1 -d '{"name": "John Doe", "email": "jhon@example.com"}'
```

#### Ogrenci Silme

```bash
curl -X DELETE http://localhost:8080/students/1
```

#### Ogrenci Listeleme

```bash
curl http://localhost:8080/students
```

#### Ogrenci Detay

```bash
curl http://localhost:8080/students/1
```

### Plan

#### Plan Ekleme

```bash
curl -X POST http://localhost:8080/students/1/plans -d '{"desc": "Math", "start_date": "2021-01-01 10:00", "end_date": "2021-01-01 12:00"}'
```

#### Plan Guncelleme

```bash
curl -X PUT http://localhost:8080/students/1/plans/1 -d '{"desc": "Math", "start_date": "2021-01-01 10:00", "end_date": "2021-01-01 12:00"}'
```

#### Plan Silme

```bash
curl -X DELETE http://localhost:8080/students/1/plans/1
```

#### Plan Listeleme

```bash
curl http://localhost:8080/students/1/plans
```

#### Plan Detay

```bash
curl http://localhost:8080/students/1/plans/1
```
