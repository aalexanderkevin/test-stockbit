-- Simple Database querying

-- Terdapat sebuah table "USER" yg memiliki 3 kolom: ID, UserName, Parent. Di mana:
-- Kolom ID adalah Primary Key
-- Kolom UserName adalah Nama User
-- Kolom Parent adalah ID dari User yang menjadi Creator untuk User tertentu.
-- eg.
-- ——————————————————————————
-- | ID | UserName | Parent |
-- ——————————————————————————
-- | 1 | Ali | 2 |
-- | 2 | Budi | 0 |
-- | 3 | Cecep | 1 |
-- —————————————————————————-
-- Tuliskan SQL Query untuk mendapatkan data berisi:
-- ——————————————————————————————————
-- | ID | UserName | ParentUserName |
-- ——————————————————————————————————
-- | 1 | Ali | Budi |
-- | 2 | Budi | NULL |
-- | 3 | Cecep | Ali |
-- ——————————————————————————————————
-- *Kolom ParentUserName adalah UserName berdasarkan value Parent

SELECT u1.ID, u1.Username, u2.Username AS ParentUserName
FROM USER u1
LEFT JOIN USER u2 ON u2.ID = u1.Parent
ORDER BY u1.ID;