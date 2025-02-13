---
create_at: 2024.01.18
update_at: 2024.01.19
draft: false
tags: 
 - Rust
---

# Mac

## Install rustup

`rustup` 是負責管理及安裝 rust 的工具。他會將所有工具安裝在 `~/.cargo/bin` 目錄當中，安裝的工具包括 `rustc`、`cargo`、`rustup`。

在 mac 中可以直接用 homebrew 進行安裝：
```shell
brew install rustup
rustup-init
. "$HOME/.cargo/env"            # For sh/bash/zsh/ash/dash/pdksh
```

或是用官網提供的方式進行安裝：

```shell
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

安裝完成後可以試試看印出版本資訊來確認是否安裝成功：

```shell
rustc --version
cargo --version
```

## Uninstall `rustup`

如果用 brew 安裝的話可以執行以下指令解除安裝

```shell
brew uninstall rustup
```

如果是用其他方法，如，`curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh` 進行安裝的話，可以試試看下面指令：

```shell
rustup self uninstall
```

# References

https://www.rust-lang.org/zh-TW/tools/install