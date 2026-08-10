# Smart Make Art (Smart) 🚀

[![Go Reference](https://pkg.go.dev/badge/github.com/extbit/smart.svg)](https://pkg.go.dev/github.com/extbit/smart)
[![Join the chat at https://gitter.im/duzy/smart](https://badges.gitter.im/duzy/smart.svg)](https://gitter.im/duzy/smart)

> **🚧 Status: Under Active Development (Beta)**
> *Smart is currently mid-way through development. While the core JIT engine and syntax are functional, features and APIs are subject to change. We are actively seeking early adopters, feedback, and contributors to help shape the future of the project.*

**Smart Make Art** is a next-generation command-line utility and scripting language written in Go. Inspired by GNU `make`, Smart is explicitly designed to solve the pain points of compiling massive, hierarchical projects (like LLVM or Bitcoin Core) by offering a modular, data-typed, and multi-dialect build environment.

📚 [Read the Official Documentation](https://github.com/extbit/smart/wiki/Smart-Construction)

---

## Why Smart?

Building projects with complex hierarchies should be easy.

While traditional `Makefile` relies on a single, easily-polluted global namespace, `smart` introduces strict modularity. Symbols and rules are safely contained within local module or project scopes, and dependencies are explicitly declared using `import` or `use` keywords.

### Key Features
* **True Modularity:** No global namespace clashes. Modules handle specific tasks and are `use`d by parent projects.
* **Data-Typed Macros:** Unlike GNU Make, Smart evaluates variables with native data types for safer, predictable scripting.
* **Multi-Dialect Recipes:** Write build recipes natively in `shell`, `python`, or generate files dynamically using `plain` text dialects.

---

## Commercial Support & Custom Plugins

**Smart** is open-source and free to use under the BSD 3-Clause License.

However, because every build environment is unique, ExtBit LLC offers commercial engineering services for teams that need advanced capabilities. If your business requires specialized integrations, we can develop and license custom proprietary plugins, including:

* Custom multi-dialect recipes for proprietary toolchains.
* Specialized CI/CD pipeline integrations.
* Advanced build caching and distributed compilation modules.

Need a custom plugin or dedicated support for your company? [Contact ExtBit LLC](mailto:biz@extbit.com) or hire us directly via our [Upwork Project Catalog](https://www.upwork.com/services/product/development-it-custom-build-automation-or-modular-makefile-migration-using-smart-2085770956140247247?ref=project_share).

---

## Quick Start

### 1. Install `smart`
Install the `smart` command-line utility directly via Go:

```shell
$ go install [github.com/extbit/smart/cmd/smart@latest](https://github.com/extbit/smart/cmd/smart@latest)
$ smart -help
