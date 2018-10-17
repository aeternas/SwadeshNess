[![Build Status](https://travis-ci.org/aeternas/SwadeshNess.svg?branch=development)](https://travis-ci.org/aeternas/SwadeshNess)

# SwadeshNess
Backend for Swadesh-like lists creation. More information about Swadesh lists on [Wikipedia](https://en.wikipedia.org/wiki/Swadesh_list?oldformat=true)

Powered by [«Яндекс.Переводчик»](http://translate.yandex.ru/) so you need to acquire an [API key](https://translate.yandex.ru/developers/keys) and setup envirnoment variable:
```
export YANDEX_API_KEY=<your key>
```

Example queries:

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
