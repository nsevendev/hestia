# 📝 Logger - Système de Logs pour Hestia

Ce document explique le fonctionnement du système de logging mis en place dans le projet **Hestia** (Go).

---

## 📦 Structure

Le logger est défini dans le fichier :
```bash
/internal/logger/logger.go
```

---

## 🎯 Objectifs du logger

- Enregistrer **tous les logs** de l'application dans un fichier.
- Supporter la **rotation journalière** (1 fichier de log par jour).
- Différencier les logs selon l'environnement (`dev`, `prod`, etc.).

---

## 🏗 Arborescence des fichiers de logs

```text
runtime/
└── logs/
    └── dev/
        └── hestia-2025-03-23.log
```

Les logs sont rangés par environnement et par jour automatiquement.

---

## ⚙️ Initialisation

Le logger est initialisé via :
```go
logger.InitFromEnv()
```

Cette fonction :
- Lit la variable `APP_ENV` (par défaut à `dev`)
- Crée le dossier de logs si besoin
- Crée (ou ouvre en append) le fichier `hestia-YYYY-MM-DD.log`
- Loggue à la fois dans le fichier **et dans la console**

> L'appel se fait dans `internal/init/env.go` (automatiquement via `import _`)

---

## 🔧 Fonctions disponibles

```go
logger.Info("Message")
logger.Warn("Attention")
logger.Error("Erreur")
logger.Fatal("Message d'erreur et exit")
logger.Fatalf("Erreur : %v", err)
// et d'autre visiter le fichier internal/logger/logger.go
```

### 🎨 Affichage console
- [✔] Timestamp + filename automatique

---

## 🧪 Exemple de log dans le fichier

```
2025/03/23 09:01:24 logger.go:34: ℹ️ [INFO] ✅ Connexion à la base de données postgres réussie
```

---

## 🚀 En Production

En mode binaire (build Go), le fichier de log est généré à côté du binaire dans `runtime/logs/prod/` automatiquement.

Pas besoin de pré-créer les dossiers : ils sont gérés dynamiquement.

---

## 📌 Remarques

- Le logger utilise `log.SetOutput(io.MultiWriter(...))` pour écrire en même temps dans le fichier **et** la console.

---

## ✅ Bonus

Tu peux utiliser les fonctions `logger.Info`, `logger.Error`, etc. partout dans le projet sans te soucier de l'environnement ou du formatage.

