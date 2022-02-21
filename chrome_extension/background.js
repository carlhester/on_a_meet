var UpdateURL = "http://localhost:8099/chupdate"
var UpdateSuffix = "?key="
var MatchURL = "meet.google.com"

chrome.runtime.onMessage.addListener(function (request, sender, sendResponse) {
    if (request.type === "UpdateCfg") {
        UpdateURL = request.URL
        MatchURL = request.Match
    }
})

function handleTabUpdated(tabId, changeInfo, tab) {
    meetActive = false
    if (changeInfo.status == 'complete') {
        doQuery()
    }
    return
};

function doQuery() {

    meetActive = false
    chrome.tabs.query({}).then((t) => {
        for (i = 0; i < t.length; i++) {
            let url = t[i].url
            if (url.includes(MatchURL)) {
                meetActive = true
            };
        }
        fetch(UpdateURL + UpdateSuffix + meetActive, {
            cache: 'no-cache',
        })
    })

}

// chrome.tabs.onCreated.addListener(updateStatus);
// chrome.tabs.onUpdated.addListener(handleTabUpdated);
chrome.tabs.onActivated.addListener(doQuery);
