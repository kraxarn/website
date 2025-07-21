(() => {
    const background = document.getElementById("background")
    if (!background) {
        return
    }
    const path = `/img/bg/${Math.floor(new Date().getHours() / 2)}.webp`
    background.style.backgroundImage = `url("${path}")`
})()