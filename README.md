[![Build Status](https://travis-ci.org/aeternas/SwadeshNess.svg?branch=master)](https://travis-ci.org/aeternas/SwadeshNess)

# SwadeshNess
Backend for Swadesh-like lists creation

Read about Swadesh lists on [Wikipedia](https://en.wikipedia.org/wiki/Swadesh_list?oldformat=true)

```
$ curl "localhost/?translate=Hello+World&group=Romance"

Bonjour Tout Le Monde
Hola Mundo
Ciao Mondo
Salut Lume
```
or process several language groups simultaneously:

```
$ curl "localhost/?translate=Hello+World&group=Romance&group=Turkic"
Bonjour Tout Le Monde
Hola Mundo
Ciao Mondo
Salut Lume

Сәлам Мир
Сәләм Донъяға
Salam Dünya
Merhaba Dünya
```

Full list of languages group could be retrieved on `/groups` endpoint

P.S.: Yep, I am aware that CJKV is about characters not language :)
