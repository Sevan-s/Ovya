\set ON_ERROR_STOP on


-- Create and chek if user exists

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'ovya') THEN
        CREATE ROLE ovya LOGIN PASSWORD 'ovya' 
        NOSUPERUSER NOCREATEDB NOCREATEROLE INHERIT;
    END IF;
END $$;


-- create and check if ovya_recrutement database with ovya owner

DROP database IF EXISTS ovya_recrutement;
CREATE DATABASE ovya_recrutement OWNER ovya ENCODING 'utf8' TEMPLATE template0;


-- 
DROP TABLE IF EXISTS acq, ccial, dossier, visite;

\c ovya_recrutement ovya

CREATE EXTENSION IF NOT EXISTS pgcrypto;
SELECT * FROM pg_extension WHERE extname = 'pgcrypto';

CREATE TABLE acq (
    id SERIAL PRIMARY KEY,
    nom TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    PASSWORD VARCHAR(64)
) WITH (
  OIDS = FALSE
);

CREATE TABLE ccial (
    id SERIAL PRIMARY KEY,
    nom TEXT NOT NULL,
    email TEXT UNIQUE
) WITH (
  OIDS = FALSE
);

CREATE TABLE dossier (
    id SERIAL PRIMARY KEY,
    date_insert TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ccial_id INT,
    FOREIGN KEY (ccial_id) REFERENCES ccial(id)
) WITH (
  OIDS = FALSE
);

CREATE TABLE visite (
    id SERIAL PRIMARY KEY,
    date_start TIMESTAMP NOT NULL,
    date_end TIMESTAMP NOT NULL,
    acq_id INT NOT NULL,
    ccial_id INT NOT NULL,
    dossier_id INT NOT NULL,
    canceled BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (acq_id) REFERENCES acq(id),
    FOREIGN KEY (ccial_id) REFERENCES ccial(id),
    FOREIGN KEY (dossier_id) REFERENCES dossier(id)
) WITH (
  OIDS = FALSE
);

CREATE TABLE dossier_historique (
    id SERIAL PRIMARY KEY,
    dossier_id INT NOT NULL,
    ccial_id INT NOT NULL,
    date_start TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    date_end TIMESTAMP, 
    FOREIGN KEY (dossier_id) REFERENCES dossier(id) ON DELETE CASCADE,
    FOREIGN KEY (ccial_id) REFERENCES ccial(id) ON DELETE CASCADE
) WITH (
  OIDS = FALSE
);

INSERT INTO acq (nom, email, password) VALUES ('Donald Knuth', 'dknuth@fsf.org', crypt('leonidasguibas', gen_salt('bf')));

INSERT INTO acq (nom, email, password)
VALUES
  ('Acquéreur 1', 'acq1@example.com', 'password1'),
  ('Acquéreur 2', 'acq2@example.com', 'password2'),
  ('Acquéreur 3', 'acq3@example.com', 'password3'),
  ('Acquéreur 4', 'acq4@example.com', 'password4'),
  ('Acquéreur 5', 'acq5@example.com', 'password5'),
  ('Acquéreur 6', 'acq6@example.com', 'password6'),
  ('Acquéreur 7', 'acq7@example.com', 'password7'),
  ('Acquéreur 8', 'acq8@example.com', 'password8'),
  ('Acquéreur 9', 'acq9@example.com', 'password9'),
  ('Acquéreur 10', 'acq10@example.com', 'password10');

INSERT INTO ccial (nom, email)
VALUES
  ('Commercial 1', 'ccial1@example.com'),
  ('Commercial 2', 'ccial2@example.com'),
  ('Commercial 3', 'ccial3@example.com'),
  ('Commercial 4', 'ccial4@example.com'),
  ('Commercial 5', 'ccial5@example.com'),
  ('Commercial 6', 'ccial6@example.com'),
  ('Commercial 7', 'ccial7@example.com'),
  ('Commercial 8', 'ccial8@example.com'),
  ('Commercial 9', 'ccial9@example.com'),
  ('Commercial 10', 'ccial10@example.com');

INSERT INTO dossier (ccial_id)
VALUES
  (1),
  (2),
  (3),
  (4),
  (5),
  (6),
  (7),
  (8),
  (9),
  (10);
\d