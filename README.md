# Nuovo Sito Poisson

Documentazione principale è raccolta in [un altro repository](https://github.com/phc-dm/documentazione/blob/master/progetti/server-poisson/README.md).

## Development

### Setup

Per clonare il progetto usare

```bash
$ git clone --recursive https://github.com/phc-dm/server-poisson
```

l'opzione `--recursive` serve a clonare anche il sottomodule relativo al blog del sito, si può anche omettere però non sarà disponibile la visualizzazione del blog in fase di development.

### Usage

Usare il seguente comando per lanciare il server su `localhost:8000`.

```bash
$ MODE=development go run main.go
```

Environment variables:

- `MODE`: di default è `production`, può anche essere `development`
- `PORT`: controlla la porta su cui viene lanciato il server, di default è `8000` 

## Materiale

```css
@import url('https://fonts.googleapis.com/css2?family=Share:wght@700&family=Ubuntu+Mono&family=Ubuntu:wght@300;700&display=swap');
```
