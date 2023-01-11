const date = new Date(Date.UTC(2023, 0, 15, 23, 0, 0))
const USER_TZ = getUserTimezone()

function getUserTimezone() {
    return Intl.DateTimeFormat().resolvedOptions().timeZone
}

function convertTimezone(timeUTC, timezone) {
    return new Date((typeof timeUTC === "string" ? new timeUTC(timeUTC) : timeUTC).toLocaleString("en-US", {timeZone: timezone}));
}

var countdownDate = convertTimezone(date).getTime()
if ($(".countdown").length) {
    var x = setInterval(function() {
        var now = new Date().getTime()
        var remainingTime = countdownDate - now;
    
        var days = Math.floor(remainingTime / (1000 * 60 * 60 * 24));
        var hours = Math.floor((remainingTime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        var minutes = Math.floor((remainingTime % (1000 * 60 * 60)) / (1000 * 60));
        var seconds = Math.floor((remainingTime % (1000 * 60)) / 1000);
    
        document.getElementById("days").innerText = checkZeros(days);
        document.getElementById("hours").innerText = checkZeros(hours);
        document.getElementById("minutes").innerText = checkZeros(minutes);
        document.getElementById("seconds").innerText = checkZeros(seconds);
    
        if (remainingTime < 0) {
            clearInterval(x);
            document.getElementById("demo").innerHTML = "EXPIRED";
        }
    }, 1000);
}


function checkZeros(number) {
    return String(number).padStart(2, '0');
}



function yearWatcher() {
    if (!$("#year-select").length) {
        return
    }
    !$("#year-select").on("change", function() {
        $.ajax({
            type: "POST",
            url: "/resultats",
            data: { "requestedYear": this.value },
        })
    })
}

yearWatcher()

document.querySelectorAll(".schedule-item").forEach(container => {
    container.addEventListener("click", function() {
            removeVisible(container)            
        container.querySelector(".content").classList.toggle("visible")
        toggleArrowDir(container)
    })
})

function removeVisible(current) {
    document.querySelectorAll(".schedule-item").forEach(container => {
        if (container != current) {
            container.querySelector(".content").classList.remove("visible")
        }
    })
}

function toggleArrowDir(current) {
    let currentArrow = current.querySelector(".down-arrow") 
    currentArrow.innerHTML = currentArrow.innerHTML == "expand_more" ? "expand_less" : "expand_more"
}

document.querySelector(".icon-menu").addEventListener("click", function() { 
    document.querySelector(".nav-links").classList.toggle("visible")
})

window.addEventListener("resize", function() {
    if (this.innerWidth > 950) {
        this.document.querySelector(".nav-links").classList.remove("visible")

    } 
    checkGPlength()
    resizeTimer()
})

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
function checkGPlength() {
    if (!$(".schedule-item").length) {
        return
    }
    document.querySelectorAll(".schedule-item").forEach(item => {
        let titleDiv = item.querySelector(".title")
        if (window.innerWidth <= 950 && /Grand Prix/i.test(titleDiv.innerText)) {
            titleDiv.innerText = titleDiv.innerText.replace(/Grand Prix/i, "GP")
        } else if (window.innerWidth > 950 && /GP/i.test(titleDiv.innerText)) {
            titleDiv.innerText = titleDiv.innerText.replace(/GP/i, "Grand Prix")
        }
    })
}

const MONTHS = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEPT", "OCT", "NOV", "DEC"]

if ($("#test").length) {
    document.getElementById("tz").innerText = getTimezoneOffset()
}

function getTimezoneOffset() {
    return new Date().toLocaleTimeString('en-us',{timeZoneName:'short'}).split(' ')[2]
}

function stringToLocalTime(date, time) {
    dateArr = date.split(" "); timeArr = time.split(":");
    let temp_date = new Date(Date.UTC(dateArr[2], MONTHS.indexOf(dateArr[1]), dateArr[0], timeArr[0], timeArr[1], 0))
    return convertTimezone(temp_date, USER_TZ)
}

function convertAllTimes(){
    document.querySelectorAll(".datetime").forEach(datetime => {
        let temp = stringToLocalTime(datetime.querySelector(".date").innerText, datetime.querySelector(".time").innerText)
        editDatetime(datetime, temp)
    })
}

function editDatetime(datetime, date) {
    datetime.querySelector(".date").innerText = `${checkZeros(date.getDate())} ${MONTHS[date.getMonth()]} ${date.getFullYear()}` 
    datetime.querySelector(".time").innerText = `${checkZeros(date.getHours())}:${checkZeros(date.getMinutes())}` 
}

convertAllTimes()
checkGPlength()

function dynamicCountdownResize() {

}