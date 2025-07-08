# Contact Management API

API ini adalah backend sederhana untuk mengelola **Users**, **Contacts**, dan **Addresses**. Dibuat dengan **Golang**, menggunakan **Gin Framework**, **GORM** sebagai ORM, dan **PostgreSQL** sebagai database.

---

## 🚀 Tech Stack

- **Golang** (Go)
- **Gin** (HTTP Framework)
- **GORM** (ORM)
- **PostgreSQL** (Database)
- **JWT** (Authentication)

---

## 📂 Project Structure

```
├── config/          # Konfigurasi DB dan environment
├── controllers/     # Handler HTTP (tempat logic-level request/response)
├── middlewares/     # Authentication, logging, dll
├── models/          # Struct GORM untuk representasi table DB
├── repositories/    # Layer query ke database
├── request/         # Struct untuk request binding/validation
├── response/        # Struct untuk shaping JSON Response
├── routes/          # Register semua routing API
├── services/        # Business logic (di luar layer controller dan repo)
├── utils/           # Helper/helper umum (contoh JWT utils, password hashing)
├── main.go          # Entry point app
```

---

## 📦 Installation

Ikuti langkah-langkah berikut untuk menjalankan project ini di lokal:

### 1. Clone Repository

```bash
git clone https://github.com/tiedsandi/project_contact-management-go.git
cd project_contact-management-go
```

### 2. Setup Environment

Buat file `.env` di root project:

```ini
APP_ENV=development
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=123
DB_NAME=contact_db
DB_PORT=5432

KEY_JWT=rahasia-super-aman
```

> 💡 Pastikan PostgreSQL sudah jalan, dan database `contact_db` sudah dibuat.

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run Migration (Opsional)

Jika kamu menggunakan AutoMigrate di dalam kode (misalnya di `main.go`), maka akan otomatis membuat tabel saat aplikasi dijalankan. Tapi kalau pakai migration tool (seperti `golang-migrate`), pastikan migration sudah dijalankan terlebih dahulu.

### 5. Run Server

```bash
go run main.go
```

> 🚀 Server akan jalan di default port `:8080` jika tidak dikonfigurasi lain.

---

## 📌 API Endpoint List

### 🧑‍💻 Users

| Method | Endpoint           | Description                         |
| ------ | ------------------ | ----------------------------------- |
| POST   | /api/users/signup  | Register user baru                  |
| POST   | /api/users/login   | Login user (get JWT Token)          |
| GET    | /api/users/profile | Get user profile (Auth required)    |
| PUT    | /api/users/profile | Update profile user (Auth required) |

### 📞 Contacts

| Method | Endpoint           | Description                                                    |
| ------ | ------------------ | -------------------------------------------------------------- |
| POST   | /api/contacts      | Create new contact                                             |
| PUT    | /api/contacts/\:id | Update contact                                                 |
| GET    | /api/contacts/\:id | Get single contact                                             |
| GET    | /api/contacts      | Search contacts (Pagination & Filtering by name, email, phone) |
| DELETE | /api/contacts/\:id | Delete contact                                                 |

### 🏠 Addresses

| Method | Endpoint                                       | Description                      |
| ------ | ---------------------------------------------- | -------------------------------- |
| POST   | /api/contact/\:contactId/addresses             | Create new address for a contact |
| PUT    | /api/contact/\:contactId/addresses/\:addressId | Update address                   |
| GET    | /api/contact/\:contactId/addresses/\:addressId | Get single address               |
| GET    | /api/contact/\:contactId/addresses             | List addresses for contact       |
| DELETE | /api/contact/\:contactId/addresses/\:addressId | Delete address                   |

---

## 📄 Example Response Structure

### ✅ User Login Response

```json
{
  "data": {
    "token": "<JWT_TOKEN>"
  }
}
```

### ✅ Get Profile Response

```json
{
  "data": {
    "username": "testing",
    "name": "Sandi"
  }
}
```

### ✅ Get Contact Response

```json
{
  "data": {
    "id": 1,
    "first_name": "Fachran",
    "last_name": "Sandi",
    "email": "test@test.com",
    "phone": "08123456789"
  }
}
```

### ✅ Get Address Response

```json
{
  "data": {
    "id": 1,
    "street": "Jl. Kenangan No.12",
    "city": "Bandung",
    "province": "Jawa Barat",
    "country": "Indonesia",
    "postal_code": "40291"
  }
}
```

### ✅ Pagination Response Example (Search Contacts)

```json
{
  "data": [
    {
      "id": 1,
      "first_name": "Fachran",
      "last_name": "Sandi",
      "email": "test@test.com",
      "phone": "08123456789"
    }
  ],
  "paging": {
    "page": 1,
    "total_page": 3,
    "total_item": 25
  }
}
```

---

## 🛠️ Authentication

Semua endpoint (kecuali Signup dan Login) membutuhkan **JWT Bearer Token** di header:

```http
Authorization: Bearer <your_token_here>
```

---

## ✅ Notes

- **Pagination** hanya berlaku untuk endpoint **GET /api/contacts**
- Data user dari JWT, jadi setiap user hanya bisa akses data miliknya
- Foreign key relation sudah diatur antara **User → Contact → Address**
