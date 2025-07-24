(() => {
    const body = document.querySelector("body");
    document.addEventListener("mousemove", (event) => {
        body.style.backgroundPosition = `left ${event.clientX}px top ${event.clientY}px`;
    });
})();
//# sourceMappingURL=spotlight.js.map