interface HitsterTrackInfo {
    id: string
    name: string
    preview_url: string
    artist_name: string
    album_release_date: string
}

enum HitsterGuessResult {
    None = 0,
    Wrong = 1,
    Good = 2,
    Perfect = 3,
}

class Hitster {
    private tracks: HitsterTrackInfo[];
    private score: HitsterGuessResult[];

    constructor() {
        const playlists = document.getElementById("playlists")
        playlists.addEventListener("click", event => this.onPlaylistsClick(event))

        const year = <HTMLInputElement>document.getElementById("year")
        year.addEventListener("input", event => this.onYearInput(event))

        const guess = <HTMLInputElement>document.getElementById("guess")
        guess.addEventListener("click", () => this.onGuessClick())
    }

    private get currentTrack(): HitsterTrackInfo {
        return this.tracks[this.score.length]
    }

    private get currentScore(): number {
        let score = 0;

        for (const result of this.score) {
            if (result === HitsterGuessResult.Good) {
                score += 1
            } else if (result === HitsterGuessResult.Perfect) {
                score += 2
            }
        }

        return score
    }

    private get audioPreview(): HTMLAudioElement {
        const element = document.getElementById("audio-preview")
        return <HTMLAudioElement>element
    }

    private async onPlaylistsClick(event: MouseEvent): Promise<void> {
        const target = <HTMLElement>event.target;

        const playlistId = target.dataset.playlistId;
        if (!playlistId) {
            return
        }

        const playlists = document.getElementById("playlists")
        playlists.style.display = "none"

        const playlistsTitle = <HTMLElement>playlists.previousElementSibling;
        playlistsTitle.style.display = "none"

        const game = document.getElementById("game")
        game.style.display = ""

        this.onYearInput()
        await this.loadSongs(playlistId)
    }

    private onYearInput(event?: Event): void {
        const target = <HTMLInputElement>(event?.target ?? document.getElementById("year"))

        const current = parseInt(target.value)
        const minYear = parseInt(target.min)
        const maxYear = parseInt(target.max)

        const yearLabel = <HTMLLabelElement>document.getElementById("year-label")
        yearLabel.innerText = target.value
        yearLabel.style.left = `${((current - minYear) / (maxYear - minYear)) * 88}%`
    }

    private updateScore() {
        const container = document.getElementById("score-container")
        container.innerText = ""

        for (let i = 0; i < 10; i++) {
            const score = document.createElement("span")
            const guess = this.score[i] ?? HitsterGuessResult.None;

            switch (guess) {
                case HitsterGuessResult.None:
                    score.textContent = "--";
                    break

                case HitsterGuessResult.Wrong:
                    score.textContent = ":(";
                    score.classList.add("wrong")
                    break

                case HitsterGuessResult.Good:
                    score.textContent = ":)"
                    score.classList.add("good")
                    break

                case HitsterGuessResult.Perfect:
                    score.textContent = ":D"
                    score.classList.add("perfect")
                    break

            }

            if (guess !== HitsterGuessResult.None) {
                const track = this.tracks[i]
                const releaseDate = new Date(track.album_release_date)
                score.title = `${track.artist_name} - ${track.name} (${releaseDate.getFullYear()})`
            }

            container.appendChild(score)
        }
    }

    private async loadSongs(playlistId: string): Promise<void> {
        const response = await fetch(`/hitster/${playlistId}`)
        const tracks: HitsterTrackInfo[] = await response.json()

        this.score = [];
        this.tracks = tracks;

        this.updateScore()
        await this.loadCurrentTrack()
    }

    private async loadCurrentTrack(): Promise<void> {
        this.audioPreview.src = this.currentTrack.preview_url;
        await this.audioPreview.play()
    }

    private async onGuessClick(): Promise<void> {
        const year = <HTMLInputElement>document.getElementById("year")
        const yearGuess = parseInt(year.value)

        const releaseDate = new Date(this.currentTrack.album_release_date)

        let result: HitsterGuessResult;

        if (yearGuess === releaseDate.getFullYear()) {
            result = HitsterGuessResult.Perfect;
        } else if (Math.floor(yearGuess / 10) === Math.floor(releaseDate.getFullYear() / 10)) {
            result = HitsterGuessResult.Good;
        } else {
            result = HitsterGuessResult.Wrong;
        }

        year.value = releaseDate.getFullYear().toString()
        this.onYearInput()

        this.score.push(result);
        this.updateScore()

        if (this.currentTrack) {
            await this.loadCurrentTrack()
        } else {
            this.audioPreview.pause()
            const scoreTitle = document.getElementById("score-title")
            scoreTitle.innerText += ` - ${this.currentScore} / 10`
        }
    }
}