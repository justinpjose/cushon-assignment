# Postgres

# Why?
## SQL vs NoSQL?
Decision was made to use SQL for following:
- We are dealing with structured data with well-defined schemas. Data can be organised into tables with predefined columns and data types.
- Provides strong transaction support and enforces ACID (Atomicity, Consistency, Isolation, Durability) properties, ensuring data integrity.

## Which SQL database?
Decided to use Postgres for the following:
- Advance Features - PostgreSQL is known for its rich feature set, including support for advanced data types (e.g., JSON, arrays, hstore), custom functions, and custom indexing methods.
- Data Integrity - PostgreSQL is known for its strict adherence to data integrity and provides various constraints and validation mechanisms.
- Concurrency Control: PostgreSQL uses Multi-Version Concurrency Control (MVCC), which can provide better handling of concurrent transactions.

# Enchancments
- This directory could be more organised by having subdirectories where each links to a table in the database

[//]: # (Reference Links)
[choosing-db]:<https://yalantis.com/blog/how-to-choose-a-database/#:~:text=The%20best%20SQL%20databases%20are,read%20and%20edit%20the%20data.>