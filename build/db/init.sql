CREATE TABLE companies
(
    id serial primary key,
    name text not null
);

CREATE TABLE account_types
(
    id serial primary key,
    name text not null
);

CREATE TABLE funds
(
    id serial primary key,
    account_type_id int not null,
    name text not null
);
CREATE INDEX idx_funds_account_type_id ON funds (account_type_id);

CREATE TABLE companies_funds
(
    company_id int not null,
    fund_id int not null,
    active boolean,
    PRIMARY KEY(company_id, fund_id)
);

CREATE TABLE customers
(
    id serial primary key,
    company_id int not null,
    name text not null
);
CREATE INDEX idx_customers_company_id ON customers (company_id);

CREATE TABLE customer_accounts
(
    customer_account_no serial primary key,
    customer_id int not null,
    account_type_id int not null
);
CREATE INDEX idx_customer_accounts_company_account_type_id ON customer_accounts (customer_id, account_type_id);

CREATE TABLE customer_accounts_funds
(
    id serial primary key,
    customer_account_no int not null,
    fund_id int not null,
    total_amount int not null
);
CREATE INDEX idx_customer_accounts_funds_customer_account_no_fund_id ON customer_accounts_funds (customer_account_no, fund_id);

CREATE TABLE transactions
(
    id serial primary key,
    customer_accounts_funds_id int not null,
    amount int not null
);
CREATE INDEX idx_transactions_customer_accounts_funds_id ON transactions (customer_accounts_funds_id);

INSERT INTO companies
    (name)
VALUES
    ('KFC'), 
    ('Heineken'), 
    ('Darwin'),
    ('Amadeus'),
    ('Retail (Direct)');

INSERT INTO account_types
    (name)
VALUES
    ('Cushon ISA'),
    ('Net Zero Pension'),
    ('Cushion LISA'),
    ('Cushion JISA'),
    ('Cushon GIA');

INSERT INTO funds
    (account_type_id, name)
VALUES
    (1, 'Cushon Equities Fund'), 
    (2, 'Cushon Pension Fund'), 
    (1, 'Third-Party Equities Fund');

INSERT INTO companies_funds
    (company_id, fund_id, active)
VALUES
    (5, 1, TRUE), 
    (5, 3, TRUE);

INSERT INTO customers
    (company_id, name)
VALUES
    (5, 'Justine Jose');

INSERT INTO customer_accounts
    (customer_id, account_type_id)
VALUES
    (1, 1);

INSERT INTO customer_accounts_funds
    (customer_account_no, fund_id, total_amount)
VALUES
    (1, 1, 0);