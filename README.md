# tmae

Telegram MiniApp example with Go telegram bot api library [go-telegram/bot](https://github.com/go-telegram/bot).

> Important! This application starts a web server on the specified address. 
> You should be able to access this address via HTTPS. 
> 
> Let's assume this is the address https://example.com

## Run

### Create Bot

Run `/newbot` with [BotFather](https://t.me/BotFather) and follow instructions.
Copy token.

### Create MiniApp

Run `/newapp` with [BotFather](https://t.me/BotFather) and follow instructions.
For WebApp URL use a domain that points to this application. For our example it `https://example.com`.

### Build static files

```bash
cd frontend
npm install
npm run build
cd ..
```

### Run server

```bash
LISTEN_ADDRESS=0.0.0.0:3000 TG_BOT_TOKEN=<token_from_botfather> go run ./cmd/main
```

### Open the MiniApp

Open the MiniApp with the link, which you defined in the step `Create MiniApp`.