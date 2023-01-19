if ($(".countdown").length) {
    var x = setInterval(function() {
        var now = new Date().getTime()
        var remainingTime = NEXT_RACE - now;
    
        var days = Math.floor(remainingTime / (1000 * 60 * 60 * 24));
        var hours = Math.floor((remainingTime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        var minutes = Math.floor((remainingTime % (1000 * 60 * 60)) / (1000 * 60));
        var seconds = Math.floor((remainingTime % (1000 * 60)) / 1000);
    
        document.getElementById("days").innerText = padStart(days);
        document.getElementById("hours").innerText = padStart(hours);
        document.getElementById("minutes").innerText = padStart(minutes);
        document.getElementById("seconds").innerText = padStart(seconds);
    
        if (remainingTime < 0) {
            clearInterval(x);
            document.querySelector(".countdown").innerHTML = "<img src='assets/img/drivers/rick.gif' class='rick' alt='Nice try :)'>";
        }
    }, 1000);
}


// Resizes the timer depending on the screen width:
function resizeTimer() {
    if ($(".time-block").length) {
        var blockWidth = this.innerWidth / 4 > 200 ? 200 : this.innerWidth / 4;
        var fontSize = blockWidth * 0.6
        var smallFontSize = blockWidth * 0.125
        this.document.querySelectorAll(".time-block").forEach(timeBlock => {
            timeBlock.style.fontSize = `${fontSize}px`;
            timeBlock.querySelector(".unit").style.fontSize = `${smallFontSize}px`;
        })
    }
}

resizeTimer()

window.addEventListener("resize", function() { 
    resizeTimer()
})