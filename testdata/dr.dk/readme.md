# dr.dk

https://dr.dk/drtv

~~~
url = https://www.dr.dk/drtv/program/le-mans-66_482510
monetization = FREE
country = Denmark
~~~

## android

https://play.google.com/store/apps/details?id=dk.dr.tvplayer

~~~
> play -i dk.dr.webplayer
details[8] = 0 USD
details[13][1][4] = 4.7.0
details[13][1][16] = Nov 18, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://www.dr.dk/service/privatlivspolitik/
downloads = 2.09 million
name = DRTV
size = 56.60 megabyte
version code = 180121
~~~

output is wrong - at least Android 7 is required to get to home screen,
Android 10 for login. install system certificate

~~~
adb install-multiple (Get-ChildItem *.apk)
~~~

HLS only

## web

1. log in
2. e-mail
   - mailsac.com
3. next
4. password
5. next
6. name
7. next
8. Can we use cookies for statistics? no thanks
9. Can we target content to you? no thanks
10. Do you want the latest news and recommendations for DR’s content? no thanks
11. confirm e-mail

HLS only
