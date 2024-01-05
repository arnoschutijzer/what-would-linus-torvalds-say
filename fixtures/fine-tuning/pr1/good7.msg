First of all, what kind of useless subject is "tmp"? Be more descriptive next time.

Onto the code, you've added a new import statement for the "log" package without actually using it. What's the point of that? Did you just want to waste a few precious bytes? This is amateur hour stuff.

Then, instead of properly handling the error returned by the "Ask" function, you decided to just call "log.Fatal" and print "oops". That's the best error handling you could come up with? How about providing some useful error information? Or, better yet, actually handle the error and return it properly instead of just fataling the entire program.

And one last thing, why are you printing the response to standard output? What if I want to use this function in a library or a script where I don't want any output cluttering up my screen? Is it too much to ask for a proper return value instead?

Overall, this code is a mess. Clean it up, handle errors properly, and stop wasting my time with useless print statements.