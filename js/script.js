getList();

window.onload = function(){
    document.querySelectorAll('.data-row').forEach(element => {
        //works only once - fix this
        element.addEventListener('click', (event) => {
            if(confirm("Delete row " + `${element.dataset.id}` + "?"))
                deleteEl(element);
        });
    }); 
}

function deleteEl(element){
    fetch(`/delete/${element.dataset.id}`)
        .then(response => {
            getList();
        })
        .catch(x => {
            console.log("fail");
        });
}

function getList() {
    fetch('/getlist')
        .then(response => response.json())
        .then(json => {
            document.getElementById('tbody').innerHTML = `
                ${json.map((item, index) => {
                    return `
                    <tr data-id="${item.id}" class="data-row">
                    <th scope="row">${index}</th>
                    <td>${item.name}</td>
                    <td>${item.quantity}</td>
                    <td>${item.conditions}</td>
                    </tr>
                    `;
                }).join(' ')}
            `;
        });
}