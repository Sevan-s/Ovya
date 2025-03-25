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

Le mot de passe pour le compte ovya est `ovya`.

#### 2. Lancer le backend

Dans le répertoire du backend, exécutez :

`cd ovya_backend `
`go mod tidy `
`go run main.go`

#### 3. Lancer le frontend

Dans le répertoire du frontend, assurez-vous d'avoir installé les dépendances et lancez le serveur Angular :

`cd ovya_frontend `
`npm install`
`ng serve`

## 4. Tester l'application via Postman

Avant de tester l'application, vous devez d'abord créer des **acquéreurs**, des **commerciaux**, et des **dossiers** dans la base de données. Vous pouvez faire cela facilement en utilisant Postman.

### Étapes pour tester l'application :

## 4. Workspace Postman

1. **Importez la collection Postman** :
Vous pouvez accéder au workspace Postman pour ce projet en cliquant sur le lien ci-dessous :

[Accéder au Workspace Postman ici](https://app.getpostman.com/join-team?invite_code=f79566e43f2fd8ab619f5452eb7141af8897c1de713d415b9a829009611c170b&target_code=6f82812c6d576f10e4ac9f1a885b4efc)
   - Cela vous permettra d'exécuter les requêtes API pour interagir avec l'application.

2. **Créez des acquéreurs** :
   - Utilisez les requêtes dans Postman pour créer des acquéreurs dans la base de données. Vous pouvez trouver la requête correspondante dans la collection acq sous **POST Create acq**.

3. **Créez des commerciaux** :
   - Créez des commerciaux via les requêtes correspondantes sous **POST Create ccial** dans Postman. Cela permettra de tester les fonctionnalités liées à la gestion des commerciaux.

4. **Créez des dossiers** :
   - Créez des dossiers en utilisant les requêtes sous **POST Create Folder** dans Postman. Ces dossiers seront utilisés pour tester les fonctionnalités de gestion des dossiers dans l'application.