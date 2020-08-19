# Nuovo Sito Poisson

Server che gestisce il nuovo sito del PHC, che verrà hostato a `https://phc.dm.*****.**`.

## Technologies

- [Echo Web Framework](https://echo.labstack.com/) &mdash; https://github.com/labstack/echo

    Sembra un framework abbastanza usato ed ha molte stelle.

- [Hugo](https://gohugo.io) &mdash; Generatore statico di blog scritto in Go.

    Tutta la parte di contenuti: blog, notizie, guide, etc. Verrà gestita da Hugo e servita sotto l'url `/blog` ed in teoria sarà sotto `content-poisson`.

## Development

Usare il comando `MODE=development go run main.go` per lanciare il server su `localhost:8000`.

## Design (?)

### Fonts

- Share 700: per le intestazioni
- Ubuntu 300, 700: per il testo.
- Ubuntu Mono: per i blocchi di codice


```css
@import url('https://fonts.googleapis.com/css2?family=Share:wght@700&family=Ubuntu+Mono&family=Ubuntu:wght@300;700&display=swap');
```
