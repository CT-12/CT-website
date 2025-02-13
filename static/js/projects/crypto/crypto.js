document.addEventListener("DOMContentLoaded", function () {
    // 初始化 Bootstrap popover
    let popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'));
    popoverTriggerList.map(function (popoverTriggerEl) {
        return new bootstrap.Popover(popoverTriggerEl);
    });
});

// SHA256 Hash Function
function hashSHA256(password) {
    return CryptoJS.SHA256(password).toString();
}

// AES Encryption
function aesEncrypt(plainText, secretKey) {
    return CryptoJS.AES.encrypt(plainText, secretKey).toString();
}

// AES Decryption
function aesDecrypt(encryptedText, secretKey) {
    let bytes = CryptoJS.AES.decrypt(encryptedText, secretKey);
    return bytes.toString(CryptoJS.enc.Utf8);
}

// 加密函數
function encryptText() {
    let password = document.getElementById("password").value;
    let plainText = document.getElementById("input-text").value;

    if (!password || !plainText) {
        alert("請輸入密碼和明文內容");
        return;
    }

    let secretKey = hashSHA256(password);
    let encryptedText = aesEncrypt(plainText, secretKey);

    document.getElementById("encrypted-text").value = encryptedText;
}

// 解密函數
function decryptText() {
    let password = document.getElementById("password-decrypt").value;
    let encryptedText = document.getElementById("decrypt-input").value;

    if (!password || !encryptedText) {
        alert("請輸入密碼和加密內容");
        return;
    }

    let secretKey = hashSHA256(password);
    let decryptedText = aesDecrypt(encryptedText, secretKey);

    if (!decryptedText) {
        alert("解密失敗，請確認密碼是否正確");
        return;
    }

    document.getElementById("decrypted-text").value = decryptedText;
}

// 複製加密文字
function copyText() {
    let encryptedText = document.getElementById("encrypted-text").value;
    if (!encryptedText) {
        alert("沒有可複製的加密內容");
        return;
    }
    navigator.clipboard.writeText(encryptedText).then(() => {
        alert("已複製加密內容！");
    });
}

// 清除所有輸入內容
function clearAll() {
    document.getElementById("input-text").value = "";
    document.getElementById("password").value = "";
    document.getElementById("encrypted-text").value = "";
    document.getElementById("decrypt-input").value = "";
    document.getElementById("password-decrypt").value = "";
    document.getElementById("decrypted-text").value = "";
}

// 讀取檔案
document.getElementById("fileInput").addEventListener("change", function (event) {
    // event.target.files 是 input[type="file"] 元素的 files 屬性，它是一個 FileList 物件，包含使用者選擇的所有檔案。
    // files[0] 代表使用者選擇的第一個檔案（如果允許多檔案上傳，則 files 會包含多個檔案）。
    // file 變數現在存放的是該檔案的資訊（如名稱、大小、類型等）。
    let file = event.target.files[0];
    if (!file) return;
    // 建立一個 FileReader 物件：
    // FileReader 是瀏覽器內建的 API，可以讀取使用者選擇的檔案（如文字、圖片、影片）。
    // 它允許我們讀取檔案的內容，然後以不同格式（純文字、Data URL、ArrayBuffer 等）返回。
    let reader = new FileReader();
    // 設定 onload 事件監聽器：
    // 當 FileReader 成功讀取檔案後，它會觸發 onload 事件。
    // e.target.result 包含了檔案的內容（以純文字格式）。
    // 這段程式碼的作用：將檔案的內容填入 <textarea> 元素，讓使用者可以看到檔案內容。
    reader.onload = function (e) {
        document.getElementById("input-text").value = e.target.result;
    };

    // 開始讀取檔案內容：
    // readAsText(file) 會以 純文字（UTF-8 編碼） 的格式讀取檔案，並觸發 onload 事件。
    // 適用於 .txt、.csv、.json 等純文字格式檔案。
    reader.readAsText(file);
});
