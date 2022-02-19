
function handleTabUpdated(tabId, changeInfo, tab) {
    meetActive = false
    if (changeInfo.status == 'complete') {
        doQuery()
    }
};

function doQuery() {
    meetActive = false
    chrome.tabs.query({}).then((t) => {
        for (i = 0; i < t.length; i++) {
            let url = t[i].url
            if (url.includes("meet.google.com")) {
                meetActive = true
                // fetch("https://google.com/echo")
                //     .then(response => response.text())
                //     .then(data => console.log(data))
            };
        }
        console.log("meetActive: ", meetActive)
    })

}

// chrome.tabs.onCreated.addListener(updateStatus);
// chrome.tabs.onUpdated.addListener(handleTabUpdated);
chrome.tabs.onActivated.addListener(doQuery);
