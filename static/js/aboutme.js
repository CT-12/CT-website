class Carousel {
    constructor(containerId) {
        this.carouselContainer = document.getElementById(containerId);
        this.carousel = this.carouselContainer.querySelector(".carousel");
        this.cards = this.carousel.querySelectorAll(".card");
        this.prevButton = this.carouselContainer.querySelector(".prev");
        this.nextButton = this.carouselContainer.querySelector(".next");

        this.init();
    }

    init() {
        // 綁定按鈕事件
        this.nextButton.addEventListener("click", () => this.nextSlide());
        this.prevButton.addEventListener("click", () => this.prevSlide());
    }

    nextSlide() {
        let firstCard = this.carousel.firstElementChild;

        // 執行滑動動畫
        this.carousel.style.transition = "transform 0.5s ease-in-out";
        this.carousel.style.transform = "translateX(-220px)";

        setTimeout(() => {
            // 把第一張移到最後
            this.carousel.appendChild(firstCard);
            // 移除動畫效果，避免瞬間跳動
            this.carousel.style.transition = "none";
            this.carousel.style.transform = "translateX(0)";
        }, 500);
    }

    prevSlide() {
        let lastCard = this.carousel.lastElementChild;

        // 立即將最後一張卡片移到最前面
        this.carousel.insertBefore(lastCard, this.carousel.firstElementChild);
        
        // 瞬間設置負偏移，讓它看起來還在原來的位置
        this.carousel.style.transition = "none";
        this.carousel.style.transform = "translateX(-220px)";

        setTimeout(() => {
            // 加入動畫效果
            this.carousel.style.transition = "transform 0.5s ease-in-out";
            this.carousel.style.transform = "translateX(0)";
        }, 10);
    }
}

// 🚀 初始化多個輪播
const educationCarousel = new Carousel("education-carousel");
const galleryCarousel = new Carousel("gallery-carousel");
