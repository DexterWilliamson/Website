/*!
* Start Bootstrap - Stylish Portfolio v6.0.6 (https://startbootstrap.com/theme/stylish-portfolio)
* Copyright 2013-2023 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-stylish-portfolio/blob/master/LICENSE)
*/
window.addEventListener('DOMContentLoaded', event => {

    const sidebarWrapper = document.getElementById('sidebar-wrapper');
    let scrollToTopVisible = false;
    // Closes the sidebar menu
    const menuToggle = document.body.querySelector('.menu-toggle');
    menuToggle.addEventListener('click', event => {
        event.preventDefault();
        sidebarWrapper.classList.toggle('active');
        _toggleMenuIcon();
        menuToggle.classList.toggle('active');
    })

    // Closes responsive menu when a scroll trigger link is clicked
    var scrollTriggerList = [].slice.call(document.querySelectorAll('#sidebar-wrapper .js-scroll-trigger'));
    scrollTriggerList.map(scrollTrigger => {
        scrollTrigger.addEventListener('click', () => {
            sidebarWrapper.classList.remove('active');
            menuToggle.classList.remove('active');
            _toggleMenuIcon();
        })
    });

    function _toggleMenuIcon() {
        const menuToggleBars = document.body.querySelector('.menu-toggle > .fa-bars');
        const menuToggleTimes = document.body.querySelector('.menu-toggle > .fa-xmark');
        if (menuToggleBars) {
            menuToggleBars.classList.remove('fa-bars');
            menuToggleBars.classList.add('fa-xmark');
        }
        if (menuToggleTimes) {
            menuToggleTimes.classList.remove('fa-xmark');
            menuToggleTimes.classList.add('fa-bars');
        }
    }

    // Scroll to top button appear
    document.addEventListener('scroll', () => {
        const scrollToTop = document.body.querySelector('.scroll-to-top');
        if (document.documentElement.scrollTop > 100) {
            if (!scrollToTopVisible) {
                fadeIn(scrollToTop);
                scrollToTopVisible = true;
            }
        } else {
            if (scrollToTopVisible) {
                fadeOut(scrollToTop);
                scrollToTopVisible = false;
            }
        }
    })
})

function fadeOut(el) {
    el.style.opacity = 1;
    (function fade() {
        if ((el.style.opacity -= .1) < 0) {
            el.style.display = "none";
        } else {
            requestAnimationFrame(fade);
        }
    })();
};

function fadeIn(el, display) {
    el.style.opacity = 0;
    el.style.display = display || "block";
    (function fade() {
        var val = parseFloat(el.style.opacity);
        if (!((val += .1) > 1)) {
            el.style.opacity = val;
            requestAnimationFrame(fade);
        }
    })();
};

function randomInterval(min, max){
    const maxNum = Math.floor(Math.random() * (max))+min;
    const minNum = Math.floor(Math.random() * (min))+1;
    return Math.floor(Math.random() * (maxNum))+min;
}
  
window.onload = setInterval(function() {
    const parent = document.getElementById("newSVG");
    const svgList = ["/assets/img/zig_zag.svg", "/assets/img/v.svg", 
                     "/assets/img/just_o.svg", "/assets/img/x.svg"];
    const childCount = parent.childElementCount;
    if (childCount < 30){
        for (let i = childCount; i < 30; i++){
            const img = document.createElement("img");
            img.src = svgList[Math.floor(Math.random() * (svgList.length - 1))]; // set the image source
            img.className = "svgSpawn"; // set an Class for the image
            img.style["-webkit-animation-duration"] = randomInterval(7, 15) + "s";
            img.style["-webkit-animation-delay"] = randomInterval(1, 5) + "s";
            parent.appendChild(img);

        };
    };
    if (childCount >= 30){
        childToUpdate = parent.children[Math.floor(Math.random() * (30 - 1))];
        childToUpdate.src= svgList[Math.floor(Math.random() * (svgList.length - 1))];
        childToUpdate.style["-webkit-animation-duration"] = randomInterval(7, 15) + "s";
        childToUpdate.style["-webkit-animation-delay"] = randomInterval(1, 5) + "s";
    };
    
    
  }, 1000);