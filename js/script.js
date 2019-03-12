getList();
function getList() {
    fetch('/getlist')
        .then(response => response.json())
        .then(json => {
            document.getElementById('tbody').innerHTML = `
                ${json.map((item, index) => {
                    if(item.conditions == undefined){item.conditions = "UNKNOWN"};
                    return `
                    <tr data-id="${item.id}" class="data-row" data-index="${index}">
                    <th scope="row">${index}</th>
                    <td>${item.name}</td>
                    <td>${item.quantity}</td>
                    <td>${item.conditions}</td>
                    </tr>
                    `;
                }).join(' ')}
            `;
        })
        .then(x => {
                document.querySelectorAll('.data-row').forEach(element => {
                element.addEventListener('click', (event) => {
                    if(confirm("Delete row " + `${element.dataset.index}` + "?"))
                        deleteEl(element);
                });
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

function addRow(){
    resName = document.getElementById("res-name").value;
    resQ = document.getElementById("res-quant").value;
    resCond = document.getElementById("res-cond").value;

    resName = resName.trim();
    document.getElementById("res-name").value = resName;
    resQ = resQ.trim();
    document.getElementById("res-quant").value = resQ;
    resCond = resCond.trim();
    document.getElementById("res-cond").value = resCond;

    if(resName == ''){
        alert("Enter resourse name");
        return;
    }

    resQ = resQ.replace('\s', '');
    if(isNaN(resQ) || resQ == ''){
        alert("Quantity must be a number!");
        return;
    }

    fetch('/', {method: 'POST', json: true, body: JSON.stringify({name: resName, quantity: resQ, conditions: resCond})})
        .then(resp => {getList();})
        .catch(x => {});

    document.getElementById("res-name").value = "";
    document.getElementById("res-quant").value = "";
    document.getElementById("res-cond").value = "";
}