implement an interface from inside a block

I want to implement an interface from inside a block. I tried this:

~~~go
package io

type reader interface {
   read() bool
}

func new_reader() reader {
   type hello struct{}
   return hello{}
}
~~~

result:

~~~
cannot use hello{} (value of type hello) as reader value in return statement:
hello does not implement reader (missing method read)
~~~

OK fair enough, so I add that:

~~~go
func new_reader() reader {
   type hello struct{}
   func (hello) read() bool { return true }
   return hello{}
}
~~~

but it fails again:

~~~
syntax error: unexpected bool at end of statement
syntax error: non-declaration statement outside function body
~~~

is what I am trying to do possible? if not, why not?
