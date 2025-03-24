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