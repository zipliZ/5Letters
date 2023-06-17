let Row = 0
let CompleteGame = false

function handleKey(letter) {
    var inputs = document.getElementsByTagName('input');
    var lastEmptyInput = null;

    let start = Row * 5
    let end = Row * 5 + 5

    for (var i = start; i < end; i++) {
        if (inputs[i].value === '') {
            lastEmptyInput = inputs[i];
            break;
        }
    }

    if (lastEmptyInput !== null) {
        lastEmptyInput.value = letter;
    }

}
function handleBackSpace() {
    var inputs = document.getElementsByTagName('input');
    var lastNotEmptyInput = null;
    let start = Row * 5
    let end = Row * 5 + 5
    for (var i = end -1; i >= start; i--) {
        if (inputs[i].value !== '') {
            lastNotEmptyInput = inputs[i];
            break;
        }
    }

    if (lastNotEmptyInput !== null) {
        lastNotEmptyInput.value = '';
    }
}
async function handleEnter(){

    var inputs = document.getElementsByTagName('input');
    let start = Row * 5
    let end = Row * 5 + 5
    let word = ''
    for (var i = start; i < end; i++) {
        word+=inputs[i].value
    }
    let responseData

    const res = await fetch("/api/GetMatches/", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        },
        body: JSON.stringify({Word: word})
    })

    if (res.ok){
        let Json = await res.json()
        responseData= Json.Response
    }
    else {
        console.log(res)
    }
    let correctWord
    for (let i = 0; i < 5; i++) {
        let letterColor= responseData[i].split("/")
            if (letterColor[1] === "-1"){
                for (let j = start; j <end; j++) {
                    inputs[j].style.backgroundColor = "red";
                }
                correctWord = false
                break
            }
            if (letterColor[1] === "1"){
                inputs[start+i].style.backgroundColor = "orange";
            }else if (letterColor[1] === "2"){
                inputs[start+i].style.backgroundColor = "green";
            }else {
                inputs[start+i].style.backgroundColor = "white";
            }
            correctWord = true
    }
    for (let r of responseData) {
        let letterInfo = r.split("/")
        if (letterInfo[1]!=="2"){
            CompleteGame = false
            break
        }
        CompleteGame = true

    }
    if (CompleteGame){
        document.querySelector('.popup-WinOverlay').style.display = 'flex';
    }
    if (correctWord){
        Row+=1
    }



}

function reloadPage(){
    window.location.reload()
}