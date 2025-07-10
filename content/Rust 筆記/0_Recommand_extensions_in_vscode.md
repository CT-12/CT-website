---
create_at: 2025.07.10
update_at: 2025.07.10
draft: false
tags: 
 - Rust
---

# Recommand extensions

## 1. rust-analyzer
![rust-analyzer](images/Recommand_extensions_in_vscode/rust_analyzer.png)

必裝的 extension，他可以給你 type hint, inline error 或查看 definition...。特別好用。

## 2. CodeLLDB
![CodeLLDB](images/Recommand_extensions_in_vscode/CodeLLDB.png)

幫助 rust 進行 debug 的工具，我還沒試過，等試過再回來分享心得。

## 3. Even Better TOML
![Even Better TOML](images/Recommand_extensions_in_vscode/Even_Better_TOML.png)

可以對 toml 檔進行 syntax highlight，除此之外好像還有其他功能，等弄到了再回來分享。

## 4. Dependi
![Dependi](images/Recommand_extensions_in_vscode/Dependi.png)

非常好用的工具，他可以直接在 rust 的 cargo.toml 檔中查看並挑選套件的版本。紅色叉叉❌表示你用的套件版本不是最新的，綠色勾勾✅則表示是最新的版本。
|||
|-|-|
|![Dependi-2](images/Recommand_extensions_in_vscode/Dependi-2.png)|![Dependi-3.](images/Recommand_extensions_in_vscode/Dependi-3.png)|

只要 hover 在版本上面就可以直接查看並挑選！

## 5. Error Lens
![Error Lens](images/Recommand_extensions_in_vscode/Error_Lens.png)

這個工具可以直接 inline 顯示錯誤訊息，蠻好看的。

## 6. Todo Tree
![Todo Tree](images/Recommand_extensions_in_vscode/Todo_Tree.png)

這個工具會 highlight 你寫的 `//TODO` 或是 `//FIXME` 註解，除此之外，因為 rust 有時候會用到 `todo!()` macro，所以他也可以被 highlight，更容易看到還有什麼地方沒完成。不過要注意的是，如果想讓 `todo!()` 也被 highlght 的話要去 vscode 的`設定`-> 輸入 `todo tree regex`->然後加上 `|todo!` 如下圖！

![Todo_tree-2](images/Recommand_extensions_in_vscode/Todo_tree-2.png)

# References

[Let's Get Rusty: Ultimate VS Code setup for Rust development (2025)
](https://www.youtube.com/watch?v=ZhedgZtd8gw)