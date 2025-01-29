# webDriver

https://github.com/mozilla/geckodriver

## proxy authentication

no support for `https_proxy`:

<https://bugzilla.mozilla.org/show_bug.cgi?id=1944213>

no support with geckoDriver:

https://github.com/mozilla/geckodriver/issues/1872

no support with geckoDriver:

<https://bugzilla.mozilla.org/show_bug.cgi?id=1395886>

## navigate to

~~~
curl -H content-type:application/json `
-d '{"url": "https://www.sky.ch"}' `
127.0.0.1:4444/session/222423a0-4da3-413c-a3fe-be1725c4c142/url
~~~

https://w3c.github.io/webdriver#navigate-to

## get all cookies

~~~
curl 127.0.0.1:4444/session/73a1041c-0bf6-4d40-a3f0-1784f8bde792/cookie
~~~

https://w3c.github.io/webdriver#cookies
