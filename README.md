# tsgen ⚡

A fast TypeScript project starter CLI built with Go.

`tsgen` helps you generate clean, scalable TypeScript project structures instantly — without repeating the same setup every time.

Built for developers who are tired of manually creating folders, configs, boilerplate files, and feature structures.

---

## Features

- ⚡ Fast CLI built with Go
- 📦 Generate TypeScript project starters instantly
- 🧱 Clean scalable project structure
- 🛠 Supports reusable templates
- 🎯 Interactive prompts
- 📁 Auto-generate folders and starter files
- 🚀 Future support for:
  - Flutter feature generators

---

## Installation

### Go Install

```bash
go install github.com/atocodes/tsgen@latest
```

### Verify Installation

```bash
tsgen --version
```

---

## Usage

### Create a New Project

```bash
tsgen create
```

or

```bash
tsgen new
```

---

## Example

```bash
tsgen create my-app
```

Example generated structure:

```txt
my-app/
├── src/
│   ├── features/
│   ├── shared/
│   ├── utils/
│   ├── types/
│   └── index.ts
├── tests/
├── .gitignore
├── tsconfig.json
├── package.json
└── README.md
```

---

## Interactive CLI

`tsgen` provides an interactive experience directly in the terminal.

Example flow:

```bash
✔ Project name: my-app
✔ Package manager: pnpm
✔ Initialize git: yes
✔ Install dependencies: yes
```

---

## Commands

| Command | Description |
|---|---|
| `tsgen create` | Create a new project |
| `tsgen new` | Alias for create |
| `tsgen version` | Show CLI version |
| `tsgen help` | Show help menu |

---

## Roadmap

- [ ] TypeScript starter templates
- [ ] Express.js starter
- [ ] FastAPI starter
- [ ] Flutter feature generator
- [ ] Next.js starter
- [ ] Template marketplace
- [ ] Plugin system
- [ ] Config file support
- [ ] AI-assisted scaffolding

---

## Why tsgen?

Most project generators are either:
- too heavy,
- too opinionated,
- or locked to a single framework.

`tsgen` aims to stay:
- fast,
- minimal,
- extensible,
- and developer-friendly.

You stay in control of your architecture.

---

## Tech Stack

- Go
- Cobra CLI
- Templates
- File generators

---

## Contributing

Contributions are welcome.

If you'd like to improve `tsgen`, feel free to:
- open issues
- suggest features
- submit pull requests

---

## Development

Clone the repository:

```bash
git clone https://github.com/atocodes/tsgen.git
```

Run locally:

```bash
go run main.go
```

Build binary:

```bash
go build -o tsgen
```

---

## Release

Create a release tag:

```bash
git tag v0.1.0
git push origin v0.1.0
```

---

## License

MIT License

---

## Author

Built by Ato Codes

- GitHub: https://github.com/atocodes

---

## Star the Repo ⭐

If `tsgen` helps you, consider starring the repository and sharing it with other developers.