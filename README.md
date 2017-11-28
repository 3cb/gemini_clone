# gemini_clone
A Gemini Bitcoin Exchange clone built on Vue/Vuex with xstream for websocket stream and Go backend.
Compare to https://gemini.com/ (API Docs: https://docs.gemini.com/websocket-api/)
> master branch connects to Gemini websocket API directly from browser (3 concurrent connections).
> go_stream branch multiplexes three websockets with Go and streams them to browser through one websocket connection.
> Note: go_stream branch will currently not work properly with multiple concurrent connections.

## Build Setup

``` bash
From static folder:

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```