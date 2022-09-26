# url-analyser

This service scrapes the HTML content of a web page using Goquery and returns following information.
- Title of the web page
- HTML version of the web page
- Headings present on the web page and their respective counts
- External link count on the web page
- Internal link count on the web page
- Inaccessible link count on the web page
- If the web page contains a login or not

# To run the service
- Navigate to project root
- go test -v ./...
- go run main.go

# Pre-requisites
- Go 1.17 or above should be installed

# Assumptions
- Success responses are considered as HTTP responses with the status code between 200 and 300.
- Assumed the service will be run locally and therefore, the base urls are not configured.
- Assumed all the href links in HTML template can be accessed via an HTTP GET method.
- External Link count and internal link count are considered to have accessible and inaccessible link counts.
<br />eg : External link count = Accessible external link count + Inaccessible external link count
- Internal links are considered only as href links with a prefix of "#".
- The HTML versions which are considered are, HTML 4.01, XHTML 1.0, XHTML 1.1 and HTML 5.

# Improvements
- Increase unit test coverage.
- Add metrics to visualize performance of the service(latency, error rate).
- Adapt the service functionalities to run parallelly.
- UI improvements.
- Dissect front end and back end services so they can be written as two micro services.
