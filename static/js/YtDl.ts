import "./main"

class YtDl {
	download = getById("download")
	downloading: boolean = false

	constructor() {
		this.download.onclick = this.downloadClicked
	}

	showError(msg: string): void {
		const error = getById("error")
		error.textContent = msg
		error.style.display = msg ? "block" : "none"
		if (msg) {
			(getById("download") as HTMLInputElement).value = "Try again"
		}
		this.downloading = false
	}

	setDownloading(isDownloading) {
		this.downloading = isDownloading as boolean
		const download = getById("download") as HTMLInputElement
		download.value = isDownloading ? "Please wait..." : "Downloading..."
	}

	downloadClicked(): void {
		if (this.downloading) {
			return
		}
		this.setDownloading(true)

		this.showError(null)

		const videoId = this.getVideoId()
		if (!videoId) {
			this.showError("Enter a valid YouTube URL first")
			return
		}

		fetch(`/yt/info/${videoId}`)
			.then(response => response.json())
			.then(json => {
				if (json.error) {
					this.showError(json.error)
					return
				}
				if (!json.audio.url) {
					this.showError("No audio found to download")
					return
				}
				this.downloadFile(`/yt/audio/${videoId}`, `${json.title}.opus`)
			})
			.catch(err => this.showError(err))
	}

	getVideoId(): string {
		const url = (getById("url") as HTMLInputElement).value

		if (url.startsWith("https://youtu.be/") && url.length === 28) {
			return url.substring(17)
		}

		if (url.startsWith("https://www.youtube.com/watch?v=") && url.length === 43) {
			return url.substring(32)
		}

		return null
	}

	downloadFile(url: string, filename: string): void {
		const link = document.createElement("a")
		link.download = filename
		link.href = url
		link.style.display = "block"
		link.click()
		window.onload = () => this.setDownloading(false)
	}
}

new YtDl()