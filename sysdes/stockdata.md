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
- Write API (for watching stocks)
- Login API (for connecting client to account)
- Database 

## Database Schema
- Tables
  - Users
    - Id
    - Username
    - Password
  - Stocks
    - Id
    - Ticker 
    - EOD Date
    - Open
    - Close
    - High 
    - Low
  - Users_Stocks
    - Users_ID
    - Stock_ID

## Bottlenecks

## Redesign

## Failures and Security
