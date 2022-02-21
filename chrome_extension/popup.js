let page = document.getElementById("buttonDiv");
let button = document.createElement("button");
let cfgURLInput = document.createElement("input");
let cfgMatchInput = document.createElement("input");

button.innerText = "Submit";

cfgURLInput.placeholder = "http://localhost:8099/chupdate"
cfgMatchInput.placeholder = "string to match"

function handleButtonClick(event) {
    setURL(cfgURLInput.value, cfgMatchInput.value)
    cfgURLInput.value = ""
    cfgMatchInput.value = ""
}

function setURL(url, matchString) {
    chrome.runtime.sendMessage({
        type: "UpdateCfg",
        URL: url,
        Match: matchString
    });
}

page.appendChild(cfgURLInput)
page.appendChild(cfgMatchInput)
page.appendChild(button)
button.addEventListener("click", handleButtonClick);