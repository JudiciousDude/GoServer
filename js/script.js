fetch('/getlist')
    .then(response => response.json())
    .then(json => {
            document.querySelector('.content').innerHTML = `
            ${json.map((item, index) => {
                return `${item.id} | ${item.name} | ${item.quantity} | ${item.conditions}`;
            }).join('<br>')}
        `;
    });
