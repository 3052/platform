from mitmproxy import ctx, http, tls
import base64
import logging

second = False

def request(flow: http.HTTPFlow) -> None:
   global second
   logging.info(flow.request.path)
   if flow.request.path == '/v3/login':
      if second:
         flow.kill()
      else:
         second = True

def running():
   ctx.options.anticomp = True
   ctx.options.showhost = True

def tls_clienthello(data: tls.ClientHelloData):
   hello = data.client_hello.raw_bytes()
   logging.info(base64.b64encode(hello))
