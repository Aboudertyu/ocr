# ROK Monster OCR (GoLang)

[![Discord](https://img.shields.io/discord/768180228710465598)](https://discord.gg/drhxwVQ) 
[![License: MIT](https://img.shields.io/github/license/xor22h/rok-monster-ocr-golang)](https://opensource.org/licenses/MIT)


---

👋 An idea for this project came from [ROK Monster OCR Tools](https://github.com/carmelosantana/rok-monster-ocr).

---

👋 Join our [Discord](https://discord.gg/drhxwVQ) for help getting started or show off your results!

---

## Kingdom Statistics

Command line tools to help collect player statistics from [Rise of Kingdoms](https://rok.lilithgames.com/en). By analyzing screenshots we can extract various data points such as governor power, deaths, kills and more. This can help with various kingdom statistics or fairly distributing [KvK](https://rok.guide/the-lost-kingdom-kvk/) rewards.

![Sample](./media/sample.png)

[![asciicast](https://asciinema.org/a/gYerprrrw0DVOXZbitOfHrPqg.svg)](https://asciinema.org/a/gYerprrrw0DVOXZbitOfHrPqg)

### Features

- Character recognition by [Tesseract](https://github.com/tesseract-ocr/tesseract)
- Fast hash based image comparison

### Limitations

- English language is preferred as coordinate information lines up most accurately with English.
- No way to merge user information from different screens.
- Requires properly defined template

## Getting started

```bash
git clone https://github.com/xor22h/rok-monster-ocr-golang
cd rok-monster-ocr-golang
go build .
./rok-monster-ocr-golang -help
```

## Community

Have a question, an idea, or need help getting started? Checkout our [Discord](https://discord.gg/drhxwVQ)!

## License

The code is licensed [MIT](https://opensource.org/licenses/MIT) and the documentation is licensed [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).