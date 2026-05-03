# binance-tui
 
A terminal UI for Binance Spot written in Go. Shows live orderbook depth, portfolio balance tracking, and real-time price monitoring — entirely from the terminal.
 
Built from scratch using raw HTTP and WebSocket connections, with no Binance SDK dependency.
 
## Features
 
- Live orderbook depth for any trading pair via WebSocket
- Authenticated portfolio balance viewer via Binance REST API
- Real-time price and percentage change for a configurable watchlist
- Fullscreen TUI with keyboard navigation

## Requirements
 
- Binance Demo Trading account — [demo-api.binance.com](https://demo-api.binance.com)

## Installation
 
```bash
git clone https://github.com/tanayarun/Binance-TUI.git
cd Binance-TUI
go mod tidy
go build -o binance-tui .
```
 
## Configuration
 
The `portfolio` command requires API keys from your Binance Demo Trading account. Set them as environment variables:
 
```bash
export BINANCE_API_KEY=your_api_key
export BINANCE_SECRET_KEY=your_secret_key
```
 
To persist across sessions, add them to your shell profile:
 
```bash
echo 'export BINANCE_API_KEY=your_api_key' >> ~/.zshrc
echo 'export BINANCE_SECRET_KEY=your_secret_key' >> ~/.zshrc
source ~/.zshrc
```
 
The `watch` and `tickers` commands use public WebSocket streams and do not require API keys.
 
---
 
## Usage
 
```bash
# Live orderbook for a symbol
./binance-tui watch --symbol BTCUSDT
 
# Portfolio balances (requires API keys)
./binance-tui portfolio
 
# Live prices for top USDT pairs
./binance-tui tickers
```
 
---
 
## Keybindings
 
| Key      | Action |
|----------|--------|
| q        | Quit   |
| ctrl+c   | Quit   |
 
