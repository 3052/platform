# YouTube

## Android Studio

https://play.google.com/store/apps/details?id=com.google.android.youtube

Android API | result
------------|-------
31          | fail
32          | pass

Then install system certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://www.youtube.com/watch?v=k5dX9sjXYVk
~~~

## Device OAuth

- https://datatracker.ietf.org/doc/html/rfc8628
- https://developers.google.com/identity/sign-in/devices

## How to get `client_id` and `client_secret`

Set User-Agent to [1]:

~~~
Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version
~~~

Open MITM Proxy or your browser tools, then visit:

https://www.youtube.com/tv

and click "Sign in". On the next page, dont bother with any of the instructions,
just go back to the captured requests, and after about five seconds you should
see a JSON request like this:

~~~
POST /o/oauth2/token HTTP/1.1
Host: www.youtube.com

{
  "client_id": "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com",
  "client_secret": "SboVhoG9s0rNafixCSGGKXAT",
  "code": "AH-1Ng14qVvEj76OeM_h14Mgklgyhchbyc67MhULhCKPY6K-0DTYJqaKng2ULVFTmTzU...",
  "grant_type": "http://oauth.net/grant_type/device/1.0"
}
~~~

1. <https://github.com/youtube/cobalt/blob/master/src/cobalt/browser/user_agent_string.cc>

## Image

Is `maxres1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/maxres1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/maxres1.jpg

Is `sd1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/sd1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/sd1.jpg

If `hq1` always available? Yes:

http://i.ytimg.com/vi/hq2KgzKETBw/hq1.jpg

## name

episode:

~~~html
aria-label="In The Heat Of The Night S2 E3 • The Family Secret 47 minutes"
~~~

https://youtube.com/watch?v=2ZcDwdXEVyI

film:

~~~html
aria-label="Gattaca by Drama • 1997 1 hour, 46 minutes" title="Gattaca">
~~~

https://youtube.com/watch?v=R9lZ8i8El4I

video:

~~~html
aria-label="Sleepygirl 11 by Yagya (IS) - Topic 9,746 views 5 years ago 6 minutes, 15 seconds"
~~~

https://youtube.com/watch?v=7KLCti7tOXE
