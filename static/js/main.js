const getById = (elementId) => document.getElementById(elementId)

getById("background").style.backgroundImage = `url("/img/bg/${Math.floor(new Date().getHours() / 2)}.webp")`
