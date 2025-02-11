<!-- 
date: 2024-01-19
update: 2024-01-19
Tag: 
    - Rust
-->

# Cargo 創建專案

Cargo 是 rust 用來管理套件以及專案的工具，首先先用 cargo 來創建一個新的專案：

```shell
cargo new <PACKAGE NAME>

# 例如：
cargo new hello_world
```

cargo 會自動幫你在當前目錄創建一個新的專案，名為 hello_world。hello_world 專案內會有以下資料夾結構：

```
hello_world
|---/src # 存放程式的資料夾
|     |--- main.rs # 主程式進入點
|
|---/target # 還不知道這是啥
|--- .gitignore 
|--- Cargo.lock # 記錄套件版本及相依性
|--- Cargo.toml # 記錄專案資訊
```

# 編譯程式

## 使用 rustc 編譯程式

rustc 的全名應該是 rust compile。創建專案後就可以進入專案並在終端輸入：

```shell
rustc src/main.rs
```

進行編譯，編譯後會產生 `target` 資料夾，以及在工作目錄會產生一個 `main` 執行檔。接著在終端機輸入：

```shell
./main
```

即可執行程式，應該會看到輸出 `Hello, world!` 字串。

## 使用 cargo 編譯程式

我更推薦使用 cargo 來編譯程式，在終端機輸入：

```shell
cargo build
```

會在工作目錄產生 `target` 資料夾，裡面有跟專案名同名的執行檔，因此可以輸入：

```shell
./target/debug/hello_world
```

來執行程式。不過大部分的人在測試程式運行的時候都會用：

```shell
cargo run
```

這句指令會幫你編譯 + 執行程式，這樣你就不用分兩步執行程式了！

# References

[1.3 Hello, Cargo](https://rust-lang.tw/book-tw/ch01-03-hello-cargo.html)

---
此筆記僅是為了紀錄學習過程，如欲完整學習 rust，請去看官方提供的文件，那裡有更加完善的學習資源。筆記內容為自己的理解因此可能存在錯誤，如有錯誤敬請見諒。