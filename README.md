<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/4726/4726010.png" alt="CSV Importer Logo" width="140" height="140" style="margin-bottom: 10px; animation: float 3s ease-in-out infinite; filter: drop-shadow(0 0 10px rgba(0, 153, 255, 0.6));" />
</p>

<style>
@keyframes float {
  0% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
  100% { transform: translateY(0); }
}
</style>

# 🇹🇭 CSV Importer (Golang) — ระบบอัปโหลดและนำเข้าข้อมูล CSV ระดับ 100,000+ แถว

ระบบนี้เป็นชุดโครงสร้างพร้อมใช้งานสำหรับ  
**อัปโหลดไฟล์ CSV / Excel → ประมวลผล → import เข้า MySQL**  
รองรับไฟล์ขนาดใหญ่ ประมวลผลเร็ว และแยกโครงสร้างแบบ Clean Architecture

---

## 🌟 คุณสมบัติเด่น

### ⚡ ประมวลผลไฟล์ CSV ได้มากกว่า 100,000+ แถว

รองรับ Batch insert, Streaming อ่านไฟล์ทีละส่วน ไม่กิน RAM

### 📤 API `/upload`

รองรับ multipart/form-data พร้อมอัปโหลดไฟล์ CSV ได้จากเว็บเบราว์เซอร์หรือ cURL

### 🧱 โครงสร้างแบบแยกเลเยอร์

แยก config / db / handler / service / model อย่างชัดเจน

### 🖥️ HTML UI สำหรับอัปโหลดไฟล์

มาพร้อมหน้าเว็บ upload + preview + ปุ่มอัปโหลดสวยงาม

### 🧪 Mock Data Generator

มี Python script สร้าง mock CSV จำนวน 100,000 แถวได้ในไม่กี่วินาที

---

## 📁 โครงสร้างโปรเจกต์ (Project Structure)

```
csv-importer/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── mysql.go
│   ├── handler/
│   │   └── upload_handler.go
│   ├── service/
│   │   └── csv_service.go
│   └── model/
│       └── task.go
│
├── uploads/
├── upload.html
├── mock.py
├── go.mod
└── go.sum
```

---

## 🛠 การติดตั้ง (Setup)

### 1️⃣ ติดตั้ง Go module

```bash
go mod init csv-importer
go mod tidy
```

### 2️⃣ ติดตั้ง MySQL Driver

```bash
go get github.com/go-sql-driver/mysql
```

---

## 🗄 ตั้งค่า Database (MySQL / MariaDB)

เปิด XAMPP → กด **Shell** แล้วพิมพ์:

```bash
mysql -u root
```

สร้าง Database:

```sql
CREATE DATABASE pando CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE pando;
```

สร้าง Table:

```sql
CREATE TABLE barcode_tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    job_id VARCHAR(100) NOT NULL,
    working_station VARCHAR(255) NOT NULL,
    operation_id VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🚀 การรัน Server

### แบบ Air (แนะนำสำหรับ dev)

```bash
air
```

### แบบ Go run ปกติ

```bash
go run cmd/server/main.go
```

ผลลัพธ์ที่ถูกต้อง:

```
Connected to MySQL
Server running on port 8080
```

---

## 🌐 HTML UI สำหรับอัปโหลดไฟล์

ไฟล์ `upload.html` สามารถเปิดผ่าน Browser ได้เลย

ฟีเจอร์:

- Preview ชื่อไฟล์ + ขนาด
- Drag & Drop
- ปุ่มอัปโหลดไป `http://localhost:8080/upload`

---

## 🧪 ทดสอบด้วย cURL

```bash
curl -X POST "http://localhost:8080/upload"   -F "file=@mock_100k.csv"
```

ผลลัพธ์ตัวอย่าง:

```json
{
  "inserted": 100000
}
```

---

## 🧬 สร้าง Mock CSV 100,000 แถว

ไฟล์: `mock.py`

```python
import csv
import random

FILENAME = "mock_100k.csv"
ROWS = 100000

stations = ["Station A", "Station B", "Station C", "Station D", "Station E"]
operations = ["OP-001", "OP-002", "OP-003", "OP-004", "OP-005"]

with open(FILENAME, "w", newline="", encoding="utf-8") as f:
    writer = csv.writer(f)
    writer.writerow(["JobID", "WorkingStation", "OperationID"])

    for i in range(1, ROWS + 1):
        job_id = 10000 + i
        ws = random.choice(stations)
        op = random.choice(operations)
        writer.writerow([job_id, ws, op])

print("Created:", FILENAME)
```

รัน:

```bash
python mock.py
```

---

## 🔄 Workflow การใช้งานทั้งหมด (ตั้งแต่เริ่ม → สำเร็จ)

1. สร้างโฟลเดอร์โปรเจกต์
2. สร้างโครงสร้างไฟล์ตามที่ให้
3. ติดตั้ง go.mod และ dependencies
4. สร้าง MYSQL DB และ TABLE
5. รัน server (air หรือ go run)
6. เปิด `upload.html`
7. เลือก CSV → Preview → อัปโหลด
8. API `/upload` ประมวลผลและ import ข้อมูล
9. ตรวจใน Database:

```sql
SELECT COUNT(*) FROM barcode_tasks;
```

---

## 🧯 การแก้ปัญหา (Troubleshooting)

### ❌ Unknown database 'pando'

> แก้:

```sql
CREATE DATABASE pando;
```

### ❌ Table doesn't exist

> แก้:  
> สร้าง table `barcode_tasks`

### ❌ 500 Internal Server Error

- มักเกิดจาก Upload ไฟล์ว่าง
- ใช้คำสั่ง curl ที่ถูกต้อง:

```
curl -F "file=@mock_100k.csv" http://localhost:8080/upload
```

---

## 📜 License

MIT © 2025 — [Chinawat](https://github.com/ChinawatDc)

---

### ❤️ ขอให้ระบบนายลื่นปรื๊ดไม่มีสะดุด!

# csv-importer
