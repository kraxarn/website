import "./main"
import "./BrowserInfo"
import "./BrowserMediaSupport"

class Index {
	showDebug = false
	changelog = getById("changelog")
	browserInfo = getById("browserInfo")

	constructor() {
		getById("debugUserAgent").textContent = navigator.userAgent
		getById("debugDoNotTrack").textContent = `Do not track: ${navigator.doNotTrack === "1"
			? "Enabled"
			: "Disabled"}`

		this.setBrowserInfo()
		this.updateBrowserWarning()
	}

	setBrowserInfo(): void {
		const browserInfo = new BrowserInfo()
		getById("browser").textContent = `You're running ${browserInfo.info}`
	}

	toggleDebug(): void {
		if (this.showDebug) {
			this.changelog.style.transform = "translate(100%, 0)"
			this.browserInfo.style.transform = "translate(-50%, -150%)"
			this.showDebug = false
		} else {
			this.changelog.style.transform = "translate(-25%, 0)"
			this.browserInfo.style.transform = "translate(-50%, 100%)"
			this.showDebug = true
			this.getChanges()
		}
	}

	updateBrowserWarning(): void {
		const supported = new BrowserMediaSupport()
		if (!supported.ogg || !supported.webm) {
			getById("warningBrowser").style.display = "block"
		}
	}

	toggleLogo(): void {
		const logo = getById("logo") as HTMLImageElement
		logo.src = logo.src.includes("logo_v7")
			? "img/logo_v5.webp"
			: "img/logo_v7_lightblue.webp"
	}

	toggleOlderChanges(): void {
		const olderChanges = getById("olderChanges")
		const showChanges = getById("showChanges")
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

	getChanges(): void {
		const latest = getById("latestChanges")
		const older = getById("olderChanges")

		fetch("/changes")
			.then(response => response.json())
			.then(json => json.forEach((item, i) => (i === 0 ? latest : older).innerHTML = this.createChanges(item)))
			.catch(err => latest.textContent = err)
	}

	createChanges(json: object): string {
		const title = document.createElement("span")
		title.className = "changelogTitle"
		title.textContent = json["name"]

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
}

new Index()