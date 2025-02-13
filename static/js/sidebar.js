function toggleSidebar() {
    let sidebar = document.querySelector(".sidebar");
    let hamburger = document.querySelector(".hamburger");
    
    sidebar.classList.toggle("active");
    hamburger.classList.toggle("active"); // 讓漢堡變成 X
}

function toggleDropdown() {
    let dropdown = document.querySelector(".dropdown");
    dropdown.classList.toggle("active");
}
