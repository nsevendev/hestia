# Structure du projet Hestia

Ce document explique rapidement l'organisation du projet **Hestia**, une application web en Go avec Gin, architecture MVC et gestion du build via Docker + Air.

---

## ⚖️ Racine du projet

| Dossier/Fichier     | Rôle                                                                 |
|----------------------|----------------------------------------------------------------------|
| `main.go`            | Point d'entrée de l'application                                      |
| `go.mod` / `go.sum`  | Modules Go                                                           |
| `Dockerfile`         | Image multi-stage (dev, build, prod)                                |
| `compose.yaml`       | Docker Compose (Go + PostgreSQL)                                    |
| `.air.toml`          | Config du hot reload (Air)                                          |
| `.gitignore`         | Fichiers/dossiers à ignorer                                          |

---

## 📂 `/app/`

Contient la "partie présentation" de l'application (MVC)

| Dossier       | Contenu                                                  |
|----------------|-----------------------------------------------------------|
| `assets/`      | Fichiers statiques : CSS, JS, images                     |
| `views/`       | Templates HTML (rendus avec `c.HTML(...)`)              |
| `controllers/` | Handlers Gin (fonctions liées aux routes)               |
| `router/`      | Configuration des routes Gin                             |

---

## 📚 `/internal/`

Contient la logique métier de l'application, privée au projet (non importable depuis l'extérieur).

Exemple possible :
```
/internal/
  database/
  logger/
  domain/
    user/
      service.go
```

---

## 📂 `/runtime/`

Tout ce qui est temporaire, automatique ou généré par le build.

| Dossier         | Rôle                                                                  |
|------------------|----------------------------------------------------------------------|
| `go-mod/`        | Cache Go pour les dépendances (via `GOMODCACHE`)                     |
| `air/`           | Binaire temporaire buildé par `air` pour le hot-reload               |
| `log/dev`        | Tous log ecris avec logger en mode APP_ENV=dev                       |
| `log/prod`        | Tous log ecris avec logger en mode APP_ENV=prod                     |

> Ce dossier est ignoré dans `.gitignore`

---

## 📂 `/runtime/`

Tout ce qui est temporaire, automatique ou généré par le build.

| Dossier         | Rôle                                                                                              |
|------------------|--------------------------------------------------------------------------------------------------|
| `env.go`        | initialise les variable d'environement dans go                                                    |
| `init.go`        | fichier principal d'initialisation (là où tous les autres sont implementer pour etre initier)    |

---

## 📄 `/docs/`

Documentation technique, guides, notes… Exemple :
- `structure.md` : ce document

---

## ✅ Récapitulatif visuel

```
/app
  assets/
  controllers/
  router/
  views/
/internal
  .gitkeep
/runtime
  go-mod/
  air/
/docs
  structure.md
main.go
Dockerfile
compose.yaml
.air.toml
.gitignore
```

---

</br>
</br>