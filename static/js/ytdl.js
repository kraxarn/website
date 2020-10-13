const download = getById("download")

const showError = msg => {
	const error = getById("error")
	error.textContent = msg
	error.style.display = msg ? "block" : "none"
	if (msg) {
		getById("download").value = "Try again"
	}
}

let downloading = false

const setDownloading = value => {
	downloading = value
	getById("download").value = value
		? "Please wait..."
		: "Downloading..."
}

download.onclick = () => {
	if (downloading) {
		return
	}
	setDownloading(true)

	showError(null)

	const videoId = getVideoId()
	if (!videoId) {
		showError("Enter a valid YouTube URL first")
		return
	}

	fetch(`/yt/info/${videoId}`)
		.then(response => response.json())
		.then(json => {
			if (!json.audio.url) {
				showError("No audio found to download")
				return
			}
			downloadFile(`/yt/audio/${videoId}`, `${json.title}.opus`)
		})
		.catch(err => showError(err))
}

const getVideoId = () => {
	const url = getById("url").value

	if (url.startsWith("https://youtu.be/") && url.length === 28) {
		return url.substring(17)
	}

	if (url.startsWith("https://www.youtube.com/watch?v=") && url.length === 43) {
		return url.substring(32)
	}

	return null
}

const downloadFile = (url, filename) => {
	const link = document.createElement("a")
	link.download = filename
	link.href = url
	link.style.display = "block"
	link.click()
	setDownloading(false)
}