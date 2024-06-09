var HitsterGuessResult;
(function (HitsterGuessResult) {
    HitsterGuessResult[HitsterGuessResult["None"] = 0] = "None";
    HitsterGuessResult[HitsterGuessResult["Wrong"] = 1] = "Wrong";
    HitsterGuessResult[HitsterGuessResult["Good"] = 2] = "Good";
    HitsterGuessResult[HitsterGuessResult["Perfect"] = 3] = "Perfect";
})(HitsterGuessResult || (HitsterGuessResult = {}));
class Hitster {
    tracks;
    score;
    constructor() {
        const playlists = document.getElementById("playlists");
        playlists.addEventListener("click", event => this.onPlaylistsClick(event));
        const year = document.getElementById("year");
        year.addEventListener("input", event => this.onYearInput(event));
        const guess = document.getElementById("guess");
        guess.addEventListener("click", () => this.onGuessClick());
    }
    get currentTrack() {
        return this.tracks[this.score.length];
    }
    get currentScore() {
        let score = 0;
        for (const result of this.score) {
            if (result === HitsterGuessResult.Good) {
                score += 1;
            }
            else if (result === HitsterGuessResult.Perfect) {
                score += 2;
            }
        }
        return score;
    }
    get audioPreview() {
        const element = document.getElementById("audio-preview");
        return element;
    }
    async onPlaylistsClick(event) {
        const target = event.target;
        const playlistId = target.dataset.playlistId;
        if (!playlistId) {
            return;
        }
        const playlists = document.getElementById("playlists");
        playlists.style.display = "none";
        const playlistsTitle = playlists.previousElementSibling;
        playlistsTitle.style.display = "none";
        const game = document.getElementById("game");
        game.style.display = "";
        this.onYearInput();
        await this.loadSongs(playlistId);
    }
    onYearInput(event) {
        const target = (event?.target ?? document.getElementById("year"));
        const current = parseInt(target.value);
        const minYear = parseInt(target.min);
        const maxYear = parseInt(target.max);
        const yearLabel = document.getElementById("year-label");
        yearLabel.innerText = target.value;
        yearLabel.style.left = `${((current - minYear) / (maxYear - minYear)) * 88}%`;
    }
    updateScore() {
        const container = document.getElementById("score-container");
        container.innerText = "";
        for (let i = 0; i < 10; i++) {
            const score = document.createElement("a");
            const guess = this.score[i] ?? HitsterGuessResult.None;
            switch (guess) {
                case HitsterGuessResult.None:
                    score.textContent = "--";
                    break;
                case HitsterGuessResult.Wrong:
                    score.textContent = ":(";
                    score.classList.add("wrong");
                    break;
                case HitsterGuessResult.Good:
                    score.textContent = ":)";
                    score.classList.add("good");
                    break;
                case HitsterGuessResult.Perfect:
                    score.textContent = ":D";
                    score.classList.add("perfect");
                    break;
            }
            if (guess !== HitsterGuessResult.None) {
                const track = this.tracks[i];
                const releaseDate = new Date(track.album_release_date);
                score.title = `${track.artist_name} - ${track.name} (${releaseDate.getFullYear()})`;
                score.href = `https://open.spotify.com/track/${track.id}`;
            }
            container.appendChild(score);
        }
    }
    async loadSongs(playlistId) {
        const response = await fetch(`/hitster/${playlistId}`);
        const tracks = await response.json();
        this.score = [];
        this.tracks = tracks;
        this.updateScore();
        await this.loadCurrentTrack();
    }
    async loadCurrentTrack() {
        this.audioPreview.src = this.currentTrack.preview_url;
        await this.audioPreview.play();
    }
    async onGuessClick() {
        const year = document.getElementById("year");
        const yearGuess = parseInt(year.value);
        const releaseDate = new Date(this.currentTrack.album_release_date);
        let result;
        if (yearGuess === releaseDate.getFullYear()) {
            result = HitsterGuessResult.Perfect;
        }
        else if (Math.floor(yearGuess / 10) === Math.floor(releaseDate.getFullYear() / 10)) {
            result = HitsterGuessResult.Good;
        }
        else {
            result = HitsterGuessResult.Wrong;
        }
        year.value = releaseDate.getFullYear().toString();
        this.onYearInput();
        this.score.push(result);
        this.updateScore();
        if (this.currentTrack) {
            await this.loadCurrentTrack();
        }
        else {
            this.audioPreview.pause();
            const scoreTitle = document.getElementById("score-title");
            scoreTitle.innerText += ` - ${this.currentScore} / 10`;
        }
    }
}
//# sourceMappingURL=hitster.js.map