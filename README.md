# Nuovo Sito Poisson

Servizio del nuovo sito del PHC

## Technologies

- [Gin Web Framework](https://gin-gonic.com/) &mdash; https://github.com/gin-gonic/gin

    Sembra un framework abbastanza usato ed ha molte stelle.

- [Hugo](https://gohugo.io) &mdash; Generatore statico di blog scritto in Go.

    Tutta la parte di contenuti: blog, notizie, guide, etc. Verrà gestita da Hugo e servita sotto l'url `/blog`.

## Development

Esegui `go run main.go` per farlo girare su `localhost:8000` e vedere le modifiche in diretta, amico.

## Design (?)

### Fonts

- Share 700: per le intestazioni
- Ubuntu 300, 700: per il testo.
- Ubuntu Mono: per i blocchi di codice 

```css
@import url('https://fonts.googleapis.com/css2?family=Share:wght@700&family=Ubuntu+Mono&family=Ubuntu:wght@300;700&display=swap');
```
