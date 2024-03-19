## HTTP 서버 테스트
* You can make use of the testing package to test your HTTP servers too. In addition, the net/http/httptest package provides some helpers that make the process painless.


Exercise 4
Using your solution to Exercise 3 above (or this one) as a starting point, write a test function that uses the net/http/httptest package to record the your server's HTTP response so you can verify the status code and body for each scenario. Use table-driven testing to iterate over the different scenarios we specified in our original expectations of the program back in Exercise 3.

Objectives
Make use of the net/http/httptest package to check response status code and body for each scenario.
Use table-driven testing technique.
What you need to know
Take a quick look at the net/http/httptest package document examples to see how it's used in practice.
Test Coverage
As you test, you may find it useful to see how much of your code you're covering with your tests. Go offers the ability to generate a coverage report after your tests are run. The coverage report for a solution to Exercise 4 above looks like this:

go test -v -cover
=== RUN   TestProverbsHandler
--- PASS: TestProverbsHandler (0.00s)
PASS
coverage: 68.2% of statements
ok      github.com/jboursiquot/go-in-3-weeks/exercises/http/e4   0.020s
Exercise 5
The Go toolchain also provides a richer coverage report experience through a generated HTML report. Your task is to figure out the command line options for producing and consuming this report. For example, the richer version of the coverage report above for the solution looks like this:

Test Coverage

Check out The cover story for an interesting read of how the cover tool is used in Go.

Objectives
Use the go toolchain to generate an HTML coverage report you can consume in your browser.
💡 Rather than running the multi-step process of generating the coverage report manually, you could use tools like Make to bundle all of the commands under one call.

💡 You may also find GoConvey useful as it provides an intuitive and pleasant-looking browser GUI that refreshes while you change the code. It uses all the same standard Go tooling behind the scenes so you can use it on new or existing projects aline.