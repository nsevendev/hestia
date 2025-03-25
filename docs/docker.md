# Docker  

- le projet fonctionne avec traefik, il faut donc le lancer en premier.
- vous pouvez utiliser le CLI nseven ou les commandes docker basic.

## DockerFile

- 4 blocs
    - base: init une image golang
    - dev: utilise base pour lancer air pour le mode dev
    - build: genere le binaire pour la prod
    - prod: lance le binaire

## les fichiers compose

- compose.yaml: contient la base des conteneurs
- compose.override.yaml: permet de surcharger la config par des ajout de config pour traefik
- compose.prod.yaml: surcharge le compose basic avec une config axé sur la prod

## variables d'environnement

- dev  
    - host traefik: permet de définir le host pour voir le site en mode dev
    - port: défini le port utiliser par l'application et par traefik