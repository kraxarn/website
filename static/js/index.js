let showDebug = false
const debug = document.getElementById('debug')
const changelog = document.getElementById('changelog')
const browserInfo = document.getElementById('browserInfo')
const text = info

/* Browser checking */
let support
switch(browser) {
	case "Chrome":
	case "Firefox":
	case "Vivaldi":
	case "Opera":
		support = "full support."
		break
	case "Internet Explorer":
		support = "partial support, no music, video, background or fonts."
		break
	case "Edge":
	case "Safari":
	case "Silk":
		support = "partial support, no music, video or background."
		break;
	default:
		support = "unknown support."
		break;
}
if (browser === "Firefox" || browser === "Chrome" || browser === "Vivaldi") {
	support = "full support."
}

document.getElementById('browser').textContent = `You're running ${browser} ${browserVer} with ${support}`
document.getElementById('debugUserAgent').textContent = navigator.userAgent;

document.getElementById('debugDoNotTrack').textContent = `Do not track: ${navigator.doNotTrack === "1" ? "Enabled" : "Disabled"}`

function toggleDebug() {
	if (showDebug) {
		changelog.style.transform = 'translate(100%, 0)'
		browserInfo.style.transform = 'translate(-50%, -150%)'
		showDebug = false
	} else {
		changelog.style.transform = 'translate(-25%, 0)'
		browserInfo.style.transform = 'translate(-50%, 100%)'
		showDebug = true
	}
}

if (!supportOgg || !supportWebm) {
	document.getElementById('warningBrowser').style.display = "block"
}

function toggleLogo()  {
	const logo = document.getElementById("logo")
	if (logo.src.includes("logo_v7")) {
		// Using new logo, switch to old one
		logo.src = "img/logo_v5.webp"
	} else {
		// Using old logo, switch to new one
		logo.src = "img/logo_v7_lightblue.webp"
	}
}

function toggleOlderChanges() {
	const olderChanges = document.getElementById("olderChanges")
	const showChanges  = document.getElementById("showChanges")
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