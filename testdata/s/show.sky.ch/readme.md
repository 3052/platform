# show.sky.ch

~~~
url = https://show.sky.ch/de/filme/76978/mollys-game
monetization = FLATRATE
country = Switzerland
~~~

## smart proxy

~~~
GET https://show.sky.ch/de/filme/76978/mollys-game HTTP/2.0
← Upstream proxy ch.smartproxy.com:29001 refused HTTP CONNECT request: 403
Forbidden
~~~

## proxy seller

1. Einen Pass hinzufügen (add a passport)
2. enable javascript
3. sky show light, Abonnieren (subscribe)
4. e-mail
   - 2024-6-2@mailsac.com
5. E-Mail bestätigen (confirm e-mail)
6. password
   - Train-Over6
7. Ich habe gelesen und akzeptiere hiermit Sky Switzerland
   Datenschutzrichtlinie (I have read and hereby accept Sky Switzerland Privacy
   Policy)
8. Mein Konto erstellen (create my account)
9. credit card

~~~
>> GET https://js.stripe.com/v3 HTTP/2.0
← Upstream proxy res.proxy-seller.com:10000 refused HTTP CONNECT request:
503 Target host denied
~~~

use 10001
