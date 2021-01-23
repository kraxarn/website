const getById = (elementId) => document.getElementById(elementId)

const setBackground = () => {
	const background = getById("background")
	if (background) {
		background.style.backgroundImage = `url("/img/bg/${Math.floor(new Date().getHours() / 2)}.webp")`
	}
}
setBackground()