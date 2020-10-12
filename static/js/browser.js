const getBrowserInfo = () => {
	// Name
	let browser
	const userAgent = navigator.userAgent
	let browserChrome = false

	if (userAgent.indexOf("Firefox") > -1) {
		browser = "Firefox"
	} else if (userAgent.indexOf("Trident") > -1) {
		browser = "rv"
	} else if (userAgent.indexOf("Safari") > -1) {
		if (userAgent.indexOf("Edge") > -1) {
			browser = "Edge"
			browserChrome = true
		} else if (userAgent.indexOf("OPR") > -1) {
			browser = "OPR"
			browserChrome = true
		} else if (userAgent.indexOf("Vivaldi") > -1) {
			browser = "Vivaldi"
			browserChrome = true
		} else if (userAgent.indexOf("Chrome") > -1) {
			browser = "Chrome"
		} else {
			browser = "Version"
		}
	} else if (userAgent.indexOf("Silk") > -1) {
		browser = "Silk"
	} else {
		browser = "Unknown"
	}

	const browserVerPre = userAgent.substring(userAgent.indexOf(browser) + (browser.length + 1))
	const browserVer = browserVerPre.substring(0, browserVerPre.indexOf("."))

	if (browser === "Version") {
		browser = "Safari"
	}
	if (browser === "rv") {
		browser = "Internet Explorer"
	}
	if (browser === "OPR") {
		browser = "Opera"
	}

	return {
		name: browser,
		version: browserVer
	}
}

const getBrowserMediaSupport = () => {
	const audio = document.createElement("audio")
	const video = document.createElement("video")

	return {
		ogg: !!audio.canPlayType("audio/ogg"),
		webm: !!video.canPlayType("video/webm"),
		opus: !!audio.canPlayType("audio/opus"),
		flac: !!audio.canPlayType("audio/x-flac")
	}
}