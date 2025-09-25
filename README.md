
# Worker boilerplate showcasing github.com/samber/do

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-%23007d9c)
![Build Status](https://github.com/samber/do-template-worker/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/do-template-worker)](https://goreportcard.com/report/github.com/samber/do)
[![License](https://img.shields.io/github/license/samber/do-template-worker)](./LICENSE)

**⚙️ A comprehensive worker template demonstrating the `github.com/samber/do` dependency injection library.**

This project showcases the full power of the `samber/do` dependency injection library in the context of a message-driven worker application. It implements a complete pub/sub worker with PostgreSQL integration and RabbitMQ messaging, demonstrating how `do` enables clean, modular, and testable worker architectures.

Perfect as a starting point for message-driven Go projects or as a learning resource for understanding dependency injection in worker applications.

**See also:**

- [do-template-api](https://github.com/samber/do-template-api)
- [do-template-cli](https://github.com/samber/do-template-cli)

## 🚀 Install

Clone the repo and install dependencies:

```bash
git clone --depth 1 --branch main https://github.com/samber/do-template-worker.git your-project-name
cd your-project-name

docker compose up -d
make deps
make deps-tools
```

## 💡 Features

- **Type-safe dependency injection** - Service registration and resolution using `samber/do`
- **Message-driven architecture** - Complete pub/sub worker with RabbitMQ consumer and producer
- **Database integration** - PostgreSQL with connection pooling and repository pattern
- **Modular architecture** - Clean separation of concerns with dependency tree visualization
- **Configuration management** - Environment-based configuration with dependency injection
- **Service lifecycle management** - Health checks and graceful shutdown handling
- **Repository pattern** - Data access layer with injected dependencies
- **Worker patterns** - Business logic with proper dependency management
- **Application lifecycle** - Health checks and graceful shutdown handling
- **Comprehensive error handling** - Structured logging and error management
- **Production-ready** - Ready to fork and customize for your next worker project
- **Extensive documentation** - Inline comments explaining every `do` library feature

## 🚀 Contributing

```sh
# install deps
make deps
make deps-tools

# compile
make build

# build with hot-reload
make watch-run

# test with hot-reload
make watch-test
```

## 🤠 `do` documentation

- [GoDoc: https://godoc.org/github.com/samber/do/v2](https://godoc.org/github.com/samber/do/v2)
- [Documentation](https://do.samber.dev/docs/getting-started)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2025 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
