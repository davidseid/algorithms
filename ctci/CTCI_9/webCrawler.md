# Web Crawler System Design

### Prompt
Design a web crawler application that recursively checks links for a particular domain, and checks images and links to other domains using HEAD requests (which don't return data). 

Make sure to avoid infinite loops.

### Considerations
- Discuss Scope, Assumptions, Components, Bottlenecks, Redesign
- Address failure, security, reliability, maintainability, and scalability
- Do back of envelope calculations and reference algorithms

#### 1. Scope

1.1 Main features
  - Given a particular URL, our web crawler service will scan the web page and recursively check all the links on that page if they link to the same domain, or at least check they are not broken if they are images or go to other domains. 
  - Purpose is unknown, perhaps indexing the pages or validating links. 

#### 2. Assumptions and Constraints

2.1 Assumptions 
  - Let's assume our web crawler is being used mainly to validate links, and perhaps will keep a record of any invalid links in a web domain. 
  - We can assume we will be given a URL to search as the basis of a crawl.
  - It is unclear what we will need to store if anything, or who our users will be. 

#### 3. Components

3.1 Client 
  - Can exist for a user to submit a URL to crawl

3.2 Web Server
  - Handles requests from a client to check the URL, collect all links from the same domain to recursively check, as well as other links to send HEAD requests. 

3.2.1 Read API
  - Scans the website and returns a list of broken links

3.3 Database
  - Not clear if we need this, as we would need to crawl each time afresh in case links break or change.

3.4 Crawl Algorithm
  - Input: URL
  - Algorithm: 
    - Identifies the domain of the URL
    - Performs a GET request to receive the HTML of that URL
    - Reads the HTML and creates a queue of URLS that have the same domain
    - Creates another queue of links (images, etc) that have different domains
    - Stores a hash of all the URLs checked so far 
    - Perform HEAD request on each of the diff domain links
    - Store a list of any link errors (such as 404)
    - Recursively call algorithm on same-domain links
    - Check each time before recursing that the link has not been visited (to avoid infinite loops)
    - When all links have been visited, return the broken links

#### 4. Calculations

4.1 Data base storage

4.2 User Traffic
- In a realistic scenario, let's say we have a very popular web crawler, but this is not really a retail service, we are only likely to receive requests from developers or website owners
- However, some of these users may automate their requests, or submit requests on very large websites.

#### 5. Bottlenecks

5.1 Failures 
  - Infinite loops are a good thing to watch out for

5.2 Bottlenecks at Scale
  - A single server having to handle too many requests

5.3 Slow links
  - Do we retry links that are slow? Perhaps we can have a number of attempts for broken or slow links

#### 6. Redesign

6.1 Handle infinite loops
  - In recursive algorithm check for previously visited URLs

6.2 Overloaded Web Server
  - Use a load balancer to delegate to a fleet of READ APIs that will receive a URL and perform the GET and HEAD requests. Each READ API server will act in a non-blocking manner, but will keep track of the number of links to avoid overloading. 

#### 7. Maintainability and Extensibility

7.1 Security
  - Important to scan links to make sure they are legitimate

7.2 Availability and Reliability
  - Important that the service is reasonably available, but likely not essential that it is always available
  - Very important that the service is reliable 

7.3 Read versus Write heavy
  - Not clear if we will do any writing, mostly reading the web








