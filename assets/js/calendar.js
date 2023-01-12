// FUNCTIONS:

// Closes all open calendar items when another one is clicked:
function removeVisibleClass(current) {
    document.querySelectorAll(".schedule-item").forEach(container => {
        if (container != current) {
            container.querySelector(".content").classList.remove("visible")
        }
    })
}

// Toggles the change of direction of calendar item divs:
function toggleArrowDir(current) {
    let currentArrow = current.querySelector(".down-arrow") 
    currentArrow.innerHTML = currentArrow.innerHTML == "expand_more" ? "expand_less" : "expand_more"
}

document.querySelectorAll(".schedule-item").forEach(container => {
    container.addEventListener("click", function() {
            removeVisibleClass(container)            
        container.querySelector(".content").classList.toggle("visible")
        toggleArrowDir(container)
    })
})

// Shortens "Grand Prix" to "GP" when screen width is small, and vice-versa:
function shortenGPName() {
    document.querySelectorAll(".schedule-item").forEach(item => {
        let titleDiv = item.querySelector(".title")
        if (window.innerWidth <= 950 && /Grand Prix/i.test(titleDiv.innerText)) {
            titleDiv.innerText = titleDiv.innerText.replace(/Grand Prix/i, "GP")
        } else if (window.innerWidth > 950 && /GP/i.test(titleDiv.innerText)) {
            titleDiv.innerText = titleDiv.innerText.replace(/GP/i, "Grand Prix")
        }
    })
}
// Runs shortenGPName() when window is resized:
window.addEventListener("resize", function() {
    shortenGPName()
})
shortenGPName()

// Converts times from a schedule divs (UTC) to user local time:
function stringToLocalTime(date, time) {
    dateArr = date.split(" "); timeArr = time.split(":");
    let temp_date = new Date(Date.UTC(dateArr[2], MONTHS.indexOf(dateArr[1]), dateArr[0], timeArr[0], timeArr[1], 0))
    return convertTimezone(temp_date, USER_TZ)
}

// Converts all available schedule divs to user local time:
function convertAllTimes(){
    document.querySelectorAll(".datetime").forEach(datetime => {
        let temp = stringToLocalTime(datetime.querySelector(".date").innerText, datetime.querySelector(".time").innerText)

        datetime.querySelector(".date").innerText = `${padStart(temp.getDate())} ${MONTHS[temp.getMonth()]} ${temp.getFullYear()}` 
        datetime.querySelector(".time").innerText = `${padStart(temp.getHours())}:${padStart(temp.getMinutes())}` 
    })
}
convertAllTimes();

// Fills the user tz div:
if ($("#tz").length) {
    document.getElementById("tz").innerText = getTimezoneOffset()
}