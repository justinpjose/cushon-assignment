# Assignment

## Decisions
- Router was based on the following documentation - https://www.alexedwards.net/blog/which-go-router-should-i-use
- Using `julienschmidt/httprouter` as it handles 404 and 405 responses correctly (including allowing you to set custom handlers for these responses), and also automatically handles OPTIONS requests for you too.