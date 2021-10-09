CREATE TABLE IF NOT EXISTS products (
                                           id SERIAL PRIMARY KEY,
                                           name VARCHAR(255) NOT NULL,
                                           vendor_code FLOAT NOT NULL,
                                           price FLOAT NOT NULL,
                                           stock_balance INT NOT NULL,
                                           description VARCHAR(2048) NOT NULL,
                                           specifications jsonb,
                                           category_id INT NOT NULL,
                                           type_id INT NOT NULL,
                                           created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at TIMESTAMPTZ,
                                           FOREIGN KEY (category_id) REFERENCES categories(id),
                                           FOREIGN KEY (type_id) REFERENCES types(id)


);