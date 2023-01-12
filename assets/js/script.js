const NEXT_RACE = convertTimezone(new Date(Date.UTC(2023, 0, 15, 23, 0, 0))).getTime()
const MONTHS = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEPT", "OCT", "NOV", "DEC"]
const USER_TZ = getUserTimezone()

// Returns full name user TZ ("Europe/Paris" for ex):
function getUserTimezone() {
    return Intl.DateTimeFormat().resolvedOptions().timeZone
}

// Return short name user TZ ("GMT+1" for ex):
function getTimezoneOffset() {
    return new Date().toLocaleTimeString('en-us',{timeZoneName:'short'}).split(' ')[2]
}

// Converts UTC date object to a given TZ:
function convertTimezone(timeUTC, timezone) {
    return new Date((typeof timeUTC === "string" ? new timeUTC(timeUTC) : timeUTC).toLocaleString("en-US", {timeZone: timezone}));
}

// Adds a given number of "0" before an int converted to a string (default=2):
function padStart(number, pad=2) {
    return String(number).padStart(pad, '0');
}


// *RESPONSIVE NAVBAR:
document.querySelector(".icon-menu").addEventListener("click", function() { 
    document.querySelector(".nav-links").classList.toggle("visible")
})

window.addEventListener("resize", function() {
    if (this.innerWidth > 950) {
        this.document.querySelector(".nav-links").classList.remove("visible")

    } 
})

// *------------------ //

//! RESULTS.HTML (technically useless so far):
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
