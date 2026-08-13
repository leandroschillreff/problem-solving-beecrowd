# 1050 - DDD

## Description

Read an integer number that is the code number for phone dialing. Then, print the destination according to the following table:

| DDD | Destination    |
| :-- | :------------- |
| 61  | Brasilia       |
| 71  | Salvador       |
| 11  | Sao Paulo      |
| 21  | Rio de Janeiro |
| 32  | Juiz de Fora   |
| 19  | Campinas       |
| 27  | Vitoria        |
| 31  | Belo Horizonte |

If the input number isn’t found in the table above, the output must be `DDD nao cadastrado`.

## Specifications

- **Time Limit:** 1 second

## Input

The input consists of a single integer number.

## Output

Print the city name corresponding to the input DDD. Print `DDD nao cadastrado` if the typed number does not exist in the table.

## Samples

| Input | Output    |
| :---- | :-------- |
| 11    | Sao Paulo |
