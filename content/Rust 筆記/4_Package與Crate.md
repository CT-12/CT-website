---
create_at: 2024.01.20
update_at: 2024.01.20
draft: false
tags: 
 - Rust
 - Package
 - Crate
---

# Package

Package 是從 `cargo new <Package Name>` 指令所產生。裡面會存放一個 `Cargo.toml` 檔案來暗示他是一個 Package。一個 Package 中會有一個或多個 crates

# Crate

Crate 是 rust 中的最小編譯單位，每個 crate 都會有一個 crate root，它是編譯時的入口點。

Crate 分成兩種：

- 第一種是`執行檔 crate（Binary crate）`。這種類型的 crate 的入口點(crate root)會是 `src/main.rs` 文件。也就是 rust 會從 `src/main.rs` 進入，並將其編譯成一個與 package 名稱相同的二進制檔案。
- 第二種是`函式庫 crate (Library crate)`。這種類型的 crate 的入口點(crate root)會是 `src/lib.rs` 文件。也就是 rust 會從 `src/lib.rs` 進入，並將其編譯成一個與 package 名稱相同的函式庫檔案。

值得注意的地方是，一個 package 可以有一個或多個 binary crate，但是最多只能有一個 library crate（換句話說就是只能有一個 `src/lib.rs` 檔案）。並且，一個 package 中至少要有一個 crate，不管是 binary 還是 library。

前面提到 binary crate 可以有一個或多個。如果 package 只有一個 `src/main.rs` 檔案，rust 就會從 `src/main.rs` 進入並編譯成一個 binary crate。但如果需要多個 binary crates (也就是多個二進制執行檔)，可以在 `src/bin/` 資料夾下新增檔案，每個檔案都會被視為獨立的 binary crate。例如，現在有個檔案結構：

```
package/
├── Cargo.toml
└── src/
    ├── main.rs
    ├── bin/
    │   ├── bin_crate_1.rs
    │   └── bin_crate_2.rs
    └── lib.rs
```

如果直接輸入：

```shell
cargo run
```

會出現錯誤：

```shell
error: `cargo run` could not determine which binary to run. Use the `--bin` option to specify a binary, or the `default-run` manifest key.
available binaries: bin_crate_1, bin_crate_2, package
```

因為 rust 不知道要編譯哪個 binary crate，所以他會建議你使用 `--bin` 來指定你要編譯的 binary crate。同時他還會列出所有可以編譯的 binary crates。

!!! note "Tips" 

    仔細看會發現他列出的其中一個 binary crate 叫做 package，它所對應的入口文件就是 main.rs！因為 main.rs 預設會編譯出與 package 同名的 binary crate。

總結：

Crate 是入口文件被編譯過後的檔案，`lib.rs` 會編譯成 library crate。`main.rs` 會編譯成 binary crate。


# Mod (模組)

首先先快速了解一下編譯器會如何編譯模組，以下例子來自 Rust 程式設計語言的 7.2 章節，想看程式碼範例最好去那邊看。

目錄結構：
```
backyard
├── Cargo.lock
├── Cargo.toml
└── src
    ├── garden
    │   └── vegetables.rs
    ├── garden.rs
    └── main.rs
```

1. 一開始編譯器會尋找 crate 的源頭（函式庫 crate 的話，通常就是 `src/lib.rs`；執行檔 crate 的話，通常就是 `src/main.rs`）來編譯程式碼。
2. 在 crate 的入口檔案會需要**宣告**將使用的模組，不是想用就用，你還得先宣告！比方說你在 main.rs（執行檔 crate 的入口檔案）裡想使用 garden 模組 `mod garden`。編譯器會去以下幾個地方尋找 garden 模組的程式碼：
    - 同檔案內用 `mod garden` 加上大括號，寫在括號內的程式碼
    - `src/garden.rs` 檔案中
    - `src/garden/mod.rs` 檔案中
3. 如果想要在 garden 模組下再宣告子模組的話，舉個例來說，可以在 garden.rs 檔案中加上 `mod vegetables;`。編譯器讀到這一行後就會去與當前模組(garden)同名的資料夾下這幾處尋找子模組的程式碼：
    - 同檔案內，直接用 `mod vegetables` 加上大括號，寫在括號內的程式碼
    - `src/garden/vegetables.rs` 檔案中
    - `src/garden/vegetables/mod.rs` 檔案中
4. 一但模組成為 crate 的一部分，就可以在任何地方使用該模組（假設都是 `pub` 的話）。舉例來說，如果現在想要用 vegetable 模組裡面的 Asparagus 型別，就可以用 `crate::garden::vegetables::Asparagus` 來找到。(最前面的 crate 代表從源頭開始)
 


# References

[7.2 定義模組來控制作用域與隱私權
](https://rust-lang.tw/book-tw/ch07-02-defining-modules-to-control-scope-and-privacy.html)

---
此筆記僅是為了紀錄學習過程，如欲完整學習 rust，請去看官方提供的文件，那裡有更加完善的學習資源。筆記內容為自己的理解因此可能存在錯誤，如有錯誤敬請見諒。