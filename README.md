# BCC

Encrypt and Decrypt text with caeser cipher cryptography's method

## Encrypt

```sh
breaker_caeser_cipher encrypt [text] [left_shift]
```

***Example*** :

```sh
breaker_caeser_cipher encrypt "Hello, world!" 3
```
=> Khoor#/zruog$

## Decrypt

```sh
breaker_caeser_cipher decrypt [text] [left_shift]
```

***Example*** :

```sh
breaker_caeser_cipher decrypt "Khoor#/zruog$" 3
```
=> Hello, world!