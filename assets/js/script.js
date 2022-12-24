const date = new Date(Date.UTC(2023, 2, 5, 15, 0, 0))

function getUserTimezone() {
    return Intl.DateTimeFormat().resolvedOptions().timeZone
}

function convertTimezone(timeUTC, timezone) {
    return new Date((typeof timeUTC === "string" ? new timeUTC(timeUTC) : timeUTC).toLocaleString("en-US", {timeZone: timezone}));
}

var countdownDate = convertTimezone(date).getTime()

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

function checkZeros(number) {
    return String(number).padStart(2, '0');
}

