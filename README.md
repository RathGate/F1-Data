# F1 Data

![Landing page of the website](https://i.postimg.cc/4yLXQK5Z/1.png)

## About

Project submitted for a front-end course, realised by [Malcom Kojok Tarraf](https://github.com/KhaledTarraf) and [RathGate](https://github.com/RathGate). As it is something that interests us both, we decided to work on our own version of the [Formula 1 Official Website](https://www.formula1.com/), displaying the content and data we, as casual F1 fans, would like to easily find before and after the races.

Technical documentation: [HERE (GDrive)](https://docs.google.com/document/d/1RphHnO_cOwu32hdEOIRDSWFri1YW21bF1sdot0vsidE/edit?usp=share_link)

How we organised our work: [HERE (Monday)](https://drive.google.com/file/d/1lpvbXwP1t2FPQ5lWueLt15Syn36XQ-qE/view?usp=share_link) & [HERE (Trello)](https://drive.google.com/file/d/1ZZOogSOiHAmlau366zSUJhc7TAC2B2BA/view?usp=share_link)

## Technical Specifications

 - Back-end: Golang 
 - Front-end: HTML, CSS, JS (mostly used to send minor
   ajax requests that would not require to reload the entire page).
 - API used: [Ergast Formula One API](http://ergast.com/mrd/)

For now, the program does not exactly use real-time API calls due to random shutdowns and overall instability of the Ergast Database (at least in December 2022). A temporary solution has been to save the most important part of the data into static .json when the API was working - therefore, the dynamic API calls have been temporarily commented and the fallback functions using the static files are always used instead.

**COMPATIBILITY:** The website has been entirely tested on the latest versions of Chrome, Firefox and Edge. It should also appear responsive on Chrome, Safari and Firefox mobile browsers !

## How to use the program

For now, the website has not been hosted anywhere, though we might find a way to host it in the future. In order to visit it on your computer, you must run the go server on your machine. 

To clone the repository:

    git clone https://github.com/RathGate/F1-Data.git

Then, launch a terminal in the `/src` folder:

    go run .

**NOTE:** `go run main.go` will not work properly as the `package main` is divided into three separate files and not just contained in `main.go`.

The website should be available on `localhost:8080`. In case of port collision, look for this line of code at the bottom of `/src/main.go` :

    preferredPort :=  ":8080"

Change the numerical value after the colon with any other port number.

Enjoy ♫

![Website Screenshot(results page)](https://i.postimg.cc/4yQRWYQn/Capture-d-cran-2023-01-14-013615.png)
