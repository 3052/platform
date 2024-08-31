from mitmproxy import ctx, http

def running():
   ctx.options.anticomp = True
   ctx.options.showhost = True

second = False

def request(f: http.HTTPFlow) -> None:
   global second
   if f.request.path == '/v3/login':
      if second:
         f.kill()
      else:
         second = True
