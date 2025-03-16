# joyn

~~~
url = https://www.joyn.de/filme/side-effects-toedliche-nebenwirkungen-bpfifg67qmh4
monetization = FLATRATE
country = Germany
~~~

## paid

1. VPN
2. joyn.de
3. email
   - mailsac.com fail
   - gmail.com pass
4. vorname (first name)

even with regular card its rejected, so it must be a country thing

cannot link privacy.com to paypal because privacy is a piece of shit

so just use paypal and cancel? fails also

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
