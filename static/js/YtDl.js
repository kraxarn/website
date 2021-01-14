class YtDl {
	constructor() {
		this.download = getById("download")
		this.downloading = false
		this.download.onclick = this.downloadClicked
	}

	showError(msg) {
		const error = getById("error")
		error.textContent = msg
		error.style.display = msg ? "block" : "none"
		if (msg) {
			getById("download").value = "Try again"
		}
		this.downloading = false
	}

	setDownloading(isDownloading) {
		this.downloading = isDownloading
		const download = getById("download")
		download.value = isDownloading ? "Please wait..." : "Downloading..."
	}

	downloadClicked() {
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

	getVideoId() {
		const url = getById("url").value
		if (url.startsWith("https://youtu.be/") && url.length === 28) {
			return url.substring(17)
		}
		if (url.startsWith("https://www.youtube.com/watch?v=") && url.length === 43) {
			return url.substring(32)
		}
		return null
	}

	downloadFile(url, filename) {
		const link = document.createElement("a")
		link.download = filename
		link.href = url
		link.style.display = "block"
		link.click()
		window.onload = () => this.setDownloading(false)
	}
}

new YtDl()
