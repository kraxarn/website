const baseUrl = "http://localhost:5000"

const id = (elementId) => document.getElementById(elementId)

const get = url => new Promise((resolve, reject) =>
	fetch(`${baseUrl}${url}`, {
		mode: "cors"
	}).then(response => response.json())
		.then(json => resolve(json))
		.catch(err => {
			console.error(`${url}: ${err}`)
			reject(err)
		}))

const post = (url, data) => new Promise((resolve, reject) =>
	fetch(`${baseUrl}${url}`, {
		mode: "cors",
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(data)
	}).then(response => response.json())
		.then(json => resolve(json))
		.catch(err => {
			console.log(`${url}: ${err}`)
			reject(err)
		}))