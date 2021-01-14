class BrowserInfo {
	constructor() {
		// Name
		let browser
		const userAgent = navigator.userAgent

		if (userAgent.includes("Firefox")) {
			browser = "Firefox"
		} else if (userAgent.includes("Trident")) {
			browser = "rv"
		} else if (userAgent.includes("Safari")) {
			if (userAgent.includes("Edge")) {
				browser = "Edge"
			} else if (userAgent.includes("OPR")) {
				browser = "OPR"
			} else if (userAgent.includes("Vivaldi")) {
				browser = "Vivaldi"
			} else if (userAgent.includes("Chrome")) {
				browser = "Chrome"
			} else {
				browser = "Version"
			}
		} else if (userAgent.includes("Silk")) {
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
		this.name = browser
		this.version = browserVer
	}

	get info() {
		return `${this.name} ${this.version}`
	}
}
