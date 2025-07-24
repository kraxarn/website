(() => {
    const loading = (element) => {
        let count = element.textContent.match(/\./g)?.length;
        if (!count && element.textContent.length > 0) {
            return;
        }
        if (count >= 4) {
            count = -1;
        }
        element.textContent = ".".repeat((count || 0) + 1);
        setTimeout(() => loading(element), 800);
    };
    const rows = document.querySelectorAll("#content table tr");
    for (const row of rows) {
        const type = row.querySelector("td:first-child");
        const status = row.querySelector("td:last-child");
        if (!type || !status) {
            continue;
        }
        loading(status);
    }
})();
//# sourceMappingURL=servers.js.map