# Stock Data System Design Prompt

## Requirements and Objective
- App will serve 1000 clients
- It will provide EOD Stock price info (open, close, high, low)
- Assume I already have that data
- How will we design the client facing service?
- Include dev, rollout, monitoring, maintenance 

## Scope
- Users include up to 1000 clients who wish to view stock data
- They will need to be able to access this information every day
- Features: 
  - See all stock data
  - Search stock by listing code
  - Potentially sort by diff between open and close
  - Just viewing, not writing any information 
  - Perhaps can filter too based on some criteria 
    - Industry
  - Can also create a log in to track particular stocks associated with that users account
  - Can view My stocks or all stocks
  - Potentially could do server side pagination, such that only 1 page exists on the client at any given time

## Assumptions
- We probably only need to update the data once a day since we aren't displaying real time changes, just the EOD stock price info
- Only 1000 clients so we don't need to worry too much about scale, although it is possible we will have to worry about all 1000 requests hitting our application at the same exact time in the worst case, potentially if everyone immediately wants to see the EOD data when the market closes. 
- Also assuming this is tied to a particular stock exchange, otherwise we would create options to filter by exchange, and it would possibly need to be updated upon each close. 
- Probably no reason to store historical data, just current EOD data

## Questions
- Can we just serve up all the info to the client initially? Need to count up how much it is.

## Component Design

- Client 
- Web Server
- Read API (for viewing particular subset of data)
  - If we want to sort by volatility, will need to index stocks by that, as well as ticker
  - User supplies simple get all, or options to filter the data
- Write API (for watching stocks)
  - User supplies a ticker or tickers to add the ids to the user table
- Login API (for connecting client to account)
  - Track cookie for login info
  - If none passed in cookie, prompt user for login and password
  - Or if new, give sign in option
- SQL Database
- Worker to Update Datastore each data

## Database Schema
- Tables
  - Users -> 16 bytes
    - Id: INT -> 4 bytes
    - Username -> 8 bytes
    - Password -> 8 bytes
  - Stocks -> 56 bytes * 5k -> 250k bytes -> 250M
    - Id -> 8 bytes
    - Ticker -> 8 bytes
    - EOD Date -> 8 bytes
    - Open -> 8 bytes
    - Close -> 8 bytes
    - Diff -> 8 bytes
    - High -> 8 bytes
    - Low -> 8 bytes
  - Users_Stocks
    - Users_ID
    - Stock_ID

## Bottlenecks
- All 1000 users pinging servers at same time
- Sorting by certain information

## Redesign
- Storing all data on the client, potentially might need pagination
- Indexing
  - Index User by username
  - Index Stocks by Ticker, Diff

## Failures and Security
- Mainly read heavy 
- Traffic at max can only be 1000 RPS, not too bad
- Database can be overrun potentially, in which case we can have duplicates or do some database sharding to partition parts of the datastore
- Availability: 
  - The information should always be available and always be up to date, althoug it is only updated once per day.
- Reliability: 
  - 100% 
- Maintenance: Load balancer, cache