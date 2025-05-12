
// Helper Function //
function parseStartTime(startTime){
    let date_time_offset_tz = startTime.split(" ")
    date_time_offset_tz[1] = date_time_offset_tz[1].substring(0, 12)
    let offset = date_time_offset_tz[2]
    offset = offset.substring(0,offset.length-2)+":"+offset.substring(offset.length-2)
    return date_time_offset_tz[0]+"T"+date_time_offset_tz[1]+offset
}

function formatElapsedTime(elapsedTime){
    let min_sec = elapsedTime.split(" ")
    let sec = min_sec[1].slice(0, min_sec[1].length-1)
    let min = min_sec[0].slice(0, min_sec[0].length-1)
    
    if (Number(sec) === 59){
        sec = "0s"
        min = Number(min)+1 + "m"
    }
    else{
        sec = Number(sec)+1 + "s"
        min += "m"
    }
    return min+" "+sec
}


// Page Setup //
function setupSignUp(){
    const playerSection = document.querySelector("#players-section")
    let playerSectionCount = 1
    const addPlayerButton = document.querySelector("#add-player")
    
    function onAddPlayer(){
        const newPlayerInfo = document.createElement("div")
        const newNameLabel = document.createElement("label") 
        const newNameInput = document.createElement("input")

        newPlayerInfo.classList.add("player-info")

        newNameLabel.setAttribute("for", "name-"+playerSectionCount)
        newNameLabel.innerText = "Name: "

        newNameInput.setAttribute("type", "text")
        newNameInput.setAttribute("id", "name"+playerSectionCount)
        newNameInput.setAttribute("name", "name"+playerSectionCount)

        playerSectionCount += 1

        newPlayerInfo.appendChild(newNameLabel)
        newPlayerInfo.appendChild(newNameInput)
        playerSection.insertBefore(newPlayerInfo, addPlayerButton)

        if (playerSectionCount >= 4){
                addPlayerButton.disabled = true
        }
    }

    addPlayerButton.addEventListener("click", onAddPlayer)
}  

function setupCourts(){

    var CourtStartTimes = document.querySelectorAll(".court-info>.start-time>.time")
    console.log(CourtStartTimes)
    CourtStartTimes.forEach((startTime)=>{
        console.log(parseStartTime(startTime.innerHTML))
        let time = new Date(parseStartTime(startTime.innerHTML))
        console.log("Hello World")
        console.log(time)
        let currentTime = new Date()
        console.log(currentTime)
        let timeElapsed = currentTime - time // in milliseconds
        timeElapsed /= 1000
        startTime.innerHTML = Math.floor(timeElapsed/60)+"m "+Math.floor(timeElapsed%60)+"s"
    
    });
    
    setInterval(()=>{
        CourtStartTimes.forEach((startTime)=>{
            startTime.innerHTML = formatElapsedTime(startTime.innerHTML)
        })
    }, 1000)
}

// Send notification game is done, update UI to ask "rematch" or "new court"
function onFinish(){

}
var path = window.location.pathname
console.log(path)
switch(true){
    case /^\/$/.test(path):
        setupSignUp()
        break
    case /^\/queue\/[a-zA-Z\-]+\/?$/.test(path):
    case /^\/playing$/.test(path):
        setupCourts()
        break
}