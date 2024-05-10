# joyn

- https://joyn.de/filme/barry-seal-only-in-america
- https://justwatch.com/de/Anbieter/joyn-plus

## free

1. https://joyn.de/mein-account
2. Alles akzeptieren (accept everything)
3. e-mail
   - invalid email passed : 2024-5-9@mailsac.com
4. Anmelden (log in)
5. Klick einfach auf den Link in der Mail, dann kann's losgehen (Just click on
   the link in the e-mail and you're ready to go)
6. password
7. Weiter (next)
8. Männlich (male)
9. day
10. month
11. year
12. Speichern (save)

## paid

1. https://joyn.de
2. PLUS kostenlos testen (test PLUS free of charge)
3. Jetzt kostenlos testen (test now free of charge)
4. Vorname (first name)
5. last name
6. credit card
7. Karteninhaber (cardholder)
8. card number
9. CVC
10. MM
11. YY
6. Ich stimme der Ausführung des Vertrages vor Ablauf der Widerrufsfrist
   ausdrücklich zu (I expressly agree to the execution of the contract before the
   expiry of the withdrawal period)
7. Zahlungspflichtig bestellen (order with obligation to pay)
8. Ihre Kreditkarte wurde abgelehnt (your credit card has been declined)

## quality

android phone:

~~~go
(type=="video"&&MaxHeight<=576)||(type=="audio"&&FourCC=="AACL")
~~~

android tv:

~~~
type=="video"&&MaxHeight<=480)||(type=="audio"&&FourCC=="AACL")
~~~

web anonymous:

~~~go
(type=="video"&&MaxHeight<=480)||(type=="audio"&&FourCC=="AACL")
~~~

web free:

~~~
(type=="video"&&MaxHeight<=480)||(type=="audio"&&FourCC=="AACL")
~~~

## prior art

<https://github.com/Maven85/plugin.video.joyn/blob/master/resources/lib/submodules/libjoyn_video.py>

looks like its here now:

<https://joyn.de/_next/static/chunks/3806.dcb7578b32019150.js>

but obfuscated:

~~~js
t[i(486) + "t"] = "5C783" + i(628) + "78646" + i(642) + i(701) + "65C78" +
i(545) + i(371) + i(336) + i(470) + "65395" + i(415) + i(438) + i(587) + i(468)
+ i(395) + i(561) + "5"
~~~

x-api-key is hard coded in JavaScript

## android

https://play.google.com/store/apps/details?id=de.prosiebensat1digital.seventv

~~~
> play -i de.prosiebensat1digital.seventv
details[6] = Joyn GmbH
details[8] = 0 USD
details[13][1][4] = 5.59.0-AOS-559092432
details[13][1][16] = May 3, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 7.0 and up
downloads = 15.76 million
name = Joyn | deine Streaming App
size = 94.26 megabyte
version code = 559092432
~~~

Create Android 7 device. Install system certificate

https://apkmirror.com/apk/joyn-gmbh/joyn-deine-streaming-app
