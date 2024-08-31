from mitmproxy import ctx, http

second = False

def request(flow: http.HTTPFlow) -> None:
   global second
   if flow.request.path == '/v3/login':
      if second:
         flow.kill()
      else:
         second = True

def running():
   ctx.options.anticomp = True
   ctx.options.showhost = True
