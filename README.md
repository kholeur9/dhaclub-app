# MIRANDE

MIRANDE est un projet backend développé en Go dont l'objectif est de construire une application de gestion de documents, notes et tâches tout en développant une compréhension profonde des fondamentaux du backend moderne.

Le nom **MIRANDE** signifie "documents" dans ma langue maternelle.

---

# Vision du projet

MIRANDE est construit avec une approche progressive :

Comprendre les fondations avant d'ajouter de la complexité.

Le projet privilégie :

- la compréhension de Go ;
- la maîtrise de l'architecture backend ;
- la conception d'API propres ;
- la gestion des données ;
- la sécurité ;
- la scalabilité.

L'objectif n'est pas seulement de faire fonctionner une application, mais de comprendre chaque décision technique derrière celle-ci.

---

# Philosophie technique

MIRANDE suit volontairement une approche simple :

- Go natif
- API REST
- SQL brut
- PostgreSQL
- Architecture modulaire
- Pas d'ORM pour garder le contrôle sur les requêtes SQL

Cette approche permet de comprendre les mécanismes internes avant d'utiliser des abstractions plus complexes.

---

# Stack technique

## Backend

- Go
- net/http
- PostgreSQL
- SQL brut
- UUID
- JWT

## Architecture

Pattern utilisé :

```
Handler
   |
   v
Service
   |
   v
Store
   |
   v
Database
```

---

# Architecture du projet

## Handler

Responsabilités :

- réception des requêtes HTTP ;
- récupération des paramètres ;
- validation basique ;
- appel des services ;
- formatage des réponses.

---

## Service

Responsabilités :

- logique métier ;
- règles de validation ;
- gestion des conflits ;
- traitement des erreurs ;
- orchestration des opérations.

---

## Store

Responsabilités :

- communication avec PostgreSQL ;
- requêtes SQL ;
- persistance des données.

---

# Fonctionnalités actuelles

## Gestion des utilisateurs

Implémenté :

- création utilisateur ;
- validation des données ;
- vérification des conflits ;
- mise à jour des informations utilisateur ;
- gestion centralisée des erreurs.

Champs actuellement gérés :

- Email
- Username
- Password (hashé)

---

# Authentification

Implémenté :

- génération JWT ;
- validation JWT ;
- middleware d'authentification ;
- protection des routes ;
- récupération de l'identité utilisateur depuis le contexte.

Flux :

```
Client
  |
  v
Login
  |
  v
JWT
  |
  v
Middleware
  |
  v
User Context
```

---

# Gestion des erreurs

MIRANDE utilise un système d'erreurs séparé :

## Erreurs internes

Utilisées pour la logique backend :

```go
ErrUserNotFound
ErrManyFieldsToUpdate
ErrNoFieldsToUpdate
```

## Messages API

Utilisés pour communiquer avec le client :

```text
User not found
No data to update
You can only update one field at a time
```

Cette séparation permet de garder un backend propre et facilement évolutif.

---

# Gestion des tâches (Todo)

Modèle actuel :

```
User
 |
 └── Todo
```

Un utilisateur possède plusieurs tâches.

Fonctionnalités :

- création ;
- lecture ;
- modification ;
- suppression.

---

# Modèle de données actuel

## User

```
User

ID
Email
Username
Password
CreatedAt
UpdatedAt
```

---

## Todo

```
Todo

ID
Description
UserID
CreatedAt
UpdatedAt
```

---

# Roadmap V1

Objectif : construire une base backend solide.

## Étape 1

Finalisation complète du module Todo :

- validations ;
- cas limites ;
- gestion des erreurs.

## Étape 2

Pagination :

- limitation des résultats ;
- navigation entre pages.

## Étape 3

Filtres et tri :

- recherche ;
- organisation des données.

## Étape 4

Tests :

- tests unitaires ;
- tests d'intégration.

## Étape 5

Authentification avancée :

- sessions ;
- refresh tokens ;
- gestion des connexions.

---

# Évolutions futures

MIRANDE pourra évoluer vers un système plus riche :

```
User

 |
 v

Notes

 |
 v

Todos
```

Une note pourra contenir plusieurs tâches.

Une collaboration multi-utilisateurs pourra être envisagée plus tard.

---

# Objectifs d'apprentissage

À travers MIRANDE, les concepts étudiés sont :

- langage Go ;
- structs ;
- interfaces ;
- pointeurs ;
- packages ;
- gestion mémoire ;
- erreurs personnalisées ;
- architecture backend ;
- sécurité ;
- bases de données ;
- conception d'API.

---

# État du projet

🚧 En développement actif

MIRANDE évolue étape par étape avec une priorité donnée à la compréhension, la qualité du code et la construction d'une architecture backend solide.
