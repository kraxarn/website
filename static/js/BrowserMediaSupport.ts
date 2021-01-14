class BrowserMediaSupport {
	audio: HTMLAudioElement
	video: HTMLAudioElement

	constructor() {
		this.audio = document.createElement("audio")
		this.video = document.createElement("video")
	}

	get ogg(): boolean {
		return !!this.audio.canPlayType("audio/ogg")
	}

	get webm(): boolean {
		return !!this.video.canPlayType("video/webm")
	}

	get opus(): boolean {
		return !!this.audio.canPlayType("audio/opus")
	}

	get flac(): boolean {
		return !!this.audio.canPlayType("audio/x-flac")
	}
}