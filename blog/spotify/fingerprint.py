from mitmproxy import ctx, http
import mitmproxy
import logging
import base64

def running():
   ctx.options.anticomp = True
   ctx.options.showhost = True

def tls_clienthello(data: mitmproxy.tls.ClientHelloData):
   hello = data.client_hello.raw_bytes()
   logging.info(base64.b64encode(hello))
   
def request(flow: http.HTTPFlow) -> None:
   logging.info(flow.request.path)
