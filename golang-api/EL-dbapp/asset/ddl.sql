CREATE TABLE customers (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(225) NOT NULL,
    phone VARCHAR(225) NOT NULL,
    address VARCHAR(225) DEFAULT ''
);

CREATE TABLE employees (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(225) NOT NULL,
    phone VARCHAR(225) NOT NULL,
    address VARCHAR(225) DEFAULT ''
);

CREATE TABLE product (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(225) NOT NULL,
    price INTEGER NOT NULL,
    unit VARCHAR(225) DEFAULT ''
);

CREATE TABLE transaction (
    id  SERIAL PRIMARY KEY,
    bill_date TIMESTAMP NOT NULL,
    entry_date TIMESTAMP NOT NULL,
    finish_date TIMESTAMP NOT NULL,
    employee_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE transaction_detail (
    id  SERIAL PRIMARY KEY,
    bill_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    product_price INTEGER NOT NULL,
    qty INTEGER NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id),
    FOREIGN KEY (bill_id) REFERENCES transaction(id)
);


