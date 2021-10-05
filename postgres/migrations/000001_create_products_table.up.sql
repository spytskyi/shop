CREATE TABLE IF NOT EXISTS products (
                                           id SERIAL PRIMARY KEY,
                                           vendor_code VARCHAR(255) NOT NULL,
                                           name VARCHAR(255) NOT NULL,
                                           category_id INT NOT NULL,
                                           type_id INT NOT NULL,
                                           description VARCHAR(255) NOT NULL,
                                           stock_balance INTEGER NOT NULL,
                                           specifications jsonb,
                                           created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
                                           updated_at TIMESTAMPTZ,
                                           FOREIGN KEY (category_id) REFERENCES categories(id),
                                           FOREIGN KEY (type_id) REFERENCES types(id)


);