# ginErrorHandler
simple package for handling custom errors to sending back to HTTP requests on gin framework .

there is a simple custom error type named GinError in the package that you can create
your own desired errors as its instants with optional error messages but with
valid http.Status codes.
by this, you can easily use your own instant as an error in the program body and when it needs to be passed to HTTP client,
you can pass it to HandleTheError func in the package to showing the custom message to the client.

 ! Attention : after using HandleTheError func, do not forget to call return just after the function call  in the scope you calling it. 

