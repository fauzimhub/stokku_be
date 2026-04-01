
TRUNCATE TABLE categories RESTART IDENTITY CASCADE;

INSERT INTO categories (name, description)
VALUES
    ('Ice Cream', 'Es Krim'),
    ('Milk', 'Susu kotak atau botol'),
    ('Candy', 'Permen atau coklat'),
    ('Snack', 'Makanan kecil seperti keripik atau semacamnya');
