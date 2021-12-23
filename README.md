# Nuovo Sito Poisson

Repo del server del nuovo sito per il PHC.

### Setup

Per clonare il progetto usare

```bash
$ git clone https://github.com/phc-dm/phc-server
```

## Development

Copiare il file `.env.dev` in `.env` per dire al server di lavorare in modalità di development e su quale indirizzo servire il sito, poi avviare il server.

```bash shell
$ cp .env.dev .env
$ go run .
```

Un comando comodo in fase di development che usa entr è

```bash shell
$ find . -type f -name '*.go' | entr -r go run .
```

### Environment Variables

- `MODE`

    Può essere `production` (default) o `development`.

- `HOST`

    Indirizzo (locale) sul quale servire il sito, di default è `localhost:8000`.

- `<SERVIZIO>_URL`

    Rappresentano link ad altri servizi forniti, è comodo impostarli così in modo da poter anche provare tutto insieme in locale su varie porte (e poi in produzione i vari url diventano link a sotto-domini del sito principale).

    Per ora ci sono solo `GIT_URL` e `FORUM_URL`.
