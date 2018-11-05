[![Build Status](https://travis-ci.org/aeternas/SwadeshNess.svg?branch=development)](https://travis-ci.org/aeternas/SwadeshNess)
[![CircleCI](https://circleci.com/gh/aeternas/SwadeshNess/tree/development.svg?style=svg)](https://circleci.com/gh/aeternas/SwadeshNess/tree/development)

# SwadeshNess
Backend for Swadesh-like lists creation. More information about Swadesh lists on [Wikipedia](https://en.wikipedia.org/wiki/Swadesh_list?oldformat=true)

Powered by [«Яндекс.Переводчик»](http://translate.yandex.ru/) so you need to acquire an [API key](https://translate.yandex.ru/developers/keys) and setup envirnoment variable:
```
export YANDEX_API_KEY=<your key>
```

Example queries:

```
$ curl "localhost/?translate=Hello+World&group=Romanic" | jq .

{
  "results": [
    {
      "name": "Romanic",
      "results": [
        {
          "name": "Latin",
          "translation": "Salve Mundi"
        },
        {
          "name": "French",
          "translation": "Bonjour Tout Le Monde"
        },
        {
          "name": "Spanish",
          "translation": "Hola Mundo"
        },
        {
          "name": "Italian",
          "translation": "Ciao Mondo"
        },
        {
          "name": "Romanian",
          "translation": "Salut Lume"
        },
        {
          "name": "Portugal",
          "translation": "Olá Mundo"
        }
      ]
    }
  ]
}
```
or process several language groups simultaneously:

```
$ curl "localhost/?translate=Hello+World&group=Romanic&group=Turkic" | jq .

{
  "results": [
    {
      "name": "Romanic",
      "results": [
        {
          "name": "Latin",
          "translation": "Salve Mundi"
        },
        {
          "name": "French",
          "translation": "Bonjour Tout Le Monde"
        },
        {
          "name": "Spanish",
          "translation": "Hola Mundo"
        },
        {
          "name": "Italian",
          "translation": "Ciao Mondo"
        },
        {
          "name": "Romanian",
          "translation": "Salut Lume"
        },
        {
          "name": "Portugal",
          "translation": "Olá Mundo"
        }
      ]
    },
    {
      "name": "Turkic",
      "results": [
        {
          "name": "Tatar",
          "translation": "Сәлам Мир"
        },
        {
          "name": "Bashkort",
          "translation": "Сәләм Донъяға"
        },
        {
          "name": "Azerbaijanian",
          "translation": "Salam Dünya"
        },
        {
          "name": "Turkish",
          "translation": "Merhaba Dünya"
        }
      ]
    }
  ]
}
```

Full list of languages group could be retrieved on `/groups` endpoint
