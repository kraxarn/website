let showDebug = false
const changelog = getById("changelog")
const browserInfo = getById("browserInfo")

getById("debugUserAgent").textContent = navigator.userAgent
getById("debugDoNotTrack").textContent = `Do not track: ${navigator.doNotTrack === "1" ? "Enabled" : "Disabled"}`

const setBrowserInfo = () => {
	const browser = getBrowserInfo()
	getById("browser").textContent = `You're running ${browser.name} ${browser.version}`
}
setBrowserInfo()

const toggleDebug = () => {
	if (showDebug) {
		changelog.style.transform = "translate(100%, 0)"
		browserInfo.style.transform = "translate(-50%, -150%)"
		showDebug = false
	} else {
		changelog.style.transform = "translate(-25%, 0)"
		browserInfo.style.transform = "translate(-50%, 100%)"
		showDebug = true
	}
}

const updateBrowserWarning = () => {
	const supported = getBrowserMediaSupport()
	if (!supported.ogg || !supported.webm) {
		getById("warningBrowser").style.display = "block"
	}
}
updateBrowserWarning()

const toggleLogo = () => {
	const logo = getById("logo")
	logo.src = logo.src.includes("logo_v7")
		? "img/logo_v5.webp"
		: "img/logo_v7_lightblue.webp"
}

const toggleOlderChanges = () => {
	const olderChanges = document.getElementById("olderChanges")
	const showChanges = document.getElementById("showChanges")
	if (olderChanges.style.display === "none" || olderChanges.style.display === "") {
		// Hidden, show them
		olderChanges.style.display = "block"
		showChanges.textContent = "Show latest changes"
	} else {
		// Already shown, hide them
		olderChanges.style.display = "none"
		showChanges.textContent = "Show all changes"
	}
}