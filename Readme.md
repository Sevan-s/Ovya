# Ovya

Ovya est une application fullstack qui permet de gérer les **visites immobilières**, les **dossiers clients**, les **commerciaux** et les **acquéreurs**, avec un suivi historique des affectations des dossiers.

## Technologies utilisées

- **Frontend** : Angular
- **Backend** : Go
- **Base de données** : PostgreSQL
- **Orchestration** : Docker & Docker Compose

## Prérequis

Vous pouvez lancer ce projet de deux manières :

### Option 1: Avec Docker & Docker Compose

- Assurez-vous d'avoir installé Docker et Docker Compose sur votre machine.

### Option 2: Sans Docker

- Go (version ≥ 1.21)
- Node.js (version ≥ 18)
- PostgreSQL (version ≥ 15)

## Lancer le projet

### Avec Docker

Dans la racine du projet, exécutez la commande suivante :

`docker-compose up --build`

Cette commande va construire et lancer tous les conteneurs nécessaires (backend, frontend et base de données).

### Sans Docker

#### 1. Démarrer PostgreSQL

- Connectez-vous à PostgreSQL :

`psql -U postgres`

- Exécutez ensuite le script SQL pour créer la base et les tables :

`\i ovya_db/create_table.sql`

La base de données s'appelle `ovya_recrutement`.

#### 2. Lancer le backend

Dans le répertoire du backend, exécutez :

`cd ovya_backend `
`go mod tidy go `
`run main.go`

#### 3. Lancer le frontend

Dans le répertoire du frontend, assurez-vous d'avoir installé les dépendances et lancez le serveur Angular :

`cd ovya_frontend `
`npm install`
`ng serve`
