// Open avatar selection
getById("avatar").onclick = () => {
    const avatarSelect = getById("avatarSelect")
	if (avatarSelect.style.display === "block") {
		avatarSelect.style.display = "none"
	} else {
		avatarSelect.style.display = "block"
	}
}

// Switching between viewing/editing username
getById("username").onclick = () => {
	getById("username").style.display = "none"
	getById("nameEntry").style.display = "flex"
}

getById("saveName").onclick = () => {
    const username = getById("username")
	username.style.display = "block"
	getById("nameEntry").style.display = "none"
	username.textContent = getById("nameInput").value
	setName()
}

// Show room name selection
getById("createRoom").onclick = () => {
	getById("createRoom").style.display = "none"
	getById("roomEntry").style.display = "flex"
}

// Save room button
getById("saveRoom").onclick = () => {
	const name = getById("roomInput").value.replace(/\s/g, "").toLowerCase()
	if (name.length < 3 || name.length > 16) {
		return
	}
	console.log("Room name: %s", name)
}

// Update avatar image
const setAvatar = name => {
	getById("avatar").src = `img/${name}.svg`
	getById("avatarSelect").style.display = "none"

	updateUserInfo({
		avatar: name
	})
}

// Update name
const setName = () => {
    const nameInput = getById("nameInput")
	if (nameInput.value.length < 3) {
		return
	}
	updateUserInfo({
		name: nameInput.value
	})
}

const updateUserInfo = body =>
	fetch("/api/user/set_info", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(body)
	})
		.then(response => response.json())
		.then(json => {
			if (json.error) {
				console.log(json.message)
			} else {
				console.log("Update successful")
			}
		})
		.catch(err => console.log(err))