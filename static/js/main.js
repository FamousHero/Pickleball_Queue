var activeCourtStartTimes = document.querySelectorAll("#active-courts>.group>.start-time>.time")

console.log(activeCourtStartTimes)
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
activeCourtStartTimes.forEach((startTime)=>{
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
    activeCourtStartTimes.forEach((startTime)=>{
        startTime.innerHTML = formatElapsedTime(startTime.innerHTML)
    })
}, 1000)
