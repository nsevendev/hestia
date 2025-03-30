# hestia

Site de labelfortaine

### Prérequis

- docker, docke compose
- le projet `https://github.com/nsevendev/infra-traefik` doit etre démarré sur votre machine
- CLI nseven (facultatif)
- copier le fichier `.env.dist` en `.env` et le compléter

## Installation

- avec le CLI nseven

```bash
# lancer le build de l'image
ns bi

# ensuite lancer les containers
ns up

# arrêter les containers
ns down
```

- sans le CLI nseven

```bash
# lancer le build de l'image
docker-compose build

# ensuite lancer les containers sans logs
docker-compose up -d

# arrêter les containers
docker-compose down
```