sql -- language for structured databases queries
================================================

## Querying

`SELECT`
- is used to select which columns to include
- can use `AS` with each column to alias its name
- uses the `ALL` keyword by default

`FROM`
- identifies tables that contain desired columns
- can use `JOIN`, `INNER JOIN`, etc.
- can use `AS` with each table to alias its name
- can use a subquery to generate a temporary table

`WHERE`
- filters rows meeting specified conditions

`GROUP BY`
- groups (combines) rows based on a common value

`ORDER BY`
- sorts rows by column(s)
- should also use `ASC` or `DESC`
- can use an expression function

## REFERENCES

Alan Beaulieu. _Learning SQL_. O'Reilly. April 27, 2009