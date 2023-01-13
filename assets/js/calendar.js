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