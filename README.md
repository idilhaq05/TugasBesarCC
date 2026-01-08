# Tugas Besar 

## Deskripsi
SemangatBelajar adalah aplikasi web berbasis Go (backend) dan React+Vite (frontend) yang menggunakan PostgreSQL sebagai database. Aplikasi ini dideploy menggunakan Docker Swarm dengan Nginx sebagai load balancer.

## Daftar Isi
- [Prasyarat](#prasyarat)
- [Konfigurasi IP Address](#konfigurasi-ip-address)
- [Cara Menjalankan Program](#cara-menjalankan-program)
  - [1. Menjalankan Secara Lokal (Tanpa Docker)](#1-menjalankan-secara-lokal-tanpa-docker)
  - [2. Menjalankan dengan Docker Compose](#2-menjalankan-dengan-docker-compose)
  - [3. Menjalankan dengan Docker Swarm](#3-menjalankan-dengan-docker-swarm)
- [Cara Menjalankan Stress Test](#cara-menjalankan-stress-test)
  - [Instalasi k6](#instalasi-k6)
  - [Konfigurasi Stress Test](#konfigurasi-stress-test)
  - [Menjalankan Stress Test](#menjalankan-stress-test)

---

## Prasyarat

### Windows
1. **Go** (versi 1.24.2 atau lebih baru)
   - Download dari: https://golang.org/dl/
   - Verifikasi instalasi: `go version`

2. **Node.js dan npm** (versi LTS terbaru)
   - Download dari: https://nodejs.org/
   - Verifikasi instalasi: `node --version` dan `npm --version`

3. **Docker Desktop**
   - Download dari: https://www.docker.com/products/docker-desktop/
   - Pastikan WSL 2 sudah terinstall untuk Docker Desktop
   - Verifikasi instalasi: `docker --version` dan `docker-compose --version`

4. **Git**
   - Download dari: https://git-scm.com/download/win
   - Verifikasi instalasi: `git --version`

5. **k6** (untuk stress testing)
   - Akan dijelaskan di bagian [Instalasi k6](#instalasi-k6)

### macOS
1. **Go** (versi 1.24.2 atau lebih baru)
   ```bash
   # Menggunakan Homebrew
   brew install go
   
   # Atau download dari: https://golang.org/dl/
   # Verifikasi instalasi
   go version
   ```

2. **Node.js dan npm** (versi LTS terbaru)
   ```bash
   # Menggunakan Homebrew
   brew install node
   
   # Atau download dari: https://nodejs.org/
   # Verifikasi instalasi
   node --version
   npm --version
   ```

3. **Docker Desktop**
   ```bash
   # Menggunakan Homebrew
   brew install --cask docker
   
   # Atau download dari: https://www.docker.com/products/docker-desktop/
   # Verifikasi instalasi
   docker --version
   docker-compose --version
   ```

4. **Git**
   ```bash
   # Git biasanya sudah terinstall di macOS
   # Jika belum, install dengan Homebrew
   brew install git
   
   # Verifikasi instalasi
   git --version
   ```

5. **k6** (untuk stress testing)
   - Akan dijelaskan di bagian [Instalasi k6](#instalasi-k6)

---

## Konfigurasi IP Address

**PENTING:** Sebelum menjalankan program, Anda perlu menyesuaikan IP address sesuai dengan perangkat Anda.

### 1. Mendapatkan IP Address Lokal Anda

#### Windows:
```cmd
ipconfig
```
Cari bagian "IPv4 Address" pada adapter jaringan yang aktif (biasanya Ethernet atau Wi-Fi).
Contoh: `192.168.1.100`

#### macOS:
```bash
ifconfig | grep "inet "
```
Atau lebih mudah:
```bash
ipconfig getifaddr en0  # untuk Wi-Fi
ipconfig getifaddr en1  # untuk Ethernet
```
Contoh: `192.168.1.100`

### 2. Update File Konfigurasi

Ganti IP address `192.168.231.128` dengan IP address lokal Anda di file-file berikut:

#### a. Frontend Environment (`.env.local`)
```bash
# File: frontend/.env.local
VITE_API_BASE_URL = http://[IP_ANDA]:8080
```

**Contoh:**
```bash
VITE_API_BASE_URL = http://192.168.1.100:8080
```

#### b. Stress Test Configuration (`stress_test.js`)
```javascript
// File: stress_test.js (baris 18)
const BASE_URL = 'http://[IP_ANDA]';
```

**Contoh:**
```javascript
const BASE_URL = 'http://192.168.1.100';
```

#### c. Docker Compose (jika menggunakan registry lokal)
```yaml
# File: docker-compose.yml
# Ganti IP registry jika Anda setup registry lokal sendiri
image: [IP_REGISTRY]:4000/backend
```

**Catatan:** Untuk development lokal tanpa Docker Swarm, Anda bisa menggunakan `localhost` atau `127.0.0.1` sebagai gantinya.

---

## Cara Menjalankan Program

### 1. Menjalankan Secara Lokal (Tanpa Docker)

#### A. Menjalankan Backend

##### Windows (Command Prompt atau PowerShell):
```cmd
cd backend

# Set environment variable (sementara untuk sesi ini)
set DATABASE_URL=your_database_url_here
set SMTP_USER=your_email@gmail.com
set SMTP_PASSWORD=your_smtp_password
set SMTP_HOST=smtp.gmail.com
set SMTP_PORT=465

# Install dependencies
go mod tidy

# Jalankan backend
go run main.go
```

**Alternatif:** Buat file `.env` di folder `backend` dengan isi:
```
DATABASE_URL=your_database_url_here
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_smtp_password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=465
```
Kemudian jalankan: `go run main.go`

**Catatan:** Ganti placeholder dengan credentials sebenarnya. Jika sudah ada file `.env` di repository, gunakan nilai yang ada di file tersebut.

##### macOS (Terminal):
```bash
cd backend

# Buat file .env atau export variabel
export DATABASE_URL="your_database_url_here"
export SMTP_USER="your_email@gmail.com"
export SMTP_PASSWORD="your_smtp_password"
export SMTP_HOST="smtp.gmail.com"
export SMTP_PORT="465"

# Install dependencies
go mod tidy

# Jalankan backend
go run main.go
```

**Catatan:** Ganti placeholder dengan credentials sebenarnya. Jika sudah ada file `.env` di repository, gunakan nilai yang ada di file tersebut.

Backend akan berjalan di: `http://localhost:8081`

#### B. Menjalankan Frontend

##### Windows (Command Prompt atau PowerShell):
```cmd
cd frontend

# Install dependencies (hanya pertama kali)
npm install

# Jalankan development server
npm run dev
```

##### macOS (Terminal):
```bash
cd frontend

# Install dependencies (hanya pertama kali)
npm install

# Jalankan development server
npm run dev
```

Frontend akan berjalan di: `http://localhost:5173` (atau port lain yang ditampilkan di terminal)

**Catatan:** Pastikan file `frontend/.env.local` sudah dikonfigurasi dengan benar agar frontend dapat berkomunikasi dengan backend.

---

### 2. Menjalankan dengan Docker Compose

Docker Compose cocok untuk development dan testing lokal.

#### Windows & macOS:
```bash
# Pastikan Docker Desktop sudah berjalan

# Build dan jalankan semua service
docker-compose up --build

# Atau jalankan di background
docker-compose up -d --build

# Melihat logs
docker-compose logs -f

# Menghentikan services
docker-compose down
```

**Akses aplikasi:**
- Frontend: `http://localhost` atau `http://[IP_ANDA]`
- Backend API: `http://localhost:8080` atau `http://[IP_ANDA]:8080`

**Troubleshooting:**
- Jika port sudah digunakan, ganti port di `docker-compose.yml`
- Jika ada masalah build, coba: `docker-compose down -v` kemudian build ulang

---

### 3. Menjalankan dengan Docker Swarm

Docker Swarm digunakan untuk production deployment dengan load balancing dan high availability.

#### A. Inisialisasi Docker Swarm

##### Windows & macOS:
```bash
# Initialize swarm mode
docker swarm init

# Atau jika ada multiple network interface, specify IP
docker swarm init --advertise-addr [IP_ANDA]
```

#### B. Setup Docker Registry (Opsional - untuk multi-node)

Jika Anda ingin deploy ke multiple nodes, setup registry lokal:

```bash
# Jalankan registry container
docker service create --name registry --publish published=4000,target=5000 registry:2

# Atau untuk single node, gunakan:
docker run -d -p 4000:5000 --restart=always --name registry registry:2
```

**Update `docker-compose.yml`** dengan IP registry Anda:
```yaml
image: [IP_ANDA]:4000/backend
image: [IP_ANDA]:4000/nginx_backend
image: [IP_ANDA]:4000/frontend
```

#### C. Build dan Push Images

##### Windows (PowerShell):
```powershell
# Build images
docker-compose build

# Tag dan push ke registry (jika menggunakan registry)
docker tag backend:latest [IP_ANDA]:4000/backend
docker push [IP_ANDA]:4000/backend

docker tag nginx_backend:latest [IP_ANDA]:4000/nginx_backend
docker push [IP_ANDA]:4000/nginx_backend

docker tag frontend:latest [IP_ANDA]:4000/frontend
docker push [IP_ANDA]:4000/frontend
```

##### macOS (Terminal):
```bash
# Build images
docker-compose build

# Tag dan push ke registry (jika menggunakan registry)
docker tag backend:latest [IP_ANDA]:4000/backend
docker push [IP_ANDA]:4000/backend

docker tag nginx_backend:latest [IP_ANDA]:4000/nginx_backend
docker push [IP_ANDA]:4000/nginx_backend

docker tag frontend:latest [IP_ANDA]:4000/frontend
docker push [IP_ANDA]:4000/frontend
```

#### D. Deploy Stack ke Swarm

##### Windows & macOS:
```bash
# Deploy stack dengan nama "semangatbelajar"
docker stack deploy -c docker-compose.yml semangatbelajar

# Cek status services
docker stack services semangatbelajar

# Cek logs service tertentu
docker service logs semangatbelajar_backend
docker service logs semangatbelajar_frontend
docker service logs semangatbelajar_nginx_backend

# Scale service (contoh: scale backend ke 4 replicas)
docker service scale semangatbelajar_backend=4

# Update service
docker service update semangatbelajar_backend

# Remove stack
docker stack rm semangatbelajar
```

#### E. Monitoring dan Management

```bash
# List semua services
docker service ls

# Detail service tertentu
docker service ps semangatbelajar_backend

# List nodes di swarm
docker node ls

# Inspect node
docker node inspect self

# Leave swarm (hati-hati!)
docker swarm leave --force
```

**Akses aplikasi:**
- Frontend: `http://localhost` atau `http://[IP_ANDA]`
- Backend API: `http://localhost:8080` atau `http://[IP_ANDA]:8080`

---

## Cara Menjalankan Stress Test

Stress test dilakukan menggunakan **k6**, sebuah tool modern untuk load testing.

### Instalasi k6

#### Windows:

**Metode 1: Menggunakan Chocolatey**
```cmd
choco install k6
```

**Metode 2: Menggunakan Windows Package Manager (winget)**
```cmd
winget install k6 --source winget
```

**Metode 3: Download Manual**
1. Download installer dari: https://github.com/grafana/k6/releases
2. Extract file zip
3. Tambahkan folder k6 ke PATH environment variable
4. Verifikasi: `k6 version`

#### macOS:

**Metode 1: Menggunakan Homebrew (Recommended)**
```bash
brew install k6
```

**Metode 2: Download Manual**
```bash
# Cek versi terbaru di: https://github.com/grafana/k6/releases
# Download binary (contoh untuk versi 0.47.0, Apple Silicon)
curl -L https://github.com/grafana/k6/releases/download/v0.47.0/k6-v0.47.0-macos-arm64.tar.gz -o k6.tar.gz

# Untuk Intel Mac, gunakan:
# curl -L https://github.com/grafana/k6/releases/download/v0.47.0/k6-v0.47.0-macos-amd64.tar.gz -o k6.tar.gz

# Extract
tar -xzf k6.tar.gz

# Move to /usr/local/bin
sudo mv k6-*/k6 /usr/local/bin/

# Verifikasi
k6 version
```

### Konfigurasi Stress Test

1. **Edit file `stress_test.js`** dan update IP address:

```javascript
// Baris 18: Ganti dengan IP address atau hostname Anda
const BASE_URL = 'http://[IP_ANDA]';

// Contoh:
const BASE_URL = 'http://192.168.1.100';
// atau untuk localhost:
const BASE_URL = 'http://localhost';
```

2. **Sesuaikan credentials** (jika diperlukan):

```javascript
// Baris 21-24: Pastikan credentials ini sesuai dengan database Anda
export function setup() {
  let res = http.post(`${BASE_URL}/api/login`, {
    email: 'admin@ecosteps.com',
    password: 'admin',
  });
  // ...
}
```

3. **Sesuaikan stages** (opsional - untuk mengatur intensitas test):

```javascript
// Baris 4-16: Konfigurasi load test stages
export let options = {
  stages: [
    { duration: '1m', target: 50 },      // Warmup: 50 users selama 1 menit
    { duration: '1m', target: 200 },     // Ramp up: 200 users
    { duration: '1m', target: 500 },     // 500 users
    { duration: '1m', target: 1000 },    // 1000 users
    { duration: '1m', target: 10000 },   // 10K users
    { duration: '1m', target: 50000 },   // 50K users
    { duration: '1m', target: 100000 },  // 100K users (peak load)
    { duration: '2m', target: 100000 },  // Sustain 100K users
    { duration: '2m', target: 0 },       // Cool down
  ],
};
```

**Catatan:** Test dengan 100K concurrent users sangat berat! Untuk testing awal, gunakan angka yang lebih kecil:

```javascript
export let options = {
  stages: [
    { duration: '30s', target: 10 },     // Warmup: 10 users
    { duration: '1m', target: 50 },      // 50 users
    { duration: '1m', target: 100 },     // 100 users
    { duration: '30s', target: 0 },      // Cool down
  ],
};
```

### Menjalankan Stress Test

#### Windows (Command Prompt atau PowerShell):

```cmd
# Pastikan aplikasi sudah berjalan terlebih dahulu

# Jalankan stress test
k6 run stress_test.js

# Jalankan dengan output ke file
k6 run --out json=test_results.json stress_test.js

# Jalankan dengan summary ke file
k6 run stress_test.js > test_summary.txt

# Jalankan dengan custom VUs (Virtual Users)
k6 run --vus 10 --duration 30s stress_test.js
```

#### macOS (Terminal):

```bash
# Pastikan aplikasi sudah berjalan terlebih dahulu

# Jalankan stress test
k6 run stress_test.js

# Jalankan dengan output ke file
k6 run --out json=test_results.json stress_test.js

# Jalankan dengan summary ke file
k6 run stress_test.js > test_summary.txt

# Jalankan dengan custom VUs (Virtual Users)
k6 run --vus 10 --duration 30s stress_test.js

# Jalankan dengan monitoring real-time (tambahan)
k6 run --out influxdb=http://localhost:8086/k6 stress_test.js
```

### Membaca Hasil Stress Test

Setelah stress test selesai, k6 akan menampilkan summary seperti:

```
     ✓ login ok
     ✓ dashboard ok
     ✓ fast < 1s

     checks.........................: 100.00% ✓ 15000      ✗ 0    
     data_received..................: 45 MB   150 kB/s
     data_sent......................: 5.0 MB  17 kB/s
     http_req_blocked...............: avg=1.2ms   min=0s   med=1ms    max=100ms
     http_req_connecting............: avg=800µs   min=0s   med=700µs  max=50ms
     http_req_duration..............: avg=450ms   min=200ms med=400ms  max=900ms
     http_req_receiving.............: avg=150µs   min=100µs med=150µs  max=500µs
     http_req_sending...............: avg=100µs   min=50µs  med=90µs   max=300µs
     http_req_waiting...............: avg=449ms   min=199ms med=399ms  max=899ms
     http_reqs......................: 5000    16.67/s
     iteration_duration.............: avg=2.45s   min=2.2s  med=2.4s   max=3s
     iterations.....................: 5000    16.67/s
     vus............................: 50      min=10       max=100
     vus_max........................: 100     min=100      max=100
```

**Interpretasi Hasil:**
- **checks**: Persentase assertion yang berhasil (harus 100%)
- **http_req_duration**: Response time rata-rata (semakin kecil semakin baik)
- **http_reqs**: Total requests dan throughput (requests/second)
- **vus**: Jumlah virtual users concurrent
- **✓/✗**: Jumlah checks yang pass/fail

### Troubleshooting Stress Test

#### 1. Connection Refused atau Timeout
```bash
# Pastikan aplikasi berjalan
docker ps  # cek container yang running
curl http://[IP_ANDA]/api/health  # test endpoint

# Pastikan IP address sudah benar di stress_test.js
```

#### 2. Login Failed (Status != 200)
```bash
# Pastikan credentials di stress_test.js sudah benar
# Pastikan user sudah terdaftar di database
# Test manual login:
curl -X POST http://[IP_ANDA]/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@ecosteps.com","password":"admin"}'
```

#### 3. Stress Test Terlalu Berat (Komputer Hang)
- Kurangi jumlah `target` di `stages`
- Kurangi `duration` di setiap stage
- Tutup aplikasi lain yang berjalan
- Gunakan komputer yang lebih powerful atau jalankan stress test dari komputer terpisah

#### 4. k6 Command Not Found
```bash
# Windows: Pastikan k6 ada di PATH
where k6

# macOS: Pastikan k6 terinstall
which k6

# Jika tidak ditemukan, install ulang sesuai instruksi di atas
```

---

## Ringkasan Perintah Cepat

### Development Lokal:
```bash
# Terminal 1 - Backend
cd backend
go run main.go

# Terminal 2 - Frontend
cd frontend
npm install
npm run dev
```

### Docker Compose:
```bash
docker-compose up --build
```

### Docker Swarm:
```bash
docker swarm init
docker stack deploy -c docker-compose.yml semangatbelajar
```

### Stress Test:
```bash
# Edit stress_test.js terlebih dahulu!
k6 run stress_test.js
```

---

## Catatan Penting

1. **Database Connection**: Aplikasi ini menggunakan Supabase PostgreSQL. Pastikan koneksi internet stabil.
   - Credentials database dan SMTP dapat ditemukan di file `.env` yang ada di repository
   - Untuk keamanan, credentials tidak ditampilkan di dokumentasi ini
   - Pastikan file `.env` tidak di-commit ke public repository

2. **SMTP Configuration**: Email verification menggunakan Gmail SMTP. Pastikan credentials valid.

3. **Port yang Digunakan**:
   - Backend: `8081` (lokal) atau `8080` (melalui nginx)
   - Frontend: `80` (Docker) atau `5173` (dev server)
   - Registry: `4000` (jika digunakan)

4. **IP Address**: Selalu update IP address di semua file konfigurasi jika berpindah network atau perangkat.

5. **Firewall**: Pastikan firewall tidak memblokir port yang digunakan aplikasi.

6. **Resource Requirements**: 
   - Minimal: 4GB RAM, 2 CPU cores
   - Recommended untuk stress test: 8GB RAM, 4+ CPU cores

---

## Troubleshooting Umum

### 1. "Cannot connect to database"
- Cek koneksi internet
- Verifikasi `DATABASE_URL` di `.env` file
- Test koneksi: `curl https://aws-0-ap-southeast-1.pooler.supabase.com`

### 2. "Port already in use"
```bash
# Windows - Cek port yang digunakan
netstat -ano | findstr :8080

# macOS - Cek port yang digunakan
lsof -i :8080

# Kill process yang menggunakan port
# Windows: taskkill /PID <PID> /F
# macOS: kill -9 <PID>
```

### 3. "docker: command not found"
- Pastikan Docker Desktop sudah terinstall dan berjalan
- Restart Docker Desktop
- Verifikasi: `docker --version`

### 4. Frontend tidak bisa connect ke Backend
- Cek `VITE_API_BASE_URL` di `frontend/.env.local`
- Pastikan backend sudah running
- Cek CORS settings di backend
- Test endpoint: `curl http://[IP_ANDA]:8080/api/health`

### 5. Docker Swarm services tidak bisa berkomunikasi
- Pastikan overlay networks sudah dibuat: `docker network ls`
- Cek service logs: `docker service logs [service_name]`
- Verifikasi service sudah running: `docker service ls`

---

## Lisensi & Kontribusi

Project ini dibuat untuk Tugas Besar Cloud Computing. Untuk kontribusi atau pertanyaan, silakan hubungi tim development.

---
