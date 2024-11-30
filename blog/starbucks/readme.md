# starbucks

- https://apkmirror.com/apk/starbucks-coffee-company/starbucks
- https://play.google.com/store/apps/details?id=com.starbucks.cn
- https://play.google.com/store/apps/details?id=com.starbucks.mobilecard

~~~
> play -i com.starbucks.mobilecard
details[8] = 0 USD
details[13][1][4] = 6.85.1
details[13][1][16] = Nov 18, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://www.starbucks.com/terms/privacy-notice/
downloads = 40.36 million
name = Starbucks
size = 50.39 megabyte
version code = 34922
~~~

create Android 8 device. install system certificate

1. virtual device? pass
2. virtual device with proxy?

## sign up

1. first name
2. last name
3. email
4. password
5. I accept the terms of use
6. join now

https://starbucks.com

## sign in

details above are wrong - Android sdkVersion 32 is required for login. might
have to try twice. install app, then install Frida:

~~~
pip install frida-tools
~~~

download and extract server:

https://github.com/frida/frida/releases

MAKE SURE TO GET `x86_64` file, example:

~~~
frida-server-16.0.0-android-x86_64.xz
~~~

then push:

~~~
adb root

adb push frida-server-16.5.7-android-x86_64 /data/local/tmp/frida-server
adb shell chmod +x /data/local/tmp/frida-server
adb shell /data/local/tmp/frida-server
~~~

then edit `config.js` with `mitmproxy-ca-cert.pem`, then start Frida:

~~~
frida -U `
-l android/android-certificate-unpinning.js `
-l config.js `
-f com.starbucks.mobilecard
~~~

https://github.com/httptoolkit/frida-interception-and-unpinning
