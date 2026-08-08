# 1049 - Animal

## Description

In this problem, your job is to read three Portuguese words. These words define an animal according to the table below, from left to right. After, print the chosen animal defined by these three words.

### Decision Tree Table

- **vertebrado**
  - **ave**
    - **carnivoro:** `aguia`
    - **onivoro:** `pomba`
  - **mamifero**
    - **onivoro:** `homem`
    - **herbivoro:** `vaca`
- **invertebrado**
  - **inseto**
    - **hematofago:** `pulga`
    - **herbivoro:** `lagarta`
  - **anelideo**
    - **hematofago:** `sanguessuga`
    - **onivoro:** `minhoca`

## Specifications

- **Time Limit:** 1 second

## Input

The input contains 3 words, one per line, that will be used to identify the animal, with all letters in lowercase.

## Output

Print the animal name according to the given input.

## Samples

| Input                               | Output  |
| :---------------------------------- | :------ |
| vertebrado<br>mamifero<br>onivoro   | homem   |
| vertebrado<br>ave<br>carnivoro      | aguia   |
| invertebrado<br>anelideo<br>onivoro | minhoca |
