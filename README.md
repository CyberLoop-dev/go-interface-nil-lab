<div align="center">

![Header](https://capsule-render.vercel.app/api?type=waving&color=0:0d1822,100:00add8&height=170&section=header&text=go-interface-nil-lab&fontColor=ffffff&fontSize=38)

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=22&pause=1000&center=true&vCenter=true&width=620&lines=Understand+typed+nil+interfaces;See+the+dynamic+type+and+value" alt="Typing title"/>

[![License](https://img.shields.io/github/license/CyberLoop-dev/go-interface-nil-lab)](LICENSE)
![Language](https://img.shields.io/github/languages/top/CyberLoop-dev/go-interface-nil-lab)
[![Stars](https://img.shields.io/github/stars/CyberLoop-dev/go-interface-nil-lab?style=social)](https://github.com/CyberLoop-dev/go-interface-nil-lab/stargazers)
![Last commit](https://img.shields.io/github/last-commit/CyberLoop-dev/go-interface-nil-lab)

<img src="assets/mark.svg" width="520" alt="Typed nil interface diagram"/>

</div>

## 🧭 Overview
A tiny Go command-line demonstration of the typed-nil interface gotcha. It shows why an interface containing a nil pointer is not itself equal to `nil`.

## ✨ Features
- 🧩 Compares a nil interface with an interface holding a typed nil pointer.
- 🔎 Prints the interface's dynamic type and underlying pointer state.
- 🛡️ Demonstrates a nil-safe pointer-receiver method.
- 📦 Supports readable and JSON output.

## 🖥️ Demo
Run the example locally with Go; no hosted demo is configured.

## 🚀 Quick Start
```bash
go run .
go run . --json
go test ./...
```

## 🧰 Tech Stack
- Go 1.22+
- Standard library: `flag`, `reflect`, `encoding/json`

## 🗂️ Folder Structure
```text
.
├── assets/mark.svg
├── main.go
├── main_test.go
├── go.mod
└── README.md
```

## 🗺️ Roadmap
- [ ] Add examples with multiple interface methods.
- [ ] Compare value and pointer receiver method sets.
- [ ] Add a short escape-analysis companion example.

## 🤝 Contributing
Add small, compilable examples with tests and explain the exact Go rule being demonstrated.

## 📄 License
MIT. See [LICENSE](LICENSE).

<div align="center">

![Footer](https://capsule-render.vercel.app/api?type=waving&color=0:00add8,100:0d1822&height=100&section=footer)

</div>
