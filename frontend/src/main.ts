import './style.css'

const tg: WebApp = window.Telegram.WebApp;

tg.ready()

fetch("/api/open?" + tg.initData)
