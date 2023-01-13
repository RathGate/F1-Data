const CONSTRUCTOR_COLORS = {
    "red_bull": "var(--color-rb)",
    "ferrari": "var(--color-ferrari)",
    "mercedes": "var(--color-mercedes)",
    "mclaren": "var(--color-mclaren)",
    "alpine_f1_team": "var(--color-alpine)",
    "aston_martin": "var(--color-aston)",
    "alphatauri": "var(--color-alphatauri)",
    "alfa_romeo": "var(--color-alfa)",
    "williams": "var(--color-williams)",
    "haas_f1_team": "var(--color-haas)",
}

document.querySelectorAll(".race-select-item").forEach(element => {

        element.addEventListener("click", function() {
            if (!element.classList.contains("active")) {
            $.ajax({
                type: "POST",
                url: "/resultats",
                data: { "requestedYear": element.id },
                success: function(data) {
                    removeClass(".race-select-item", "active")
                    element.classList.add("active")
                    document.querySelector("#content").innerHTML = data
                    shortenGPName()
                }
            })
        }
        })
    }
)

function removeClass(selector, classname) {
    document.querySelectorAll(selector).forEach(container => {
        container.classList.remove(classname)
    })
}

$(document).ready(function(){
    if (!$("#race-select-container").length || !$('.active')) {
        return
    }
    $("#race-select-container").animate({ 
        scrollLeft: $('.active').offset().left 
    }, 2500, "swing");
});