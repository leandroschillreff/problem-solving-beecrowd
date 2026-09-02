# 1051 - Taxes

## Description

Read a value with 2 digits after the decimal point, equivalent to the salary of a Lisarb inhabitant. Then print the due value that this person must pay of taxes, according to the table below:

| Salary Range              | Tax Rate      |
| :------------------------ | :------------ |
| From 0.00 to 2,000.00     | 8% (Wait, 0%) |
| From 2,000.01 to 3,000.00 | 8%            |
| From 3,000.01 to 4,500.00 | 18%           |
| Above 4,500.00            | 28%           |

_(Note: The brackets apply incrementally over the portion within each range)._

## Specifications

- **Time Limit:** 1 second

## Input

The input contains only a float-point number, with 2 digits after the decimal point.

## Output

Print the message `"R$ "` followed by a blank space and the total tax to be paid, with two digits after the decimal point. If the value is up to 2000, print the message `"Isento"`.

## Samples

| Input   | Output    |
| :------ | :-------- |
| 3002.00 | R$ 80.36  |
| 1701.12 | Isento    |
| 4520.00 | R$ 355.60 |
