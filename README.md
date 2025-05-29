# 🌾 SeedLine - CLI Farm Game

A simple terminal-based farming simulation game written in Go. Manage your crops, plant seeds, wait for them to grow, harvest, and sell them to earn coins — all from your command line!

## Table of Contents

* [About](#about)
* [Features](#features)
* [Get Started](#get-started)
* [Usage Example](#usage-example)
* [Contributing](#contributing)
* [License](#license)

## About

**SeedLine** is a minimalist text-based farming simulator designed for the command line. You begin with a few coins and a plot of land. Plant seeds, wait for them to grow over time, harvest your crops, and sell them for profit. The game runs in real time, making for a casual yet engaging gameplay experience.

## Features

* [x] Real-time crop growth with ticker-based updates
* [x] Basic economy: plant, harvest, and sell crops
* [x] Simple text-based UI with ASCII rendering of the farm
* [x] Multiple crop types with varying growth/sell values
* [ ] Save/load game state
* [ ] Weather effects or random events
* [ ] Crop upgrade or fertilizer system

## Get Started

### Prerequisites

* [Go](https://golang.org/doc/install) 1.20 or higher
* Git

### Running the project

1. **Clone the repository**

```bash
git clone https://github.com/username/seedline.git
cd seedline
```

2. **Building and Running the game**

```bash
make build
make run
```

> Ensure your terminal supports ANSI escape codes for screen clearing.

## Usage Example

When you run the game, you'll see a UI like this:

```
🧺 Basket: 0  🪙 Coins: 0

🌱🌱🌱🌱🟫🟫🟫🟫🟫🟫🟫🟫
🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫
🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫
🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫
🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫
🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫🟫

Farm Game
1. Plant Seed
2. Harvest
3. Sell All
4. Exit

Time passed. Crops grew!

> Choose an action:
```

### Game Actions:

* `1`: **Plant Seed** – Buy and plant new seeds on available plots.
* `2`: **Harvest** – Collect all mature crops and move them to your basket.
* `3`: **Sell All** – Sell everything in your basket for coins.
* `4`: **Exit** – Quit the game.

⏱ Crops grow automatically every 20 seconds in real time!

## Contributing

Contributions are welcome! If you have any suggestions or improvements, please open an issue or a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
