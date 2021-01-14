class BrowserMediaSupport {
	constructor() {
		this.audio = document.createElement("audio")
		this.video = document.createElement("video")
	}

	get ogg() {
		return !!this.audio.canPlayType("audio/ogg")
	}

	get webm() {
		return !!this.video.canPlayType("video/webm")
	}

	get opus() {
		return !!this.audio.canPlayType("audio/opus")
	}

	get flac() {
		return !!this.audio.canPlayType("audio/x-flac")
	}
}
