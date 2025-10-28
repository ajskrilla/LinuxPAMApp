
# PAM Okta Helper

A lightweight Linux authentication agent that integrates with Okta and Active Directory for secure login, user profile management, and sudo access — without relying on `/etc/passwd` or `/etc/sudoers`.

---

## 🚀 Project Goals

- Authenticate Linux users against Okta (supporting MFA)
- Dynamically create temporary user profiles using a local SQLite cache
- Allow secure privilege elevation **without touching sudoers**
- Fully configurable, encrypted, and extensible
- Written in **Go** and **C** for performance and system integration

---

## 📁 Project Structure

```text
pam-okta-helper/
├── main.go                # Entrypoint — handles login flow
├── go.mod / go.sum        # Go modules
├── config/
│   └── config.go          # Environment/config loading
├── db/
│   ├── db.go              # SQLite DB bootstrap/init
│   └── schema.sql         # User table schema
├── model/
│   └── user.go            # User struct (mapped to DB)
├── oktaauth/
│   └── auth.go            # Okta login and MFA logic
├── util/
│   ├── logger.go          # Shared logger across modules
├── pam/
│   └── pam_module.c       # [WIP] PAM integration in C
└── README.md              # You're here!
```

---

## 🛠️ Setup

### Requirements

- Go 1.20+
- SQLite3
- A free Okta Developer account (or AD federated with Okta)

### Build

```bash
go mod tidy
go build -o pam-okta-helper
```

### Run

```bash
./pam-okta-helper
```

Follow prompts to enter your Okta credentials.

---

## 🔐 Security

- All user credentials are verified directly with Okta.
- Local cache (SQLite) uses encrypted values (TBD).
- Logger is configurable and does not leak secrets.
- MFA handled via Okta Verify and TOTP (in progress).

---

## 🧪 Testing

Want to test login without PAM?

```bash
go run main.go
```

To simulate PAM authentication, a `pam_module.c` file is provided and will be integrated later.

---

## 📌 Roadmap

- [x] Okta authentication with session token
- [x] Local DB cache for authenticated users
- [x] Project-wide logger injection
- [ ] MFA Challenge for TOTP and Okta Verify
- [ ] PAM integration
- [ ] NSS module to simulate `/etc/passwd` entries
- [ ] Privilege escalation without sudoers

---

## 🤝 Contributing

Pull requests and forks welcome! Please follow standard Go conventions.

```bash
git clone https://github.com/yourusername/pam-okta-helper.git
cd pam-okta-helper
```

Open issues or enhancements via GitHub issues.
