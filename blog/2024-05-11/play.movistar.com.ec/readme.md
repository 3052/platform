# play.movistar.com.ec

country = Ecuador

command:

~~~
curl -m 9 -v -i -H x-forwarded-for:186.68.0.0 `
https://play.movistar.com.ec/details/movie/20820131
~~~

result:

~~~
*   Trying 213.140.60.178:443...
* Connection timed out after 9006 milliseconds
* Closing connection 0
curl: (28) Connection timed out after 9006 milliseconds
~~~
