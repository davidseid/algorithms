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
    - 





