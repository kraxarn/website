(() => {
    const body = document.querySelector("body")
    document.addEventListener("mousemove", (event: MouseEvent) => {
        body.style.backgroundPosition = `left ${event.clientX}px top ${event.clientY}px`
    })
})()