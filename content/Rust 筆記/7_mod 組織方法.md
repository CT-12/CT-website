---
create_at: 2025.07.11
update_at: 2025.07.11
draft: false
tags: 
 - Rust
 - mod
---

假設現在有以下專案結構：

```lua
my_crate/
├── Cargo.toml
├── src/
|   ├── main.rs
│   ├── lib.rs
│   ├── math/
│   │   ├── mod.rs        
│   │   ├── arithmetic.rs <-- 子模組
│   │   └── algebra.rs    <-- 子模組
│   └── utils/
│       └── mod.rs
├── tests/
│   └── integration_test.rs

```

在 `lib.rs` 裡面要先暴露 mod，所以在 `lib.rs` 裡面會這樣：
```rust
pub mod math;
pub mod utils;
```

在這裡的 `mod math`，就是在告訴 rust 編譯器有一個 mod 叫 math，要記得編譯它！然後 rust 編譯器就會去以下三個地方找這個 mod：

1. 同個文件中（這裡是 `lib.rs`），有沒有定義 `mod math{ ... }`
2. 同個目錄下（這裡是 `src/`），有沒有一個文件名叫做 `math.rs`
3. 同個目錄下（這裡是 `src/`），有沒有一個資料夾叫做 `math` 並且底下有一個文件叫做 `mod.rs`（i.e. `src/math/mod.rs`）

只要在其中一個地方有找到就沒問題了！而如果還想要在模組中定義子模組，這很常見，就可以繼續如上圖專案結構所示，先在 `math/mod.rs` 在定義 math 的子模組：

```rust
//
// src/math/mod.rs
//
pub mod arithmetic;
pub mod algebra;

// 可選：你也可以在 math 模組裡提供轉發函數
pub use arithmetic::add;
pub use algebra::square;
```

然後將 mod 暴露出來，而 arithmatic 和 algebra 這兩個模組跟 mod.rs 在同一層目錄中，所以 rust 編譯器能夠順利找到他們。至於要不要轉發函數就是個人選擇。轉發函數大概像這樣：

```rust
//
// src/math/arithmetic.rs
//
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

pub fn subtract(a: i32, b: i32) -> i32 {
    a - b
}

//
// src/math/algebra.rs
//
pub fn square(x: i32) -> i32 {
    x * x
}

```
```rust
//
// src/main.rs
//
use my_crate::math::add; // 可以這樣使用，原本要這樣使用 -> use my_crate::math::arithmetic::add

```