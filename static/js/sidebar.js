function toggleSidebar() {
    let sidebar = document.querySelector(".sidebar");
    let hamburger = document.querySelector(".hamburger");
    
    sidebar.classList.toggle("active");
    hamburger.classList.toggle("active"); // 讓漢堡變成 X

    // 滑到回到頁面頂部
    window.scrollTo(
        {
            top: 0,
            behavior: "smooth" // 平滑滾動
        }
    )
}
