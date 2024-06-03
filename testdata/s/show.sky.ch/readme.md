# show.sky.ch

~~~
url = https://show.sky.ch/de/filme/76978/mollys-game
monetization = FLATRATE
country = Switzerland
~~~

## smart proxy

~~~
> curl -x username:password@ch.smartproxy.com:29001
> https://show.sky.ch/de/filme/76978/mollys-game
curl: (56) Received HTTP code 403 from proxy after CONNECT
~~~

## proxy seller

~~~
curl -x socks5://username:password@res.proxy-seller.com:10000 `
https://js.stripe.com/v3
~~~

but MITM Proxy does not support:

https://github.com/mitmproxy/mitmproxy/issues/211

workaround:

https://mitmproxy.org/posts/har-support

then:

1. Einen Pass hinzufügen (add a passport)
2. enable javascript
3. sky show light, Abonnieren (subscribe)
4. e-mail
   - 2024-6-3@mailsac.com
5. E-Mail bestätigen (confirm e-mail)
6. password
   - Train-Over6
7. Ich habe gelesen und akzeptiere hiermit Sky Switzerland
   Datenschutzrichtlinie (I have read and hereby accept Sky Switzerland Privacy
   Policy)
8. Mein Konto erstellen (create my account)
9. credit card
10. Ich akzeptiere hiermit die allgemeinen Geschäftsbedingungen Sky Switzerland
   und ich bin volljährig (I hereby accept the Sky Switzerland General Terms and
   Conditions and I am of legal age)
11. Bezahlen (pay)

Es ist nicht genug Geld auf Ihrer Karte (there is not enough money on your card)
