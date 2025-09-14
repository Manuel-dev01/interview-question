# interview-question
1. Question: Tell me about a recent bug you diagnosed. What were the symptoms? Did you identify the root cause? What was it? and how did you fix this bug?. Max 3 sentences

  answer: While testing my Go-snippet project under moderate load, the app started throwing "too many connections" error, it turned out i hadn't configuresd Go SQL driver connection pool limits, so after tuning the pool with db.SetMaxOpenConns() and db.SetMaxIdleConns(), it helped solve the error by eliminating connection exhaustion.

2. Question: Choose any programming language of your choice. Write a function that returns the index of the first non-repeating character in a string, or -1 if none.

  answer: Refer to main.go file
