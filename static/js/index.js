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
		getChanges()
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

const getChanges = () => {
	const latest = getById("latestChanges")
	const older = getById("olderChanges")

	fetch("/changes")
		.then(response => response.json())
		.then(json => json.forEach((item, i) => (i === 0 ? latest : older).innerHTML = createChanges(item)))
		.catch(err => latest.textContent = err)
}

const createChanges = json => {
	const title = document.createElement("span")
	title.className = "changelogTitle"
	title.textContent = json.name

	const ul = document.createElement("ul")
	json["changes"].forEach(item => {
		const li = document.createElement("li")
		li.textContent = item
		ul.appendChild(li)
	})

	const div = document.createElement("div")
	div.appendChild(title)
	div.appendChild(ul)
	return div.outerHTML
}