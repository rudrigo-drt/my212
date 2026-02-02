# My212

My212 is a personal web application designed to track, analyze, and visualize investment data from a Trading 212 account in a customized way.

The goal of this project is to provide deeper insights, historical analysis, and metrics that go beyond what is available in the native Trading 212 app.

---

## ✨ Features

TBD

---

## 🏗️ Architecture

My212 follows a simple and secure architecture:

Frontend (Web App) <br>
↓ <br>
Backend API <-> Database<br>
↓ <br>
T212 API

- The frontend never communicates directly with Trading 212 or the database
- All Trading 212 API calls are handled by the backend
- Investment data is persisted to enable historical analysis

---

## 🧰 Tech Stack

> Final stack decisions are still in progress.

Planned components:
- Backend: Go
- Frontend: TBD
- Database: TBD (PostgreSQL)

---

## 🚧 Project Status

🚀 Early development / architecture phase

This is a personal project built for learning and experimentation purposes.

---

## ⚠️ Disclaimer

This project is **not affiliated with Trading 212**.

My212 is a personal tool intended for private use.  
All investment data belongs to the account owner and is accessed via a personal API key.

---

## 📄 License

TBD
