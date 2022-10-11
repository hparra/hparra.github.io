sql -- language for structured databases queries
================================================

## Querying

`SELECT`
- is used to select which columns to include
- can use `AS` operator with each column to alias its name
- uses the `ALL` keyword by default

`FROM`
- identifies tables that contain desired columns
- can use `JOIN`, `INNER JOIN`, etc.
- can use `AS` operator with each table to alias its name
- can use a subquery to generate a temporary table

`WHERE`
- filters rows meeting specified conditions

`GROUP BY`
- groups (combines) rows based on a common value

`ORDER BY`
- sorts rows by column(s)
- should also use `ASC` or `DESC`
- can use an expression function

## Altering

- `TABLE`
  - `RENAME`
  - `ALTER`
    - `COLUMN`
      - `ADD`
      - `RENAME`
      - `DROP`
      - `ALTER`
        - `SET DEFAULT`
        - `SET NULL`

## Updating

`UPDATE`

`SET`

`FROM`

`WHERE`

`RETURNING` (Postgres-only)

## TOOLS

### psql

- `\list`: list all databases
- `\c` connect to a database
- `\dt` list tables in current database
- `\d+` describe table

### Other

- [DB Browser for SQLite](http://sqlitebrowser.org)

## REFERENCES

- [SQL](https://en.wikipedia.org/wiki/SQL). Wikipedia.
- [SQL Statement Syntax](https://dev.mysql.com/doc/refman/5.7/en/sql-syntax.html). MySQL Reference.
- [The SQL Language](https://www.postgresql.org/docs/9.5/static/sql.html). PostgresSQL Manual.
- [SQL As Understood By SQLite](https://www.sqlite.org/lang.html). SQLite.
- [Query Reference](https://cloud.google.com/bigquery/query-reference). Google BigQuery Documentation.
- [Streaming SQL Concepts](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/streaming-sql-concepts.html). Amazon Kinesis Analytics Developer Guide.
- [Modern SQL](http://modern-sql.com/). Markus Winand.
- [SQLite Tutorial](http://www.sqlitetutorial.net/). 

### Books

- Learning SQL. Alan Beaulieu. O'Reilly. April 27, 2009.